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
}

type Client struct {
  Conn net.Conn
}

func NewServer() *Server {
  var server *Server
  server = &Server{make(map[string]*Handler)}
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
  server.Handlers[handler.Method+"."+handler.Format] = handler
}
func (server *Server) ServeConn(conn net.Conn) error {
  var req *proboscis.Request
  var rep *proboscis.Response = nil
  var err error
  
  req, err = DecodeRequest(conn)
  if err != nil {
    // TODO: Maybe send a response here?
    return err
  }
  
  // fmt.Printf("req: %#v\n", req)
  // fmt.Printf("err: %#v\n", err)
  
  var handler *Handler
  handler = server.Handlers[req.Method+"."+req.Format]
  
  if handler == nil {
    rep = CreateMethodNotFoundResponse(req)
  } else {
    rep = handler.Function(req)
  }
  
  // fmt.Printf("rep: %#v\n", rep)
  
  err = EncodeResponse(rep, conn)
  if err != nil {
    // TODO: And also maybe here?
    return err
  }
  
  return conn.Close()
}

func CreateMethodNotFoundResponse(req *proboscis.Request) *proboscis.Response {
  var rep *proboscis.Response
  rep = proboscis.NewResponse()
  rep.Status = "404"
  rep.Format = "text"
  data_string := fmt.Sprintf("Method %s.%s not found", req.Method, req.Format)
  rep.Data = []byte(data_string)
  return rep
}

// CLIENT ---------------------------------------------------------------------

func (client *Client) CallRequest(req *proboscis.Request) (*proboscis.Response, error) {
  EncodeRequest(req, client.Conn)
  
  var rep *proboscis.Response
  var err error
  rep, err = DecodeResponse(client.Conn)
  
  return rep, err
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
