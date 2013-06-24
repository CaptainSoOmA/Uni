package char

// (c) Christian Maurer   v. 121215 - license see murus.go

import
  . "murus/obj"
type
  Character interface {

  Editor
  Stringer
  Printer
  Valuator

  Equiv (y Object) bool
  Def (b byte) // -> Defined (b byte) bool // ?
  ByteVal () byte
}
