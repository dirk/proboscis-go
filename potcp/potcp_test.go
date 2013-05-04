package potcp

import (
  "testing"
  // "fmt"
  "proboscis-go"
  "net"
)

func TestClientServer(t *testing.T) {
  var handler_function HandlerFunction
  handler_function = func(req *proboscis.Request) *proboscis.Response {
    res := req.MakeResponse()
    res.Format = "text"
    res.Data = req.Data
    return res
  }
  
  var server *Server
  server = NewServer()
  
  var handler *Handler
  handler = NewHandler("echo", "text", handler_function)
  
  server.Register(handler)
  
  listener, err := net.Listen("tcp", "localhost:9999")
  defer listener.Close()
  if err != nil { panic(err) }
  conn, err := listener.Accept()
  if err != nil { panic(err) }
  
  server.ServeConn(conn)
}

func TestHandlerRegistration(t *testing.T) {
  server := NewServer()
  if len(server.Handlers) != 0 {
    t.Fatal("Server should have no handlers; currently:", len(server.Handlers))
  }
  
  handler_function := func(req *proboscis.Request) *proboscis.Response {
    return req.MakeResponse()
  }
  handler := NewHandler("test", "text", handler_function)
  server.Register(handler)
  
  if len(server.Handlers) != 1 {
    t.Fatal("Server should have 1 handler; currently:", len(server.Handlers))
  }
}
