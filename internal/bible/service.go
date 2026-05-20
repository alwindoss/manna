package bible

import "io"

func NewBibleService() *BibleService {
	bs := new(BibleService)

	return bs
}

type BibleService struct {
	bibles []Bible
	notes  []Note
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
