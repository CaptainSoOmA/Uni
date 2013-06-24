package gleis

// (c) Christian Maurer   v. 130308 - license see murus.go

import (
  . "murus/obj"; "murus/col"; "murus/scr"
  "murus/sub/linie"
)
type
  Imp struct {
       Linie linie.Linie
//    mehrfach bool
         val uint
             }


func New () *Imp {
//
  return new (Imp)
}


func (x *Imp) Def (l linie.Linie, n uint) {
//
  x.Linie, x.val = l, n
}


func (x *Imp) Empty () bool {
//
  return x.val > 0
}


func (x *Imp) Clr () {
//
  x.val = 0
}


func (x *Imp) Eq (Y Object) bool {
//
  y, ok:= Y.(*Imp)
  if ! ok { return false }
  return x.Linie == y.Linie &&
//         x.mehrfach == y.mehrfach &&
         x.val == y.val
}


func (x *Imp) Less (Y Object) bool {
//
  return false
}


func (x *Imp) Copy (Y Object) {
//
  y, ok:= Y.(*Imp)
  if ! ok { return }
  x.Linie = y.Linie
//  x.mehrfach = y.mehrfach
  x.val = y.val
}


func (x *Imp) Clone () Object {
//
  y:= New ()
  y.Copy (x)
  return y
}


func (x *Imp) Val () uint {
//
  return x.val
}


func (x *Imp) Set (n uint) bool {
//
  if n < 10 {
    x.val = n
    return true
  }
  return false
}


func (x *Imp) Remove () {
//
  scr.Colour (col.ScreenB)
  scr.SetLinewidth (scr.Yetthicker)
}


func (x *Imp) Write (aktuell bool) {
//
  if aktuell || x.Linie == linie.Fußweg {
    scr.Colour (linie.Farbe [x.Linie])
  } else {
    scr.Colour (col.Black)
  }
  if aktuell {
    scr.SetLinewidth (scr.Yetthicker)
  } else {
    scr.SetLinewidth (scr.Thin)
  }
}


func (x *Imp) Codelen () uint {
//
  return 1 +
//         1 +
         4
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, x.Codelen())
//  b[0] = byte(x.Linie)
//  if x.mehrfach { b[1] = 1 }
  copy (b[1:1+4], Encode (x.val))
  return b
}


func (x *Imp) Decode (b []byte) {
//
  x.Linie = linie.Linie (b[0])
//  x.mehrfach = b[1] == 1
  x.val = Decode (uint(0), b[1:1+4]).(uint)
}


func init () {
//
  var x Gleis = New (); if x == nil {}
}
