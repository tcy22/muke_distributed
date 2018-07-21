package parser

import (
	"muke_distributed/crawier/engine"
	"regexp"
	"muke_distributed/crawier/model"
)

var num = regexp.MustCompile(`<h2 class="l">([^<]+)</h2>`)
var score = regexp.MustCompile(`<span class="meta">综合评分</span><span class="meta-value">([^<]+)</span>`)
var comment  = regexp.MustCompile(`<p class="content">([^<]+)</p>`)
var idUrlRe = regexp.MustCompile(`https://www.imooc.com/coursescore/([\d]+)`)

func ParserProfile(contents []byte,url string) engine.ParseResult{
	profile := model.Profile{}

	profile.Num = "人数:"+extractString(contents,num)
	profile.Score = "分数:"+extractString(contents,score)
	profile.Comments = extractString2(contents,comment)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Type:"muke",
				Id:extractString([]byte(url),idUrlRe),
				Payload:profile,
			},
		},
	}
	return result
}

func extractString(contents []byte,re *regexp.Regexp) string{
	match := re.FindSubmatch(contents)
	if len(match) >=2 {
		return string(match[1])
	}else{
		return ""
	}
}

func extractString2(contents []byte,re *regexp.Regexp) []string{
	var Items []string
	matches := re.FindAllSubmatch(contents,-1)
	for _,m := range matches {
		Items = append(Items,"评论:"+string(m[1]))
	}
	return Items
}