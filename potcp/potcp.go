package potcp

import (
  "proboscis-go"
  "net"
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

// SERVER ---------------------------------------------------------------------

func (server *Server) Register(handler *Handler) {
  server.Handlers[handler.Method] = handler
}
func (server *Server) ServeConn(conn net.Conn) {
  server.Conn = conn
  // FIXME: Make this work
}
