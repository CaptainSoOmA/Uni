package stk

// (c) Christian Maurer   v. 130326 - license see murus.go

import (
  . "murus/obj"
  "murus/seq"
)
type
  Imp struct {
             seq.Sequence
             }


func New (a Any) *Imp {
//
  return &Imp { seq.New (a) }
}


func (x *Imp) Push (a Any) {
//
  x.Seek (0)
  x.Ins (a)
}


func (x *Imp) Pop () {
//
  if x.Empty () { return }
  x.Seek (0)
  x.Del()
}


func (x *Imp) Top () Any {
//
  if x.Empty () { return nil }
  x.Seek (0)
  return x.Get()
}


func init () { var _ Stack = New (0) }
