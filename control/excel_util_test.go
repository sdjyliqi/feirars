package control

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"testing"
)

func Test_CreateExcelFile(t *testing.T) {
	f, err := os.Create("test.xls")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w := csv.NewWriter(f)
	w.Write([]string{"编号", "姓名", "年龄"})
	w.Write([]string{"1", "张三", "23"})
	w.Write([]string{"2", "李四", "24"})
	w.Write([]string{"3", "王五", "25"})
	w.Write([]string{"4", "赵六", "26"})
	w.Flush()
}

func Test_CreateExcelFile2(t *testing.T) {
	fileName := "liqitest.csv"
	buf := new(bytes.Buffer)
	r2 := csv.NewWriter(buf)
	line := []string{"编号", "姓名", "性别"}
	r2.Write(line)
	for i := 0; i < 10; i++ {
		line := []string{fmt.Sprintf("%d", i), "liqi", "M"}
		r2.Write(line)

	}
	r2.Flush()
	fmt.Print(buf)

	fout, err := os.Create(fileName)
	fout.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	if err != nil {
		t.Log(err)
	}
	fout.Write(buf.Bytes())
	defer fout.Close()
}
