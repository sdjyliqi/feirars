package models

import (
	"github.com/sdjyliqi/feirars/testutil"
	"testing"
)

func Test_NewsDetailGetItemsByPage(t *testing.T) {
	testNewsDetail := NewsDetail{}
	items, count, err := testNewsDetail.GetItemsByPage(testutil.TestMysql, 1, 10)
	t.Log(items, count, err)
	for _, v := range items {
		t.Log(v.EventDay, v.Pv)
	}
}
