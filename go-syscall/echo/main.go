package echo

import "syscall"

func Echo(s string) {
	_, err := syscall.Write(1, []byte(s))
	if err != nil {
		syscall.Write(2, []byte("error: write stdout : "+err.Error()+"\n"))
		syscall.Exit(1)
	}
}
