package potcp

import (
  "proboscis-go"
  "io"
  "fmt"
  "bufio"
  "strconv"
)

var period_byte byte = byte('.')
var colon_byte  byte = byte(':')
var period_byte_slice []byte = []byte{period_byte}
var colon_byte_slice  []byte = []byte{colon_byte}

func EncodeResponse(rep *proboscis.Response, w io.Writer) error {
  w.Write([]byte(rep.Status))
  w.Write(colon_byte_slice)
  w.Write([]byte(rep.Format))
  w.Write(colon_byte_slice)
  w.Write([]byte(strconv.Itoa(len(rep.Data))))
  w.Write(colon_byte_slice)
  w.Write(rep.Data)
  return nil
}
func DecodeResponse(reader io.Reader) (*proboscis.Response, error) {
  var r *bufio.Reader
  r = bufio.NewReader(reader)
  
  var rep *proboscis.Response
  var status, format, length_string string
  var length, read int
  var data []byte
  
  status, _ = r.ReadString(colon_byte)
  status = status[0:len(status) - 1]
  
  format, _ = r.ReadString(colon_byte)
  format = format[0:len(format) - 1]
  
  length_string, _ = r.ReadString(colon_byte)
  length_string = length_string[0:len(length_string) - 1]
  
  length, _ = strconv.Atoi(length_string)
  
  data = make([]byte, length)
  read, _ = r.Read(data)
  
  if read != length {
    return nil, fmt.Errorf(
      "Error reading data (expected %d bytes, read %d)", length, read,
    )
  }
  
  rep = proboscis.NewResponse()
  rep.Status = status
  rep.Format = format
  rep.Length = length
  rep.Data   = data
  return rep, nil
}

func EncodeRequest(req *proboscis.Request, w io.Writer) error {
  w.Write([]byte(req.Method))
  w.Write(period_byte_slice)
  w.Write([]byte(req.Format))
  w.Write(colon_byte_slice)
  w.Write([]byte(strconv.Itoa(len(req.Data))))
  w.Write(colon_byte_slice)
  w.Write(req.Data)
  return nil
}
func DecodeRequest(reader io.Reader) (*proboscis.Request, error) {
  var r *bufio.Reader
  r = bufio.NewReader(reader)
  
  var req *proboscis.Request
  var method, format, length_string string
  var length, read int
  var data []byte
  
  // TODO: At least some sanity-checking
  method, _ = r.ReadString(period_byte)
  method = method[0:len(method) - 1]
  
  format, _ = r.ReadString(colon_byte)
  format = format[0:len(format) - 1]
  
  length_string, _ = r.ReadString(colon_byte)
  length_string = length_string[0:len(length_string) - 1]
  
  length, _ = strconv.Atoi(length_string)
  
  data = make([]byte, length)
  read, _ = r.Read(data)
  
  if read != length {
    return nil, fmt.Errorf(
      "Error reading data (expected %d bytes, read %d)", length, read,
    )
  }
  
  req = proboscis.NewRequest()
  req.Method = method
  req.Format = format
  req.Length = length
  req.Data   = data
  return req, nil
}
