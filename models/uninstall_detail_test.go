package models

import (
	"github.com/sdjyliqi/feirars/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_UninstallDetailGetItemsByPage(t *testing.T) {
	testInstallDetail := UninstallDetail{}
	items, count, err := testInstallDetail.GetItemsByPage(testutil.TestMysql, "all,BZ", 1, 10, 0, time.Now().Unix())
	assert.Nil(t, err)
	t.Log(items, count, err)
}

func Test_UninstallDetailGetAllChannels(t *testing.T) {
	testActiveDetail := PreserveDetail{}
	items, err := testActiveDetail.GetAllChannels(testutil.TestMysql)
	assert.Nil(t, err)
	t.Log(items, err)
}
