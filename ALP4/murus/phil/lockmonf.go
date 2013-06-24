package phil

// (c) Christian Maurer   v. 130526 - license see murus.go

// >>> Fair solution with Monitor due to Dijkstra
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 163

import (
  . "murus/obj"
  "murus/mon"
)
type
  LockMonf struct {
                  mon.Monitor
                  }


func NewLockMonf () *LockMonf {
//
  var m mon.Monitor
  f:= func (a Any, i uint) Any {
        p:= a.(uint)
        if i == lock {
          if stat[left(p)] == dining && stat[right(p)] == satisfied ||
             stat[left(p)] == satisfied && stat[right(p)] == dining {
            changeStatus (p, starving)
          }
          for stat[left(p)] == dining || stat[left(p)] == starving ||
            stat[right(p)] == dining || stat[right(p)] == starving {
            m.Wait (p)
          }
        } else { // unlock
          if stat[left(p)] == hungry || stat[left(p)] == starving {
            m.Signal (left(p))
          }
          if stat[right(p)] == hungry || stat[right(p)] == starving {
            m.Signal (right(p))
          }
        }
        return nil
      }
  m = mon.New (nPhilos, f)
  return &LockMonf { m }
}


func (x *LockMonf) Lock (p uint) {
//
  changeStatus (p, hungry)
  x.F (p, lock)
  changeStatus (p, dining)
}


func (x *LockMonf) Unlock (p uint) {
//
  changeStatus (p, satisfied)
  x.F (p, unlock)
}
