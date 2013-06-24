package obj

// (c) Christian Maurer   v. 130115 - license see murus.go

// TODO DeepClone

import (
//  "fmt"
  "reflect"
)
type
  Equaler interface { // x denotes the calling object.

// Returns true, iff the x has the same type as y
// and coincides with it in its value[s].
  Eq (y Object) bool

// If y has the same type as x, then x.Eq(y) (y is unchanged).
  Copy (y Object)

// Returns a clone of x, i.e. x.Eq(x.Clone()).
  Clone () Object
}


func TypeEq (a, b Any) bool {
//
  x, y:= reflect.TypeOf (a), reflect.TypeOf (b)
  if x != y { TypeNotEqPanic (x, y) }
  return x == y
}


func Eq (a, b Any) bool {
//
//  println ("obj.Eq")
  if ! TypeEq (a, b) {
    return false
  }
  if X, ok:= a.(Object); ok {
    return X.Eq (b.(Object))
  }
  if x, ok:= a.([]byte); ok {
    for i:= 0; i < len (x); i++ {
      if a.([]byte)[i] != b.([]byte)[i] {
        return false
      }
      return true
    }
  }
  if Atomic (a) {
    return a == b
  }
  return reflect.DeepEqual (a, b)
}


func Clone (a Any) Any {
//
//  println ("obj.Clone")
  if Atomic (a) {
    return a
  }
  switch a.(type) { case Object:
    xx:= a.(Object).Clone () // .(Object)
    return xx
/*
  case bool:
    return a.(bool)
  case string:
    return a.(string)
  case uint8:
    return a.(uint8)
  case uint16:
    return a.(uint16)
  case uint32:
    return a.(uint32)
  case uint:
    return = a.(uint)
  case uint64:
    return a.(uint64)
  case int8:
    return a.(int8)
  case int16:
    return a.(int16)
  case int32:
    return a.(int32)
  case int:
    return a.(int)
  case int64:
    return a.(int64)
  case float32:
    return a.(float32)
  case float64:
    return a.(float64)
  case complex64:
    return a.(complex64)
  case complex128:
    return a.(complex128)
*/
  case []byte:
    b:= make ([]byte, len (a.([]byte)))
    copy (b, a.([]byte))
    return b
  default:
    // TODO
  }
  return nil
}
