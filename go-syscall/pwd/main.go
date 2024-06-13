package pwd

import (
	syscall "golang.org/x/sys/unix"
)

func Pwd() {
	buf := make([]byte, 1024)
	_, err := syscall.Getcwd(buf)
	if err != nil {
		syscall.Write(2, []byte("error: getcwd : "+err.Error()+"\n"))
		syscall.Exit(1)
	}
	buf = append(buf, byte(0x0A))
	_, err = syscall.Write(1, buf)
	if err != nil {
		syscall.Write(2, []byte("error: write stdout : "+err.Error()+"\n"))
		syscall.Exit(1)
	}
}
