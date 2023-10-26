package loadbalancer

import (
	"container/heap"
	"log"
	"network/primitives"
	"network/worker"
)

type LoadBalancer struct {
	Workers   []*worker.Worker
	Requests  chan primitives.Request
	Responses chan primitives.Response
}

func Create(workerCount int) *LoadBalancer {
	lb := LoadBalancer{}
	for i := 0; i < workerCount; i++ {
		lb.Workers = append(lb.Workers, worker.CreateWorker(i+1))
		//lb.Workers[i].RequestCount = uint64(4 - (i + 1))
	}
	lb.Requests = make(chan primitives.Request, 10)
	lb.Responses = make(chan primitives.Response, 10)
	heap.Init(&lb)
	go lb.processedRequests()
	return &lb
}

func (l *LoadBalancer) processedRequests() {
	for response := range l.Responses {
		log.Println("LoadBalancer got response ", response)
	}
}

func (l *LoadBalancer) AddRequest(request primitives.Request) {
	// get the least occupied worker from the worker pool
	leastOccupiedWorker := heap.Pop(l).(*worker.Worker)
	// send the request to that worker
	leastOccupiedWorker.AddRequest(request)
	// add the leastOccupiedWorker to the workers
	heap.Push(l, leastOccupiedWorker)
}

// implementing the sort and heap interface
func (lb LoadBalancer) Len() int { return len(lb.Workers) }
func (lb LoadBalancer) Less(i, j int) bool {
	return lb.Workers[i].RequestCount < lb.Workers[j].RequestCount
}
func (lb LoadBalancer) Swap(i, j int) {
	lb.Workers[i], lb.Workers[j] = lb.Workers[j], lb.Workers[i]
}

func (lb *LoadBalancer) Push(req any) {
	lb.Workers = append(lb.Workers, req.(*worker.Worker))
}

func (lb *LoadBalancer) Pop() any {
	res := lb.Workers[len(lb.Workers)-1]
	lb.Workers = lb.Workers[:len(lb.Workers)-1]
	return res
}
