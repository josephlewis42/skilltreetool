VERSION?=v0.0.0

.PHONY: all
all: build wasm test

.PHONY: build
build: build-dir skilltreetool

.PHONY: skilltreetool
skilltreetool: build-dir
	go build -ldflags "-X main.version=$(VERSION)" -o build/skilltreetool main.go

.PHONY: test
test:
	go test -cover ./...

.PHONY: wasm
wasm:
	GOOS=js GOARCH=wasm go build -o build/skilltreetool.wasm pkg/browser/wasm/js.go

.PHONY: build-dir
build-dir:
	mkdir -p build

.PHONY: clean
clean:
	rm -r build