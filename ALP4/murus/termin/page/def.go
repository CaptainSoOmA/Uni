package page

// (c) Christian Maurer   v. 130510 - license see murus.go

import (
  . "murus/obj"
  "murus/day"
)
type
  Page interface {

  Editor
  Printer

  SetFormat (p day.Period)

  Set (d *day.Imp)
  Day () Any

  HasWord () bool

  Terminate ()
}
