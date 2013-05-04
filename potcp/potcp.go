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


func (server *Server) Register(handler *Handler) {
  server.Handlers[handler.Method] = handler
}
func (server *Server) ServeConn(conn net.Conn) {
  // FIXME: Make this work
}
