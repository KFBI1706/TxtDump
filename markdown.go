package main

import (
	"html/template"

	"github.com/shurcooL/github_flavored_markdown"
)

func parse(input string) template.HTML {
	//input = bluemonday.UGCPolicy().Sanitize(input)
	output := github_flavored_markdown.Markdown([]byte(input))
	return template.HTML(string(output))
}
