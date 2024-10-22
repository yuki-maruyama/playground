package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"

	"golang.org/x/sys/unix"
)

var (
	clients   = make(map[int]net.Conn)
	clientsMu sync.Mutex
)

func main() {
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Chat server started on port 12345")

	epollFd, err := unix.EpollCreate1(0)
	if err != nil {
		fmt.Println("Error creating epoll instance:", err)
		os.Exit(1)
	}
	defer unix.Close(epollFd)

	file, err := listener.(*net.TCPListener).File()
	if err != nil {
		fmt.Println("Error getting file descriptor:", err)
		os.Exit(1)
	}
	listenerFd := int(file.Fd())
	event := unix.EpollEvent{Events: unix.EPOLLIN | unix.EPOLLET, Fd: int32(listenerFd)}
	if err := unix.EpollCtl(epollFd, unix.EPOLL_CTL_ADD, listenerFd, &event); err != nil {
		fmt.Println("Error adding listener to epoll instance:", err)
		os.Exit(1)
	}

	events := make([]unix.EpollEvent, 1000)
	for {
		n, err := unix.EpollWait(epollFd, events, -1)
		if err != nil {
			fmt.Println("Error during epoll wait:", err)
			continue
		}

		for i := 0; i < n; i++ {
			if int(events[i].Fd) == listenerFd {
				conn, err := listener.Accept()
				if err != nil {
					fmt.Println("Error accepting connection:", err)
					continue
				}
				connFile, err := conn.(*net.TCPConn).File()
				if err != nil {
					fmt.Println("Error getting file descriptor:", err)
					conn.Close()
					continue
				}
				connFd := int(connFile.Fd())
				event := unix.EpollEvent{Events: unix.EPOLLIN, Fd: int32(connFd)}
				if err := unix.EpollCtl(epollFd, unix.EPOLL_CTL_ADD, connFd, &event); err != nil {
					fmt.Println("Error adding connection to epoll instance:", err)
					conn.Close()
					continue
				}
				clientsMu.Lock()
				clients[connFd] = conn
				clientsMu.Unlock()
				fmt.Println("Client connected:", conn.RemoteAddr())
			} else {
				connFd := int(events[i].Fd)
				clientsMu.Lock()
				conn := clients[connFd]
				clientsMu.Unlock()
				handleClient(connFd, conn, epollFd)
			}
		}
	}
}

func handleClient(connFd int, conn net.Conn, epollFd int) {
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Client disconnected:", conn.RemoteAddr())
		unix.EpollCtl(epollFd, unix.EPOLL_CTL_DEL, connFd, nil)
		conn.Close()
		clientsMu.Lock()
		delete(clients, connFd)
		clientsMu.Unlock()
		return
	}

	// fmt.Printf("Received message from %s: %s", conn.RemoteAddr(), message)
	broadcastMessage(connFd, message)
}

func broadcastMessage(senderFd int, message string) {
	clientsMu.Lock()
	defer clientsMu.Unlock()
	for fd, conn := range clients {
		if fd != senderFd {
			writer := bufio.NewWriter(conn)
			writer.WriteString(message)
			writer.Flush()
		}
	}
}
