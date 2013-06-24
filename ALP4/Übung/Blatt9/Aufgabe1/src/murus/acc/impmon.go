package acc

// (c) Christian Maurer   v. 130115 - license see murus.go

import (
  . "murus/obj"
  "murus/euro"
  "murus/mon"
)
type
  ImpMon struct {
        balance *euro.Imp
                *mon.Imp
                }


func NewMon () *ImpMon {
//
  x:= new (ImpMon)
  x.balance = euro.New ()
  x.balance.Set2 (0, 0)
  p:= func (a Any, i uint) bool {
        if i == 1 { // draw
          return a.(*euro.Imp).Eq (x.balance) || a.(*euro.Imp).Less (x.balance)
        }
        return true // deposit
      }
  f:= func (a Any, i uint) Any {
        if i == 0 { // deposit
          x.balance.Plus (a.(*euro.Imp))
        } else { // draw
          x.balance.Minus (a.(*euro.Imp))
        }
        return x.balance
      }
  x.Imp = mon.NewC (2, f, p)
  return x
}


func (x *ImpMon) bal () *euro.Imp {
//
  return x.balance.Clone ().(*euro.Imp)
}


func (x *ImpMon) Deposit (e *euro.Imp) *euro.Imp {
//
  return x.Imp.F (e, 0).(*euro.Imp)
}


func (x *ImpMon) Draw (e *euro.Imp) *euro.Imp {
//
  return x.Imp.F (e, 1).(*euro.Imp)
}
