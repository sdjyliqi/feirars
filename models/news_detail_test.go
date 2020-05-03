package models

import (
	"github.com/sdjyliqi/feirars/testutil"
	"testing"
	"time"
)

func Test_NewsDetailGetItemsByPage(t *testing.T) {
	testNewsDetail := NewsDetail{}
	items, count, err := testNewsDetail.GetItemsByPage(testutil.TestMysql, "", 1, 10, 0, time.Now().Unix())
	t.Log(items, count, err)
	for _, v := range items {
		t.Log(v.EventDay, v.Pv)
	}
}
