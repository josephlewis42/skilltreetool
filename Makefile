VERSION?=v0.0.0

.PHONY: all
all: build test 

.PHONY: build
build: build-dir skilltreetool frontend

.PHONY: skilltreetool
skilltreetool: build-dir
	go build -ldflags "-X main.version=$(VERSION)" -o build/skilltreetool main.go

.PHONY: test
test:
	go test -cover ./...

.PHONY: wasm
wasm:
	GOOS=js GOARCH=wasm go build -o build/skilltreetool.wasm pkg/browser/wasm/js.go
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" build/wasm_exec.js

.PHONY: frontend
frontend: build-dir wasm
	cd frontend; yarn install
	cd frontend; yarn build
	mv frontend/dist build/frontend

.PHONY: build-dir
build-dir:
	mkdir -p build

.PHONY: clean
clean:
	rm -r build