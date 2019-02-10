GO_FILES := $(shell find . -name "*.go")


test:
	go test -v ./...


build: psi_collector.out


psi_collector.linux.out: $(GO_FILES)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -ldflags '-w' -o $@ -v .

psi_collector.out: $(GO_FILES)
	go build -o $@ -v -i .
