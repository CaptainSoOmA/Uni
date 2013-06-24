package obj

// (c) Christian Maurer   v. 120909 - license see murus.go

type
  Format byte
type
  Formatter interface {

// x has the format f.
  SetFormat (f Format)
}
