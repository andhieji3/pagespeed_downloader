package googleSheetLibrary

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/sheets/v4"
	"log"
	"pagespeed_reader/pkg/global_functions"
	"pagespeed_reader/pkg/googleauth"
	"strings"
	"time"
)

func ConfigureNewFile(fileId string) {
	fmt.Println("Configuring data for file " + fileId)
	ctx := context.Background()
	month := int(time.Now().Month())
	vRange := &sheets.ValueRange{}
	cellRange := "Google PageSpeed Summary!A4:A34"

	resp, err := googleAuthLibrary.SpreadSheetsService.Spreadsheets.Values.Get(fileId, cellRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		for _, row := range resp.Values {
			splitRowString := strings.Split(row[0].(string), "-")
			newValue := splitRowString[0] + "-" + global_functions.IntToString(month) + "-" + splitRowString[2]
			rowValue := []interface{}{newValue}
			vRange.Values = append(vRange.Values, rowValue)
		}
	}

	googleAuthLibrary.SpreadSheetsService.Spreadsheets.Values.Update(fileId, cellRange, vRange).ValueInputOption("RAW").Context(ctx).Do()
	fmt.Println("Success configuring data for file " + fileId)
}

func WriteReport(fileId string, report map[string]map[int]string) {
	fmt.Println("Write report to file " + fileId)
	ctx := context.Background()
	vRange := &sheets.ValueRange{}
	day := int(time.Now().Day())
	column := global_functions.GetCurrentColumn()
	cellRange := global_functions.IntToString(day) + "!" + column + "11:" + column + "14"

	for _, data := range report {
		for _, value := range data {
			convertedValue, err := global_functions.StringToFloat(value)

			if err != nil {
				log.Fatalf("Error when converting data", err)
			}
			rowValue := []interface{}{convertedValue}
			vRange.Values = append(vRange.Values, rowValue)
		}
	}

	googleAuthLibrary.SpreadSheetsService.Spreadsheets.Values.Update(fileId, cellRange, vRange).ValueInputOption("RAW").Context(ctx).Do()
	fmt.Println("Success report to file " + fileId)
}
