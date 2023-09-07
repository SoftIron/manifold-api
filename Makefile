.PHONY: build
build:
	cd example && go build .

.PHONY: clean
clean:
	rm -f example/example

.PHONY: generate
generate:
	buf generate
	go generate ./...
	go generate -tags swag .

.PHONY: nuke
nuke:
	rm -rf gen

.PHONY: format
format:
	buf format -w


.PHONY: lint
lint:
	cd proto && buf lint

.PHONY: test
test:
	go test -v ./...


