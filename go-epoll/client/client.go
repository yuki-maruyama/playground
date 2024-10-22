package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
	"time"
)

const (
	serverAddress = "localhost:12345"
	numClients    = 100
	message       = "Hello from client %d\n"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	totalMessages := 0

	for i := 0; i < numClients; i++ {
		wg.Add(1)
		go func(clientID int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", serverAddress)
			if err != nil {
				fmt.Println("Error connecting to server:", err)
				return
			}
			defer conn.Close()

			reader := bufio.NewReader(conn)
			writer := bufio.NewWriter(conn)

			go func() {
				for {
					_, err := reader.ReadString('\n')
					if err != nil {
						fmt.Println("Error reading from server:", err)
						return
					}
					// fmt.Printf("Client %d received: %s", clientID, response)
				}
			}()

			for {
				_, err := writer.WriteString(fmt.Sprintf(message, clientID))
				if err != nil {
					fmt.Println("Error writing to server:", err)
					return
				}
				writer.Flush()
				mu.Lock()
				totalMessages++
				mu.Unlock()
				time.Sleep(1 * time.Millisecond)
			}
		}(i)
	}

	time.Sleep(10 * time.Second)
	fmt.Printf("Total messages sent in 10 seconds: %d\n", totalMessages)
	panic("done")
	// wg.Wait()
}
