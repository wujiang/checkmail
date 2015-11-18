package main

import (
	"os"
	"path/filepath"
)

type MailDir struct {
	name string
	path string
}

func (md MailDir) Unseen() ([]string, error) {
	fp, err := os.Open(filepath.Join(md.path, "new"))
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	names, err := fp.Readdirnames(0)
	if err != nil {
		return nil, err
	}
	return names, err
}
