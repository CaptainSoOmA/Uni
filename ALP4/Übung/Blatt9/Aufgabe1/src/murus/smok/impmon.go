package smok

// (c) Christian Maurer   v. 130424 - license see murus.go

// >>> Solution with a monitor

import (
  . "murus/obj"; "murus/mon"
)
type
  ImpMon struct {
                mon.Monitor
                }


func NewMon () *ImpMon {
//
  var m mon.Monitor
  var avail [3]bool
  f:= func (a Any, k uint) Any {
        if k < 3 {
          if a == nil { // AgentOut
            m.Wait (3)
          } else { // AgentIn
            avail[(k + 1) % 3], avail[(k + 2) % 3] = true, true
            m.Signal (k)
          }
        } else { // k == 3
          if a == nil { // SmokerOut
            m.Signal (3)
          } else { // SmokerIn
            u:= a.(uint)
            u1, u2:= (u + 1) % 3, (u + 2) % 3
            if ! avail[u1] || ! avail[u2] {
              m.Wait (u)
            }
            avail[u1], avail[u2] = false, false
          }
        }
        return a
      }
  m = mon.New (3 + 1, f)
  return &ImpMon { m }
}


func (x *ImpMon) AgentIn (u uint) {
//
  x.F (0, u)
}


func (x *ImpMon) AgentOut () {
//
  x.F (nil, 0)
}


func (x *ImpMon) SmokerIn (u uint) {
//
  x.F (u, 3)
}


func (x *ImpMon) SmokerOut () {
//
  x.F (nil, 3)
}
