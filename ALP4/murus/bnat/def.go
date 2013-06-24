package bnat

// (c) Christian Maurer   v. 130304 - license see murus.go

import
  . "murus/obj"
type
  Natural interface { // natural numbers < some power of 10

// New (n uint) returns a new Object, that can hold natural
// numbers with at most d digits, where d = nat.Len (n).
// String() has always leading zeros, iff n % 10 == 0.

  Editor
  Stringer
  Valuator
  Printer

  Startval () uint
// Adder, Multiplier, further arithmetics ?
}
