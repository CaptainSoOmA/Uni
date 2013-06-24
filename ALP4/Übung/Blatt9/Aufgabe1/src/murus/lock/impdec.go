package lock

// (c) Christian Maurer   v. 120225 - license see murus.go

type
  ImpDEC struct {
              n int32
                }


func NewDEC () *ImpDEC {
//
  return &ImpDEC { n: int32(1) }
}


func (L *ImpDEC) Lock() {
//
  for DecrementInt32 (&L.n) {
    null ()
  }
}


func (L *ImpDEC) Unlock() {
//
  L.n = 1
}
