package phil

// (c) Christian Maurer   v. 130526 - license see murus.go

// >>> Unsymmetric solution with message-passing
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 187

type
  LockCh1 struct {
        cl, cll,
        cu, cul []chan bool
                 }


func NewLockCh1 () *LockCh1 {
//
  x:= new (LockCh1)
  x.cl  = make ([]chan bool, nPhilos)
  x.cll = make ([]chan bool, nPhilos)
  x.cu  = make ([]chan bool, nPhilos)
  x.cul = make ([]chan bool, nPhilos)
  for p:= uint(0); p < nPhilos; p++ {
    x.cl [p] = make (chan bool)
    x.cll[p] = make (chan bool)
    x.cu [p] = make (chan bool)
    x.cul[p] = make (chan bool)
  }
  for p:= uint(0); p < nPhilos; p++ {
    go func (i uint) {
         for {
           select {
           case <-x.cl[i]:
             <-x.cu[i]
           case <-x.cll[i]:
             <-x.cul[i]
           }
         }
       }(p)
  }
  return x
}


func (x *LockCh1) Lock (p uint) {
//
  changeStatus (p, hungry)
  if p % 2 == 0 {
    x.cll[left(p)] <- true
    changeStatus (p, hasLeftFork)
    x.cl[p] <- true
  } else {
    x.cl[p] <- true
    changeStatus (p, hasRightFork)
    x.cll[left(p)] <- true
  }
  changeStatus (p, dining)
}


func (x *LockCh1) Unlock (p uint) {
//
  changeStatus (p, satisfied)
  x.cul[left(p)] <- true
  x.cu[p] <- true
}
