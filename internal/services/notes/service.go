package notes

import (
	"io/fs"

	"github.com/wailsapp/wails/v3/pkg/application"
)

func NewNotesService(app *application.App, dataFS fs.FS) *NotesService {
	ns := new(NotesService)

	return ns
}

type NotesService struct {
}

func (ns *NotesService) AddNote() {

}

func (ns *NotesService) GetAllNotesForTheVerse(chNum, verNum int) {

}
