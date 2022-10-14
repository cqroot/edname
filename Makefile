.PHONY: build
build:
	go build -o vinamer cmd/main.go

.PHONY: install
install: build
	cp ./vinamer $$GOPATH/bin/
