package main

import (
	"fmt"
	"sync"
	"time"
)

type Request struct {
	i int
}

func process(request Request) {
	time.Sleep(5 * time.Second)
	fmt.Printf("Process completed for request = %+v\n", request.i)
}

type Executor struct {
	channel      chan Request
	wg           sync.WaitGroup
	maxWorkers   int
	stopChannels []chan struct{}
}

func NewExecutor(maxWorkers int) *Executor {
	return &Executor{
		channel:      make(chan Request),
		maxWorkers:   maxWorkers,
		stopChannels: make([]chan struct{}, maxWorkers),
	}
}

func (e *Executor) startWorker(workerID int) {
	e.wg.Add(1)
	stop := make(chan struct{})
	e.stopChannels[workerID] = stop

	go func() {
		defer e.wg.Done()
		for {
			select {
			case request, ok := <-e.channel:
				if !ok {
					return
				}
				process(request)
			case <-stop:
				return
			}
		}
	}()
}

func (e *Executor) Start() {
	for i := 0; i < e.maxWorkers; i++ {
		e.startWorker(i)
	}
}

func (e *Executor) Stop() {
	for _, stop := range e.stopChannels {
		close(stop)
	}
	e.wg.Wait()
	close(e.channel)
}

func (e *Executor) ExecuteWorkflow(workflowRequest Request) {
	go func() {
		e.channel <- workflowRequest
	}()
}

func main() {
	executor := NewExecutor(512)
	executor.Start()

	startTime := time.Now()

	for i := 1; i <= 5000; i++ {
		req := Request{i}
		executor.ExecuteWorkflow(req)
	}
	fmt.Println("Completed submitting to request ", time.Now().Sub(startTime))

	//executor.Stop()
	time.Sleep(5 * time.Minute)
}


