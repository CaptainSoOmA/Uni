package lock2

// (c) Christian Maurer   v. 111126 - license see murus.go

type
  ImpDekker struct {
        interested [2]bool
          favoured uint
                   }


func NewDekker () *ImpDekker {
//
  return new (ImpDekker)
}


func (L *ImpDekker) Lock (p uint) {
//
  if p > 1 { return }
  L.interested[p] = true
  for L.interested [1-p] {
    if L.favoured == 1 - p {
      L.interested [p] = false
      for L.favoured == 1 - p { /* do nothing */ }
      L.interested [p] = true
    }
  }
}


func (L *ImpDekker) Unlock (p uint) {
//
  if p > 1 { return }
  L.favoured = 1 - p
  L.interested[p] = false
}
