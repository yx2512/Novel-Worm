package parser

import (
	"regexp"

	"github.com/yx2512/crawler/engine"
)

const addrPrefix = `http://www.xz23.com`

const bookListRe = `<li><a href="(/[a-z.]+)"[^>]*>([^<]*)</a></li>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(bookListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}

	for _, content := range matches {
		result.Items = append(result.Items, "category "+string(content[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        addrPrefix + string(content[1]),
			ParserFunc: ParseCity})
	}
	return result
}
