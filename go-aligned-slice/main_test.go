package main

import (
	"os"
	"testing"
)

func Benchmark_makeByteSlice(b *testing.B) {
	pageSize := uintptr(os.Getpagesize())
	blockSize := uintptr(5)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		makeByteSlice(pageSize, blockSize)
	}
}

func Benchmark_alignedByteSlice(b *testing.B) {
	pageSize := uintptr(os.Getpagesize())
	blockSize := uintptr(5)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		alignedByteSlice(pageSize, blockSize)
	}
}
