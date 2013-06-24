package font

// (c) Christian Maurer   v. 120909 - license see murus.go

type
  Font byte; const (
  Roman = iota
  Bold
  Slanted
  Italic
  NFonts
)
type
  Size byte; const (
  Tiny = iota
  Small
  Normal
  Big
  Huge
  NSizes
)
const
  M = 6 // len names
var
  Name []string

// Returns a string of len 2, that uniquely defines f and s.
//  Code (f Font, s Size) string
