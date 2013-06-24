package lockp

// (c) Christian Maurer   v. 111111 - license see murus.go

// >>> Bakery-Algorithm of Lamport

type
  ImpBakery struct {
        nProcesses uint
            number []uint
             draws []bool
                   }


func (L *ImpBakery) max () uint {
  m:= uint(0)
  for i:= uint(1); i <= L.nProcesses; i++ {
    if L.number[i] > m {
      m = L.number[i]
    }
  }
  return m
}


func (L *ImpBakery) less (i, k uint) bool {
  if L.number[i] < L.number[k] {
    return true
  }
  if L.number[i] == L.number[k] {
    return i < k
  }
  return false
}


func NewBakery (n uint) *ImpBakery {
//
  if n < 2 { return nil }
  L:= new (ImpBakery)
  L.nProcesses = n
  L.number = make ([]uint, n)
  L.draws = make ([]bool, n)
  return L
}


func (L *ImpBakery) Lock (p uint) {
//
  if p >= L.nProcesses { return }
  L.draws[p] = true
  L.number[p] = L.max() + 1
  L.draws[p] = false
  for a:= uint(1); a <= L.nProcesses; a++ {
    for L.draws[a] { /* do nothing */ }
    for L.number[a] > 0 && L.less (a, p) { /* do nothing */ }
  }
}


func (L *ImpBakery) Unlock (p uint) {
//
  if p >= L.nProcesses { return }
  L.number[p] = 0
}
