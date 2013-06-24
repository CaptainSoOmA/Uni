package lock2

// (c) Christian Maurer   v. 130210 - license see murus.go

type
  ImpKessels struct {
         interested [2]bool
           favoured [2]uint // < 2
                    }


func NewKessels () *ImpKessels {
//
  return new (ImpKessels)
}


func (x *ImpKessels) Lock (p uint) {
//
  if p > 1 { return }
  x.interested[p] = true
  x.favoured[p] = (p + x.favoured[1 - p]) % 2
  for x.interested[1 - p] && x.favoured[p] == (p + x.favoured[1 - p]) % 2 { /* do nothing */ }
}


func (x *ImpKessels) Unlock (p uint) {
// 
  if p > 1 { return }
  x.interested[p] = false
}
