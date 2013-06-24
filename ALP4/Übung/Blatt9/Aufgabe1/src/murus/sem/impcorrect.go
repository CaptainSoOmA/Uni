package sem

// (c) Christian Maurer   v. 120910 - license see murus.go

// >>> corrected naive solution
//     Nichtsequentielle Programmierung mit Go 1 kompakt, S. 71

import
  "sync"
type
  ImpCorrect struct {
                val int
             binsem,
              mutex sync.Mutex
                   }


func NewCorrect (n uint) *ImpCorrect {
//
  x:= new (ImpCorrect)
  x.val = int(n)
  x.binsem.Lock ()
  return x
}


func (x *ImpCorrect) P () {
//
  x.mutex.Lock ()
  x.val--
  if x.val < 0 {
    x.mutex.Unlock()
    x.binsem.Lock ()
  }
  x.mutex.Unlock ()
}


func (x *ImpCorrect) V () {
//
  x.mutex.Lock ()
  x.val++
  if x.val <= 0 {
    x.binsem.Unlock ()
  } else {
    x.mutex.Unlock ()
  }
}
