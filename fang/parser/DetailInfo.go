package parser

import (
	"crowler/engine"
	"crowler/model"
	"log"
	"regexp"
)


var houstnameRe = regexp.MustCompile(`<h1><strong>(.*)</strong></h1>`)

func ParseDetailInfo(contents []byte) engine.ParseResult {
	log.Println("commin in  ParseDetailInfo....................")
	houseProfile := model.HouseProfile{}

	houseName := extractString(contents,houstnameRe)
	houseProfile.HouseName = string(houseName)


	log.Println("comman in  ParseDetailInfo...........end.........")

	result := engine.ParseResult{
		Items: []interface{} {houseProfile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}else {
		return ""
	}
}
