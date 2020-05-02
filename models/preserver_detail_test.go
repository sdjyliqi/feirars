package models

import (
	"github.com/sdjyliqi/feirars/testutil"
	"testing"
)

func Test_PreserveDetailGetItemsByPage(t *testing.T) {
	testPreserveDetail := PreserveDetail{}
	items, count, err := testPreserveDetail.GetItemsByPage(testutil.TestMysql, 1, 10)
	t.Log(items, count, err)
	for _, v := range items {
		t.Log(v.EventTime, v.Uv)
	}
}
