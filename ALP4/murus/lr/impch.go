package lr

// (c) Christian Maurer   v. 120330 - license see murus.go

// >>> left/right problem: implementation with channels
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. TODO

type
  ImpCh struct {
        lI, lO,
        rI, rO,
          done chan int
               }


func NewCh () *ImpCh {
//
  x:= new (ImpCh)
  x.lI, x.lO = make (chan int), make (chan int)
  x.rI, x.rO = make (chan int), make (chan int)
  x.done = make (chan int)
  go func () {
    var nL, nR int
    for {
//      if _, ok:= <-x.done; ok { break }
      if nL == 0 {
        if nR == 0 {
          select { case <-x.lI:
            nL ++
          case <-x.rI:
            nR ++
          }
        } else { // nR > 0
          select { case <-x.rI:
            nR ++
          case <-x.rO:
            nR --
          }
        }
      } else { // nL > 0
        select { case <-x.lI:
          nL ++
        case <-x.lO:
          nL --
        }
      }
    }
  }()
  return x
}


func (x *ImpCh) LeftIn () {
//
  x.lI <- 0
}


func (x *ImpCh) LeftOut () {
//
  x.lO <- 0
}


func (x *ImpCh) RightIn () {
//
  x.rI <- 0
}


func (x *ImpCh) RightOut () {
//
  x.rO <- 0
}


func (x *ImpCh) Terminate () {
//
  x.done <- 0
}
