package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Configuration struct {
	Pagespeed_Url              string   `json:"Pagespeed_Url"`
	Pagespeed_API_Key          string   `json:"Pagespeed_API_Key"`
	Googlesheets_Client_ID     string   `json:"Googlesheets_Client_ID"`
	Googlesheets_Client_Secret string   `json:"Googlesheets_Client_Secret"`
	WEB_URL                    []string `json:"WEB_URL"`
	WAP_URL                    []string `json:"WAP_URL"`
}

func GetConfiguration() Configuration {
	jsonFile, err := os.Open(os.Getenv("GOPATH") + "/src/pagespeed_reader/configs/development.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var configs Configuration
	json.Unmarshal(byteValue, &configs)

	return configs
}
