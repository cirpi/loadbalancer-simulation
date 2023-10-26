package primitives

import (
	"time"
)

type Request struct {
	Request          func()
	Val              int
	ResponseReceiver chan Response
}

func MakeRequests(ch chan<- Request, doneCh chan bool) {
	for i := 1; i <= 500; i++ {
		//val := rand.Intn(2)
		i := i
		ch <- Request{Request: func() {
			time.Sleep(time.Second * time.Duration(0))
		}, Val: i}
	}
	close(ch)
	//done <- true
}
