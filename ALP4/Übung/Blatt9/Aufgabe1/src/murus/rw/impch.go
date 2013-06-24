package rw

// (c) Christian Maurer   v. 120330 - license see murus.go

//     readers/writers problem, solution with client-server-paradigma
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 182

type
  ImpCh struct {
        rI, rO,
        wI, wO,
          done chan int
               }


func NewCh () *ImpCh {
//
  x:= new (ImpCh)
  x.rI, x.rO = make (chan int), make (chan int)
  x.wI, x.wO = make (chan int), make (chan int)
  x.done = make (chan int)
  go func () {
    var nR, nW uint // active readers, writers
    for {
// if _, ok:= <-x.done; ok { break }
/*
      if nR == 0 {
        if nW == 0 {
          select { case <-x.rI:
            nR ++
          case <-x.wI:
            nW = 1
          }
        } else { // nW == 1
          select { case <-x.wO:
            nW = 0
          }
        }
      } else { // nR > 0
        select { case <-x.rI:
          nR ++
        case <-x.rO:
          nR --
        }
      }
*/
      if nW == 0 {
        if nR == 0 {
          select {
          case <-x.rI:
            nR ++
          case <-x.wI:
            nW = 1
          }
        } else { // nR > 0
          select {
          case <-x.rI:
            nR ++
          case <-x.rO:
            nR --
          }
        }
      } else { // nW == 1
        select { case <-x.wO:
          nW = 0
        }
      }
    }
  }()
  return x
}


func (x *ImpCh) ReaderIn () {
//
  x.rI <- 0
}


func (x *ImpCh) ReaderOut () {
//
  x.rO <- 0
}


func (x *ImpCh) WriterIn () {
//
  x.wI <- 0
}


func (x *ImpCh) WriterOut () {
//
  x.wO <- 0
}


func (x *ImpCh) TerminateCh () {
//
  x.done <- 0
}
