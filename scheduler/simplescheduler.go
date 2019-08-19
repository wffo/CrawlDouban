package scheduler

import "jy/douban/engine"

type SimpleScheduler struct {
	workerChan chan  engine.Request
	ProxyChan chan string
}

func (s *SimpleScheduler)Submit(r engine.Request)  {
	go func() {
		s.workerChan <- r
	}()
}
func (s *SimpleScheduler) ConfigurWorkerChan(c chan engine.Request)  {
	s.workerChan = c
}