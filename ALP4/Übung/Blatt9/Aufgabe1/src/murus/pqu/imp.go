package pqu

// (c) Christian Maurer   v. 120909 - license see murus.go

import (
  "murus/ker"; . "murus/obj"
  "murus/pqu/heap"
)
const
  pack = "pqu"
type
  Imp struct {
       object Any
       anchor *heap.Imp
          num uint
              }


func New (a Any) *Imp {
//
  x:= new (Imp)
  x.object = Clone (a)
  x.anchor = heap.New()
  return x
}


func (x *Imp) Num () uint {
//
  return x.num
}


func (x *Imp) Ins (a Any) {
//
  if ! TypeEq (a, x.object) { ker.Stop (pack, 1) }
  x.num ++
  x.anchor = x.anchor.Ins (a, x.num).(*heap.Imp)
  x.anchor.Lift (x.num)
}


func (x *Imp) Get () Any {
//
  if x.num == 0 {
    return nil
  }
  return x.anchor.Get ()
}


func (x *Imp) Del () Any {
//
  if x.num == 0 {
    return nil
  }
  if x.num == 1 {
    a:= x.anchor.Get ()
    x.anchor = heap.New()
    x.num = 0
    return a
  }
  y, a:= x.anchor.Del (x.num)
  x.anchor = y.(*heap.Imp)
  x.num --
  if x.num > 0 {
    x.anchor.Sift (x.num)
  }
  return a
}
