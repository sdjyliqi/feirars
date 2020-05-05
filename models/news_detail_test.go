package models

import (
	"github.com/sdjyliqi/feirars/testutil"
	"github.com/stretchr/testify/assert"
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

func Test_NewsDetailGetAllChannels(t *testing.T) {
	testActiveDetail := NewsDetail{}
	items, err := testActiveDetail.GetAllChannels(testutil.TestMysql)
	assert.Nil(t, err)
	t.Log(items, err)
}
