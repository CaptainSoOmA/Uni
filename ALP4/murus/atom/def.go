package atom

// (c) Christian Maurer   v. 130115 - license see murus.go

import
  . "murus/obj"
type
  Atomtype byte; const (
  Enumerator = Atomtype(iota) // enum
  TruthValue     // tval
  Character      // char
  Text           // text
  Natural        // bnat
  Real           // breal
  Clocktime      // clk
  Calendarday    // day
  Euro           // euro
  Country        // cntry
  Person         // pers
  PhoneNumber    // phone
  Address        // addr
  NAtomtypes
)
type
  Atom interface {
// Returns a new Atom of the type of o, where o is an object
// of one of the types corresponding to the above constants.
//  New (o Object) *Imp

  Object
  Formatter
  Printer

// Returns true, iff x and Y have the same type.
  Equiv (Y Object) bool

// Returns the Atomtype of x.
  Type () Atomtype
}
