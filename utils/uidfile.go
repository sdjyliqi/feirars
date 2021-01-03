package utils

import (
	"encoding/json"
	"github.com/golang/glog"
	"io/ioutil"
)

func GetUIDsFromFile(path string) ([]string, error) {
	var uids []string
	b, err := ioutil.ReadFile(path) // just pass the file name
	if err != nil {
		glog.Errorf("Read the file %s failed,err:%+v", path, err)
		return nil, err
	}
	err = json.Unmarshal(b, &uids)
	if err != nil {
		glog.Errorf("Unmarshal the content  failed,err:%+v", err)
		return nil, err
	}
	return uids, err
}
