package buf

// (c) Christian Maurer   v. 120909 - license see murus.go

// >>> Implementation with asynchronous message passing

import
  . "murus/obj"
type
  ImpCh struct {
             c chan Any
               }


func NewCh (a Any, n uint) *ImpCh {
//
  if a == nil || n == 0 { return nil }
  return &ImpCh { make (chan Any, n) }
}


func (x *ImpCh) Num () uint {
//
  return 0 // senseless
}


func (x *ImpCh) Full () bool {
//
  return false // senseless
}


func (x *ImpCh) Ins (a Any) {
//
  x.c <- a
}


func (x *ImpCh) Get () Any {
//
  return Clone (<-x.c)
}
