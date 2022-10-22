PROJ_NAME=vina

.PHONY: build
build:
	go build -o $(PROJ_NAME)      cmd/main.go

.PHONY: clean
clean:
	rm -f ./$(PROJ_NAME)

.PHONY: install
install: build
	cp ./$(PROJ_NAME) $${GOPATH}/bin/

.PHONY: uninstall
uninstall:
	rm -f $${GOPATH}/bin/$(PROJ_NAME)

.PHONY: gen-testdata
gen-testdata:
	sh "$(CURDIR)/scripts/gen-testdata.sh"

.PHONY: test
test:
	go test -v ./...
