.PHONY: build
build:
	go build -o vnamer     cmd/vnamer/main.go
	go build -o vnamerdiff cmd/vnamerdiff/main.go

.PHONY: install
install: build
	cp ./vnamer     $$GOPATH/bin/
	cp ./vnamerdiff $$GOPATH/bin/

.PHONY: clean
clean:
	rm -rf ./vnamer            ./vnamerdiff
	rm -rf $$GOPATH/bin/vnamer $$GOPATH/bin/vnamerdiff
