package engine

type ParserFunc func(contents []byte,url string) ParseResult

type Parser interface {
	Parser(contents []byte,url string) ParseResult  //解析获得有用信息
	Serialize() (name string,args interface{})      //处理函数
}    //调用这个函数，填充SerializedParser

type Request struct {
	Url        string
	Parser     Parser    //接口，接口里面有方法。
}

/*type SerializedParser struct {
	Name string
	Args interface{}
}
//{"ParseCityList",nil},{"ProfileParse",userName}*/

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Type         string
	Id           string
	Payload      interface{}
}

type NilParser struct {}

func (NilParser) Parser(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser",nil
}

type FuncParser struct {
	parser ParserFunc       //放一个函数
	name string            //函数的名字
}

func (f *FuncParser) Parser(contents []byte,url string) ParseResult {
	return f.parser(contents,url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil

}

func CreateFuncParser(p ParserFunc,name string) *FuncParser{
	return &FuncParser{
		parser :p,        //函数
		name: name,       //函数名字
	}
}


