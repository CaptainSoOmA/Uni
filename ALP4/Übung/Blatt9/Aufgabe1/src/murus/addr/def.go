package addr

// (c) Christian Maurer   v. 130115 - license see murus.go

import
  . "murus/obj"
type
  Address interface {

  Editor
  Printer

  Equiv (Y Object) bool
}
