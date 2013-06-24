package puls

// (c) Christian Maurer   v. 121030 - license see murus.go

import
  . "murus/obj"
type
  Pulse interface {

// New (host []string, port []uint, nb[]uint, f FuncSpectrum, a Any) *Imp

// TODO Spec
  Go (i uint)

// TODO Spec
  F (i uint) Any

// TODO Spec
  Terminate ()
}
