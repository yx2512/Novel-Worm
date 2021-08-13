package parser

import (
	"regexp"
	"strconv"

	"github.com/yx2512/crawler/engine"
	"github.com/yx2512/crawler/model"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([0-9]+岁)</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([未离]{1}[婚异]{1})</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([0-9]+cm)</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([0-9]+kg)</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>(月收入:[0-9\-.][万千]{1})</div>`)
var occupationRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>(工作地:.+)</div>`)
var residentRe = regexp.MustCompile(`<div class="m-btn pink"[^>]*>(籍贯:.+)</div>`)
var constellationRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^座]*座[^<]*)</div>`)
var houseRe = regexp.MustCompile(`<div class="m-btn pink"[^>]*>(.{1}购房)</div>`)
var carRe = regexp.MustCompile(`<div class="m-btn pink"[^>]*>(.{1}买车)</div>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}

	profile.Name = name
	if age, err := strconv.Atoi(extractString(contents, ageRe)); err == nil {
		profile.Age = age
	}
	profile.Height = extractString(contents, heightRe)
	profile.Weight = extractString(contents, weightRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Residence = extractString(contents, residentRe)
	profile.Constellation = extractString(contents, constellationRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)

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
