package main

import (
	"html/template"

	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"
)

func parse(input []byte) template.HTML {
	output := blackfriday.Run(input)
	html := bluemonday.UGCPolicy().SanitizeBytes(output)
	return template.HTML(string(html))
}
