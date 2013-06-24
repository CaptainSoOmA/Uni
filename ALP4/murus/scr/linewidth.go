package scr

// (c) Christian Maurer   v. 120909 - license see murus.go

import
  "murus/xker"
type
  application func (int, int)
var
  actualLinewidth Linewidth


func setLinewidth (w Linewidth) {
//
  actualLinewidth = w
  if underX {
    xker.SetLinewidth (1 + uint(w))
  } else {
    // TODO
  }
}
