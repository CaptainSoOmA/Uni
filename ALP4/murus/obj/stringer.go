package obj

// (c) Christian Maurer   v. 120909 - license see murus.go

type
  Stringer interface {

// Returns a string representation of x.
  String () string

// If s represents an object, x is that object,
// otherwise x is unchanged.
  Defined (s string) bool
}
