package date

// (c) Christian Maurer   v. 130105 - license see murus.go

import (
  . "murus/obj"
  "murus/day"; "murus/clk"
)
type
  DayTime interface {
// Pairs (d, t) for day d and clktime t.
// (M/O, 2) means (last Sunday in March/October, 2.00:00)

  Editor
  Stringer
  Printer

// x is (d, t).
  Set (d *day.Imp, t *clk.Imp)

// Returns the day/the time of x.
  Day() *day.Imp
  Time() *clk.Imp

// Returns true, iff x is not empty and x < (O, 2) or (M, 2) <= x.
  Normal () bool

// x has the format d for its day and t for its clktime.
  SetFormat (d, t Format)

// If x is Normal, then y equals x, otherwise y equals x + 1 hour.
  Actualize (y DayTime)

// spec TODO
  Normalize ()

// // If t is of type *clk.Imp, then
// If x is not empty, x is increased by t.
  Inc (y *clk.Imp)

// If x is not empty, x is decreased by y.
  Dec (y *clk.Imp)
}
