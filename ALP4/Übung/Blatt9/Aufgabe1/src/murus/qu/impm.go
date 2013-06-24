package qu

// (c) Christian Maurer   v. 130125 - license see murus.go

import (
  . "murus/obj"
  "sync"
)
type
  ImpM struct {
              *Imp
     notEmpty,
        mutex sync.Mutex
              }


func NewM (a Any) *ImpM {
//
  x:= new (ImpM)
  x.Imp = New (a)
  x.notEmpty.Lock ()
  return x
}


func (x *ImpM) Num () uint {
//
  x.mutex.Lock ()
  defer x.mutex.Unlock ()
  return x.Imp.Num ()
}


func (x *ImpM) Ins (a Any) {
//
  x.mutex.Lock ()
  x.Imp.Ins (a)
  x.mutex.Unlock ()
  x.notEmpty.Unlock ()
}


func (x *ImpM) Get () Any {
//
  x.notEmpty.Lock ()
  x.mutex.Lock ()
  defer x.mutex.Unlock ()
  return x.Imp.Get ()
}


func (x *ImpM) Del () Any {
//
  x.notEmpty.Lock ()
  x.mutex.Lock ()
  defer x.mutex.Unlock ()
  return x.Imp.Del ()
}
