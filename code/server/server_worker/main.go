package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type server struct {
	router  *http.ServeMux
	server  *http.Server
	counter int
	inc     chan chan int
}

func (s *server) runCounter() {
	go func() {
		for resultCh := range s.inc {
			s.counter++
			resultCh <- s.counter
		}
	}()
}

func (s *server) handleCount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resultCh := make(chan int)
		s.inc <- resultCh
		c := <-resultCh
		io.WriteString(w, fmt.Sprintf("counter: %d", c))
	}
}

func main() {
	s := &server{
		router: http.NewServeMux(),
		server: &http.Server{
			Addr:           ":8080",
			ReadTimeout:    3 * time.Second,
			WriteTimeout:   3 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
	s.runCounter()
	s.router.HandleFunc("/", s.handleCount())
	s.server.Handler = s.router
	fmt.Println("Server started at :8080")
	err := s.server.ListenAndServe()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
