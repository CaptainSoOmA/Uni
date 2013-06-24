package integ

// (c) Christian Maurer   v. 130312 - license see murus.go

import (
  . "murus/ker"; "murus/str"
  "murus/col"; "murus/box"; "murus/errh"; "murus/nat"
)
const (
  max = 11 // sign plus maximal 10 digits
  m2 = uint(- MinInt) // == uint(MaxInt + 1)
)
var (
  bx *box.Imp = box.New ()
  width uint
)


func defined (z *int, s string) bool {
//
  str.Move (&s, true)
  negative:= s[0] == '-'
  var n uint
  if negative {
    n = uint(len (s))
    s = str.Part (s, 1, n - 1)
  }
  if str.Empty (s) {
    return false
  }
  if nat.Defined (&n, s) {
    if negative {
      if n < m2 {
        *z = - int(n)
        return true
      } else if n == m2 {
        *z = MinInt
        return true
      }
    } else if n <= MaxInt {
      *z = int(n)
      return true
    }
  }
  return false
}


func wd (z int) uint {
//
  if z < 0 { z = - z }
  return 1 + nat.Wd (uint(z))
}


func string_(z int) string {
//
  s:= ""
  if z < 0 {
    s = "-"
    z = -z
  }
  return s + nat.String (uint(z))
}


func stringFmt (z int, l uint) string {
//
  a:= " "; if z < 0 { a = "-"; z = -z }
  w:= Wd (z)
  if l < w { l = w }
  return a + nat.StringFmt (uint(z), l - 1, false)
}


func setColours (f, b col.Colour) {
//
  bx.Colours (f, b)
}


func write (z int, l, c uint) {
//
  w:= Wd (z)
  if w > c + 1 { return }
  bx.Wd (w)
//  scr.SwitchFontsize (scr.Normal)
  bx.Write (StringFmt (z, w), l, c + 1 - w)
}


func setWd (w uint) {
//
  if w == 0 {
    width = 2
  } else if w > max {
    width = max
  } else {
    width = w
  }
}


func edit (z *int, l, c uint) {
//
  w:= Wd (*z)
  if width > w { w = width }
  bx.Wd (w)
  s:= StringFmt (*z, w)
  for {
    bx.Edit (&s, l, c + 1 - w)
    if defined (z, s) {
      break
    } else {
      errh.Error ("keine Zahl", 0) // , l + 1, c)
    }
  }
  bx.Write (StringFmt (*z, w), l, c + 1 - w)
}


func init () {
//
//  bx.SetNumerical() // TODO
}
