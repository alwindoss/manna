package bible

import "io"

func NewBibleService() *BibleService {
	bs := new(BibleService)
	bs.bibleBooks = bibleBooks
	bs.bibleChapters = bibleChapters
	bs.bibleVerses = bibleVerses
	return bs
}

type BibleService struct {
	bibleBooks    []string
	bibleChapters map[string]int
	bibleVerses   map[string][]int
	bibles        []Bible
	notes         []Note
}

func (bs *BibleService) GetBooksOfTheBible() []string {
	return bs.bibleBooks
}

func (bs *BibleService) GetChaptersInTheBook(book string) int {
	c := bibleChapters[book]
	return c
}

func (bs *BibleService) GetVersesInTheChapter(book string, chapter int) int {
	b := bibleVerses[book]
	v := b[chapter-1]
	return v
}

type ImportResult struct {
}

func (bs *BibleService) Import(f *io.Reader) error {
	return nil
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
