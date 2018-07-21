package parser

import (
	"regexp"
	"muke_distributed/crawier/engine"
)

const courseListRe  = `<a href="(/course/list\?c=[0-9a-zA-Z]+)" [^>]*>([^<]*)</a>`
var urlp = `https://www.imooc.com`

//课程列表解析器
func ParseCourseList(contents []byte,_ string) engine.ParseResult{
	re := regexp.MustCompile(courseListRe)
	matches := re.FindAllSubmatch(contents,-1)     //[][][]byte
	result := engine.ParseResult{}
	limit := 10  //只获取5个城市
	for _,m := range matches {
        url := urlp + string(m[1])
		result.Requests = append(result.Requests,engine.Request{
			Url:        url,
			Parser:     engine.CreateFuncParser(ParseCourse,"ParseCourse"),
		})
		limit--
		if limit ==0 {
			break
		}
	}
	return result
}
