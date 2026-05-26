package bible

import (
	"fmt"
	"os"
	"strconv"

	biblepkg "github.com/dailymanna/manna/pkg/bible"
	"github.com/wailsapp/wails/v3/pkg/application"
)

func NewBibleService(app *application.App) *BibleService {
	bs := new(BibleService)
	bs.app = app
	bs.bibleBooks = bibleBooks
	bs.bibleChapters = bibleChapters
	bs.bibleVerses = bibleVerses
	bs.translationsAvailable = []string{"KJV"}
	err := bs.importOpenSongBible("./data/bibles/KJV.xmm", "KJV")
	if err != nil {
		panic(err)
	}

	return bs
}

type BibleService struct {
	app                   *application.App
	bibleBooks            []string
	bibleChapters         map[string]int
	bibleVerses           map[string][]int
	bibles                []*Bible
	notes                 []Note
	translationsAvailable []string
}

func (bs *BibleService) GetBooksOfTheBible() []string {
	return bs.bibleBooks
}

func (bs *BibleService) GetCountOfChaptersInTheBook(book string) int {
	c := bibleChapters[book]
	return c
}

func (bs *BibleService) GetCountOfVersesInTheChapter(book string, chapter int) int {
	b := bibleVerses[book]
	v := b[chapter-1]
	return v
}

type ImportResult struct {
}

func (bs *BibleService) importOpenSongBible(path string, version string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	osb, err := biblepkg.ParseOpenSongBible(f)
	if err != nil {
		return err
	}
	bbl := new(Bible)
	bbl.Version = version
	for bkIndex, book := range osb.B {
		bk := new(Book)
		bk.Name = book.N
		for _, chapter := range book.C {
			ch := new(Chapter)
			chNum, _ := strconv.Atoi(chapter.N)
			ch.Number = chNum
			for _, verse := range chapter.V {
				v := new(Verse)
				bkID := bkIndex + 1
				v.BookID = bkID
				v.Chapter = chNum
				v.Verse, _ = strconv.Atoi(verse.N)
				v.Text = verse.Text
				ch.Verses = append(ch.Verses, v)
			}
			bk.Chapters = append(bk.Chapters, ch)
		}
		bk.ChapterCount = len(bk.Chapters)
		bbl.Books = append(bbl.Books, bk)
	}
	bs.bibles = append(bs.bibles, bbl)

	return nil
}

func (bs *BibleService) ShowNotification(title, message string) {
	infoDialog := bs.app.Dialog.Warning()
	infoDialog.Title = title
	infoDialog.Message = message
	infoDialog.Show()

}

type GetVersesResponse struct {
	Num  string `json:"num"`
	Text string `json:"text"`
	// Verses []*Verse1 `json:"verses"`
}

func (bs *BibleService) isVersionAvailable(v string) bool {
	versionAvailable := false
	for _, ver := range bs.translationsAvailable {
		if ver == v {
			versionAvailable = true
			return versionAvailable
		}
	}
	return versionAvailable
}

func (bs *BibleService) GetVerses(version string, book string, chapter int) ([]*GetVersesResponse, error) {
	if version == "" || !bs.isVersionAvailable(version) {
		return nil, fmt.Errorf("version of bible not found")
	}
	var verses []*GetVersesResponse
	for _, bbl := range bs.bibles {
		if bbl.Version == version {
			for _, bk := range bbl.Books {
				if bk.Name == book {
					for _, chap := range bk.Chapters {
						if chap.Number == chapter {
							// verses = chap.Verses
							for i, v := range chap.Verses {
								gvr := &GetVersesResponse{
									Num:  strconv.Itoa(i + 1),
									Text: v.Text,
								}
								verses = append(verses, gvr)
							}
						}
					}
				}
			}
		}
	}
	return verses, nil
}

func (bs *BibleService) AddNote(text string, book, chapter, verse int) error {
	return nil
}

func (bs *BibleService) UpdateNote(id int, text string, book, chapter, verse int) error {
	return nil
}

func (bs *BibleService) DeleteNote(id int, text string, book, chapter, verse int) error {
	return nil
}

func (bs *BibleService) GetNoteByID(id int) ([]*Note, error) {
	return nil, nil
}

func (bs *BibleService) GetNoteByVerse(book, chapter, verse int) ([]*Note, error) {
	return nil, nil
}

func (bs *BibleService) GetNotesByBook(book int) ([]*Note, error) {
	return nil, nil
}

func (bs *BibleService) GetNotesByChapter(book, chapter int) ([]*Note, error) {
	return nil, nil
}

func (bs *BibleService) GetAllNotes() ([]*Note, error) {
	return nil, nil
}
