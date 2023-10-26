package worker

import (
	"fmt"
	"log"
	"network/primitives"
	"sync"
)

type Worker struct {
	Id           uint
	RequestCount uint64
	Requests     chan primitives.Request
	Responses    chan primitives.Response
	lock         sync.Mutex
}

func CreateWorker(id int) *Worker {
	worker := &Worker{}
	worker.Id = uint(id)
	worker.Requests = make(chan primitives.Request, 100)
	worker.Responses = make(chan primitives.Response, 100)
	// Start se requests
	go worker.processedRequests()
	return worker
}

func (w *Worker) String() string {
	return fmt.Sprintf(`Worker{Id: %v, count: %v}`, w.Id, w.RequestCount)
}

func (w *Worker) processedRequests() {
	for response := range w.Responses {
		log.Println("Finished by Worker ", w.Id)
		response.Receiver <- response
	}
}

func (w *Worker) AddRequest(request primitives.Request) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.RequestCount++
	log.Println("Adding Request to worker ", w.Id)
	// handle the request
	go w.handleRequest(request)
}

func (w *Worker) handleRequest(request primitives.Request) {
	request.Request()
	w.lock.Lock()
	w.RequestCount--
	defer w.lock.Unlock()
	msg := fmt.Sprintf("Request %v.", request.Val)
	response := primitives.Response{Response: msg, Receiver: request.ResponseReceiver}
	w.Responses <- response
}
