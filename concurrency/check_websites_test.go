package concurrency

import (
	"reflect"
	"testing"
)

func MockWebSitesChecker(url string) bool {
  return url != "go.dev"
}

func TestCheckerWebsites(t *testing.T) {
  webSites := []string{
    "google.com",
    "gmail.com",
    "go.dev",
  }

  want := map[string]bool{
    "google.com": true,
    "gmail.com": true,
    "go.dev": false,
  }

  got := CheckerWebsites(MockWebSitesChecker, webSites)

  if !reflect.DeepEqual(got, want) {
    t.Fatalf("wanted: %v, got: %v", want, got)
  }
}
