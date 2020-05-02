package control

//相关封装函数的UT
import (
	"github.com/sdjyliqi/feirars/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

var TestUtil = CreatePingbackCenter(&testutil.TestCfg)

func Test_GetActiveDetailItems(t *testing.T) {
	item,count,err := TestUtil.GetActiveDetailItems(1,10)
	assert.Nil(t,err)
	t.Log(item,count)
}

func Test_GetFeirarDetailItems(t *testing.T) {
	items,count,err := TestUtil.GetFeirarDetailItems(1,10)
	assert.Nil(t,err)
	t.Log(items,count)
}

