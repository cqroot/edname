PROJ_NAME=edname
BUILD_DIR=$(CURDIR)/.build

.PHONY: build
build:
	go build -o "$(BUILD_DIR)/$(PROJ_NAME)" main.go

.PHONY: run
run: build
	"$(BUILD_DIR)/$(PROJ_NAME)"

.PHONY: clean
clean:
	rm -f "$(BUILD_DIR)"

.PHONY: install
install: build
	cp "$(BUILD_DIR)/$(PROJ_NAME)" $${GOPATH}/bin/

.PHONY: uninstall
uninstall:
	rm -f $${GOPATH}/bin/$(PROJ_NAME)

.PHONY: gen-testdata
gen-testdata:
	sh "$(CURDIR)/scripts/gen-testdata.sh"

.PHONY: test
test: gen-testdata
	go test -v ./...
