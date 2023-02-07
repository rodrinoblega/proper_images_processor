package usecases

import "sync"

type SemaphoredWaitGroup struct {
	Sem chan bool
	Wg  sync.WaitGroup
}

func (s *SemaphoredWaitGroup) Add(delta int) {
	s.Wg.Add(delta)
	s.Sem <- true
}

func (s *SemaphoredWaitGroup) Done() {
	<-s.Sem
	s.Wg.Done()
}

func (s *SemaphoredWaitGroup) Wait() {
	s.Wg.Wait()
}
