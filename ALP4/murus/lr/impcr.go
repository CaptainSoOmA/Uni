package lr

// (c) Christian Maurer   v. 130409 - license see murus.go

// >>> left/right problem: implementation with critical resources

import (
  "murus/cr"
)
type
  ImpCR struct {
               cr.CriticalResource
               }


func NewCR () *ImpCR {
//
  const nc = 2
  x:= &ImpCR { cr.New (nc, 1) }
  m:= make ([][]uint, nc)
  for i:= uint(0); i < nc; i++ { m[i] = make ([]uint, 1) }
  m[0][0], m[1][0] = 5, 3
  x.Limit (m)
  return x
}


func (x *ImpCR) LeftIn () {
//
  _ = x.Enter (left)
}


func (x *ImpCR) LeftOut () {
//
  x.Leave (left)
}


func (x *ImpCR) RightIn () {
//
  _ = x.Enter (right)
}


func (x *ImpCR) RightOut () {
//
  x.Leave (right)
}
