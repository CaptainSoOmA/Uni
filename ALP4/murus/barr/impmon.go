package barr

// (c) Christian Maurer   v. 121030 - license see murus.go

// >>> Implementation with a monitor
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 162

import (
  . "murus/obj"
  "murus/mon"
)
type
  ImpMon struct {
                mon.Monitor
                }


func NewMon (n uint) *ImpMon {
//
  if n < 2 { return nil }
  var m mon.Monitor
  involved:= n
  waiting:= uint(0)
  f:= func (a Any, k uint) Any {
        waiting++
        if waiting < involved {
          m.Wait (0)
        } else {
/*
          m.SignalAll (0) // simple solution with broadcast
          waiting = 0
*/
          for waiting > 0 { // standard solution
            waiting --
            if waiting > 0 {
              m.Signal (0)
            }
          }
        }
        return nil
      }
  m = mon.New (1, f)
  return &ImpMon { m }
}


func (x *ImpMon) Wait () {
//
  _ = x.F (nil, 0)
}
