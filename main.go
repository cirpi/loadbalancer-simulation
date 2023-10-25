package main

import (
	"fmt"
	"log"
	"math/rand"
	"network/primitives"
	"network/worker"
	"os"
	"time"
)

// used to check if all the requests were completed

func main() {
	w := worker.CreateWorker(1)
	req := make(chan primitives.Request, 10)
	makeRequests(req)
	close(req)
	log.SetOutput(logFile())
	go func() {
		for r := range req {
			log.Println("Adding request ", r.Val)
			w.AddRequest(r)
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println("Count: ", w.RequestCount)
		}
	}()
	time.Sleep(time.Second * 70)
}

func makeRequests(req chan primitives.Request) {
	for i := 0; i < 10; i++ {
		i := i
		s := rand.Intn(5)
		curr := primitives.Request{
			Val: i + 1,
			Request: func() {
				time.Sleep(time.Second * time.Duration(s))
			},
		}
		req <- curr
		fmt.Println("time for ", i+1, " is ", s)
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
