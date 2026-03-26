package bible_test

import (
	"os"
	"testing"

	"github.com/dailymanna/manna/internal/bible"
)

func TestImport(t *testing.T) {
	pathToBible := "/home/alwin/Sandbox/data/bibles/KJV.xmm"
	content, err := os.Open(pathToBible)
	if err != nil {
		t.Errorf("expected the error to be nil but was %v", err)
	}
	err = bible.Import(content)
	if err != nil {
		t.Errorf("expected the error to be nil but was %v", err)
	}
	t.FailNow()
}
