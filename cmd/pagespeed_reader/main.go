package main

import (
	"pagespeed_reader/pkg/configs"
	"pagespeed_reader/pkg/googleauth"
	"pagespeed_reader/pkg/googlesheets"
	"pagespeed_reader/pkg/helper"
)

func main() {
	result := make(map[string]map[string]string)
	configuration := configs.GetConfiguration()
	var webUrl []string
	var wapUrl []string

	result["desktop"] = map[string]string{}
	result["mobile"] = map[string]string{}

	webUrl = configuration.WEB_URL
	wapUrl = configuration.WAP_URL

	result["desktop"] = helper.GetPagespeedData(webUrl, configuration.Pagespeed_Url, configuration.Pagespeed_API_Key, "desktop")
	result["mobile"] = helper.GetPagespeedData(wapUrl, configuration.Pagespeed_Url, configuration.Pagespeed_API_Key, "mobile")

	googleauth.InitializeSpreadSheets()
	googlesheets.TestUseService()
}
