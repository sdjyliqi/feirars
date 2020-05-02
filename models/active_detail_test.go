package models

import (
	"github.com/sdjyliqi/feirars/testutil"
	"testing"
)



func Test_ActiveDetailGetItemsByPage(t *testing.T){
      testActiveDetail:= ActiveDetail{}
      items,count,err:= testActiveDetail.GetItemsByPage(testutil.TestMysql, 10,1)
      t.Log(items,count,err)
      for _,v := range items {
      	t.Log(v.EventDay,v.Pv)
	  }
}
