package parser

import (
	"regexp"

	"github.com/yx2512/crawler/engine"
)

var nameRe = regexp.MustCompile(`<span class="s2"><a href="(/[a-z/]+)"[^>]*>([^<]+)</a></span>`)
var urlPrefix = `http://www.xz23.com`

func ParseCity(contents []byte) engine.ParseResult {
	matches := nameRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, content := range matches {
		title := string(content[2])
		url := urlPrefix + string(content[1])
		result.Items = append(result.Items, "Title "+string(content[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(contents []byte) engine.ParseResult {
				return ParseProfile(contents, title)
			}})
	}
	return result
}
