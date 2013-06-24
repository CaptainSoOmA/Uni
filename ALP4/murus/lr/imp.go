package lr

// (c) Christian Maurer   v. 120330 - license see murus.go

// >>> Left/Right problem: Simple Solution with mutexes
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 79

import
  . "sync"
type
  Imp struct {
      nL, nR int
      mL, mR,
          lr Mutex
             }


func New () *Imp {
//
  return new (Imp)
}


func (x *Imp) LeftIn () {
//
  x.mL.Lock ()
  x.nL++
  if x.nL == 1 {
    x.lr.Lock ()
  }
  x.mL.Unlock ()
}


func (x *Imp) LeftOut () {
//
  x.mL.Lock ()
  x.nL--
  if x.nL == 0 {
    x.lr.Unlock ()
  }
  x.mL.Unlock ()
}


func (x *Imp) RightIn () {
//
  x.mR.Lock ()
  x.nR++
  if x.nR == 1 {
    x.lr.Lock ()
  }
  x.mR.Unlock ()
}


func (x *Imp) RightOut () {
//
  x.mR.Lock ()
  x.nR--
  if x.nR == 0 {
    x.lr.Unlock ()
  }
  x.mR.Unlock ()
}
