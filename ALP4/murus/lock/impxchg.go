package lock

// (c) Christian Maurer   v. 120225 - license see murus.go

type
  ImpXCHG struct {
               n uint32
                 }


func NewXCHG () *ImpXCHG {
//
  return &ImpXCHG { 1 }
}


func (L *ImpXCHG) Lock() {
//
  local:= uint32(0)
  for ExchangeUint32 (&L.n, local) == 0 {
    null ()
  }
}


func (L *ImpXCHG) Unlock() {
//
//  local:= uint32(1); local = ExchangeUint32 (&L.n, local)
  L.n = 1
}
