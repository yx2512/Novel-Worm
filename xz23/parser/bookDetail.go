package parser

import (
	"regexp"
	"strconv"

	"github.com/yx2512/crawler/engine"
	"github.com/yx2512/crawler/model"
)

var statusRe = regexp.MustCompile(`<meta property="og:novel:status" content="(.+)" />`)
var updateTimeRe = regexp.MustCompile(`<meta property="og:novel:update_time" content="(.+)" />`)
var newestChapterRe = regexp.MustCompile(`<meta property="og:novel:latest_chapter_name" content="(.+)" />`)
var cntRe = regexp.MustCompile(`<meta property="og:novel:click_cnt" content="([0-9]+)" />`)
var recommendationRe = regexp.MustCompile(`<div id="listtj">&nbsp;推荐阅读：<a href="/[a-z/]+"[^>]*>([^<]*)</a>、<a href="/jipinquannengxuesheng/" title="极品全能学生" target="_blank">极品全能学生</a>、<a href="/zuiqiangyisheng/" title="最强医圣" target="_blank">最强医圣</a>、<a href="/wulingzhichuanchengjipintaxifu/" title="五零之穿成极品他媳妇" target="_blank">五零之穿成极品他媳妇</a>、<a href="/xiewangzhuiqifeicainitianxiaojieyishiqingcheng/" title="邪王追妻：废材逆天小姐（一世倾城）" target="_blank">邪王追妻：废材逆天小姐（一世倾城）</a>、<a href="/dushiqimenyisheng/" title="都市奇门医圣" target="_blank">都市奇门医圣</a>、<a href="/pingbuqingyun/" title="平步青云" target="_blank">平步青云</a>、<a href="/yixu/" title="医婿" target="_blank">医婿</a>、<a href="/dushiyaoxieyi/" title="都市逍遥邪医" target="_blank">都市逍遥邪医</a>、<a href="/zuiqiangkuangbing/" title="最强狂兵" target="_blank">最强狂兵</a>、<a href="/hansanqiansuyingxiaquanwenmianfeiyuedu/" title="韩三千苏迎夏全文免费阅读" target="_blank">韩三千苏迎夏全文免费阅读</a>、<a href="/daqingyinlong/" title="大清隐龙" target="_blank">大清隐龙</a>、<a href="/guaiyishengshouyexuan/" title="怪医圣手叶皓轩" target="_blank">怪医圣手叶皓轩</a>、<a href="/jueshigaoshou/" title="绝世高手" target="_blank">绝世高手</a>、<a href="/sanjietaobaodian/" title="三界淘宝店" target="_blank">三界淘宝店</a></div>
`)

func ParseProfile(contents []byte, author string) engine.ParseResult {
	profile := model.Profile{}

	profile.Author = author

	if cnt, err := strconv.Atoi(extractString(contents, cntRe)); err == nil {
		profile.Cnt = cnt
	}

	profile.UpdateTime = extractString(contents, updateTimeRe)
	profile.NewestChapter = extractString(contents, newestChapterRe)
	profile.Status = extractString(contents, statusRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
