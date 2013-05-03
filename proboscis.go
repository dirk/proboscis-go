package proboscis

type Request struct {
  Method string
  Format string
  Length uint32
  Data   []byte
}

type Response struct {
  Status string
  Format string
  Length uint32
  Data   []byte
}

func NewRequest() *Request {
  var req *Request
  req = &Request{"", "", 0, []byte{}}
  return req
}
func (req *Request) MakeResonse() *Response {
  rep := NewResponse()
  rep.Status = "200"
  rep.Format = req.Format
  return rep
}

func NewResponse() *Response {
  var rep *Response
  rep = &Response{"", "", 0, []byte{}}
  return rep
}
