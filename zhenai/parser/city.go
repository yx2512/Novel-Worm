package parser

import (
	"regexp"

	"github.com/yx2512/crawler/engine"
)

var nameRe = regexp.MustCompile(`<th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a></th>`)

func ParseCity(contents []byte) engine.ParseResult {
	matches := nameRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, content := range matches {
		name := string(content[2])
		url := string(content[1])
		result.Items = append(result.Items, "User "+string(content[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(contents []byte) engine.ParseResult {
				return ParseProfile(contents, name)
			}})
	}
	return result
}
