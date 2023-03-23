.PHONY: install
install:
	go install .

.PHONY: gen-testdata
gen-testdata:
	bash "$(CURDIR)/scripts/gen-testdata.sh"

.PHONY: test
test: gen-testdata
	go test -v ./...

.PHONY: check
check:
	@echo '******************************'
	golangci-lint run
	@echo
	@echo '******************************'
	gofumpt -l .
