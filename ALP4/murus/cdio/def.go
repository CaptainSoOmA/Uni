package cdio

// (c) Christian Maurer   v. 130224 - license see murus.go

import
  "murus/cdker"

// Returns (n, true), iff the mouse is within the representation of the
// static data of a track on the screen; in this case n is that track.
func TrackUnderMouse () (uint8, bool) { return trackUnderMouse() }

// Returns (b, n, true), iff the mouse is within the dynamic representation
// of a track on the screen; in this n is the actual time in seconds,
// for b == true/false from beginning of the CD/of the actual track.
func TimeUnderMouse () (bool, uint, bool) { return timeUnderMouse() }

// Returns (n, true), iff the mouse is within the representation of c
// on the screen; in this case n is the value of c, otherwise n == 0.
func ControlUnderMouse (c *cdker.Controller) (uint, bool) { return controlUnderMouse(c) }

// The static informations of the CD are written to the screen.
func WriteMask () { writeMask() }

// The actual state of the CD is written to the screen.
func Write () { write() }
