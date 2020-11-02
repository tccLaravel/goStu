package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	//把 request 请求放入调度器里
	Submit(Request)
	//初始化 workChan,给 workChan 赋值
	ConfigureMasterWorkerChan(chan Request)
	//
	WokerReady(chan Request)
	//
	Run()
}

func (e *ConcurrentEngine) Run (seeds ...Request)  {
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in,out)
	}
	//把 request 送入 调度器 scheduler 里
	for _, r := range  seeds{
		e.Scheduler.Submit(r)
	}
	var count = 0
	for {
		result := <- out
		//for _,item := range result.Items{
		//	log.Printf("Got item: %v \n",item)
		//}
		for _,request := range result.Requests {
			log.Printf("Got Url: %v \n",request.Url)
			count ++
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request,out chan ParseResult)  {
	go func() {
		for  {
			//tell schduler I'm ready
			request := <- in
			result ,err := woker(request)
			if err != nil{
				continue
			}
			out <- result
		}
	}()
}