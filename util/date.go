package util

import (
	"strings"
	"time"
)

func FormatIndonesianDate(date string) (string, error) {
	monthMap := map[string]string{
		"January":   "Januari",
		"February":  "Februari",
		"March":     "Maret",
		"April":     "April",
		"May":       "Mei",
		"June":      "Juni",
		"July":      "Juli",
		"August":    "Agustus",
		"September": "September",
		"October":   "Oktober",
		"November":  "November",
		"December":  "December",
	}

	dateObj, err := time.Parse("2006-01-02", date)

	if err != nil {
		return "", err
	}

	formattedDate := dateObj.Format("2 January 2006")

	for englishMonth, indonesianMonth := range monthMap {
		if strings.Contains(formattedDate, englishMonth) {
			formattedDate = strings.Replace(formattedDate, englishMonth, indonesianMonth, 1)
			break
		}
	}

	return formattedDate, nil
}
