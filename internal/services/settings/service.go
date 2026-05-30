package settings

import (
	"archive/zip"
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type SettingsServiceConfig struct {
	App    *application.App
	DataFS fs.FS
	DB     *sql.DB
}

func NewSettingsService(cfg *SettingsServiceConfig) *SettingsService {
	ss := new(SettingsService)
	ss.db = cfg.DB
	return ss
}

type SettingsService struct {
	app    *application.App
	dataFS fs.FS
	db     *sql.DB
}

func (ss *SettingsService) SetupOfflineCrossReferences() error {
	tableName := "cross_references"
	if !ss.IsTableEmpty(tableName) {
		return fmt.Errorf("The table %s already has data", tableName)
	}
	crURL := "https://github.com/dailymanna/database/releases/download/v0.0.1-alpha3/cross_references.zip"
	resp, err := http.Get(crURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	tempDir, err := os.MkdirTemp("", "manna-temp-cache-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempDir)
	tempZip, err := os.CreateTemp(tempDir, "download-*.zip")
	if err != nil {
		log.Fatalf("Failed to create temp file: %v", err)
	}
	writtenCount, err := io.Copy(tempZip, resp.Body)
	if err != nil {
		return err
	}
	log.Printf("Wrote %d to file %s", writtenCount, tempZip.Name())

	// zip.NewReader requires a file descriptor, which is why we saved it first
	zipReader, err := zip.OpenReader(tempZip.Name())
	if err != nil {
		log.Fatalf("Failed to open zip archive: %v", err)
	}
	defer zipReader.Close()

	var sqlFile *zip.File
	for _, file := range zipReader.File {
		if file.Name == "cross_references.sql" {
			sqlFile = file
			break
		}
	}

	if sqlFile == nil {
		log.Fatalf("Target file 'cross_references.sql' not found inside the zip archive")
	}
	sqlStream, err := sqlFile.Open()
	if err != nil {
		log.Fatalf("Failed to open stream for SQL file: %v", err)
	}
	defer sqlStream.Close()

	err = ss.executeSQLiteStream(sqlStream)
	if err != nil {
		return err
	}

	return nil
}

func (ss *SettingsService) IsTableEmpty(tableName string) bool {
	tableIsEmpty := true
	var exists int
	checkQuery := fmt.Sprintf("SELECT 1 FROM %s LIMIT 1;", tableName)

	// QueryRowContext automatically manages timeouts if passed from your HTTP/main handler
	err := ss.db.QueryRow(checkQuery).Scan(&exists)
	if err == nil {
		log.Printf("Table '%s' already has data. Skipping download and execution to prevent duplicates.", tableName)
		tableIsEmpty = false
		return tableIsEmpty
	}

	return tableIsEmpty
}
func (ss *SettingsService) executeSQLiteStream(sqlStream io.Reader) error {
	scanner := bufio.NewScanner(sqlStream)
	var chunkBuilder strings.Builder

	// 8MB threshold keeps RAM flat while sending thousands of insert lines at a time
	const maxChunkSize = 8 * 1024 * 1024
	tx, err := ss.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	lineNum := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lineNum++

		// Skip comments and entirely empty lines
		if line == "" || strings.HasPrefix(line, "--") || strings.HasPrefix(line, "/*") {
			continue
		}

		// Append the line to our current memory chunk block
		chunkBuilder.WriteString(line)
		chunkBuilder.WriteString("\n")

		// Safely break chunks ONLY when a line ends with a semicolon.
		// This protects multi-line block statements like the 'CREATE TABLE' definition.
		if chunkBuilder.Len() >= maxChunkSize && strings.HasSuffix(line, ";") {
			queryChunk := chunkBuilder.String()
			chunkBuilder.Reset()
			// Execute directly on ss.db because the SQL file manages its own BEGIN/COMMIT
			_, err := tx.Exec(queryChunk)
			if err != nil {
				return fmt.Errorf("failed to execute SQL chunk: %w", err)
			}
		}
	}

	// Flush any remaining INSERT statements left in the buffer
	if chunkBuilder.Len() > 0 {
		_, err := tx.Exec(chunkBuilder.String())
		if err != nil {
			return fmt.Errorf("failed to execute final SQL chunk: %w", err)
		}
	}

	// Catch any internal reading or streaming errors
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading dump stream: %w", err)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
