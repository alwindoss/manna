// Package bible implements functionality that are essential to work with a bible,
// this package is intended to be used only as an internal package therefore it will not be exposed
package bible

import (
	"encoding/xml"
	"fmt"
	"io"

	"golang.org/x/net/html/charset"
)

type Bible struct {
	XMLName xml.Name `xml:"bible"`
	Text    string   `xml:",chardata"`
	B       []struct {
		Text string `xml:",chardata"`
		N    string `xml:"n,attr"`
		C    []struct {
			Text string `xml:",chardata"`
			N    string `xml:"n,attr"`
			V    []struct {
				Text string `xml:",chardata"`
				N    string `xml:"n,attr"`
			} `xml:"v"`
		} `xml:"c"`
	} `xml:"b"`
}

type OpenSongBible struct {
}

func Import(content io.Reader) error {
	dc := xml.NewDecoder(content)
	dc.CharsetReader = charset.NewReaderLabel
	b := new(Bible)
	err := dc.Decode(b)
	if err != nil {
		return err
	}
	fmt.Printf("%s--> %s\n", b.B[0].C[0].V[0].N, b.B[0].C[0].V[0].Text)
	return nil
}
