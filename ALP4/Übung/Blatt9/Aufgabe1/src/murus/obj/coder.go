package obj

// (c) Christian Maurer   v. 131115 - license see murus.go

import (
//  "math"
  "unsafe"
  "reflect"
  "strconv"
  "murus/ker"
)
type
  Coder interface {

// Returns the number of bytes, that are needed
// to serialize x uniquely revertibly.
  Codelen () uint

//// Returns the minimal number of bytes,
//  that are needed to code any object. // TODO
//  MinCodelen () uint

// x.Eq (x.Decode (x.Encode())
  Encode () []byte

// Pre: b is result of y.Encode() for some y of the type of x.
// x.Eq(y); x.Encode() == b, i.e. those slices coincide.
  Decode (b []byte)
}


func Codelen (a Any) uint {
//
  var n uint
  switch a.(type) { case Object:
    n = (a.(Object)).Codelen ()
  case string:
    n = uint(len (a.(string)))
  case []byte:
    n = uint(len (a.([]byte)))
  default:
    n = uint(reflect.TypeOf(a).Size())
  }
  return n
}


func Encode (a Any) []byte {
//
  var b []byte
  switch a.(type) { case Object:
    b = a.(Object).Encode ()
  case byte:
    b = make ([]byte, 1)
    b[0] = a.(byte)
/*/
/*/
  case float64:
    b = make ([]byte, 8)
    f:= a.(float64)
//    u:= math.Float64bits (f)
    u:= *(*uint64)(unsafe.Pointer(&f))
    for i:= 7; i >= 0; i-- {
      b[i] = byte(u)
      u >>= 8
    }
/*/
/*/
  case string:
    b = ([]byte)(a.(string))
  case []byte:
    b = make ([]byte, len (a.([]byte)))
    copy (b, a.([]byte))
  default:
    s:= reflect.TypeOf (a).Size ()
    b = make ([]byte, s)
    n:= uintptr(unsafe.Sizeof(a)) // 8 (386) resp. 16 (amd64)
    if s < n {
      n = uintptr(n / 2)
    } else {
      n = uintptr(0)
    }
    addr:= uintptr(unsafe.Pointer(&a))
    for i:= uintptr(0); i < s; i++ {
      b[i] = *(*byte)(unsafe.Pointer(addr + n + i))
    }
  }
  return b
}


func Decode (a Any, b []byte) Any {
//
  switch a.(type) { case Object:
    a.(Object).Decode (b)
    return a
  case byte:
    a = b[0]
/*/
/*/
  case float64: // workaround (bug in Go ?)
    u:= uint64 (0)
    for i:= 0; i < 8; i++ {
      u <<= 8
      u += uint64(b[i])
    }
//    return math.Float64frombits (u)
    return *(*float64)(unsafe.Pointer(&u))
/*/
/*/
  case string:
    return string(b)
  case []byte:
    copy (a.([]byte), b)
    return a
  default:
    n:= uintptr(unsafe.Sizeof(a)) // 8 (386) resp. 16 (amd64)
/*/
    t:= reflect.TypeOf (a)
    s:= t.Size ()
/*/
    s:= reflect.TypeOf (a).Size ()
    if s != uintptr(len (b)) {
      ker.Panic ("obj.Decode error: s == " + strconv.Itoa (int(s)) +
                 ", len(b) == " + strconv.Itoa(int(uintptr(len (b)))))
    }
    if s < n {
      n = uintptr(n / 2)
    } else {
      n = uintptr(0)
    }
    addr:= uintptr(unsafe.Pointer(&a))
    for i:= uintptr(0); i < s; i++ {
      *(*byte)(unsafe.Pointer(addr + n + i)) = b[i]
    }
  }
  return a
}
