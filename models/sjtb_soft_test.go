package models

import (
	"encoding/json"
	"github.com/sdjyliqi/feirars/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_SjtbSoftGetChartItems(t *testing.T) {
	testMod := SjtbSoft{}
	items, err := testMod.GetChartItems(testutil.TestMysql, "all", time.Now().Unix()-3600*24*7, time.Now().Unix())
	assert.Nil(t, err)
	t.Log(items, err)
	content, _ := json.Marshal(items)
	t.Log(string(content))
}
