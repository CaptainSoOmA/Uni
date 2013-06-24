package pbox

// (c) Christian Maurer   v. 121119 - license see murus.go

import (
  "murus/str"
  "murus/font"; "murus/prt"
)
type
  Imp struct {
        font font.Font
             }


func NLines () uint {
//
  return prt.NLines()
}


func NColumns () uint {
//
  return prt.NColumns()
}


func New () *Imp {
//
  return &Imp { font.Roman }
}


func (x *Imp) SetFont (f font.Font) {
//
  x.font = f
}


func (x *Imp) Font () font.Font {
//
  return x.font
}


func (x *Imp) Print (s string, l, c uint) {
//
  if l >= prt.NLines() || c >= prt.NColumns() {
    return
  }
  str.RemSpaces (&s)
  if len (s) == 0 {
    return
  }
  prt.Print (s, l, c, x.font)
}


func (x *Imp) PageReady () {
//
  prt.GoPrint()
}
