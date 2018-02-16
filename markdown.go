package main

import (
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"
)

func parse(input []byte) string {
	output := blackfriday.Run(input)
	html := bluemonday.UGCPolicy().SanitizeBytes(output)
	return string(html)
}
