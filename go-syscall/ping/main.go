package ping

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	syscall "golang.org/x/sys/unix"
)

func Ping(addr string) {
	sendData := make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(sendData, time.Now().UnixMilli())
	ipaddr, err := net.ResolveIPAddr("ip", addr)
	if err != nil {
		syscall.Write(2, []byte("error: resolve address : "+err.Error()+"\n"))
		syscall.Exit(1)
	}
	pid := syscall.Getpid()
	icmpMsg := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   pid & 0xffff,
			Seq:  0,
			Data: sendData,
		},
	}
	icmpMsgBytes, err := icmpMsg.Marshal(nil)
	if err != nil {
		syscall.Write(2, []byte("error: marshal icmp message : "+err.Error()+"\n"))
		syscall.Exit(1)
	}

	sockFd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_ICMP)
	if err != nil {
		syscall.Write(2, []byte("error: create socket : "+err.Error()+"\n"))
		syscall.Exit(1)
	}
	defer syscall.Close(sockFd)
	err = syscall.Connect(sockFd, &syscall.SockaddrInet4{Port: 0, Addr: [4]byte(ipaddr.IP)})
	if err != nil {
		syscall.Write(2, []byte("error: socket connect : "+err.Error()+"\n"))
		syscall.Exit(1)
	}
	err = syscall.Sendto(sockFd, icmpMsgBytes, 0, nil)
	if err != nil {
		syscall.Write(2, []byte("error: sendto : "+err.Error()+"\n"))
		syscall.Exit(1)
	}
	recvBuf := make([]byte, 65534)
	recvBytes, _, err := syscall.Recvfrom(sockFd, recvBuf, 0)
	if err != nil {
		syscall.Write(2, []byte("error: recvfrom : "+err.Error()+"\n"))
		syscall.Exit(1)
	} else {
		rm, err := icmp.ParseMessage(ipv4.ICMPTypeEcho.Protocol(), recvBuf[:recvBytes])
		if err == nil && rm.Type == ipv4.ICMPTypeEchoReply {
			echo, ok := rm.Body.(*icmp.Echo)
			if !ok {
				syscall.Write(2, []byte("error: not echo reply\n"))
				syscall.Exit(1)
			} else {
				t, _ := binary.Varint(echo.Data)
				syscall.Write(1, []byte(fmt.Sprintf("%dms\n", time.Now().UnixMilli()-t)))
			}
		} else {
			syscall.Write(2, []byte("error: packet parse : "+err.Error()+"\n"))
			syscall.Exit(1)
		}
	}
}
