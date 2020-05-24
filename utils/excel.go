package utils

import (
	"bytes"
	"encoding/csv"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func CreateExcelFile(excelFile *excelize.File, cols []map[string]string, items [][]string) error {
	buf := new(bytes.Buffer)
	r2 := csv.NewWriter(buf)
	var titleCols []string
	for _, v := range cols {
		titleCols = append(titleCols, v["name"])
	}
	r2.Write(titleCols)
	r2.Flush()
	//for _,v := range items {
	//	r2.(v)
	//	r2.Flush()
	//}

	r2.WriteAll(items)
	r2.Flush()

	//excelFile.w
	//sheet := excelFile.NewSheet("加班餐费")
	//sheet("加班餐费明细")
	//sheet.WriteRow("月份", "2018年06月", "姓名", "wwww", "部门", "研发部")
	//sheet.WriteRow("序号", "日期", "事由", "中餐/晚餐", "金额", "备注")
	//sheet.WriteRow("1", "2018-06-01", "加班", "晚餐", "20", "21:05")
	//sheet.WriteRow("2", "2018-06-02", "加班", "晚餐", "20", "21:05")
	//sheet.WriteRow("3", "2018-06-01", "加班", "晚餐", "20", "21:05")
	//sheet.WriteRow("4", "2018-06-02", "加班", "晚餐", "20", "21:05")
	//sheet.WriteRow("", "", "", "", "金额合计", "80.00", "")
	//sheet.Apply(excel.NewExcelStyle(10, -1, false));

	return excelFile.Write(buf)
}
