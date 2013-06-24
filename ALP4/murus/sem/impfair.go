package sem

// (c) Christian Maurer   v. 120217 - license see murus.go

// Uddings Algorithm
// Ben-Ari p. 129-131

import
  . "sync"
type
  ImpFair struct {
           gate1,
           gate2,
         onlyOne Mutex
          nGate1,
          nGate2 uint
                 }


func NewFair () *ImpFair {
//
  x:= new (ImpFair)
  x.gate2.Lock()
  return x
}


func (x *ImpFair) P() {
//
  x.gate1.Lock()
  x.nGate1 ++
  x.gate1.Unlock()
  x.onlyOne.Lock()
  x.gate1.Lock()
  x.nGate2 ++
  if x.nGate1 > 0 {
    x.gate1.Unlock()
  } else {
    x.gate2.Unlock()
  }
  x.onlyOne.Unlock()
  x.gate2.Lock()
  x.nGate2--
}


func (x *ImpFair) V() {
//
  if x.nGate2 > 0 {
    x.gate2.Unlock()
  } else {
    x.gate1.Unlock()
  }
}

// Udding Inform. Process. Lett. 23 (1986), 159

/*
func () Lock() {
  eu.Lock()
  ne++
  eu.Unlock()
  qu.Lock()
  eu.Lock()
  nm++
  ne--
  if ne > 0 {
    eu.Unlock()
  } else if ne == 0 {
    mu.Unlock()
  }
  qu.Unlock()
  mu.Lock()
  nm--
}

func () V() {
  if nm > 0 {
    mu.Unlock()
  } else if nm == 0 {
    eu.Unlock()
  }
}
*/

// Morris Inform. Process. Lett 8 (1979) 89

/*
func () P() {
  em.Lock()
  ne++
  em.Unlock()
  qm.Lock()
  nm++
  em.Lock()
  ne--
  if ne > 0 {
    em.Unlock()
    qm.Unlock()
  } else if ne == 0 {
    em.Unlock()
    mm.Unlock()
  }
  mm.Lock()
  nm--
}


func () V() {
  if nm > 0 {
    mm.Unlock()
  } else if nm == 0 {
    qm.Unlock()
  }
}
*/
