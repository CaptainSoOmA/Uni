package mol

// (c) Christian Maurer   v. 130115 - license see murus.go

import (
  . "murus/obj"
  "murus/atom"
)
type
  Molecule interface {

  Editor
  Printer

  Ins (a *atom.Imp, l, c uint)
  Del (n uint)
  Num () uint
  Component () Any

// Equiv (Y Object) bool
// Sort ()
}
