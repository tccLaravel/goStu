package scheduler

import (
	"goStu/crawler/engine"
)

type SimpleScheduler struct {
	wokerChan chan engine.Request
}

//结构体里面的,需要先申明,然后才能使用
//我们先要 configure  wokerChan ，告诉他我们用哪个channel来收任务，然后才能用submit对它发送任务。
//configure 里面是给 workerChan 赋值，用了 =
func (s *SimpleScheduler) ConfigureMasterWorkerChan(r chan engine.Request) {
	s.wokerChan = r
}

//submit里面是给已经赋值的workerChan送任务进去，用了<-
func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.wokerChan <- r
	}()
}

