package scheduler

import (
	"github.com/lightjameslyy/lt.go/imooclearngo/crawler/engine"
)

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	// send request down to worker chan
	// s.workerChan <- r

	// prevent from dead-lock
	go func() {
		s.workerChan <- r
	}()
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}
