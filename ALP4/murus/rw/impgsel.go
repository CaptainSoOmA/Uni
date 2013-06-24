package rw

// (c) Christian Maurer   v. 120909 - license see murus.go

//     readers/writers problem: solution with "guarded selective waiting"
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 185

import
  . "murus/obj"
type
  ImpGSel struct {
          iR, oR,
          iW, oW chan Any
            done chan int
                 }


func NewGSel () *ImpGSel {
//
  x:= new (ImpGSel)
  x.iR, x.oR = make (chan Any), make (chan Any)
  x.iW, x.oW = make (chan Any), make (chan Any)
  x.done = make (chan int)
  go func () {
    var nR, nW uint // active readers, writers
//    loop:
    for {
      select {
//      case <-x.done: break loop
      case <-When (nW == 0, x.iR):
        nR++
      case <-When (nR > 0, x.oR):
        nR--
      case <-When (nR == 0 && nW == 0, x.iW):
        nW = 1
      case <-When (nW == 1, x.oW):
        nW = 0
      }
    }
  }()
  return x
}


func (x *ImpGSel) ReaderIn () {
//
  x.iR <- 0
}


func (x *ImpGSel) ReaderOut () {
//
  x.oR <- 0
}


func (x *ImpGSel) WriterIn () {
//
  x.iW <- 0
}


func (x *ImpGSel) WriterOut () {
//
  x.oW <- 0
}


func (x *ImpGSel) TerminateGSel () {
//
  x.done <- 0
}
