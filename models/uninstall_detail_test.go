package models

import (
	"github.com/sdjyliqi/feirars/testutil"
	"testing"
)

func Test_uninstallDetailGetItemsByPage(t *testing.T) {
	testInstallDetail := UninstallDetail{}
	items, count, err := testInstallDetail.GetItemsByPage(testutil.TestMysql, 1, 10)
	t.Log(items, count, err)
	for _, v := range items {
		t.Log(v.EventDay, v.Pv)
	}
}
