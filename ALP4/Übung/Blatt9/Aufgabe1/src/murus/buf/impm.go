package buf

// (c) Christian Maurer   v. 120909 - license see murus.go

import (
  "sync"
  "murus/ker"; . "murus/obj"
)
type
  ImpM struct {
            b *buffer
     notEmpty,
      notFull,
       meProd,
       meCons sync.Mutex
              }


func NewM (a Any, n uint) *ImpM {
//
  if n == 0 { return nil }
  x:= new (ImpM)
  x.b = newBuffer (a, n)
  x.notEmpty.Lock ()
  return x
}


func (x *ImpM) Num () uint {
//
  ker.Stop ("buf mutex", 1) // pointless to be called
  return 0
}


func (x *ImpM) Full () bool {
//
  ker.Stop ("buf mutex", 2) // pointless to be called
  return false
}


func (x *ImpM) Ins (a Any) {
//
  x.notFull.Lock ()
  x.meProd.Lock ()
  x.b.Ins (a)
  x.meProd.Unlock ()
  x.notEmpty.Unlock ()
}


func (x *ImpM) Get () Any {
//
  x.notEmpty.Lock ()
  x.meCons.Lock ()
  defer x.notFull.Unlock ()
  defer x.meCons.Unlock ()
  return x.b.Get ()
}
