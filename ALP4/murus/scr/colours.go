package scr

// (c) Christian Maurer   v. 120909 - license see murus.go

import (
  "murus/xker"; "murus/col"
)
const
  colC = esc1 + "3%d;4%d"
type
  coloring func (col.Colour)
var (
  colourdepth uint // 0..4 in Byte
  setColours coloring
)


func colours (F, B col.Colour) {
//
  col.Actualize (F, B)
  if underX {
    xker.Colours (F, B) // !
  }
}


func colour (F col.Colour) {
//
  colours (F, col.ScreenB)
}


func nColours () uint {
//
  return col.Number ()
}


func nixfaerben (F col.Colour) {
//
// do nothing
}


func actualize (c coloring) {
//
  setColours = c
}
