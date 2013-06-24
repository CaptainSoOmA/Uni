package lock2

type
  ImpUeb struct {
     interested [2]bool
                }

/*
func NewUeb () *ImpUeb {
//
  return &ImpUeb { [2]bool { true, true } }
}


func (L *ImpUeb) Lock (p uint) {
//
  for {
//    if interested[p] { interested[1-p] = false } else { interested[1-p] = true }
    L.interested[1-p] = ! L.interested[p]
    if L.interested[p] && ! L.interested[1-p] { break }
  }
}

func (L *ImpUeb) Unlock (p uint) {
//
  L.interested[1-p] = true
}
*/

func (L *ImpUeb) Lock (p uint) {
//
  for L.interested[p] {
    if L.interested[1-p] {
      L.interested[p] = false
    }
  }
}

func (L *ImpUeb) Unlock (p uint) {
//
  L.interested[p] = true
}
