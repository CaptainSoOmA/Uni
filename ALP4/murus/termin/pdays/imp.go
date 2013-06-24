package pdays

// (c) Christian Maurer   v. 130115 - license see murus.go

import (
  . "murus/obj"; "murus/str"
  "murus/set"; "murus/pseq"; "murus/day"
)
const (
  pack = "pdayset"
  suffix = "day"
)
type
  Imp struct {
      dayset *set.Imp
        file *pseq.Imp
        name string
     changed bool
             }


func imp (X Object) *day.Imp {
//
  x, ok:= X.(*day.Imp)
  if ! ok { TypePanic() }
  return x
}


func New () *Imp {
//
  x:= new (Imp)
  d:= day.New ()
  x.file = pseq.New (d)
  x.dayset = set.New (d)
  return x
}


func (x *Imp) Name (s string) {
//
  x.name = s
  str.RemSpaces (&x.name)
  x.file.Name (x.name + "." + suffix)
  x.dayset.Clr ()
  x.file.Trav (func (a Any) { x.dayset.Ins (a.(*day.Imp)) })
}


func (x *Imp) Rename (N string) {
//
  x.name = N
  x.file.Rename (x.name + "." + suffix)
}


func (x *Imp) Empty () bool {
//
  return x.dayset.Empty ()
}


func (x *Imp) Clr () {
//
  x.dayset.Clr ()
}


func (x *Imp) Num () uint {
//
  return x.dayset.Num ()
}


func (x *Imp) Ex (Y Object) bool {
//
  return x.dayset.Ex (imp (Y))
}


func (x *Imp) Ins (Y Object) {
//
  d:= imp (Y)
  if ! x.dayset.Ex (d) {
    x.dayset.Ins (d)
    x.changed = true
  }
}


func (x *Imp) Del (Y Object) {
//
  d:= imp (Y)
  x.changed = x.dayset.Ex (d)
  if x.changed {
    x.dayset.Del ()
  }
}


func (x *Imp) Terminate () {
//
  if x.changed {
    x.file.Clr ()
    x.dayset.Trav (func (a Any) { x.file.Ins (a.(*day.Imp)) })
    x.file.Terminate ()
  }
}
