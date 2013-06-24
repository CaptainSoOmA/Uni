package breal

// (c) Christian Maurer   v. 130115 - license see murus.go

// TODO: more than 2 digits after the decimal point

import (
//  "math"
  . "murus/obj"
  "murus/col"; "murus/str"; "murus/box"; "murus/errh"; "murus/font"; "murus/pbox"
  "murus/nat"; "murus/real"
)
const
  m = 9 // 
type
  Imp struct {
            r float64
      pre, wd uint
      invalid float64
       cF, cB col.Colour
           fo font.Font
             }
var (
  bx *box.Imp = box.New()
  pbx *pbox.Imp = pbox.New()
)


func (x *Imp) imp (Y Object) *Imp {
//
  y, ok:= Y.(*Imp)
  if ! ok || x.wd != y.wd { TypeNotEqPanic (x, Y) }
  return y
}


func exp (n uint) float64 {
//
  if n == 0 { return 1 }
  return 10 * exp (n - 1)
}


func New (n uint) *Imp {
//
  x:= new (Imp)
  if n == 0 { n = 1 }
  if n > m { n = m }
  x.pre, x.wd = n, 1 + n + 1 + 2
  x.invalid = exp (n)
  x.cF, x.cB = col.ScreenF, col.ScreenB
  return x
}


func (x *Imp) Empty () bool {
//
  return x.r == x.invalid
}


func (x *Imp) Clr () {
//
  x.r = x.invalid
}


func (x *Imp) Copy (Y Object) {
//
  x.r = x.imp (Y).r
}


func (x *Imp) Clone () Object {
//
  y:= New (x.pre)
  y.r = x.r
  return y
}


func (x *Imp) Eq (Y Object) bool {
//
  return x.r == x.imp (Y).r
}


func (x *Imp) Less (Y Object) bool {
//
  y:= x.imp (Y)
  if x.r == x.invalid || y.r == y.invalid {
    return false
  }
  return x.r < y.r
}


func (x *Imp) Codelen () uint {
//
  return 8
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, x.Codelen())
  copy (b[:8], Encode (x.r))
  return b
}


func (x *Imp) Decode (b []byte) {
//
  x.r = float64(Decode (float64(0), b[:x.Codelen()]).(float64))
}


func (x *Imp) Defined (s string) bool {
//
  if uint(len (s)) > x.wd { return false }
  str.RemAllSpaces (&s)
//  n:= x.wd / 2
//  P, L:= make ([]uint, n), make ([]uint, n)
//  n = nat.NDigitSequences (s, &P, &L)
  n, t, p, l:= nat.DigitSequences (s)
  if n == 0 || n > 2 || l[0] > x.pre {
    return false
  }
  if n == 2 {
    c:= s[p[1] - 1]
    if l[1] > 2 || ! (c == '.' || c == ',') {
      return false
    }
  }
  var n1 uint
  nat.Defined (&n1, t[0])
  x.r = float64(n1)
  if n == 2 {
    nat.Defined (&n, t[1])
    if n < 10 { n *= 10 }
    x.r = x.r + float64(n) / 100
  }
  if s[0] == '-' { x.r = - x.r }
  return true
}


func (x *Imp) String () string {
//
  if x.r == x.invalid {
    return str.Clr (x.wd)
  }
  s:= real.String (x.r)
  str.Move (&s, true)
  str.RemSpaces (&s)
  str.Norm (&s, x.wd)
  str.Move (&s, false)
  return s
}


func (x *Imp) SetColours (f, b col.Colour) {
//
  x.cF, x.cB = f, b
}


func (x *Imp) Write (l, c uint) {
//
  bx.Wd (x.wd)
  bx.Colours (x.cF, x.cB)
  bx.Write (x.String(), l, c)
}


func (x *Imp) Edit (l, c uint) {
//
  x.Write (l, c)
  s:= x.String()
  for {
    bx.Edit (&s, l, c)
    if x.Defined (s) {
      break
    } else {
      errh.ErrorPos ("Eingabe falsch", 0, l + 1, c)
    }
  }
  x.Write (l, c)
}


func (x *Imp) SetFont (f font.Font) {
//
  x.fo = f
}


func (x *Imp) Print (l, c uint) {
//
  pbx.SetFont (x.fo)
  pbx.Print (x.String(), l, c)
}


func (x *Imp) RealVal () float64 {
//
  return x.r
}


func (x *Imp) SetReal (r float64) bool {
//
  if r < x.invalid {
    x.r = r
    return true
  }
  return false
}


func init () { var _ Real = New (4) }
