package sem

// (c) Christian Maurer   v. 120909 - license see murus.go

import
  . "murus/obj"
type
  ImpGSel struct {
            p, v chan Any
                 }


func NewGSel (n uint) *ImpGSel {
//
  x:= new (ImpGSel)
  x.p, x.v = make (chan Any), make (chan Any)
  go func () {
    val:= n
    for {
      select {
      case <-When (val > 0, x.p):
        val--
      case <-x.v:
        val++
      }
    }
  }()
  return x
}


func (x *ImpGSel) P() {
//
  x.p <- 0
}


func (x *ImpGSel) V() {
//
  x.v <- 0
}
