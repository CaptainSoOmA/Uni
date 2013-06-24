package enumbase

// (c) Christian Maurer   v. 130115 - license see murus.go

import
  . "murus/obj"
const ( // Format
  Short = iota
  Long
  NFormats
)
type
  Enumbase interface {
// New (b) *Imp returns a new Object of type t,
// where b == byte(e) for e == one of enum/Enum.

  Formatter
  Editor
  Stringer
  Printer

// Returns the type of x.
  Typ () byte

// Returns the number of elements of Enum (common for all elements).
  Num () uint

// Returns the order number of x.
  Ord () uint

// Returns the width of the string representation of x (common for all elements).
  Wd () uint

// Returns true, iff there is an n-th element in Enum.
// In this case x is that element, otherwise x is empty.
  Set (n uint) bool
}
