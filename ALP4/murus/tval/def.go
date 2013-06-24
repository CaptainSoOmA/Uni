package tval

// (c) Christian Maurer   v. 130115 - license see murus.go

// >>> TODO logical operations

import
  . "murus/obj"
type
  TruthValue interface { // truth values "true", "false" and "undecidable"

  Editor
  Stringer
  Printer

// Pre: len(f) == len(t) > 0.
// false/true as strings are represented by f/t;
// undecidable by an empty string of the same length.
  SetFormat (f, t string)

// The value of x is set to b.
  Set (b bool)
}
