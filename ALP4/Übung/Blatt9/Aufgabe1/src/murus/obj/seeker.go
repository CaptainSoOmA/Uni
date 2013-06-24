package obj

// (c) Christian Maurer   v. 130115 - license see murus.go

type
  Seeker interface { // makes only sense for objects of type Collector

// Returns Num(), iff Offc(); returns otherwise
// the position of the actual object of x (starting at 0).
  Pos () uint

// The actual object of x is its p-th object, iff p < Num();
// otherwise Offc() == true.
  Seek (p uint)
}


type
  SeekerIterator interface {

  Iterator
  Seeker
}
