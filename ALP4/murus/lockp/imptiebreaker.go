package lockp

// (c) Christian Maurer   v. 111126 - license see murus.go

// >>> Tiebreaker-Algorithm of Peterson

type
  ImpTiebreaker struct {
            nProcesses uint
              achieved,
                  last []uint
                       }


func NewTiebreaker (n uint) *ImpTiebreaker {
//
  if n < 2 { return nil }
  L:= new (ImpTiebreaker)
  L.nProcesses = n
  L.achieved = make ([]uint, n)
  L.last = make ([]uint, n)
  return L
}


func (L *ImpTiebreaker) Lock (p uint) {
//
  if p >= L.nProcesses { return }
  for e:= uint(0); e < L.nProcesses - 1; e++ {
    L.achieved[p] = e
    L.last[e] = p
    for a:= uint(0); a < L.nProcesses; a++ {
      if p != a {
        for e <= L.achieved[a] && p == L.last[e] { /* do nothing */ }
      }
    }
  }
}


func (L *ImpTiebreaker) Unlock (p uint) {
//
  if p >= L.nProcesses { return }
  L.achieved[p] = 0
}
