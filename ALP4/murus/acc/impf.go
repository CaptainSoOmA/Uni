package acc

// (c) Christian Maurer   v. 130115 - license see murus.go

import (
  . "murus/obj"
  "murus/euro"
  "murus/fmon"
)
type
  ImpF struct {
              *fmon.Imp
              }


func NewF ( s string, p uint) *ImpF {
//
  x:= new (ImpF)
  balance:= euro.New ()
  balance.Set2 (0, 0)
  c:= func (a Any, i uint) bool {
        if i == 1 { // draw
          return a.(*euro.Imp).Eq (balance) || a.(*euro.Imp).Less (balance)
        }
        return true // deposit
      }
  f:= func (a Any, i uint) Any {
        balance.Write (0, 0)
        if i == 0 { // deposit
          balance.Plus (a.(*euro.Imp))
        } else { // draw
          balance.Minus (a.(*euro.Imp))
        }
        balance.Write (1, 0)
        return balance
      }
  x.Imp = fmon.New (balance, 2, f, c, s, p)
  return x
}


func (x *ImpF) Deposit (e *euro.Imp) *euro.Imp {
//
  return x.Imp.F (e, 0).(*euro.Imp)
}


func (x *ImpF) Draw (e *euro.Imp) *euro.Imp {
//
  return x.Imp.F (e, 1).(*euro.Imp)
}
