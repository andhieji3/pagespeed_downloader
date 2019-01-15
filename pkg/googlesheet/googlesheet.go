package googleSheetLibrary

import (
	"fmt"
	"pagespeed_reader/pkg/googleauth"
)

func TestUseService() {
	googleAuthLibrary.SpreadSheetsService = nil
	fmt.Println("Write Data to 201901 - Pagespeed Monitoring Report - Sheet (11)")
}
