package models

import (
	"github.com/sdjyliqi/feirars/testutil"
	"testing"
)



func Test_FeirarDetailGetItemsByPage(t *testing.T){
      testFeirarDetail:= FeirarDetail{}
      items,count,err:= testFeirarDetail.GetItemsByPage(testutil.TestMysql, 1,10)
      t.Log(items,count,err)
      for _,v := range items {
      	t.Log(v.EventDay,v.Pv)
	  }
}
