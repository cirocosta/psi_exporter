GO_FILES := $(shell find . -name "*.go")


test:
	go test -v ./...


build: psi_collector.out


psi_collector.out: $(GO_FILES)
	go build -o $@ -v -i .
