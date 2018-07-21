package parser

import (
	"regexp"
	"muke_distributed/crawier/engine"
	"strings"
)
const urlC  = `https://www.imooc.com`
const courseRe  = `<a target="_blank" href="(/learn/[0-9]+)"[^>]*>`

//课程分类解析器
func ParseCourse(contents []byte,_ string) engine.ParseResult{
	re := regexp.MustCompile(courseRe)
	matches := re.FindAllSubmatch(contents,-1)     //[][][]byte

	result := engine.ParseResult{}
	for _,m := range matches {
		url := urlC + string(m[1])
		urlT := strings.Replace(url,"learn","coursescore",1)
		result.Requests = append(result.Requests,engine.Request{
			Url:        urlT,
			Parser:     engine.CreateFuncParser(ParserProfile,"ParserProfile"),
		})
	}
	return result
}

