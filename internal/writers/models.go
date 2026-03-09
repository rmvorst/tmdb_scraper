package writers

import "encoding/xml"

type UniqueID struct {
	Type    string `xml:"type,attr"`
	Default bool   `xml:"default,attr"`
	Value   int    `xml:",chardata"`
}

type metadata struct {
	XMLName xml.Name `xml:"tvshow"`

	ID      UniqueID `xml:"uniqueid"`
	Title   string   `xml:"title"`
	Season  int      `xml:"season"`
	Episode int      `xml:"episode"`
	Plot    string   `xml:"plot"`
}
