package buf

// (c) Christian Maurer   v. 120909 - license see murus.go

// >>> Implementation with synchronous message passing and guarded selective waiting

import (
  "murus/ker"; . "murus/obj"
)
type
  ImpGSel struct {
          cI, cG chan Any
                 }


func NewGSel (a Any, n uint) *ImpGSel {
//
  if n == 0 { return nil }
  x:= new (ImpGSel)
  x.cI, x.cG = make (chan Any), make (chan Any)
  go func () {
    buffer:= make ([]Any, n)
    var in, out, num uint
    for {
      select {
      case buffer [in] = <-When (num < n, x.cI):
        in = (in + 1) % n
        num ++
      case When (num > 0, x.cG) <- buffer [out]:
        out = (out + 1) % n
        num --
      }
    }
  }()
  return x
}


func (x *ImpGSel) Num () uint {
//
  ker.Stop ("buf gsel", 1) // pointless to call
  return 0
}


func (x *ImpGSel) Full () bool {
//
  ker.Stop ("buf gsel", 2) // pointless to call
  return false
}


func (x *ImpGSel) Ins (a Any) {
//
  x.cI <- a
}


func (x *ImpGSel) Get () Any {
//
  return Clone (<-x.cG)
}
