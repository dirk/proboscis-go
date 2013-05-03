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
func EncodeRequest(req *Request, w io.Writer) error {
  w.Write([]byte(req.Method))
  w.Write(period_byte_slice)
  w.Write([]byte(req.Format))
  w.Write(colon_byte_slice)
  w.Write([]byte(strconv.Itoa(len(req.Data))))
  w.Write(colon_byte_slice)
  w.Write(req.Data)
  return nil
}
func DecodeRequest(r bufio.Reader) (*Request, error) {
  var req *Request
  req = NewRequest()
  
  // TODO: At least some sanity-checking
  method, _ := r.ReadString(period_byte)
  method = method[0:len(method) - 2]
  
  format, _ := r.ReadString(colon_byte)
  format = format[0:len(format) - 2]
  
  length_string, _ := r.ReadString(colon_byte)
  length_string = length_string[0:len(length_string) - 2]
  
  length, _ := strconv.Atoi(length_string)
  
  data := make([]byte, length)
  read, _ := r.Read(data)
  
  if read != length {
    return nil, fmt.Errorf(
      "Error reading data (expected %d bytes, read %d)", length, read,
    )
  }
  
  req.Method = method
  req.Format = format
  req.Length = uint32(length)
  req.Data   = data
  return req, nil
}
  rep.Status = "200"
  rep.Format = req.Format
  return rep
}

func NewResponse() *Response {
  var rep *Response
  rep = &Response{"", "", 0, []byte{}}
  return rep
}
