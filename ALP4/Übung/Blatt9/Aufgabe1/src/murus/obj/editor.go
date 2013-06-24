package obj

// (c) Christian Maurer   v. 120909 - license see murus.go

import (
  "murus/col"
)
// Objects, that can be written at a particular position of
// a screen, given by logical or pixeloriented coordinates, i.e.
// by pairs of unsigned integers (line,column) or integers (x,y)
// (where (0,0) denotes the top left corner of the screen),
// and that can be changed by interaction with a user
// (e.g. by means of a keyboard), hence is
type
  Editor interface {

  Object

// x has the colours fore-/background f/b.
  SetColours (f, b col.Colour)

// Pre: l, c have to be "small enough", i.e.
//      l + height (object) < scr.NoLines,
//      c + width (object) < scr.NoColums.
// x is in its colours written to the screen
// with its left top corner at line, column = l, c.
  Write (l, c uint)

// Pre: see Write.
// x has the value, that was edited at line/column l/c.
// Hint: A "new" object is "read" by editing an empty one.
  Edit (l, c uint)


// >>>  eventually new version:

// x has the colours given by the parameters.
//  Colour (... col.Colour)

// Pre: If there are position parameters p[0], p[1],
//      then they have to be "small enough", i.e.
//      p[0] + height (object) < scr.NoLines,
//      p[1] + width (object) < scr.NoColums.
// x is in its colours written to the screen
// [ with its left top corner at line/column p[0]/p[1]) ].
//  Write (... uint)

// Precondition: see Write.
// x has the value, that was edited
// [ at line/column p[0]/p[1] (see Write) ].
// Hint: A "new" object is "read" by editing an empty one.
//  Edit (... uint)
}
