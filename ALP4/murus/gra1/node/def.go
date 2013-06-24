package node

// (c) Christian Maurer   v. 121215 - license see murus.go

import (
  . "murus/obj"
  "murus/col"
)
var
  transparent bool // at the beginning false
type
// Nodes with names, ..., and positions on the screen
  Node interface {

//// x has the namelength n.
//  New (n uint) { // n < 22

  Object // Empty means no name

// Pre: len (s) <= namelength of x; C is one of '<', '>', '^' or '_'; x < scr.NX, y < scr.NY.
// x ... .
  Def (s string, c byte, x, y uint)

// x hat die Position der Maus.
  Locate ()

// Returns the name of x.
  String () string

// Returns the coordinates of x.
  Pos () (uint, uint)

// Returns 4, if the namelength of x > 2 ist, otherwise ... .
  Radius () uint

// Returns true, iff the mouse has the position of x.
//  UnderMouse (a Any) bool

// f is the normal and a the actual Colour of the Nodes.
  SetColours (f, a col.Colour)

// x is written to the screen at its position (evtl. transparent)
// if inv, the colour of the pixels of x on the screen are inverted.
// else, if ! vis, is written with the bgcolour of the screen.
  Write (vis, inv bool)

// x is written at its position to the screen, for u == true in its actual, else in its normal colour.
  WriteCond (u bool)

// x has the name, that the user has edited.
  Edit ()
}
