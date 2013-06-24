package pdays

// (c) Christian Maurer   v. 120909 - license see murus.go

import
  . "murus/obj"
type
  PersistentDaySet interface {

  Clearer
  Persistor

  Ex (x Object)
  Ins (x Object)
  Del (x Object)

  Terminate ()
}
