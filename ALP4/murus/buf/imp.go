package buf

// (c) Christian Maurer   v. 120909 - license see murus.go

import (
  "sync"
  . "murus/obj"
)
type
  Imp struct {
           b *buffer
           m sync.Mutex
             }


func New (a Any, n uint) *Imp {
//
  if a == nil || n == 0 { return nil } // TODO Panic
  x:= new (Imp)
  x.b = newBuffer (a, n)
  return x
}


func (x *Imp) Num () uint {
//
  return x.b.Num()
}


func (x *Imp) Empty () bool {
//
  return x.b.Empty ()
}


func (x *Imp) Full () bool {
//
  return x.b.Full ()
}


func (x *Imp) Ins (a Any) {
//
  x.m.Lock()
  x.b.Ins (a)
  x.m.Unlock()
}


func (x *Imp) Get () Any {
//
  x.m.Lock()
  defer x.m.Unlock()
  return x.b.Get ()
}
