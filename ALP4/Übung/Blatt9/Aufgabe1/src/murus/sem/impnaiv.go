package sem

// (c) Christian Maurer   v. 120910 - license see murus.go

// >>> naive incorrect solution
//     Nichtsequentielle Programmierung mit Go 1 kompakt, S. 70

import
  "sync"
type
  ImpNaive struct {
              val int
           binsem,
            mutex sync.Mutex
                 }


func NewNaive (n uint) *ImpNaive {
//
  S:= new (ImpNaive)
  S.val = int(n)
  S.binsem.Lock ()
  return S
}


func (S *ImpNaive) P () {
//
  S.mutex.Lock ()
  S.val--
  if S.val < 0 {
    S.mutex.Unlock()
    S.binsem.Lock ()
  } else {
    S.mutex.Unlock ()
  }
}


func (S *ImpNaive) V () {
//
  S.mutex.Lock ()
  S.val++
  if S.val <= 0 {
    S.binsem.Unlock ()
  }
  S.mutex.Unlock ()
}
