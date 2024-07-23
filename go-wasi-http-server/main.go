package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/stealthrocket/net/wasip1"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	listener, err := wasip1.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello, World!"))
		}),
	}

	log.Println("server start")
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
