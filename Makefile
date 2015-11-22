.PHONY: build doc local

default: build

build:
	GOOS=linux GOARCH=amd64 go build -o checkmail

doc:
	godoc -http=:8765

local:
	go build -o checkmail
