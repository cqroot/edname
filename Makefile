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
