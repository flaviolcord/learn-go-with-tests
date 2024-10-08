package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
  t.Run("compare speeds of servers", func(t *testing.T) {
    fastServer := makeDelayedServer(0 * time.Millisecond)
    slowServer := makeDelayedServer(20 * time.Millisecond)

    defer slowServer.Close()
    defer fastServer.Close()

    slowUrl := slowServer.URL 
    fastUrl := fastServer.URL

    want := fastUrl
    got, _ := Racer(slowUrl, fastUrl)

    if got != want {
      t.Errorf("got: %q, want: %q", got, want)
    }
  })

  t.Run("return error if a server takes more than 10s to respond", func(t *testing.T) {
    server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
  })
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
  return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    time.Sleep(delay)
    w.WriteHeader(http.StatusOK)
  }))
}
