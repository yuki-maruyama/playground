//go:build !js

package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

func main() {
	flag.Parse()
	width, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		width = 160
	}
	height, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		height = 48
	}
	timeStart := time.Now()
	result := generateMandelbrot(width, height)
	timeEnd := time.Now()
	fmt.Print(string(result))
	fmt.Printf("Time taken: %v\n", timeEnd.Sub(timeStart))
}
