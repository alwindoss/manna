package bible

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html/charset"
)

type Bible struct {
}

type OpenSongBible struct {
	XMLName xml.Name `xml:"bible" json:"bible,omitempty"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
	B       []struct {
		Text string `xml:",chardata" json:"text,omitempty"`
		N    string `xml:"n,attr" json:"n,omitempty"`
		C    []struct {
			Text string `xml:",chardata" json:"text,omitempty"`
			N    string `xml:"n,attr" json:"n,omitempty"`
			V    []struct {
				Text string `xml:",chardata" json:"text,omitempty"`
				N    string `xml:"n,attr" json:"n,omitempty"`
			} `xml:"v" json:"v,omitempty"`
		} `xml:"c" json:"c,omitempty"`
	} `xml:"b" json:"b,omitempty"`
}

func (osb *OpenSongBible) PrettyPrint() {
	separator := strings.Repeat("*", 30)
	fmt.Println(separator)
	fmt.Println(separator)

	for _, book := range osb.B {
		fmt.Printf("Book Name: %s\n", book.N)
		for _, chapter := range book.C {
			fmt.Printf("Chapter: %s\n", chapter.N)
			for _, verse := range chapter.V {
				fmt.Printf("%s:\t%s\n", verse.N, verse.Text)
			}
			fmt.Println(separator)
		}
		fmt.Println(separator)
	}
	fmt.Println(separator)

}

func ParseOpenSongBible(r io.Reader) (*OpenSongBible, error) {
	osb := new(OpenSongBible)
	decoder := xml.NewDecoder(r)
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(osb)
	if err != nil {
		return nil, err
	}
	return osb, nil
}
