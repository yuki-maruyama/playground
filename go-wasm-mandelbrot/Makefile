GO := go
TINYGO := tinygo
GOARCH := wasm
GOOS := js
OUTPUT := main.wasm

all: build tinygo-build build-wasi tinygo-build-wasi

# Build target
build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build -o $(OUTPUT) -tags=js --ldflags="-s -w" .

build-wasi:
	GOOS=wasip1 GOARCH=wasm $(GO) build -o wasi.$(OUTPUT) --ldflags="-s -w" .

# TinyGo build target
tinygo-build:
	$(TINYGO) build -o tinygo.$(OUTPUT) -target wasm -no-debug .

tinygo-build-wasi:
	$(TINYGO) build -o tinygo.wasi.$(OUTPUT) -target wasi -no-debug .

# Clean target
clean:
	rm -f *.wasm