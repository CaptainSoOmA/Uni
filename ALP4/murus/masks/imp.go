package masks

// (c) Christian Maurer   v. 130115 - license see murus.go

import (
  . "murus/obj"; "murus/str"
  "murus/col"; "murus/scr"
  "murus/font"; "murus/pbox"
//  "murus/errh"
)
type
  Imp struct {
        mask []string
        l, c []uint
          l0,
         num uint
      cF, cB col.Colour
//         fnt font.Font
             }
var
  pbx *pbox.Imp = pbox.New ()


func (x *Imp) imp (Y Object) *Imp {
//
  y, ok:= Y.(*Imp)
  if ! ok { TypeNotEqPanic (x, Y) }
  if y == nil { DivBy0Panic() }
//  if len (x.mask) != len (y.mask) { TypeNotEqPanic (x, Y) }
  return y
}


func New () *Imp {
//
  x:= new (Imp)
  x.cF, x.cB = col.ScreenF, col.ScreenB
  return x
}


func (x *Imp) Line (n uint) {
//
  x.l0 = n
}


func (x *Imp) Ins (s string, l, c uint) {
//
  x.mask = append (x.mask, s)
  x.l, x.c = append (x.l, l), append (x.c, c)
  x.num ++
}


func (x *Imp) Del (n uint) {
//
  if n >= x.num { return }
  for i:= uint(n); i + 1 < x.num; i++ {
    x.mask[i] = x.mask[i + 1]
    x.l[i], x.c[i] = x.l[i + 1], x.c[i + 1]
  }
  x.num --
}


func (x *Imp) Empty () bool {
//
  return x.num == 0
}


func (x *Imp) Clr () {
//
  x.mask = nil
  x.l, x.c = nil, nil
  x.num = 0
}


func (x *Imp) Eq (Y Object) bool {
//
  y:= x.imp (Y)
  for i:= uint(0); i < x.num; i++ {
    if x.mask[i] != y.mask[i] || x.l[i] != y.l[i] || x.c[i] != y.c[i] {
      return false
    }
  }
  return false
}


func (x *Imp) Less (Y Object) bool {
//
  return false
}


func (x *Imp) Copy (Y Object) {
//
  y:= x.imp (Y)
  x.num = y.num
  x.mask = make ([]string, x.num)
  for i:= uint(0); i < x.num; i++ {
    x.mask[i] = y.mask[i]
    x.l[i], x.c[i] = y.l[i], y.c[i]
  }
  x.cF, x.cB = y.cF, y.cB
}


func (x *Imp) Clone () Object {
//
  y:= new (Imp)
  y.Copy (x)
  return y
}


func (x *Imp) SetColours (f, b col.Colour) {
//
  x.cF, x.cB = f, b
}


func (x *Imp) Write (l, c uint) {
//
  if x.num == 0 { return }
  scr.Colours (x.cF, x.cB)
  l += x.l0
  for i:= uint(0); i < x.num; i++ {
    scr.Write (x.mask[i], l + x.l[i], c + x.c[i])
  }
}


func (x *Imp) Edit (l, c uint) {
//
// TODO
}


func (x *Imp) SetFont (f font.Font) {
//
//  x.font = f
}


func (x *Imp) Print (l, c uint) {
//
  if x.num == 0 { return }
//  pbx.SetFont (x.font)
  for i:= uint(0); i < x.num; i++ {
    pbx.Print (x.mask[i], x.l[i], x.c[i])
  }
}


func (x *Imp) Codelen () uint { // TODO cF, cB
//
  c:= uint(4)
  for k:= uint(0); k < x.num; k++ {
    c += 4
    c += uint(len (x.mask[k]))
    c += 4
    c += 2 * col.Codelen()
  }
  return c
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, x.Codelen())
  i, a:= uint(0), uint(4)
  copy (b[i:i+a], Encode (x.num))
  i += a
  for k:= uint(0); k < x.num; k++ {
    a = uint(4)
    n:= uint(len (x.mask[k]))
    copy (b[i:i+a], Encode (n))
    i += a
    a = uint(n)
    copy (b[i:i+a], []byte(x.mask[k]))
    i += a
    a = uint(4)
    n = x.l[k] + 256 * x.c[k]
    copy (b[i:i+a], Encode (n))
    i += a
    a = col.Codelen()
    copy (b[i:i+a], Encode (x.cF))
    i += a
    copy (b[i:i+a], Encode (x.cB))
    i += a
  }
  return b
}


func (x *Imp) Decode (b []byte) { // TODO cF, cB
//
  i, a:= uint(0), uint(4)
  x.num = Decode (uint(0), b[i:i+a]).(uint)
  i += a
  x.mask = make ([]string, x.num)
  for k:= uint(0); k < x.num; k++ {
    a = uint(4)
    n:= Decode (uint(0), b[i:i+a]).(uint)
    i += a
    a = n
    x.mask[i] = Decode (str.Clr (n), b[i:i+a]).(string)
    i += a
    a = uint(4)
    n = Decode (uint(0), b[i:i+a]).(uint)
    x.l[i], x.c[i] = n % 256, n / 256
    i += a
    a = col.Codelen()
    x.cF = Decode (col.Black, b[i:i+a]).(col.Colour)
    i += a
    x.cB = Decode (col.Black, b[i:i+a]).(col.Colour)
    i += a
  }
}


func init () { var _ MaskSequence = New() }
