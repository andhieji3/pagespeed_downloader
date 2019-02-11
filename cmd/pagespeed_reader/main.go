package main

import (
	"pagespeed_reader/pkg/configs"
	"pagespeed_reader/pkg/googleauth"
	"pagespeed_reader/pkg/googledrive"
	"pagespeed_reader/pkg/googlesheet"
	"pagespeed_reader/pkg/helper"
)

func main() {
	result := make(map[string]map[int]string)
	configuration := configs.GetConfiguration()
	var webUrl []string
	var wapUrl []string

	result["desktop"] = map[int]string{}
	result["mobile"] = map[int]string{}

	webUrl = configuration.WEB_URL
	wapUrl = configuration.WAP_URL

	result["desktop"] = helper.GetPagespeedData(webUrl, configuration.Pagespeed_Url, configuration.Pagespeed_API_Key, "desktop")
	result["mobile"] = helper.GetPagespeedData(wapUrl, configuration.Pagespeed_Url, configuration.Pagespeed_API_Key, "mobile")

	googleAuthLibrary.InitializeSpreadSheets()
	googleAuthLibrary.InitializeDrive()

	fileId := googleDriveLibrary.GetFileId()
	googleSheetLibrary.WriteReport(fileId, result)
}
