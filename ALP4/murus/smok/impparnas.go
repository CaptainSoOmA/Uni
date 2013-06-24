package smok

// (c) Christian Maurer   v. 130424 - license see murus.go

// >>> Solution with helper processes due to D. L. Parnas:
//     On a Solution to the Cigarette Smoker's Problem
//     Comm. ACM 18 (1975), 181-183

import
  "sync"
type
  ImpParnas struct {
             avail [3]bool
             mutex,
             agent sync.Mutex
         layedDown,
          maySmoke [3]sync.Mutex
                   }


func (x *ImpParnas) help (u uint) {
//
  var first bool
  for {
    x.layedDown[u].Lock ()
    x.mutex.Lock ()
    u1, u2:= (u + 1) % 3, (u + 2) % 3
    first = true
    if x.avail[u1] {
      first = false
      x.avail[u1] = false
      x.maySmoke[u2].Unlock ()
    }
    if x.avail[u2] {
      first = false
      x.avail[u2] = false
      x.maySmoke[u1].Unlock () }
    if first {
      x.avail[u] = true
    }
    x.mutex.Unlock ()
  }
}


func NewParnas () *ImpParnas {
//
  x:= new (ImpParnas)
  for u:= uint(0); u < 3; u++ {
    x.layedDown[u].Lock ()
    x.maySmoke[u].Lock ()
  }
  x.agent.Lock()
  for u:= uint(0); u < 3; u++ {
    go x.help (u)
  }
  return x
}


func (x *ImpParnas) AgentIn (u uint) {
//
  x.agent.Lock ()
  x.layedDown[(u + 1) % 3].Unlock ()
  x.layedDown[(u + 2) % 3].Unlock ()
}


func (x *ImpParnas) AgentOut () {
//
}


func (x *ImpParnas) SmokerIn (u uint) {
//
  x.maySmoke[u].Lock ()
}


func (x *ImpParnas) SmokerOut () {
//
  x.agent.Unlock ()
}
