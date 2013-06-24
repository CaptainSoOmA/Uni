package audio

// (c) Christian Maurer   v. 130526 - license see murus.go

import (
  . "murus/obj"; "murus/col"; "murus/enum"; "murus/text";
  "murus/masks"; "murus/atom"; "murus/mol"
)
const (
  lenWerk = 25
  lenOrch = 25
  lenName = 25
)
type
  Imp struct {
             *mol.Imp
             }


func (x *Imp) imp (Y Object) *mol.Imp {
//
  y, ok:= Y.(*Imp)
  if ! ok { NotCompatiblePanic() }
  return y.Imp
}


func New () *Imp {
//
  x:= new (Imp)
  x.Imp = mol.New ()
  a:= atom.New (enum.New (enum.Composer))
  a.SetColours (col.Yellow, col.Black)
  x.Ins (a, 0, 11)
  a = atom.New (text.New (lenWerk))
  a.SetColours (col.LightRed, col.Black)
  x.Ins (a, 1, 11)
  a = atom.New (text.New (lenOrch)) // Orchester
  x.Ins (a, 2, 11)
  a = atom.New (text.New (lenName)) // Dirigent
  a.SetColours (col.Cyan, col.Black)
  x.Ins (a, 3, 11)
  a = atom.New (text.New (lenName)) // Solist
  a.SetColours (col.LightBlue, col.Black)
  x.Ins (a, 4, 11)
  a = atom.New (text.New (lenName)) // Solist
  a.SetColours (col.LightBlue, col.Black)
  x.Ins (a, 5, 11)
  a = atom.New (enum.New (enum.RecordLabel))
  a.SetColours (col.LightCyan, col.Black)
  x.Ins (a, 6, 11)
  a = atom.New (enum.New (enum.AudioMedium))
  x.Ins (a, 7, 11)
  a = atom.New (enum.New (enum.SparsCode))
  x.Ins (a, 8, 11)

  var m *masks.Imp = masks.New ()
  m.Ins ("Komponist:", 0, 0)
  m.Ins ("     Werk:", 1, 0)
  m.Ins ("Orchester:", 2, 0)
  m.Ins (" Dirigent:", 3, 0)
  m.Ins (" Solist 1:", 4, 0)
  m.Ins (" Solist 2:", 5, 0)
  m.Ins ("    Firma:", 6, 0)
  m.Ins ("   Platte:", 7, 0)
  m.Ins ("       ad:", 8, 0)
  x.SetMask (m)

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


func (x *Imp) Clone () Object { // otherwise *Imp <--> *mol.Imp:
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
//  ___.RotOrder ()
}


func init () { var _ Indexer = New () }
