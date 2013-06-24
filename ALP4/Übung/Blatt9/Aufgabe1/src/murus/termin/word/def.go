package word

// (c) Christian Maurer   v. 130113 - license see murus.go

import
  . "murus/obj"
type
  Word interface {

  Editor
  Stringer
  Marker
  Printer
}
