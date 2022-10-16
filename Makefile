PROJ_NAME=vina

.PHONY: build
build:
	go build -o $(PROJ_NAME)      cmd/$(PROJ_NAME)/main.go
	go build -o $(PROJ_NAME)diff  cmd/$(PROJ_NAME)diff/main.go

.PHONY: clean
clean:
	rm -f \
		./$(PROJ_NAME) \
		./$(PROJ_NAME)diff \

.PHONY: install
install: build
	cp ./$(PROJ_NAME)      $${GOPATH}/bin/
	cp ./$(PROJ_NAME)diff  $${GOPATH}/bin/

.PHONY: uninstall
uninstall:
	rm -f \
		$${GOPATH}/bin/$(PROJ_NAME) \
		$${GOPATH}/bin/$(PROJ_NAME)diff
