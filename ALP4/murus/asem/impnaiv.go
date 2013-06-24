package asem

// (c) Christian Maurer   v. 130326 - license see murus.go

// >>> incorrect naive representation
//     Nichtsequentielle Programmierung mit Go 1 kompakt, S. 100

import
  "murus/sem"
type
  ImpNaiv struct {
               p uint32 // number of processes allowed to use the critical section,
                        // which shall be protected by the semaphore
               s sem.Semaphore
                 }


func NewNaiv (n uint) *ImpNaiv {
//
  return &ImpNaiv { p: uint32(n), s: sem.New (n) }
}


func (x *ImpNaiv) Padd (n uint) {
//
  for n > 0 {
    x.s.P()
    n--
  }
}


func (x *ImpNaiv) Vadd (n uint) {
//
  for n > 0 {
    x.s.V()
    n--
  }
}
