//go:build !js
// +build !js

package main

import (
	"fmt"
	"time"
)

func main() {
	const width, height = 160, 48
	timeStart := time.Now()
	result := generateMandelbrot(width, height)
	timeEnd := time.Now()
	fmt.Print(string(result))
	fmt.Printf("Time taken: %v\n", timeEnd.Sub(timeStart))
}
