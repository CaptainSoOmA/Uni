package qu

// (c) Christian Maurer   v. 130228 - license see murus.go

import (
  . "murus/obj"; "murus/seq"
)
type
  Imp struct {
             seq.Sequence
             }


func New (a Any) *Imp {
//
  return &Imp { seq.New (a) }
}


func (x *Imp) Ins (a Any) {
//
  x.Seek (x.Num ())
  x.Sequence.Ins (a)
}


func (x *Imp) Get () Any {
//
  if x.Empty () { return nil }
  x.Seek (0)
  return x.Sequence.Get ()
}


func (x *Imp) Del () Any {
//
  if x.Empty () {
    return nil
  }
  x.Seek (0)
  defer x.Sequence.Del ()
  return x.Get ()
}


func init () { var _ Queue = New (0) }
