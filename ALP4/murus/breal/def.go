package breal

// (c) Christian Maurer   v. 130115 - license see murus.go

import
  . "murus/obj"
type
  Real interface { // real numbers < some power of 10

// Returns a new Object, that can hold real numbers
// with at most d digits , where d = nat.Len (n).
//  func New (v, n uint) *Imp

  Editor
  Stringer
  Printer

  RealVal () float64
// Adder, Multiplier, further arithmetics ?
}
