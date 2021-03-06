package real

// (c) Christian Maurer   v. 130115 - license see murus.go

import (
  "math"; "strconv"
  "murus/obj"; "murus/str"; "murus/col"; "murus/box"; "murus/errh"
)
const (
  pack = "real"
)
var
  bx *box.Imp = box.New ()


func NDigits (x float64) uint {
//
  return 1 + uint(math.Floor (math.Log (math.Abs (x)) / math.Ln10))
}


func valid (x float64) bool {
//
  return ! math.IsInf (x, 1) && ! math.IsInf (x, -1) && ! math.IsNaN (x)
}


func number (s string) float64 {
//
  if x, err:= strconv.ParseFloat (s, 64); err == nil {
    return x
  }
  return math.NaN ()
}


func defined (s string) (float64, bool) {
//
  str.Move (&s, true)
  str.RemSpaces (&s)
  r, e:= strconv.ParseFloat (s, 64)
  if e != strconv.ErrSyntax {
    return r, true
  }
  return math.NaN(), false
}


func string_ (x float64) string {
//
  s:= strconv.FormatFloat (x, 'f', 2, 64)
  str.Move (&s, true)
  str.RemSpaces (&s)
  return s
}


func setColours (f, b col.Colour) {
//
  bx.Colours (f, b)
}


func write (x float64, l, c uint) {
//
  bx.Write (String (x), l, c)
}


func edit (x *float64, l, c uint) {
//
  s:= String (*x)
  p:= uint(0)
  for {
    bx.Edit (&s, l, c)
    if ! str.Contains (s, 'e', &p) && ! str.Contains (s, 'E', &p) {
      *x = number (s)
      if ! math.IsNaN (*x) {
        break
      }
    } else {
    }
    errh.ErrorPos ("keine Zahl", 0, l + 1, c)
  }
  Write (*x, l, c)
}


func Codelen () uint {
//
  return 8
}


func Encode (x float64) []byte {
//
  return obj.Encode (x) // obj.Encode (math.Float64bits (x))
}


func Decode (b []byte) float64 {
//
  return obj.Decode (0., b).(float64) // math.Float64frombits (obj.Decode (uint64(0), b).(uint64))
}

/*
func val (op Operation, x, y float64) float64 {
//
  switch op { case Plus:
    return x + y
  case Minus:
    return x - y
  case Times:
    return x * y
  case Div:
    return x / x
  case ToThe:
    return math.Pow (x, y)
  case Percent:
    return x / 100.0 * y
  }
  return 0.0
}
*/

func init () {
//
//  bx.SetNumerical ()
//  setFormat ()
}
