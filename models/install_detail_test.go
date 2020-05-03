package models

import (
	"github.com/sdjyliqi/feirars/testutil"
	"testing"
	"time"
)

func Test_InstallDetailGetItemsByPage(t *testing.T) {
	testInstallDetail := InstallDetail{}
	items, count, err := testInstallDetail.GetItemsByPage(testutil.TestMysql, "", 1, 10, 0, time.Now().Unix())
	t.Log(items, count, err)
	for _, v := range items {
		t.Log(v.EventDay, v.Pv)
	}
}
