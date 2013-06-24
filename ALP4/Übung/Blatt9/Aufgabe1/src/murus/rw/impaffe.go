package rw

// (c) Christian Maurer   v. 130525 - license see murus.go

//     Affenversuch

import
  . "murus/obj"
type
  ImpAffe struct {
          cI, cO []chan Any
                 }


func NewAffe (n uint, p /* Pred */ CondSpectrum, e, l /* Op */ StmtSpectrum) *ImpAffe {
//
  x:= new (ImpAffe)
  x.cI, x.cO = make ([]chan Any, n), make ([]chan Any, n)
  for i:= uint(0); i < n; i++ {
    x.cI[i], x.cO[i] = make (chan Any), make (chan Any)
  }
  go func () {
    for {
      for i:= uint(0); i < n; i++ {
        select {
        case <-When (p (/* a, */ i), x.cI[i]):
          _ = x.cI[i] // b:= x.cI[i]
          e (/* b, */ i)
        case <-x.cO[i]:
          l (/* a, */ i)
        default:
        }
      }
    }
  }()
  return x
}


func (x *ImpAffe) In (a Any, k uint) {
//
  x.cI[k] <- a
}


func (x *ImpAffe) Out (a Any, k uint) {
//
  x.cO[k] <- a
}
