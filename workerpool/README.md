package main

This line specifies that the code belongs to the main package, which is the entry point for executing Go programs.

import (
	"fmt"
	"sync"
	"time"
)


Here we import necessary packages. fmt is used for formatted I/O, sync is used for managing goroutines and synchronization, and time is used for dealing with time-related operations.


type Request struct {
	i int
}


This defines a struct type called Request with a single integer field i. This struct represents a request that will be processed by the executor.

func process(request Request) {
	time.Sleep(5 * time.Second)
	fmt.Printf("Process completed for request = %+v\n", request)
}


The process function represents the time-consuming processing that will be performed on each request. It uses time.Sleep to simulate a 5-second delay and then prints a message indicating the completion of the process along with the request details.

type Executor struct {
	channel      chan Request
	wg           sync.WaitGroup
	maxWorkers   int
	stopChannels []chan struct{}
}


This defines the Executor struct type, which represents the executor for processing requests. It has the following fields:

channel: This is a channel of type Request used to receive requests for processing.
wg: This is an instance of sync.WaitGroup used to wait for all worker goroutines to finish.
maxWorkers: This is an integer field representing the maximum number of worker goroutines.
stopChannels: This is a slice of channels of type struct{} used to signal the worker goroutines to stop.


func NewExecutor(maxWorkers int) *Executor {
	return &Executor{
		channel:      make(chan Request),
		maxWorkers:   maxWorkers,
		stopChannels: make([]chan struct{}, maxWorkers),
	}
}


The NewExecutor function is a constructor for creating a new Executor instance. It takes the maxWorkers value as a parameter and returns a pointer to the Executor struct. Inside the function, it initializes the channel with a new channel of type Request, and it creates a slice of stopChannels with length equal to maxWorkers, initializing each channel with the zero value of struct{}.


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

The startWorker method starts a worker goroutine with the specified workerID. 
It increments the sync.WaitGroup counter using e.wg.Add(1) to indicate that a
new worker goroutine has started. It creates a new stop channel and assigns it 
to the stopChannels slice at the corresponding workerID.

Inside the goroutine, there is an infinite loop that uses a select statement to 
receive from two channels: e.channel and stop. If a request is received from 
e.channel, it invokes the process function to handle the request. 
If a value is received from stop, indicating that the worker should stop, 
the goroutine returns.

func (e *Executor) Start() {
	for i := 0; i < e.maxWorkers; i++ {
		e.startWorker(i)
	}
}

The Start method starts all the worker goroutines. It calls the startWorker method e.maxWorkers times, passing the worker ID from 0 to e.maxWorkers-1. This ensures that the desired number of worker goroutines is created and started.

func (e *Executor) Stop() {
	for _, stop := range e.stopChannels {
		close(stop)
	}
	e.wg.Wait()
}


The Stop method stops all the worker goroutines and waits for them to finish. 
It iterates over the stopChannels slice and closes each channel to signal the
respective worker goroutine to stop. It then calls e.wg.Wait() to wait for all 
the worker goroutines to complete using the sync.WaitGroup.

func (e *Executor) ExecuteWorkflow(workflowRequest Request) {
	e.channel <- workflowRequest
}


The ExecuteWorkflow method is used to submit a workflowRequest for processing. 
It simply sends the workflowRequest to the channel, which will be picked up by 
one of the worker goroutines for processing.

func main() {
	executor := NewExecutor(512)
	executor.Start()

	for i := 1; i <= 100; i++ {
		req := Request{i}
		executor.ExecuteWorkflow(req)
	}

	executor.Stop()
}


In the main function, we create a new Executor instance with a maximum of 512 worker 
goroutines. We then start the executor using executor.Start().

Next, we have a loop that submits 100 workflow requests for processing. 
Each request is created as a Request struct with the corresponding i value, and 
it is submitted using executor.ExecuteWorkflow(req).

Finally, we stop the executor using executor.Stop(), which stops all the worker 
goroutines and waits for them to finish before the program exits.
