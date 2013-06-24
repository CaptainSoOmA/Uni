package puls

// (c) Christian Maurer   v. 121102 - license see murus.go

import (
  . "murus/obj"
  "murus/scr"
  "murus/fmon"
)
type
  Imp struct {
        host []string
        port []uint
           h uint // len (host)
          nb []uint
           n uint // len (nb)
         mon []*fmon.Imp
           f FuncSpectrum
      object Any
             }


func New (host []string, port []uint, nb[]uint, f FuncSpectrum, a Any) *Imp {
//
  x:= new (Imp)
  x.host, x.port, x.h = host, port, uint(len (host))
  x.nb, x.n = nb, uint(len (nb))
  x.mon = make ([]*fmon.Imp, x.n)
  x.f = f
  x.object = Clone (a)
  for i:= uint(0); i < x.n; i++ {
    x.mon[i] = fmon.New0 (x.object, 1, NilSp, TrueSp, x.host[x.nb[i]], x.port[x.nb[i]])
  }
  return x
}


func (x *Imp) Go (me uint) {
//
  ego:= fmon.New0 (x.object, 1, x.f, TrueSp, x.host[me], x.port[me])
  go ego.Go ()
//  go fmon.New (x.object, 1, x.f, TrueSp, x.host[me], x.port[me]) // TODO
  for i:= uint(0); i < x.n; i++ {
    x.mon[i].Go ()
  }
}


func (x *Imp) F (i uint) Any {
//
  for k:= uint(0); k < x.n; k++ {
    x.object = x.f (x.mon[k].F (x.object, i), i)
    scr.WriteNat (k, 40, 0)
  }
  return x.object
}


func (x *Imp) Terminate () {
//
  for i:= uint(0); i < x.n; i++ {
    x.mon[i].Terminate ()
  }
}


func init () { var _ Pulse = New (nil, nil, nil, NilSp, nil) }
