package scheduler

import (
	"goStu/crawler/engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	workChan chan chan engine.Request
}

func (q *QueuedScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

func (q *QueuedScheduler)WokerReady(w chan engine.Request)  {
	q.workChan <- w
}

func (q *QueuedScheduler) ConfigureMasterWorkerChan(chan engine.Request) {
	panic("implement me")
}

func (q *QueuedScheduler)Run()  {
	q.workChan = make(chan chan engine.Request)
	q.requestChan = make(chan engine.Request)
	go func() {
		var requestQ  []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWoker chan engine.Request
			if len(requestQ)>0 && len(workerQ)>0{
				activeRequest = requestQ[0]
				activeWoker = workerQ[0]
			}
			select {
			case r := <- q.requestChan:
				requestQ = append(requestQ,r)
			case w := <- q.workChan:
				workerQ = append(workerQ,w)
			case activeWoker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}


