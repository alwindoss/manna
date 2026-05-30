package app

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/dailymanna/manna/internal/services/bible"
	"github.com/pressly/goose/v3"

	"github.com/wailsapp/wails/v3/pkg/application"

	_ "modernc.org/sqlite"

	_ "github.com/dailymanna/manna/migrations"
)

func init() {
	// Register a custom event whose associated data type is string.
	// This is not required, but the binding generator will pick up registered events
	// and provide a strongly typed JS/TS API for them.
	application.RegisterEvent[string]("time")
}

type Config struct {
	FS fs.FS
}

func NewMannaApp(cfg *Config) *application.App {
	dbPath := "./tmp/manna/data/manna.db"
	appEnv, appEnvfound := os.LookupEnv("APP_ENV")
	developmentMode := true
	if !appEnvfound {
		appEnv = "production"
		developmentMode = false
	}
	log.Printf("Running manna in %s mode", appEnv)
	if developmentMode {
		if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
			log.Fatalf("failed to create db directory: %v", err)
		}
	}
	dbPathWithOptions := fmt.Sprintf("%s?_pragma=journal_mode(WAL)&_pragma=synchronous(NORMAL)&_pragma=cache_size(-64000)", dbPath)
	db, err := sql.Open("sqlite", dbPathWithOptions)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	defer db.Close()
	if err := goose.SetDialect("sqlite3"); err != nil {
		log.Fatalf("failed to set goose dialect: %v", err)
	}
	goose.SetBaseFS(cfg.FS)

	log.Println("Running database migrations...")
	if err := goose.UpContext(context.Background(), db, "migrations"); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	log.Println("Migrations completed successfully!")

	// Create a new Wails application by providing the necessary options.
	// Variables 'Name' and 'Description' are for application metadata.
	// 'Assets' configures the asset server with the 'FS' variable pointing to the frontend files.
	// 'Bind' is a list of Go struct instances. The frontend has access to the methods of these instances.
	// 'Mac' options tailor the application when running an macOS.
	app := application.New(application.Options{
		Name:        "manna",
		Description: "A demo of using raw HTML & CSS",
		// Services: []application.Service{
		// 	application.NewService(bible.NewBibleService()),
		// },
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(cfg.FS),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})
	bibleSvcCfg := &bible.BibleServiceConfig{
		App:    app,
		DataFS: cfg.FS,
		DB:     db,
	}
	bibleSvc := bible.NewBibleService(bibleSvcCfg)

	app.RegisterService(application.NewService(bibleSvc))

	m := newMenu(app)

	// Set the application menu
	app.Menu.Set(m)

	// Create a new window with the necessary options.
	// 'Title' is the title of the window.
	// 'Mac' options tailor the window when running on macOS.
	// 'BackgroundColour' is the background colour of the window.
	// 'URL' is the URL that will be loaded into the webview.
	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:              "Manna - Renew your mind",
		UseApplicationMenu: true,
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar: application.MacTitleBar{
				UseToolbar:                true,
				AppearsTransparent:        false,
				Hide:                      false,
				HideTitle:                 false,
				FullSizeContent:           false,
				HideToolbarSeparator:      false,
				ShowToolbarWhenFullscreen: false,
				ToolbarStyle:              application.MacToolbarStyleAutomatic,
			},
			// TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(27, 38, 54),
		Width:            1080,
		MinWidth:         1080,
		URL:              "/",
	})

	// Create a goroutine that emits an event containing the current time every second.
	// The frontend can listen to this event and update the UI accordingly.
	go func() {
		for {
			now := time.Now().Format(time.RFC1123)
			app.Event.Emit("time", now)
			time.Sleep(time.Second)
		}
	}()

	return app
}

func newMenu(a *application.App) *application.Menu {
	menu := a.NewMenu()
	// Add standard menus (platform-appropriate)
	if runtime.GOOS == "darwin" {
		menu.AddRole(application.AppMenu) // macOS only
	}
	// Add a top-level menu
	fileMenu := menu.AddSubmenu("File")

	// Add menu items
	fileMenu.Add("New").SetAccelerator("CmdOrCtrl+N").OnClick(func(ctx *application.Context) {
		// Handle New
	})

	fileMenu.Add("Open").SetAccelerator("CmdOrCtrl+O").OnClick(func(ctx *application.Context) {
		// Handle Open
	})

	fileMenu.AddSeparator()

	fileMenu.Add("Import...").OnClick(func(ctx *application.Context) {})
	fileMenu.Add("Export...").OnClick(func(ctx *application.Context) {})

	fileMenu.AddSeparator()

	fileMenu.Add("Quit").SetAccelerator("CmdOrCtrl+Q").OnClick(func(ctx *application.Context) {
		a.Quit()
	})
	// menu.AddRole(application.FileMenu)
	menu.AddRole(application.EditMenu)
	menu.AddRole(application.WindowMenu)
	menu.AddRole(application.HelpMenu)

	// fileMenu := menu.FindByRole(application.FileMenu).GetSubmenu()
	// fileMenu.Add("Import...").OnClick(func(ctx *application.Context) {})
	// fileMenu.Add("Export...").OnClick(func(ctx *application.Context) {})

	return menu
}
