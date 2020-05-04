package models

import (
	"github.com/sdjyliqi/feirars/testutil"
	"testing"
	"time"
)

func Test_PreserveDetailGetItemsByPage(t *testing.T) {
	testPreserveDetail := PreserveDetail{}
	items, count, err := testPreserveDetail.GetItemsByPage(testutil.TestMysql, "all,BZ", 1, 10, 0, time.Now().Unix())
	t.Log(items, count, err)
	for _, v := range items {
		t.Log(v.EventTime, v.Uv)
	}
}
