package conn

// (c) Christian Maurer   v. 121105 - license see murus.go

import
  . "murus/obj"
type
  Connection interface {

  Object
  Formatter // Format of "murus/host"
  Stringer
  Defined2 (s string, p uint) bool
}
