package pset

// (c) Christian Maurer   v. 130115 - license see murus.go

import
  . "murus/obj"
type
  PersistentSet interface {

  Collector
  ExGeq (a Object) bool
  Trav (o Op)
  Persistor
}
