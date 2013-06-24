package lockp

// (c) Christian Maurer   v. 111126 - license see murus.go

// >>> Bakery-Algorithm of Lamport, corrected version

type
  ImpBakery1 struct {
         nProcesses uint
             number []uint
              draws []bool
                    }


func (L *ImpBakery1) max () uint {
  m:= uint(0)
  for i:= uint(1); i <= L.nProcesses; i++ {
    if L.number[i] > m {
      m = L.number[i]
    }
  }
  return m
}


func (L *ImpBakery1) less (i, k uint) bool {
  if L.number[i] < L.number[k] {
    return true
  }
  if L.number[i] == L.number[k] {
    return i < k
  }
  return false
}


func NewBakery1 (n uint) *ImpBakery1 {
//
  if n < 2 { return nil }
  L:= new (ImpBakery1)
  L.nProcesses = n
  L.number = make ([]uint, n)
  L.draws = make ([]bool, n)
  return L
}


func (L *ImpBakery1) Lock (p uint) {
//
  if p >= L.nProcesses { return }
  L.number[p] = 1
  L.number[p] = L.max() + 1
  for a:= uint(1); a <= L.nProcesses; a++ {
    if a != p {
      for L.number[a] > 0 && L.less (a, p) { /* nichts */ }
    }
  }
}


func (L *ImpBakery1) Unlock (p uint) {
//
  if p >= L.nProcesses { return }
  L.number[p] = 0
}
