package worker

import (
	"fmt"
	"log"
	"network/primitives"
)

type Worker struct {
	Id           uint
	RequestCount uint64
	Requests     chan primitives.Request
	Responses    chan primitives.Response
}

func CreateWorker(id int) *Worker {
	worker := &Worker{}
	worker.Requests = make(chan primitives.Request, 100)
	worker.Responses = make(chan primitives.Response, 100)
	// Start se requests
	go worker.processedRequests()
	return worker
}

func (w *Worker) processedRequests() {
	for response := range w.Responses {
		log.Println(response)
	}
}

func (w *Worker) AddRequest(request primitives.Request) {
	w.RequestCount++
	log.Println("Adding Request to worker ", w.Id)
	// handle the request
	go w.handleRequest(request)
}

func (w *Worker) handleRequest(request primitives.Request) {
	request.Request()
	w.RequestCount--
	msg := fmt.Sprintf("Request %v.", request.Val)
	response := primitives.Response{Response: msg}
	w.Responses <- response
}
