package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	ts := &server{
		router: http.NewServeMux(),
		server: &http.Server{
			Addr:           ":8080",
			ReadTimeout:    3 * time.Second,
			WriteTimeout:   3 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
		mtx: &sync.Mutex{},
	}
	handler := ts.handleCount()
	getCounter := func(out chan []byte) {
		req := httptest.NewRequest("GET", "http://test.xy/foo", nil)
		w := httptest.NewRecorder()
		handler(w, req)
		resp := w.Result()
		body, _ := io.ReadAll(resp.Body)
		out <- body
	}
	out := make(chan []byte)
	go getCounter(out)
	go getCounter(out)
	go getCounter(out)
	go getCounter(out)
	fmt.Println(string(<-out))
	fmt.Println(string(<-out))
	fmt.Println(string(<-out))
	fmt.Println(string(<-out))
}
