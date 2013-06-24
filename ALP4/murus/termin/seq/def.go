package seq

// (c) Christian Maurer   v. 130127 - license see murus.go

import (
  . "murus/obj"
  "murus/day"
)
type
  Sequence interface {

  Editor
  Printer

  SetFormat (p day.Period)

  HasWord () bool
}
