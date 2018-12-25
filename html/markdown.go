package html

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
func getIMG(md template.HTML) string {
	if strings.Contains(string(md), "<img") {
		imgString := strings.Split(strings.Trim(string(md), `<p>img src="" "  </p>`), `"`)
		if len(imgString) > 0 && strings.Contains(imgString[0], "htt") {
			return imgString[0]
		}
	}
	return ""
}
