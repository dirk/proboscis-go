package potcp

import (
  "testing"
  // "fmt"
  "proboscis-go"
)

func TestHandlerRegistration(t *testing.T) {
  server := NewServer()
  if len(server.Handlers) != 0 {
    t.Fatal("Server should have no handlers; currently:", len(server.Handlers))
  }
  
  handler_function := func(req *proboscis.Request) *proboscis.Response {
    return req.MakeResonse()
  }
  handler := NewHandler("test", "text", handler_function)
  server.Register(handler)
  
  if len(server.Handlers) != 1 {
    t.Fatal("Server should have 1 handler; currently:", len(server.Handlers))
  }
}
