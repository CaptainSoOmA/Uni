package acc

// (c) Christian Maurer   v. 130115 - license see murus.go

import (
  "sync"
  "murus/euro"
)
type
  Imp struct {
             *euro.Imp "balance"
          mE sync.Mutex
         pos *sync.Cond
             }


func New() *Imp {
//
  x:= new (Imp)
  x.Imp = euro.New ()
  x.Imp.Set2 (0, 0)
  x.pos = sync.NewCond (&x.mE)
  return x
}


func (x *Imp) Deposit (e *euro.Imp) *euro.Imp {
//
  x.mE.Lock()
  x.Imp.Plus (e)
  x.pos.Signal()
  x.mE.Unlock()
  return x.Imp.Clone().(*euro.Imp)
}


func (x *Imp) Draw (e *euro.Imp) *euro.Imp {
//
  x.mE.Lock()
  for x.Imp.Less (e) {
    x.pos.Wait()
  }
  x.Imp.Minus (e)
  x.pos.Signal()
  x.mE.Unlock()
  return x.Imp.Clone().(*euro.Imp)
}
