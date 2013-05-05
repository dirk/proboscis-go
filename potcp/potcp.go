package potcp

import (
  "proboscis-go"
  "net"
  "fmt"
)

type HandlerFunction func(*proboscis.Request) *proboscis.Response

type Handler struct {
  Method string
  Format string
  Function HandlerFunction
}

type Server struct {
  Handlers map[string]*Handler
  Conn     net.Conn
}

type Client struct {
  Conn net.Conn
}

func NewServer() *Server {
  var server *Server
  server = &Server{make(map[string]*Handler), nil}
  return server
}
func NewHandler(method string, format string, hf HandlerFunction) *Handler {
  var handler *Handler
  handler = &Handler{method, format, hf}
  return handler
}
func NewClient(conn net.Conn) *Client {
  var client *Client
  client = &Client{conn}
  return client
}

// SERVER ---------------------------------------------------------------------

func (server *Server) Register(handler *Handler) {
  server.Handlers[handler.Method] = handler
}
func (server *Server) ServeConn(conn net.Conn) {
  server.Conn = conn
  // FIXME: Make this work
  
  var req *proboscis.Request
  var err error
  
  req, err = DecodeRequest(conn)
  
  fmt.Printf("req: %#v\n", req)
  fmt.Printf("err: %#v\n", err)
}

// CLIENT ---------------------------------------------------------------------

func (client *Client) CallRequest(req *proboscis.Request) (*proboscis.Response, error) {
  EncodeRequest(req, client.Conn)
  
  return nil, nil
}

func (client *Client) CallString(method, format, data string) (*proboscis.Response, error) {
  var req *proboscis.Request
  req = proboscis.NewRequest()
  req.Method = method; req.Format = format
  
  req.Data = []byte(data)
  
  var rep *proboscis.Response
  var err error
  rep, err = client.CallRequest(req)
  return rep, err
}
