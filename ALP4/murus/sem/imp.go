package sem

// (c) Christian Maurer   v. 120330 - license see murus.go

type
  Imp struct {
           c chan bool
             }


func New (n uint) *Imp {
//
  x:= new (Imp)
  x.c = make(chan bool, n)
  for i:= uint(0); i < n; i++ {
    x.c <- true
  }
  return x
}


func (x *Imp) P() {
//
  <-x.c
}


func (x *Imp) V() {
//
  x.c <- true
}
