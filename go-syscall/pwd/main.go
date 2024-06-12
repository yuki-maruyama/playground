package pwd

import (
	"syscall"
)

func Pwd() {
	s, err := syscall.Getwd()
	if err != nil {
		syscall.Write(2, []byte("error: getcwd : "+err.Error()+"\n"))
		syscall.Exit(1)
	}
	_, err = syscall.Write(1, []byte(s+"\n"))
	if err != nil {
		syscall.Write(2, []byte("error: write stdout : "+err.Error()+"\n"))
		syscall.Exit(1)
	}
}
