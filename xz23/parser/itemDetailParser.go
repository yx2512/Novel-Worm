package parser

import (
	"github.com/yx2512/crawler/engine"
	"github.com/yx2512/crawler/model"
	"regexp"
)

var authorRe = regexp.MustCompile(`<meta property="og:novel:author" content="([^"]+)"/>`)
var statusRe = regexp.MustCompile(`<meta property="og:novel:status" content="([^"]+)"/>`)
var updateTimeRe = regexp.MustCompile(`<p>更新时间：([^&]+)`)
var newestChapterRe = regexp.MustCompile(`<meta property="og:novel:latest_chapter_name" content="([^"]+)"/>`)
var recommendationRe = regexp.MustCompile(`<a href="(https://www.xbiquge.so/book/[0-9]+/)"[^>]*>([^<]+)</a>`)

func ParseItemProfile(contents []byte, title string) engine.ParseResult {
	profile := model.Profile{}

	profile.Title = title

	profile.Author = extractString(contents, authorRe)

	profile.UpdateTime = extractString(contents, updateTimeRe)
	profile.NewestChapter = extractString(contents, newestChapterRe)
	profile.Status = extractString(contents, statusRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	extractStringArr(contents, recommendationRe, &result)

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}

func extractStringArr(contents []byte, re *regexp.Regexp, result *engine.ParseResult) {
	allSubmatch := re.FindAllSubmatch(contents, -1)

	for _, item := range allSubmatch {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(item[1]),
			ParserFunc: func(content []byte) engine.ParseResult {
				return ParseItemProfile(content, string(item[2]))
			},
		})
	}
}
