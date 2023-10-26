package main

import (
	"fmt"
	"log"
	"math/rand"
	"network/loadbalancer"
	"network/primitives"
	"os"
	"time"
)

// used to check if all the requests were completed

func main() {
	lb := loadbalancer.Create(4)
	log.SetOutput(logFile())
	fmt.Println(lb)
	req := make(chan primitives.Request, 10000)
	makeRequests(req, lb)
	close(req)
	for request := range req {
		lb.AddRequest(request)
	}
	time.Sleep(time.Minute)
}

func makeRequests(req chan primitives.Request, lb *loadbalancer.LoadBalancer) {
	for i := 0; i < 10000; i++ {
		rand.Seed(time.Now().Unix())
		s := rand.Intn(5)
		curr := primitives.Request{
			Val: i + 1,
			Request: func() {
				time.Sleep(time.Second * time.Duration(s))
			},
			ResponseReceiver: lb.Responses,
		}
		req <- curr
		//fmt.Println("time for ", i+1, " is ", s)
	}
}

func logFile() *os.File {
	file, er := os.OpenFile("out.log", os.O_RDWR, 0222)
	if er != nil {
		fmt.Println(er)
	}
	fmt.Println(file)
	return file
}
