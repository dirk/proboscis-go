package potcp

import (
  "testing"
  "fmt"
  "proboscis-go"
  "net"
)

func SetupServer() *Server {
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
  return server
}

func TestClientServer(t *testing.T) {
  events := make(chan string, 1)
  
  addr := "localhost:9999"
  
  // Server stuff
  go func() {
    var server *Server
    server = SetupServer()
    
    listener, err := net.Listen("tcp", addr)
    defer listener.Close()
    if err != nil { panic(err) }
    
    fmt.Println("INFO TestClientServer/Server accepting...")
    
    conn, err := listener.Accept()
    if err != nil { panic(err) }
    
    server.ServeConn(conn)
    conn.Close()
    
    // events <- "Server done"
  }()
  
  // Client stuff
  go func() {
    conn, err := net.Dial("tcp", addr)
    if err != nil { panic(err) }
    
    var client *Client
    client = NewClient(conn)
    
    
  }
  
  var event string
  event = <- events
  fmt.Printf("INFO TestClientServer/event: %s\n", event)
  
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
