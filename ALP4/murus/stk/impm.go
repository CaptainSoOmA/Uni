package stk

// (c) Christian Maurer   v. 120909 - license see murus.go

import (
  "sync"
  . "murus/obj"
)
type
  ImpM struct {
            s *Imp
     notEmpty,
           mE sync.Mutex
              }


func NewM (n uint) *ImpM {
//
  if n == 0 { return nil }
  x:= new (ImpM)
  x.s = New (n)
  x.notEmpty.Lock ()
  return x
}


func (x *ImpM) Empty () bool {
//
  x.mE.Lock ()
  defer x.mE.Unlock ()
  return x.s.Empty ()
}


func (x *ImpM) Push (a Any) {
//
  x.mE.Lock ()
  x.s.Push (a)
  x.mE.Unlock ()
  x.notEmpty.Unlock ()
}


func (x *ImpM) Pop () {
//
  x.notEmpty.Lock ()
  x.mE.Lock ()
  x.s.Pop ()
  x.mE.Unlock ()
}


func (x *ImpM) Top () Any {
//
  x.notEmpty.Lock ()
  x.mE.Lock ()
  defer x.mE.Unlock ()
  return x.s.Top ()
}
