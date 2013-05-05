
// Package proboscis provides interfaces and implementations of the Proboscis
// RPC meta-framework.
// 
package proboscis

var ValidTypes []string = []string{"text", "json", "proto", "thrift"}

type Request struct {
  Method string
  Format string
  Length int
  Data   []byte
}

type Response struct {
  Status string
  Format string
  Length int
  Data   []byte
}

func NewRequest() *Request {
  var req *Request
  req = &Request{"", "", 0, make([]byte, 0)}
  return req
}



func (req *Request) MakeResponse() *Response {
  var rep *Response
  rep = NewResponse()
  rep.Status = "200"
  rep.Format = req.Format
  return rep
}
// func (req *Request) Encode(w io.Writer) error {
//   return EncodeRequest(req, w)
// }

func NewResponse() *Response {
  var rep *Response
  rep = &Response{"", "", 0, make([]byte, 0)}
  return rep
}
