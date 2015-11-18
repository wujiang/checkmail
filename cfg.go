package main

import (
	"encoding/json"
	"io/ioutil"
)

type config struct {
	Mailboxes map[string]string `json:"mailboxes"`
}

func (cfg *config) parse(fn string) error {
	bt, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bt, cfg)
	return err
}
