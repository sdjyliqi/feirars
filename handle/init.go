package handle

import (
	"github.com/sdjyliqi/feirars/conf"
	"github.com/sdjyliqi/feirars/control"
	"sync"
)

var onceControl sync.Once
var PingbackCenter control.PingbackCenter

func init() {
	onceControl.Do(func() {
		PingbackCenter = control.CreatePingbackCenter(&conf.DefaultConfig)
	})
}
