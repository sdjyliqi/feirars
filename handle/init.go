package handle

import (
	"sync"
	"github.com/sdjyliqi/feirars/conf"
	"github.com/sdjyliqi/feirars/control"
)

var once         sync.Once
var UserCenter    control.UserCenter

func init(){
	once.Do(func() {
		UserCenter = control.CreateUserCenter(&conf.DefaultConfig)
	})
}

