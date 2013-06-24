package lr

// (c) Christian Maurer   v. 120909 - license see murus.go

// >>> left/right problem: implementation with "guarded selective waiting"
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. TODO


import
  . "murus/obj"
type
  ImpGSel struct {
          iL, oL,
          iR, oR chan Any
            done chan int
                 }


func NewGSel () *ImpGSel {
//
  x:= new (ImpGSel)
  x.iL, x.oL = make (chan Any), make (chan Any)
  x.iR, x.oR = make (chan Any), make (chan Any)
  x.done = make (chan int)
  go func () {
    var nL, nR uint // active lefts, rights
//    loop:
    for {
      select {
//      case <-x.done: break loop
      case <-When (nR == 0, x.iL):
        nL++
      case <-When (nL > 0, x.oL):
        nL--
      case <-When (nL == 0, x.iR):
        nR++
      case <-When (nR > 0, x.oR):
        nR--
      }
    }
  }()
  return x
}


func (x *ImpGSel) LeftIn () {
//
  x.iL <- 0
}


func (x *ImpGSel) LeftOut () {
//
  x.oL <- 0
}


func (x *ImpGSel) RightIn () {
//
  x.iR <- 0
}


func (x *ImpGSel) RightOut () {
//
  x.oR <- 0
}


func (x *ImpGSel) TerminateGSel () {

  x.done <- 0
}
