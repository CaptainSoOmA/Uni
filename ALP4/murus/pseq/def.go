package pseq

// (c) Christian Maurer   v. 131015 - license see murus.go

import
  . "murus/obj"
type
  PersistentSequence interface {

  Eq (Y Object) bool
//  Sorter
  SeekerIterator
  Persistor
}
