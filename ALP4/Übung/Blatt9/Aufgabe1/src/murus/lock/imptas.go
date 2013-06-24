package lock

// (c) Christian Maurer   v. 111125 - license see murus.go

type
  ImpTAS struct {
         locked bool
                }


func NewTAS () *ImpTAS {
//
  return new (ImpTAS)
}


func (L *ImpTAS) Lock() {
//
  for TestAndSet (&L.locked) {
    null ()
  }
}


func (L *ImpTAS) Unlock() {
//
  L.locked = false
}
