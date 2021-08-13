package main

import (
	"github.com/yx2512/crawler/engine"
	"github.com/yx2512/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
