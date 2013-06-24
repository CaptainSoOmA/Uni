package seq

// (c) Christian Maurer   v. 130115 - license see murus.go

import
  . "murus/obj"
type
  Sequence interface {

  Equaler
  Coder
  Sorter
  SeekerIterator
}
