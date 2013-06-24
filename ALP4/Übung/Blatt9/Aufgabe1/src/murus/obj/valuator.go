package obj

// (c) Christian Maurer   v. 130115 - license see murus.go

import
  "math"
type
  Valuator interface {

// Returns the value of x.
  Val () uint

// Returns true, iff x is defined with x.Val() == n.
  Set (n uint) bool
}


func nat (r float64) uint {
//
  if r + 0.5 >= float64(1<<32 - 1) {
    return 1<<32 - 1
  }
  return uint(r + 0.5)
}


func Val (a Any) uint {
//
  var n uint = 1
  switch a.(type) { case Valuator:
    n = (a.(Valuator)).Val ()
  case bool:
    if ! a.(bool) {
      n = 0
    }
  case int8:
    i:= a.(int8)
    if i < 0 { i = -i }
    n = uint(i)
  case int16:
    i:= a.(int16)
    if i < 0 { i = -i }
    n = uint(i)
  case int32:
    i:= a.(int32)
    if i < 0 { i = -i }
    n = uint(i)
  case int:
    i:= a.(int)
    if i < 0 { i = -i }
    n = uint(i)
  case byte:
    n = uint(a.(byte))
  case uint16:
    n = uint(a.(uint16))
  case uint32:
    n = uint(a.(uint32))
  case uint:
    n = a.(uint)
  case float32:
    n = nat (math.Trunc (float64(a.(float32) + 0.5)))
  case float64:
    n = nat (math.Trunc (a.(float64) + 0.5))
  case complex64:
    c:= a.(complex64)
    n = nat (math.Trunc (math.Sqrt(float64(real(c) * real(c) + imag(c) * imag(c))) + 0.5))
  case complex128:
    c:= a.(complex128)
    n = nat (math.Trunc (math.Sqrt(real(c) * real(c) + imag(c) * imag(c)) + 0.5))
  case string:
    // TODO sum of bytes of the string ? Hash-Code ?
  }
  return n
}
