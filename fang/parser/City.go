package parser

import (
	"crowler/engine"
	"fmt"
	"regexp"
)
//<a target="_blank" data-yd="" onclick="yidiCityCookie(&quot;泉山&quot;);" href="//xinghushangyuan0516.fang.com/">
//										星湖尚苑
//					</a>
//const CityRe = "<a  target=\"_blank\" href=\"(//[a-z+].fang.com/)\">"
const CityRe = "<a target=\"_blank\" data-yd=\"\" onclick=yidiCityCookie(\".*\");  href=\"(//[a-z0-9+].fang.com/)\">"
//解析页面，返回next Url和页面信息
func ParseCity(contents []byte) engine.ParseResult {
	var houstList2Re = regexp.MustCompile(CityRe)
	matchs := houstList2Re.FindAllSubmatch(contents,-1)

	//每一个url生成一个新的request
	result := engine.ParseResult{}
	for _,m := range matchs {
		result.Items = append(result.Items,string(m[2]))
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: engine.NilParser,
			})
		fmt.Printf(",Url:%s\n",m[1])
	}
	return result
}
