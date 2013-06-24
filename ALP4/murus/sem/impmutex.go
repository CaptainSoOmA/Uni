package sem

// (c) Christian Maurer   v. 120910 - license see murus.go

// >>> solution by Barz
//     Nichtsequentielle Programmierung mit Go 1 kompakt, S. 73

import
  "sync"
type
  ImpMutex struct {
               cs,
               me sync.Mutex
              val int
                  }


func NewMutex (n uint) *ImpMutex {
//
  x:= new (ImpMutex)
  x.val = int(n)
  if x.val == 0 {
    x.cs.Lock ()
  }
  return x
}


func (x *ImpMutex) P () {
//
  x.cs.Lock ()
  x.me.Lock ()
  x.val--
  if x.val > 0 {
    x.cs.Unlock ()
  }
  x.me.Unlock ()
}


func (x *ImpMutex) V () {
//
  x.me.Lock ()
  x.val++
  if x.val == 1 {
    x.cs.Unlock ()
  }
  x.me.Unlock ()
}
