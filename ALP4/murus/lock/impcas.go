package lock

// (c) Christian Maurer   v. 111125 - license see murus.go

import
  . "sync/atomic"
type
  ImpCAS struct {
              n uint32
                }


func NewCAS () *ImpCAS {
//
  return new (ImpCAS)
}


func (L *ImpCAS) Lock() {
//
  for ! CompareAndSwapUint32 (&L.n, 0, 1) {
    null ()
  }
}


func (L *ImpCAS) Unlock() {
//
  L.n = 0
}
