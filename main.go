package main

import (
	"github.com/yx2512/crawler/engine"
	"github.com/yx2512/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.xz23.com",
		ParserFunc: parser.ParseCityList,
	})
}
