package schedular

import "muke_distributed/crawier/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
	//一个request对应一个channel
	//每个worker去选择一个存request的channel
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request{
	return make(chan engine.Request)
}

//往scheduler的requestchan里面发任务
func (s *QueuedScheduler) Submit(r engine.Request){
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request){
	s.workerChan <- w
}

func (s *QueuedScheduler) Run(){
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func(){
		var requestQ []engine.Request    //请求队列
		var workerQ  []chan engine.Request //work队列，此队列中每个元素是可以放request的channel
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ)>0 && len(workerQ)>0{//用一个request类型的channel去接request
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ,r)  //收到request,排队
			case w := <-s.workerChan:
				workerQ = append(workerQ,w)    //收到worker，排队
			case activeWorker <- activeRequest: //用一个request类型的channel去接request,如果有闲的Worker和request，就可以处理。
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}