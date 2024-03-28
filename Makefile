.PHONY: build
build:
	cd example && go build .

.PHONY: clean
clean:
	rm -f example/example

PBFILES = gen/go/proto/snapper/v1/snapper.pb.go

.PHONY: generate
generate:
	@echo "Generating protobuf files"
	@buf generate && for f in $(PBFILES); do sed -i "" 's:,omitempty::g' $$f; done
	@echo "Generating go code"
	@go generate ./...

.PHONY: nuke
nuke:
	rm -rf gen

.PHONY: format
format:
	buf format -w


.PHONY: lint
lint:
	cd proto && buf lint
	golangci-lint run

.PHONY: test
test:
	go test -v ./...


