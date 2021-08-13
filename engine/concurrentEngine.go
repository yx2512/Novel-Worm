package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}

func (ce *ConcurrentEngine) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParseResult)
	ce.Scheduler.ConfigureMasterWorkerChan(in)

	for _, r := range seeds {
		ce.Scheduler.Submit(r)
	}

	for i := 0; i < ce.WorkerCount; i++ {
		createWorker(in, out)
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

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}

			out <- result
		}
	}()
}
