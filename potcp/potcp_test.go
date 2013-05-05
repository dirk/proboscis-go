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
  var event string
  
  addr := "localhost:9999"
  
  // Server stuff
  go func() {
    var server *Server
    server = SetupServer()
    
    listener, err := net.Listen("tcp", addr)
    defer listener.Close()
    if err != nil { panic(err) }
    
    fmt.Println("INFO TestClientServer/Server accepting")
    events <- "Server accepting"
    
    conn, err := listener.Accept()
    if err != nil { panic(err) }
    
    server.ServeConn(conn)
    conn.Close()
    
    events <- "Server done"
  }()
  
  event = <- events
  if event != "Server accepting" {
    fmt.Printf("INFO TestClientServer/event: %s\n", event)
    t.Fatal("Unexpected event: %s", event)
    return
  }
  
  // Client stuff
  go func() {
    conn, err := net.Dial("tcp", addr)
    if err != nil { panic(err) }
    defer conn.Close()
    
    var client *Client
    client = NewClient(conn)
    
    var rep *proboscis.Response
    // var req *proboscis.Request
    
    fmt.Println("INFO TestClientServer/client.CallString")
    message := "Hello world!"
    rep, err = client.CallString("echo", "text", message)
    
    // fmt.Printf("rep: %#v\n", rep)
    // fmt.Printf("err: %#v\n", err)
    
    if err != nil {
      t.Fatal("Error: %s", err)
      events <- "Client error"
      return
    }
    if rep.Status != "200" {
      t.Fatal("Status not 200 (%q)", rep.Status)
      events <- "Client error"
      return
    }
    if string(rep.Data) != message {
      t.Fatal(
        "Messages don't match (%q != %q)", message, string(rep.Data),
      )
    }
    
    events <- "Client done"
  }()
  
  event = <- events
  fmt.Printf("INFO TestClientServer/event: %s\n", event)
  
  event = <- events
  fmt.Printf("INFO TestClientServer/event: %s\n", event)
  
  fmt.Println("PASS TestClientServer")
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
