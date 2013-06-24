package lock2

// (c) Christian Maurer   v. 111126 - license see murus.go

type
  ImpPeterson struct {
          interested [2]bool
            favoured uint
                     }

func NewPeterson () *ImpPeterson {
//
  return new (ImpPeterson)
}


func (L *ImpPeterson) Lock (p uint) {
//
  if p > 1 { return }
  L.interested[p] = true
  L.favoured = 1-p
  for L.interested[1-p] && L.favoured == 1-p { /* do nothing */ }
}


func (L *ImpPeterson) Unlock (p uint) {
//
  if p > 1 { return }
  L.interested[p] = false
}
