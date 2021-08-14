package engine

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/yx2512/crawler/model"
)

var ctx = context.Background()

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{}
}

type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	WorkerReady(chan Request)
	Run()
}

func (ce *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	ce.Scheduler.Run()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	for _, r := range seeds {
		if intCmd := rdb.SAdd(ctx, "url", r.Url); intCmd.Val() == 1 {
			ce.Scheduler.Submit(r)
		}
	}

	for i := 0; i < ce.WorkerCount; i++ {
		createWorker(ce.Scheduler.WorkerChan(), out, ce.Scheduler)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			auxItem := item

			if _, ok := auxItem.(model.Profile); ok {
				go func() { ce.ItemChan <- auxItem }()
			}
		}

		for _, r := range result.Requests {
			if intCmd := rdb.SAdd(ctx, "url", r.Url); intCmd.Val() == 1 {
				ce.Scheduler.Submit(r)
			}
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, scheduler Scheduler) {
	go func() {
		for {
			scheduler.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
