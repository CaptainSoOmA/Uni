package piset

// (c) Christian Maurer   v. 130115 - license see murus.go

import
  . "murus/obj"
type
  PersistentIndexedSet interface {

  Sorter
  Iterator
  Persistor
}
