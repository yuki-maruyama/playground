package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
	"unsafe"
)

func main() {
	verbose := flag.Bool("v", false, "verbose")
	flag.Parse()

	aligned := 0
	loops, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		loops = 100
	}

	pageSize := uintptr(os.Getpagesize())
	fmt.Printf("pagesize : %v\n", pageSize)

	start := time.Now()
	for i := 0; i < loops; i++ {
		b := makeByteSlice(pageSize, 1)
		ptr := uintptr(unsafe.Pointer(unsafe.SliceData(b)))
		if *verbose {
			fmt.Printf("%v, %d\n", ptr, ptr%pageSize)
		}
		if ptr%pageSize == 0 {
			aligned++
		}
	}
	fmt.Printf("---get slice using make---\ntime: %v\naligned byte: %d/%d\n", time.Since(start), aligned, loops)

	aligned = 0
	start = time.Now()
	for i := 0; i < loops; i++ {
		b := alignedByteSlice(pageSize, 1)
		ptr := uintptr(unsafe.Pointer(unsafe.SliceData(b)))
		if *verbose {
			fmt.Printf("%v, %d\n", ptr, ptr%pageSize)
		}
		if ptr%pageSize == 0 {
			aligned++
		}
	}
	fmt.Printf("---get aligned slice---\ntime: %v\naligned byte: %d/%d\n", time.Since(start), aligned, loops)
}

func makeByteSlice(aliginSize uintptr, blockSize uintptr) []byte {
	return make([]byte, aliginSize*blockSize)
}

func alignedByteSlice(aliginSize uintptr, blockSize uintptr) []byte {
	b := make([]byte, aliginSize*(blockSize+2))
	a := uintptr(unsafe.Pointer(&b[0])) & (aliginSize - 1)
	var offset uintptr = 0
	if a != 0 {
		offset = aliginSize - a
	}
	return b[offset : offset+(aliginSize*blockSize)]
}
