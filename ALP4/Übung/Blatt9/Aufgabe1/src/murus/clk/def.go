package clk

// (c) Christian Maurer   v. 121216 - license see murus.go

import
  . "murus/obj"
const (        // Format
  Hh_mm = iota // e.g. "07.32"
  Hh_mm_ss     // e.g. "13.45:27"
  Mm_ss        // e.g. "04:19"
  NFormats
)
type
  Clocktime interface { // given by a triple of uints h.m:s with h < 24 and m, s < 60.

  Editor
  Stringer
  Formatter
  Printer

// x is equal to the system time.
  Actualize ()

// Returns true, iff x lies before the system time.
  Elapsed () bool

// Returns the absolute value of the seconds between X and x.
  Distance (X Clocktime) uint

// Returns the Distance between 00.00:00 and x.
  NSeconds () uint

// Returns the value of the hour, the minutes, the seconds of x.
  Hours () uint
  Minutes () uint
  Seconds () uint

// x is increased by X mod 24 hours.
  Inc (X Clocktime)

// x is decreased by X mod 24 hours.
  Dec (X Clocktime)

// If h < 24, m < 60 and s < 60, x is set to the time h:m:s; otherwise x is empty.
  Set (h, m, s uint)
}
