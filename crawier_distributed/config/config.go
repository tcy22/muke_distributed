package config

const (
	//Parser names
	ParseCourse	  = "ParseCourse"
	ParseCourseList = "ParseCourseList"
	ParserProfile  = "ParserProfile"
	NilParse      = "NilParse"


	//Service ports
	ItemSavePort = ":2222"
	WorkPort0 = ":9002"

	//ElasticSearch
	ElasticIndex = "muke_profile22"
	ElasticUrl = "http://192.168.163.134:9200"

	//RPC Endpoints
	ItemSaverRpc = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"

)

