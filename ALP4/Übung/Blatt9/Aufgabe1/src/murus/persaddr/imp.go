package persaddr

// (c) Christian Maurer   v. 130526 - license see murus go

import (
  . "murus/obj"; "murus/col"
  "murus/pers"; "murus/addr"
  "murus/atom"; "murus/mol"
)
type
  Imp struct {
             *mol.Imp
             }


func (x *Imp) imp (Y Object) *mol.Imp {
//
  y, ok:= Y.(*Imp)
  if ! ok { NotCompatiblePanic () }
  return y.Imp
}


func New () *Imp {
//
  x:= new (Imp)
  x.Imp = mol.New ()
  a:= atom.New (pers.New ())
  a.SetFormat (pers.LongTB)
  a.SetColours (col.Yellow, col.Black)
  x.Ins (a, 0, 0)
  a = atom.New (addr.New ())
  a.SetColours (col.LightGreen, col.Black)
  x.Ins (a, 2, 0)
  return x
}


func (x *Imp) Eq (Y Object) bool {
//
  return x.Imp.Eq (x.imp (Y))
}


func (x *Imp) Copy (Y Object) {
//
  x.Imp.Copy (x.imp (Y))
}


func (x *Imp) Clone () Object { // otherwise *Imp <--> *mol.Imp
//
  y:= New ()
  y.Copy (x)
  return y
}


func (x *Imp) Less (Y Object) bool {
//
  return x.Imp.Less (x.imp (Y))
}


func (x *Imp) Index () ObjectFunc {
//
  return func (X Object) Object {
    x, ok:= X.(*Imp)
    if ! ok { TypePanic() }
    return x.Component (0).(*atom.Imp)
  }
}


func (x *Imp) RotOrder () {
//
  pers.RotOrder ()
}


func init () { var _ Indexer = New () }
