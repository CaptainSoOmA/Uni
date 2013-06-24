package masks

// (c) Christian Maurer   v. 130116 - license see murus.go

import
  . "murus/obj"
type
  MaskSequence interface {

  Object
  Write (l, c uint)
  Printer
  Line (n uint)
  Ins (m string, l, c uint)
}
