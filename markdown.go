package main

import (
	"html/template"
	"io/ioutil"

	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"
)

func parse(input []byte) template.HTML {
	output := blackfriday.Run(input)
	html := bluemonday.UGCPolicy().SanitizeBytes(output)
	return template.HTML(string(html))
}
func readFileForTest(filename string) ([]byte, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}
