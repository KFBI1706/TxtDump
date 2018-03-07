package main

import (
	"html/template"
	"strings"

	"github.com/shurcooL/github_flavored_markdown"
)

func parse(input string) template.HTML {
	//input = bluemonday.UGCPolicy().Sanitize(input)
	output := github_flavored_markdown.Markdown([]byte(input))
	return template.HTML(string(output))
}
func getMDHeader(md template.HTML) string {
	raw := string(md)
	if strings.Contains(raw, "</h1>") {
		split := strings.SplitAfter(raw, "</h1>")
		return split[0]
	}
	return ""
}
