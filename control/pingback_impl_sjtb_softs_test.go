package control

import (
	"encoding/json"
	"github.com/sdjyliqi/feirars/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_GetSjtbSoftItems(t *testing.T) {
	util := CreatePingbackCenter(&testutil.TestCfg)
	ctime := time.Now()
	//chn string, pageID, pageCount int, tsStart, tsEnd int64
	items, count, err := util.GetSjtbSoftItems("all", 1, 100, ctime.Unix()-3600*24*10, time.Now().Unix())
	assert.Nil(t, err)
	for _, v := range items {
		content, _ := json.Marshal(v)
		t.Log(string(content))
	}
	t.Log(len(items), count)

	cols := util.GetSjtbSoftCols()
	t.Log(cols)

	chns, err := util.GetSjtbSoftChannel("admin")
	assert.Nil(t, err)
	t.Log(chns)
}
