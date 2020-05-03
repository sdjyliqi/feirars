package handle

import (
	"github.com/sdjyliqi/feirars/conf"
	"github.com/sdjyliqi/feirars/control"
	"sync"
)

var onceControl sync.Once
var PingbackCenter control.PingbackCenter
var UCenter control.UserCenter

func init() {
	onceControl.Do(func() {
		PingbackCenter = control.CreatePingbackCenter(&conf.DefaultConfig)
		UCenter = control.CreateUserCenter(&conf.DefaultConfig)
	})
}
