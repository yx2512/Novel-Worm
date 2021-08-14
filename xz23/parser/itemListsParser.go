package parser

import (
	"regexp"

	"github.com/yx2512/crawler/engine"
)

var bookRe = regexp.MustCompile(`<span class="s2"><a href="(https://www.xbiquge.so/book/[0-9]+/)"[^>]*>([^<]+)</a></span>`)

func ParseItems(contents []byte) engine.ParseResult {
	matches := bookRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, content := range matches {
		title := string(content[2])
		url := string(content[1])
		result.Items = append(result.Items, "Title "+title)
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseItemProfile(c, title)
			}})
	}
	return result
}
