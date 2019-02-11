package pagespeed_downloader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"pagespeed_reader/pkg/global_functions"
)

func Download(pagespeedUrl string, apiKey string, checkUrl string, strategy string) string {
	fullUrl := pagespeedUrl + "?url=" + checkUrl + "&strategy=" + strategy + "&key=" + apiKey
	response, err := http.Get(fullUrl)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)

	return string(contents)
}

func GetScore(downloaded string) string {
	contents := make(map[string]interface{})
	result := "0"

	json.Unmarshal([]byte(downloaded), &contents)

	getValue, err := contents["error"].(map[string]interface{})

	if err != false {
		displayError(getValue)
	}

	result = generateResult(contents)

	return result
}

func generateResult(contents map[string]interface{}) string {
	result := "0"
	converted := contents["lighthouseResult"].(map[string]interface{})
	converted = converted["categories"].(map[string]interface{})
	converted = converted["performance"].(map[string]interface{})
	result = global_functions.FloatToString(converted["score"].(float64) * 100)
	return result
}

func displayError(errorMessage map[string]interface{}) {
	convertPace1 := errorMessage["errors"].([]interface{})
	convertPace2 := convertPace1[0].(map[string]interface{})

	domain := convertPace2["domain"].(string)
	reason := convertPace2["reason"].(string)
	message := convertPace2["message"].(string)

	fmt.Println("Error happen with domain : " + domain)
	fmt.Println("Reason : " + reason)
	fmt.Println("Message : " + message)
	os.Exit(0)
}
