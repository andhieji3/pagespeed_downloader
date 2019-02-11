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
	Googlesheets_Template_ID   string   `json:"Googlesheets_Template_ID"`
	GoogleDrive_HTTP_URL       string   `json:"GoogleDrive_HTTP_URL"`
	GoogleDrive_Main_Directory string   `json:"GoogleDrive_Main_Directory"`
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
