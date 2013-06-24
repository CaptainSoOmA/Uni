package pbar

// (c) Christian Maurer   v. 130307 - license see murus.go

import
  "murus/col"
type
  Progressbar interface {

// New (h bool) *Imp returns for h == true/false a bar with horizontal
// resp. vertical fill direction, capacity 100 and fill degree 0.
// X denotes the calling progress bar.

// Pre: x + w <= scr.NX, y + h <= scr.NY.
// X has the top left corner (x, y), width w and height h.
  Locate (x, y, w, h uint)

// X has capacity c and fill degree 0.
  Def (c uint)

// Pre: i <= capacity of X.
// X has fill degree i, i.e. X is filled up to i/k % (c = capacity of B).
  Fill (i uint)

// Returns the fill degree of X.
  Filldegree () uint

// X has the fore-/backgroundcolour f resp. b.
  SetColours (f, b col.Colour)

// X is written to the screen as rectangle with its top left corner
// and its width and height, the fraction of the capacity of B, that
// corresponds to the fill degree of B (with horizontal fill direction
// the left, otherwise the bottom part of the rectangle)
// in its foregroundcolour, the rest in its backgroundcolour.
  Write ()

// Undocumented (rts).
  Edit (i *uint)
}
