package bible

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	startTime := time.Now()
	f, err := os.Open("../../data/bibles/KJV.xmm")
	assert.Nilf(t, err, "expected err to be nil but was %v", err)
	b, err := ParseOpenSongBible(f)
	assert.Nilf(t, err, "expected err to be nil but was %v", err)
	fmt.Println("Completed Parsing after:", time.Since(startTime))
	b.PrettyPrint()
	fmt.Println("Completed Printing after:", time.Since(startTime))
	t.FailNow()
}
