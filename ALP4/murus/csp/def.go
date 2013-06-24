package csp

// (c) Christian Maurer   v. 130525 - license see murus.go

import
  . "murus/obj"
type
  CriticalSectionP interface {

  Blocked (k uint) bool

  Enter (k uint, a Any)

  Leave (k uint, a Any)
}
