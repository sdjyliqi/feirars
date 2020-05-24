package utils

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"testing"
)

//func CreateExcelFile(excelFile *excelize.File, cols []map[string]string,items [][]string) error{
//	buf := new(bytes.Buffer)
//	r2 := csv.NewWriter(buf)
//	var titleCols []string
//	for _,v := range cols {
//		titleCols = append(titleCols,v["name"])
//	}
//	r2.Write(titleCols)
//	for _,v := range items {
//		r2.Write(v)
//	}
//	return excelFile.Write(buf)
//}

func Test_CreateExcelFile(t *testing.T) {
	xlsx := excelize.NewFile()
	var cols []map[string]string
	colEventDay := map[string]string{
		"name":  "日期",
		"key":   "event_day",
		"click": "0",
	}
	cols = append(cols, colEventDay)

	colPv := map[string]string{
		"name":  "pv",
		"key":   "pv",
		"click": "0",
	}
	cols = append(cols, colPv)
	colUv := map[string]string{
		"name":  "uv",
		"key":   "uv",
		"click": "0",
	}
	cols = append(cols, colUv)
	items := [][]string{
		[]string{"2020-05-10", "30", "12"},
		[]string{"2020-05-11", "40", "12"},
		[]string{"2020-05-12", "50", "12"},
	}
	CreateExcelFile(xlsx, cols, items)

	xlsx.SaveAs("aaaaaaaaaaa.csv")
}
