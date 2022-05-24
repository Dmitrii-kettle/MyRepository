package poj

import (
	"MyServer/mlog"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
)

func Parser() *user.User {
	c := &user.User{}
	content, err := FileGetContents("get.json")
	if err != nil {
		mlog.Error("open file error: %v", err)
	}
	err = json.Unmarshal([]byte(content), c)
	if err != nil {
		mlog.Error("Error: %v", err)
	}
	return c
}

func FileGetContents(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}
