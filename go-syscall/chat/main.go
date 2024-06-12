package chat

import (
	"syscall"
)

func Chat(sv bool) {
	if sv {
		syscall.Pipe()
	}
}
