package pt

// (c) Christian Maurer   v. 121215 - license see murus.go

import (
  . "murus/obj"
  "murus/col"
  "murus/vect"
)
const
  pack = "pt"
type
  Imp struct {
       class Class
      number uint
      colour col.Colour
  vect, norm *vect.Imp
             }


func New () *Imp {
//
  x:= new (Imp)
  x.class = None
  x.colour = col.ScreenF
  x.vect, x.norm = vect.New (), vect.New ()
  return x
}


func (x *Imp) Eq (X Object) bool {
//
  y, ok:= X.(*Imp)
  if ! ok { return false }
  return x.class == y.class && // x.number == y.number && // ?
         x.vect.Eq (y.vect) && x.norm.Eq (y.norm)
}


func (x *Imp) Less (Y Object) bool {
//
  return false
}


func (x *Imp) Copy (X Object) {
//
  y, ok:= X.(*Imp)
  if ! ok { panic ("point.Copy: ach du grüne Neune"); return }
  x.class = y.class
  x.number = y.number
  x.colour = y.colour
  x.vect.Copy (y.vect)
  x.norm.Copy (y.norm)
}


func (x *Imp) Clone () Object {
//
  y:= New ()
  y.Copy (x)
  return y
}


func (x *Imp) Empty () bool {
//
  return x.class == None
}


func (x *Imp) Clr () {
//
  x.class = None
}


func (x *Imp) Set (c Class, a uint, f col.Colour, v, n *vect.Imp) {
//
  x.class = c
  x.number = a
  x.colour = f
  x.vect.Copy (v)
  x.norm.Copy (n)
}


func (x *Imp) ClassOf () Class {
//
  return x.class
}


func (x *Imp) Number () uint {
//
  return x.number
}


func (x *Imp) Colour () col.Colour {
//
  return x.colour
}


func (x *Imp) Write (i uint) {
//
  x.vect.SetColours (x.colour, col.ScreenB); x.vect.Write (i,  0)
  x.norm.SetColours (x.colour, col.ScreenB); x.norm.Write (i, 20)
}


func (x *Imp) Read () *vect.Imp {
//
  return x.vect.Clone ().(*vect.Imp)
}


func (x *Imp) Read2 () (*vect.Imp, *vect.Imp) {
//
  return x.vect.Clone ().(*vect.Imp), x.norm.Clone ().(*vect.Imp)
}


var
  cluint = Codelen (uint(0))


func (x *Imp) Codelen () uint {
//
  return 1 + cluint + col.Codelen () + 2 * x.vect.Codelen ()
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, x.Codelen())
  b[0] = byte(x.class)
  i, a:= uint(1), cluint
  copy (b[i:i+a], Encode (x.number))
  i += a
  a = col.Codelen()
  copy (b[i:i+a], Encode (x.colour))
  i += a
  a = x.vect.Codelen()
  copy (b[i:i+a], x.vect.Encode ())
  i += a
  copy (b[i:i+a], x.norm.Encode ())
  return b
}


func (x *Imp) Decode (b []byte) {
//
  x.class = Class(b[0])
  i, a:= uint(1), cluint
  x.number = Decode (uint(0), b[i:i+a]).(uint)
  i += a
  a = col.Codelen()
  col.Decode (&x.colour, b[i:i+a])
  i += a
  a = x.vect.Codelen()
  x.vect.Decode (b[i:i+a])
  i += a
  x.norm.Decode (b[i:i+a])
}
