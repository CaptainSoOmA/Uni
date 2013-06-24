package index

// (c) Christian Maurer   v. 130115 - license see murus.go

import (
  . "murus/obj"
  "murus/col"
)
type
  Imp struct {
       empty,
      object Object
         pos uint
             }


func (x *Imp) imp (Y Object) *Imp {
//
  y, ok:= Y.(*Imp)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}


func New (a Object) *Imp {
//
  x:= new (Imp)
  x.empty, x.object = a.Clone (), a.Clone ()
  return x
}


func (x *Imp) Set (a Object, n uint) {
//
  x.object = a.Clone ()
  x.pos = n
}


func (x *Imp) Empty () bool {
//
  return Eq (x.object, x.empty)
}


func (x *Imp) Clr () {
//
  x.object = x.empty.Clone ()
  x.pos = 0
}


func (x *Imp) Copy (Y Object) {
//
  y:= x.imp (Y)
  x.empty = y.empty.Clone ()
  x.Set (y.object, y.pos)
}


func (x *Imp) Clone () Object {
//
  y:= New (x.empty)
  y.Copy (x)
  return y
}


func (x *Imp) Eq (Y Object) bool {
//
  return Eq (x.object, x.imp (Y).object)
}


func (x *Imp) Less (Y Object) bool {
//
  return Less (x.object, x.imp (Y).object)
}


func (x *Imp) Pos () uint {
//
  return x.pos
}


func editor (X Any) Editor {
//
  x, ok:= X.(Editor)
  if ! ok { NotCompatiblePanic() }
  return x
}


func (x *Imp) SetColours (f, b col.Colour) {
//
  editor (x.object).SetColours (f, b)
}


func (x *Imp) Write (l, c uint) {
//
  editor (x.object).Write (l, c)
}


func (x *Imp) Codelen () uint {
//
  return Codelen (x.object) + 4
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, x.Codelen())
  n:= int(Codelen (x.object))
  copy (b[:n], Encode (x.object))
  copy (b[n:n+4], Encode (x.pos))
  return b
}


func (x *Imp) Decode (b []byte) {
//
  n:= Codelen (x.object)
  Decode (x.object, b[:n])
  x.pos = Decode (uint(0), b[n:n+4]).(uint)
}
