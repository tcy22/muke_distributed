package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
	RequestProcessor Processor
}

type Processor func(Request) (ParseResult,error)

type Scheduler interface {
	Submit(Request)        //往scheduler的chan里面发任务
	WorkerChan()   chan Request   //我有一个work,请问给我哪个channel。
	WorkerReady(chan Request)
	Run()
}

type ReadyNotifier interface{
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ... Request){
	out := make(chan ParseResult)
	e.Scheduler.Run()   //等待新的任务的到来

	for i:=0; i < e.WorkerCount; i++ {
		//每个work goroutine都有自己的in channel,这里开了10个
		e.createWorker(e.Scheduler.WorkerChan(),out,e.Scheduler)
	}
    //把对engine的请求全部交给scheduler
	for _,r := range seeds {
		e.Scheduler.Submit(r)
	}
	for{
		result := <- out
		for _,item := range result.Items{
			go func(it Item){//此处若为func(),则go func里面的item与for里面的不对应。
				e.ItemChan <- it
			}(item)
		}
		//这里直接save()不行，因为拿到worker要尽快脱手，拿到request要尽快脱手。不能在这里花费太长时间去save
		//解决1、go save(item)
		//解决2、go func(){itemChan <- item}
		for _,request := range result.Requests {
			log.Println("request:",request)
			e.Scheduler.Submit(request)
		}
	}
}

func (e *ConcurrentEngine)createWorker(in chan Request,out chan ParseResult,s Scheduler){
	go func() {
		for{
			// 每个work goroutine都有自己的in channel，从channel里取出要执行的request
			s.WorkerReady(in)
			Request := <- in         //可以处理请求
			result,err := e.RequestProcessor(Request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}