//go:build js && wasm

package main

import (
	"syscall/js"
)

func generateMandelbrotWrapper(this js.Value, p []js.Value) interface{} {
	width := p[0].Int()
	height := p[1].Int()
	return generateMandelbrot(width, height)
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("generateMandelbrot", js.FuncOf(generateMandelbrotWrapper))
	<-c
}
