package lockp

// (c) Christian Maurer   v. 111126 - license see murus.go

// >>> Algorithm of Dijkstra

type
  ImpDijkstra struct {
          nProcesses,
            favoured uint
           willenter,
            critical []bool
                     }


func NewDijkstra (n uint) *ImpDijkstra {
//
  if n < 2 { return nil }
  L:= new (ImpDijkstra)
  L.nProcesses = n
  L.favoured = 0
  L.willenter = make ([]bool, n + 1) // n + 1 ???
  L.critical = make ([]bool, n)
  return L
}


func (L *ImpDijkstra) Lock (p uint) {
//
  if p >= L.nProcesses { return }
  L.willenter [p] = true
//  var andererKritisch bool
  for {
    L.critical [p] = false
    for L.favoured != p {
      if ! L.willenter [L.favoured] {
        L.favoured = p
      }
    }
    L.critical [p] = true
    andererKritisch:= false
    for a:= uint(1); a <= L.nProcesses; a++ {
      if a != p {
        andererKritisch = andererKritisch || L.critical [a]
      }
    }
    if ! andererKritisch {
      break
    }
  }
}


func (L *ImpDijkstra) Unlock (p uint) {
//
  if p >= L.nProcesses { return }
  L.willenter [p] = false
  L.critical [p] = false
  L.favoured = 0
}
