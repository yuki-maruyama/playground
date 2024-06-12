package main

import (
	"flag"

	"github.com/yuki-maruyama/playground/go-syscall/cat"
	"github.com/yuki-maruyama/playground/go-syscall/echo"
)

func main() {
	flag.Parse()
	switch flag.Arg(0) {
	case "cat":
		if flag.Arg(1) != "" {
			cat.Cat(flag.Arg(1), false)
		} else {
			cat.Cat(flag.Arg(1), true)
		}

	case "echo":
		echo.Echo(flag.Arg(1))

	default:
		return
	}
}
