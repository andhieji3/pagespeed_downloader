package googlesheets

import (
	"fmt"
	"pagespeed_reader/pkg/googleauth"
)

func TestUseService() {
	googleauth.SpreadSheetsService = nil
	fmt.Println("Write Data to 201901 - Pagespeed Monitoring Report - Sheet (11)")
}
