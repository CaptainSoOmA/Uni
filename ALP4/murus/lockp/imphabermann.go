package lockp

// (c) Christian Maurer   v. 111126 - license see murus.go

// >>> Algorithm of Habermann

type
  ImpHabermann struct {
           nProcesses,
             favoured uint
            willenter,
             critical []bool
                      }


func NewHabermann (n uint) *ImpHabermann {
//
  if n < 2 { return nil }
  L:= new (ImpHabermann)
  L.nProcesses = n
  L.favoured = 0
  L.willenter = make ([]bool, n + 1) // n + 1 ???
  L.critical = make ([]bool, n)
  return L
}


func (L *ImpHabermann) Lock (p uint) {
//
  if p >= L.nProcesses { return }
  var (
    b uint
    andererKritisch, andererEintrittswillig bool
  )
  for {
    L.willenter[p] = true
    for {
      L.critical[p] = false
      b = L.favoured
      andererEintrittswillig = false
      for b != p {
        andererEintrittswillig = L.willenter[b] || andererEintrittswillig
        if b < L.nProcesses {
          b++
        } else {
          b = 1
        }
      }
      if ! andererEintrittswillig {
        break
      }
    }
    L.critical[p] = true
    andererKritisch = false
    for a:= uint(1); a <= L.nProcesses; a++ {
      if a != p {
        andererKritisch = andererKritisch || L.critical[a]
      }
    }
    if ! andererKritisch {
      break
    }
  }
  L.favoured = p
}


func (L *ImpHabermann) Unlock (p uint) {
//
  if p >= L.nProcesses { return }
  i:= p
  for {
    i = i % L.nProcesses + 1
    if L.willenter[i] || i == p {
      break
    }
  }
  L.favoured = i
  L.critical[p] = false
  L.willenter[p] = false
}
