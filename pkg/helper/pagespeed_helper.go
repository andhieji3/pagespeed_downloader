package helper

import (
	"fmt"
	"pagespeed_reader/pkg/pagespeed_downloader"
)

func GetPagespeedData(urls []string, pagespeedUrl string, apiKey string, platform string) map[int]string {
	var value string
	result := make(map[int]string)

	for index, url := range urls {
		downloaded := pagespeed_downloader.Download(pagespeedUrl, apiKey, url, platform)
		value = pagespeed_downloader.GetScore(downloaded)
		result[index] = value
		fmt.Println(url + " : " + value)
	}

	return result
}
