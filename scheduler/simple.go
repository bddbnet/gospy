package scheduler

import "github.com/bddbnet/gospy/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
	//panic("implement me")
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

// 添加任务
func (s *SimpleScheduler) Submit(r engine.Request) {
	// send request down to worker chan
	go func() { s.workerChan <- r }()
}
