package global_functions

import (
	"time"
)

func GetCurrentColumn() string {
	column := "C"
	hour := time.Now().Hour()

	if hour >= 0 && hour < 6 {
		column = "C"
	} else if hour >= 6 && hour < 12 {
		column = "D"
	} else if hour >= 12 && hour < 18 {
		column = "E"
	} else if hour >= 18 && hour < 23 {
		column = "F"
	}

	return column
}
