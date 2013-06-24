package lock

// (c) Christian Maurer   v. 111111 - license see murus.go

type
  ImpFA struct {
             n uint32
               }


func NewFA () *ImpFA {
//
  return new (ImpFA)
}


func (L *ImpFA) Lock() {
//
  for FetchAndAddUint32 (&L.n, uint32(1)) != 0 {
    null ()
  }
}


func (L *ImpFA) Unlock() {
//
  L.n = uint32(0)
}
