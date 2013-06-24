package smok

// (c) Christian Maurer   v. 130424 - license see murus.go

// >>> Solution with a critical section

import (
  . "murus/obj"; "murus/cs"
)
type
  ImpCS struct {
               cs.CriticalSection
               }


func NewCS () *ImpCS {
//
  var avail [3]bool
  smokerOut:= true
  c:= func (k uint) bool {
        if k == 3 { // Agent
          return smokerOut
        }
        return avail[(k + 1) % 3] && avail[(k + 2) % 3] // Smoker
      }
  e:= func (a Any, k uint) {
        if k == 3 { // Agent
          u:= a.(uint)
          avail[(u + 1) % 3], avail[(u + 2) % 3] = true, true
        } else { // Smoker
          smokerOut = false
          avail[(k + 1) % 3], avail[(k + 2) % 3] = false, false
        }
      }
  l:= func (a Any, k uint) {
        if k == 3 { // AgentOut
        } else { // SmokerOut
          smokerOut = true
        }
      }
  return &ImpCS { cs.New (3 + 1, c, e, l) }
}


func (x *ImpCS) AgentIn (u uint) {
//
  x.Enter (3, u)
}


func (x *ImpCS) AgentOut () {
//
  x.Leave (3, nil) // effectless
}


func (x *ImpCS) SmokerIn (u uint) {
//
  x.Enter (u, 0)
}


func (x *ImpCS) SmokerOut () {
//
  x.Leave (0, nil)
}
