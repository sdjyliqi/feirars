package models

import (
	"github.com/sdjyliqi/feirars/testutil"
	"testing"
	"time"
)

func Test_ActiveDetailGetItemsByPage(t *testing.T) {
	testActiveDetail := ActiveDetail{}
	items, count, err := testActiveDetail.GetItemsByPage(testutil.TestMysql, "all,BZ", 1, 20, 0, time.Now().Unix())
	t.Log(items, count, err)
	for _, v := range items {
		t.Log(v.EventDay, v.Pv)
	}
}
