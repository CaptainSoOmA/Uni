package bnat

// (c) Christian Maurer   v. 130309 - license see murus.go

import (
  . "murus/obj"
  "murus/col"; "murus/str"; "murus/box"; "murus/errh"; "murus/font"; "murus/pbox"
  "murus/nat"
)
type
  Imp struct {
        start,
            n uint
           wd,
      invalid uint
         zero bool
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



func New (n uint) *Imp {
//
  x:= new (Imp)
  x.start = n
  x.wd = nat.Wd (n)
  x.zero = n % 10 == 0
  if x.wd > 9 { x.wd = 9 }
  switch { case x.wd <= 2:
    x.invalid = 1e2
  case x.wd <= 4:
    x.invalid = 1e4
  default:
    x.invalid = 1e9
  }
  x.n = x.invalid
  x.cF, x.cB = col.ScreenF, col.ScreenB
  return x
}


func (x *Imp) Empty () bool {
//
  return x.n == x.invalid
}


func (x *Imp) Clr () {
//
  x.n = x.invalid
}


func (x *Imp) Copy (Y Object) {
//
  y:= x.imp (Y)
  x.start, x.wd, x.invalid = y.start, y.wd, y.invalid
  x.zero = y.zero
  x.n = y.n
}


func (x *Imp) Clone () Object {
//
  y:= New (x.start)
  y.Copy (x)
  return y
}


func (x *Imp) Eq (Y Object) bool {
//
  return x.n == x.imp (Y).n
}


func (x *Imp) Less (Y Object) bool {
//
  y:= x.imp (Y)
  if x.n == x.invalid || y.n == y.invalid {
    return false
  }
  return x.n < y.n
}


func (x *Imp) Codelen () uint {
//
  switch { case x.wd <= 2:
    return 1
  case x.wd <= 4:
    return 2
  }
  return 4
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, x.Codelen())
  switch { case x.wd <= 2:
    copy (b[:x.Codelen()], Encode (byte(x.n)))
  case x.wd <= 4:
    copy (b[:x.Codelen()], Encode (uint16(x.n)))
  default:
    copy (b[:x.Codelen()], Encode (x.n))
  }
  return b
}


func (x *Imp) Decode (b []byte) {
//
  switch { case x.wd <= 2:
    x.n = uint(Decode (uint8(0), b[:x.Codelen()]).(uint8))
  case x.wd <= 4:
    x.n = uint(Decode (uint16(0), b[:x.Codelen()]).(uint16))
  default:
    x.n = Decode (uint(0), b[:x.Codelen()]).(uint)
  }
}


func (x *Imp) Defined (s string) bool {
//
  var n uint
  if nat.Defined (&n, s) {
    x.n = n
    return true
  }
  return false
}


func (x *Imp) String () string {
//
  n:= uint(x.n)
  if n == x.invalid {
    return str.Clr (x.wd)
  }
  return nat.StringFmt (n, x.wd, x.zero)
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
      errh.Error ("keine Zahl", 0)
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


func (x *Imp) Val () uint {
//
  return x.n
}


func (x *Imp) Set (n uint) bool {
//
  if n < x.invalid {
    x.n = n
    return true
  }
  return false
}


func (x *Imp) Startval () uint {
//
  return x.start
}


// func init () { var _ Natural = New (2) }
