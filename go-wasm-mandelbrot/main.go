package main

import (
	"bytes"
	"math/cmplx"
)

func mandelbrot(c complex128) int {
	const iterations = 1000
	const escapeRadius = 2.0
	var z complex128
	for n := 0; n < iterations; n++ {
		z = z*z + c
		if cmplx.Abs(z) > escapeRadius {
			return n
		}
	}
	return iterations
}

func generateMandelbrot(width, height int) string {
	const xmin, ymin, xmax, ymax = -2.0, -1.5, +1.0, +1.5
	var art bytes.Buffer

	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			c := complex(x, y)
			m := mandelbrot(c)
			var ch byte
			switch {
			case m < 10:
				ch = ' '
			case m < 20:
				ch = '.'
			case m < 40:
				ch = '*'
			case m < 100:
				ch = 'o'
			case m < 200:
				ch = 'O'
			default:
				ch = '@'
			}
			art.WriteByte(ch)
		}
		art.WriteByte('\n')
	}
	return art.String()
}
