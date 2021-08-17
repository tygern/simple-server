package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func Create(port string) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong\n"))
	})
	mux.HandleFunc("/pong", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ping\n"))
	})

	return &http.Server{
		Addr:    port,
		Handler: mux,
	}
}

func Start(server *http.Server) {
	log.Printf("starting server at %s\n", server.Addr)
	log.Println(server.ListenAndServe())
}

func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	fmt.Println("shutting down server")
	err := server.Shutdown(ctx)
	if err != nil {
		fmt.Println(err)
	}
}
