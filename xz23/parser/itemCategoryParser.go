package parser

import (
	"regexp"

	"github.com/yx2512/crawler/engine"
)

const addrPrefix = `https://www.xbiquge.so`

var bookCategoryRe = regexp.MustCompile(`<a href="(/[a-z/]+)"[^>]*>(.{2}小说)</a>`)

func ParseCategory(contents []byte) engine.ParseResult {
	matches := bookCategoryRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}

	for _, content := range matches {
		url := string(content[1])
		category := string(content[2])
		result.Items = append(result.Items, "category "+category)
		result.Requests = append(result.Requests, engine.Request{
			Url:        addrPrefix + url,
			ParserFunc: ParseItems})
	}
	return result
}
