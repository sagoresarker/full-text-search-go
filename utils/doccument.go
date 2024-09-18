package utils

import (
	"compress/gzip"
	"encoding/xml"
	"os"
)

type doccument struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int    `xml:"id"`
}

func LoadDoccuments(path string) ([]doccument, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	dec := xml.NewDecoder(gz)
	dump := struct {
		Doccuments []doccument `xml:"doc"`
	}{}

	if err := dec.Decode(&dump); err != nil {
		return nil, err
	}
	docs := dump.Doccuments
	for i := range docs {
		docs[i].ID = i
	}
	return docs, nil
}
