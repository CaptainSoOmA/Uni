package text

// (c) Christian Maurer   v. 130116 - license see murus.go

import
  . "murus/obj"
type
  Text interface {

  Editor
  Stringer
  Printer

// Specs see str/def.go.
  Equiv (Y Object) bool

  IsPart (Y Object) bool
  IsEquivalentPart (Y Object) bool

  Len () uint
  ProperLen () uint

  IsCap0 () bool
  ToUpper ()
  ToLower ()
}
