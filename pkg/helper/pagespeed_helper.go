package helper

import (
	"fmt"
	"pagespeed_reader/pkg/pagespeed_downloader"
)

func GetPagespeedData(urls []string, pagespeedUrl string, apiKey string, platform string) map[string]string {
	var value string
	result := make(map[string]string)

	for _, url := range urls {
		downloaded := pagespeed_downloader.Download(pagespeedUrl, apiKey, url, platform)
		value = pagespeed_downloader.GetScore(downloaded)
		result[url] = value
		fmt.Println(url + " : " + value)
	}

	return result
}
