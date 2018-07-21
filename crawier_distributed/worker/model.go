package worker

import (
	"muke_distributed/crawier/engine"
	"muke_distributed/crawier_distributed/config"
	"muke_distributed/crawier/muke/parser"
	"github.com/pkg/errors"
	"log"
	"fmt"
)

type RequestC struct {    //可以在网上传递的request
	Url string
	Parser SerializedParser
}

type ParseResult struct {  //可以在网上传递的Result
	Items []engine.Item
	Requests []RequestC
}

type SerializedParser struct {
	Name string      //函数名
	Args interface{} //参数，如profileParser中的name
}
//{"ParseCityList",nil},{"ProfileParse",userName}


/*
将engine.Request，engine.ParseResult与可以在网上传递的Request，Result转换。
*/
func SerializeRequest(r engine.Request) RequestC {
	name,args := r.Parser.Serialize()
	return RequestC{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult{
	result := ParseResult{
		Items:r.Items,
	}
	for _,req := range r.Requests {
		fmt.Println("SerializeResult",req)
		result.Requests = append(result.Requests,SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r RequestC) (engine.Request,error) {
	parser,err := deserializeParse(r.Parser)
	if err != nil {
		fmt.Println(err)
		return engine.Request{},nil
	}
	return engine.Request{
		Url: r.Url,
		Parser:parser,
	},nil
}

func DeserializeResult(r ParseResult) engine.ParseResult{
	result := engine.ParseResult{
		Items:r.Items,
	}
	for _,req := range r.Requests {
		engineReq,err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserialilzing request:%v",err)
			continue
		}
		//if engineReq.Parser != nil {
			result.Requests = append(result.Requests, engineReq)
		//}
		//fmt.Printf("req:%v,engineReq:%v",req,engineReq)
		fmt.Println()
	}
	return result
}

func deserializeParse(p SerializedParser) (engine.Parser,error){
	switch p.Name {
		case config.ParseCourseList:
			return engine.CreateFuncParser(parser.ParseCourseList,config.ParseCourseList),nil
		case config.ParseCourse:
			return engine.CreateFuncParser(parser.ParseCourse,config.ParseCourse),nil
		case config.NilParse:
			return engine.NilParser{},nil
		case config.ParserProfile:
			return engine.CreateFuncParser(parser.ParserProfile,config.ParserProfile),nil
		default:
			return nil,errors.New("unknown parser name")
	}
}
