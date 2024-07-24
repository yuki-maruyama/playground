package main

import (
	"math/big"
	"syscall/js"
)

func calcPi(this js.Value, inputs []js.Value) interface{} {
	itr := int64(inputs[0].Int())
	prec := inputs[1].Int()

	two := big.NewRat(2, 1)
	pi := big.NewRat(1.0, 1.0)
	a := big.NewRat(1.0, 1.0)
	b := big.NewRat(1.0, 2.0)
	x := new(big.Rat)
	for i := int64(2); i < itr; i += 2 {
		y := big.NewRat(i, i+1)
		a.Mul(a, y)
		x.Mul(a, b)
		pi.Add(pi, x)
		b.Quo(b, two)
	}
	return pi.Mul(pi, two).FloatString(prec)
}

func main() {
	done := make(chan struct{}, 0)
	js.Global().Set("calcPi", js.FuncOf(calcPi))
	<-done
}
