package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
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

	for _, r := range seeds {
		ce.Scheduler.Submit(r)
	}

	for i := 0; i < ce.WorkerCount; i++ {
		createWorker(ce.Scheduler.WorkerChan(),out,ce.Scheduler)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item: %v", item)
		}

		for _, request := range result.Requests {
			ce.Scheduler.Submit(request)
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
