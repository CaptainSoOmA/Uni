package pair

// (c) Christian Maurer   v. 130115 - license see murus.go

import (
  . "murus/obj"
  "murus/str"
)
const
  pack = "files/pair"
type
  Imp struct {
        name string
         typ byte
             }


func New () *Imp {
//
  return new (Imp)
}


func (x *Imp) Eq (Y Object) bool {
//
  y, ok:= Y.(*Imp)
  if ! ok { NotCompatiblePanic() }
  return x.name == y.name && x.typ == y.typ
}


func (x *Imp) Less (Y Object) bool {
//
  return false
}


func (x *Imp) Copy (Y Object) {
//
  y, ok:= Y.(*Imp)
  if ! ok { NotCompatiblePanic() }
  x.name, x.typ = y.name, y.typ
}


func (x *Imp) Clone () Object {
//
  y:= New()
  y.Copy (x)
  return y
}


func (x *Imp) Empty () bool {
//
  return str.Empty (x.name)
}


func (x *Imp) Clr () {
//
  x.name = ""
  x.typ = 0
}


func (x *Imp) Codelen () uint {
//
  return uint(len (x.name)) + 1
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, x.Codelen())
  n:= uint(len (x.name))
  copy (b[0:n], x.name)
  b[n] = x.typ
  return b
}


func (x *Imp) Decode (b []byte) {
//
  n:= uint(len (b))
  x.name = string(b[0:n])
  x.typ = b[n]
}


func (x *Imp) Set (s string, b byte)  {
//
  x.name, x.typ = s, b
}


func (x *Imp) Name () string {
//
  return x.name
}


func (x *Imp) Typ () byte {
//
  return x.typ
}


func init () { var _ Pair = New () }
