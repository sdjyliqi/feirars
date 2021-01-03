package utils

import (
	"fmt"
	"github.com/golang/glog"
	"testing"
)

func Test_GetUIDsFromFile(t *testing.T) {
	UIDs, err := GetUIDsFromFile("D:\\gowork\\src\\feirars\\all_2021-01-02.txt")
	glog.Info(len(UIDs), err)
	fmt.Println(len(UIDs), err)
}
