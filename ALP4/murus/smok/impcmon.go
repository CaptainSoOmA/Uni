package smok

// (c) Christian Maurer   v. 130424 - license see murus.go

// >>> Solution with a conditioned monitor

import (
  . "murus/obj"; "murus/mon"
)
const (
  agent = iota; smoker
)
type
  ImpCMon struct {
                 mon.Monitor
                 }


func NewCMon () *ImpCMon {
//
  var m mon.Monitor
  var avail [3]bool
  smokerOut:= true
  c:= func (a Any, k uint) bool {
        if a == nil { // AgentOut, SmokerOut
          return true
        }
        if k == agent { // AgentIn
          return smokerOut
        }
        u:= a.(uint) // SmokerIn
        return avail[(u + 1) % 3] && avail[(u + 2) % 3]
      }
  f:= func (a Any, k uint) Any {
        if k == agent {
          if a == nil { // AgentOut
            // nixtun
          } else { // AgentIn
            u:= a.(uint)
            avail[(u + 1) % 3], avail[(u + 2) % 3] = true, true
          }
        } else {
          if a == nil { // SmokerOut
            smokerOut = true
          } else { // SmokerIn
            u:= a.(uint)
            avail[(u + 1) % 3], avail[(u + 2) % 3] = false, false
            smokerOut = false
          }
        }
        return a
      }
  m = mon.NewC (2, f, c)
  return &ImpCMon { m }
}


func (x *ImpCMon) AgentIn (u uint) {
//
  x.F (u, agent)
}


func (x *ImpCMon) AgentOut () {
//
  x.F (nil, agent)
}


func (x *ImpCMon) SmokerIn (u uint) {
//
  x.F (u, smoker)
}


func (x *ImpCMon) SmokerOut () {
//
  x.F (nil, smoker)
}
