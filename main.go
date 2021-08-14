package main

import (
	"github.com/yx2512/crawler/engine"
	"github.com/yx2512/crawler/scheduler"
	"github.com/yx2512/crawler/xz23/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:        "http://www.xz23.com",
		ParserFunc: parser.ParseCityList,
	})
}
