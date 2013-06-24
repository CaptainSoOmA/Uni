package lock

// (c) Christian Maurer   v. 120115 - license see murus.go

// >>> Imp with asynchronous message passing

type
  ImpChan struct {
              ch chan bool
                 }


func NewChan () *ImpChan {
//
  L:= new (ImpChan)
  L.ch = make(chan bool, 1)
  L.ch <- true
  return L
}


func (L *ImpChan) Lock() {
//
  <-L.ch
}


func (L *ImpChan) Unlock() {
//
  L.ch <- true
}
