package pbox

// (c) Christian Maurer   v. 120909 - license see murus.go

import
  "murus/font"
type
  Printbox interface {

  SetFont (f font.Font)
  Font () font.Font
  Print (s string, l, c uint)
  PageReady ()
}
