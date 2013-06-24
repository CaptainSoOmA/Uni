package phil

// (c) Christian Maurer   v. 130526 - license see murus.go

// >>> Speisende Philosophen mit universellen kritischen Abschnitten
//     Beseitigung der Unfairness durch "aging"

import (
  . "murus/obj"
  . "murus/cs"
)
type
  LockJulian struct {
                    CriticalSection
                    }


func NewLockJulian () *LockJulian {
//
  var (
    eating []bool = make ([]bool, nPhilos)
    hungry []uint = make ([]uint, nPhilos)
  )
  x:= new (LockJulian)
  c:= func (p uint) bool {
        starving:= true
        for i:= uint(0); i < nPhilos; i++ {
          starving = starving && hungry[p] >= hungry[i]
        }
        return starving &&
               ! eating[left(p)] &&
               ! eating[right(p)]
     }
  l:= func (a Any, p uint) {
        eating[p] = true
      }
  u:= func (a Any, p uint) {
        eating[p] = false
        hungry[p] = 0
        for i:= uint(0); i < nPhilos; i++ {
          if x.Blocked(i) { hungry[i]++ }
        }
      }
  x.CriticalSection = New (nPhilos, c, l, u)
  return x
}


func (x *LockJulian) Lock (p uint) {
//
  changeStatus (p, hungry)
  x.Enter (p, nil)
  changeStatus (p, dining)
}


func (x *LockJulian) Unlock (p uint) {
//
  changeStatus (p, satisfied)
  x.Leave (p, nil)
}
