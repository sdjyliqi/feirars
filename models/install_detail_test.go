package models

import (
	"github.com/sdjyliqi/feirars/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_InstallDetailGetItemsByPage(t *testing.T) {
	testInstallDetail := InstallDetail{}
	items, count, err := testInstallDetail.GetItemsByPage(testutil.TestMysql, "all,BZ", 1, 10, 0, time.Now().Unix())
	t.Log(items, count, err)
	for _, v := range items {
		t.Log(v.EventDay, v.Pv)
	}
}

func Test_InstallDetailGetAllChannels(t *testing.T) {
	testActiveDetail := InstallDetail{}
	items, err := testActiveDetail.GetAllChannels(testutil.TestMysql)
	assert.Nil(t, err)
	t.Log(items, err)
}
