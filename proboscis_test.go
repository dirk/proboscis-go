package proboscis

import (
  "testing"
  // "fmt"
)

func TestNewRequest(t *testing.T) {
  req := NewRequest()
  if len(req.Data) != 0 {
    t.Fatal("Request should have empty data; currently:", len(req.Data))
  } else {
    // fmt.Println("PASS TestNewRequest/len(req.Data)")
  }
}
