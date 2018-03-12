package main

import (
	"html/template"
	"strings"

	"github.com/shurcooL/github_flavored_markdown"
)

func parse(input string) template.HTML {
	output := github_flavored_markdown.Markdown([]byte(input))
	return template.HTML(string(output))
}
func getMDHeader(md template.HTML) string {
	if strings.Contains(string(md), "</h1>") {
		split := strings.SplitAfter(string(md), "</h1>")
		return split[0]
	}
	return ""
}
