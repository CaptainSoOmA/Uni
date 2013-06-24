package phil

// (c) Christian Maurer   v. 130526 - license see murus.go

// >>> Unfair monitor solution due to Dijkstra
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 97

import (
  . "murus/obj"
  "murus/mon"
)
type
  LockMonuf struct {
                   mon.Monitor
                   }


func NewLockMonuf () *LockMonuf {
//
  var m mon.Monitor
  f:= func (a Any, k uint) Any {
        p:= a.(uint)
        if k == lock {
          changeStatus (p, starving)
          for stat[left(p)] == dining || stat[right(p)] == dining {
            m.Wait (p)
          }
        } else { // k == unlock
          m.Signal (left(p))
          m.Signal (right(p))
        }
        return nil
      }
  m = mon.New (nPhilos, f)
  return &LockMonuf { m }
}


func (x *LockMonuf) Lock (p uint) {
//
  changeStatus (p, hungry)
  x.F (p, lock)
  changeStatus (p, dining)
}


func (x *LockMonuf) Unlock (p uint) {
//
  changeStatus (p, satisfied)
  x.F (p, unlock)
}
