package main

import (
	"github.com/yx2512/crawler/engine"
	"github.com/yx2512/crawler/persist"
	"github.com/yx2512/crawler/scheduler"
	"github.com/yx2512/crawler/xz23/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 20,
		ItemChan:    persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:        "https://www.xbiquge.so/",
		ParserFunc: parser.ParseCategory,
	})
}
