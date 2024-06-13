package cat

import (
	syscall "golang.org/x/sys/unix"
)

func Cat(dir string, readFromStdin bool) {
	var fid int = 0
	if !readFromStdin {
		fd, err := syscall.Open(dir, syscall.O_RDONLY, 0)
		if err != nil {
			syscall.Write(2, []byte("error: file cannot open : "+err.Error()+"\n"))
			syscall.Exit(1)
		}
		fid = fd
		defer syscall.Close(fd)
	}

	for {
		buff := make([]byte, 1024)
		n, err := syscall.Read(fid, buff)
		if err != nil {
			syscall.Write(2, []byte("error: file read : "+err.Error()+"\n"))
			syscall.Exit(1)
		}
		if n == 0 {
			break
		}
		_, err = syscall.Write(1, buff)
		if err != nil {
			syscall.Write(2, []byte("error: write stdout : "+err.Error()+"\n"))
			syscall.Exit(1)
		}
	}
}
