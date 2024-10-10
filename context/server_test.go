package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
  data := "hello, world"

  t.Run("return data from the store", func(t *testing.T) {
    store := &SpyStore{response: data, t: t}
    srv := Server(store)

    request := httptest.NewRequest(http.MethodGet, "/", nil)
    response := httptest.NewRecorder()

    srv.ServeHTTP(response, request)

    if response.Body.String() != data {
      t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
    }

  })

  t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
    store := &SpyStore{response: data, t: t}
    srv := Server(store)

    request := httptest.NewRequest(http.MethodGet, "/", nil)
  
    cancellingCtx, cancel := context.WithCancel(request.Context())
    time.AfterFunc(5 * time.Millisecond, cancel)
    request = request.WithContext(cancellingCtx)

    response := &SpyResponseWriter{} 

    srv.ServeHTTP(response, request)

    if response.written {
      t.Error("a response should not have been written")
    }
  })
}

// SpyStore definition

type SpyStore struct {
	response  string
	t         *testing.T
}

func (ss *SpyStore) Fetch(ctx context.Context) (string, error) {
  data := make(chan string, 1)

  // simulate the delay to get the response
  go func() {
    var result string
    for _, c := range ss.response {
      select {
      case <- ctx.Done():
        log.Println("Context is closed")
        return 
      default:
        time.Sleep(10 * time.Millisecond)
        result += string(c)
      }
    }

    data <- result
  }()

  select {
  case <- ctx.Done():
    return "", ctx.Err() 
  case res := <-data: 
    return res, nil
  }
}

// SpyResponseWriter definition
type SpyResponseWriter struct {
  written bool
}

func (srw *SpyResponseWriter) Header() http.Header {
  srw.written = true
  return nil
}

func (srw *SpyResponseWriter) Write([]byte) (int, error) {
  srw.written = true
  return 0, errors.New("not implemented")
}

func (srw *SpyResponseWriter) WriteHeader(statusCode int) {
  srw.written = true
}
