package phil

// (c) Christian Maurer   v. 130526 - license see murus.go

// >>> Solution with message-passing
//     Ben-Ari: Principles of Concurrent and Distributed Programming 2nd edition
//     modified to be unsymmetric to avoid deadlocks
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 187

type
  LockCh struct {
             ch []chan bool
                }


func NewLockCh () *LockCh {
//
  x:= new (LockCh)
  x.ch = make ([]chan bool, nPhilos)
  for p:= uint(0); p < nPhilos; p++ {
    x.ch[p] = make (chan bool)
  }
  for p:= uint(0); p < nPhilos; p++ {
    go func (i uint) {
         for {
           x.ch[i] <- true
           <-x.ch[i]
         }
       }(p)
  }
  return x
}


func (x *LockCh) Lock (p uint) {
//
  changeStatus (p, hungry)
  if p % 2 == 0 {
    <-x.ch[left (p)]
    changeStatus (p, hasLeftFork)
    <-x.ch[p]
  } else {
    <-x.ch[p]
    changeStatus (p, hasRightFork)
    <-x.ch[left (p)]
  }
  changeStatus (p, dining)
}


func (x *LockCh) Unlock (p uint) {
//
  x.ch[p] <- true
  x.ch[left (p)] <- true
  changeStatus (p, satisfied)
}
