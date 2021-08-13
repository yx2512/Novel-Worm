package parser

import (
	"regexp"

	"github.com/yx2512/crawler/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}

	limit := 10

	for _, content := range matches {
		result.Items = append(result.Items, "City "+string(content[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(content[1]),
			ParserFunc: ParseCity})
		limit--
		if limit == 0 {
			break
		}
	}
	return result
}
