package phone

// (c) Christian Maurer   v. 130115 - license see murus.go

import (
  . "murus/obj"; "murus/str"
  "murus/col"; "murus/box"; "murus/errh"; "murus/nat"
  "murus/font"; "murus/pbox"
)
const
  width = 16
type
  Imp struct {
      prefix uint16
      number uint
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
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}


func New () *Imp {
//
  x:= new (Imp)
  x.cF, x.cB = col.ScreenF, col.ScreenB
  return x
}


func (x *Imp) Empty () bool {
//
  return x.prefix == 0 && x.number == 0
}


func (x *Imp) Clr () {
//
  x.prefix, x.number = 0, 0
}


func (x *Imp) Copy (Y Object) {
//
  y:= x.imp (Y)
  x.prefix, x.number = y.prefix, y.number
}


func (x *Imp) Clone () Object {
//
  y:= New ()
  y.Copy (x)
  return y
}


func (x *Imp) Eq (Y Object) bool {
//
  y:= x.imp (Y)
  return x.prefix == y.prefix && x.number == y.number
}


func (x *Imp) Less (Y Object) bool {
//
  y:= x.imp (Y)
  if x.prefix < y.prefix { return true }
  if x.prefix == y.prefix { return x.number < y.number }
  return false
}


func (x *Imp) Defined (s string) bool {
//
  if str.Empty (s) {
    x.Clr ()
    return true
  }
  i:= uint(0)
  str.Move (&s, true)
  l:= str.ProperLen (s)
  if str.Contains (s, ' ', &i) && s[0] == '0' {
    n:= uint(0)
    if nat.Defined (&n, s[1:i]) { // i <= l
      x.prefix = uint16(n)
      if l == i {
        return false
      }
      s = s[i:l]
    } else {
      return false
    }
  } else {
    x.prefix = 0
  }
  str.Move (&s, true)
  str.RemAllSpaces (&s)
  if s == "" {
    x.number = uint(x.prefix)
    x.prefix = 0
    return true
  }
  if nat.Defined (&x.number, s) {
    return true
  } else {
    x.prefix = 0
    x.number = 0
  }
  return false
}


func (x *Imp) String () string {
//
  s:= ""
  if x.prefix > 0 {
    s = nat.String (uint(x.prefix))
    s = "0" + s
  }
  if x.number > 0 {
    t:= nat.String (x.number)
    n:= len (t)
    switch n { case 4, 5:
      t = t[0:n-2] + " " + t[n-2:]
    case 6, 7:
      t = t[0:n-4] + " " + t[n-4:n-2] + " " + t[n-2:]
    case 8, 9:
      t = t[0:n-5] + " " + t[n-5:n-3] + " " + t[n-3:]
    }
    if x.prefix == 0 {
      s = t
    } else {
      s = s + " " + t
    }
  }
  str.Norm (&s, width)
  return s
}


func (x *Imp) SetColours (f, b col.Colour) {
//
  x.cF, x.cB = f, b
}


func (x *Imp) Write (l, c uint) {
//
  bx.Wd (width)
  bx.Colours (x.cF, x.cB)
  bx.Write (x.String(), l, c)
}


func (x *Imp) Edit (l, c uint) {
//
  bx.Wd (width)
  s:= x.String()
  for {
    bx.Edit (&s, l, c)
    if x.Defined (s) {
      break
    } else {
      errh.Error ("keine Telefonnummer", 0)
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


func (x *Imp) Codelen () uint {
//
  return 2 + // Codelen (x.prefix)
         4   // Codelen (x.number)
}


func (x *Imp) Encode () []byte {
//
  bs:= make ([]byte, x.Codelen())
  copy (bs[0:2], Encode (x.prefix))
  copy (bs[2:6], Encode (x.number))
  return bs
}


func (x *Imp) Decode (bs []byte) {
//
  x.prefix = Decode (x.prefix, bs[0:2]).(uint16)
  x.number = Decode (x.number, bs[2:6]).(uint)
}
