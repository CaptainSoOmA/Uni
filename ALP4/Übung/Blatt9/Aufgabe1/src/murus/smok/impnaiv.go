package smok

// (c) Christian Maurer   v. 130424 - license see murus.go

// >>> Naive solution with deadlock danger

import
  "sync"
type
  ImpNaiv struct {
       smokerOut sync.Mutex
           avail [3]sync.Mutex
                 }


func NewNaiv () *ImpNaiv {
//
  x:= new (ImpNaiv)
  x.smokerOut.Lock ()
  for u:= uint(0); u < 3; u++ { x.avail[u].Lock () }
  return x
}


func (x *ImpNaiv) AgentIn (u uint) {
//
  x.smokerOut.Lock()
  x.avail[(u + 1) % 3].Unlock ()
  x.avail[(u + 2) % 3].Unlock ()
}


func (x *ImpNaiv) AgentOut () {
//
}


func (x *ImpNaiv) SmokerIn (u uint) {
//
  x.avail[(u + 1) % 3].Lock ()
  x.avail[(u + 2) % 3].Lock ()
}


func (x *ImpNaiv) SmokerOut () {
//
  x.smokerOut.Unlock ()
}
