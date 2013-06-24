package v

// (c) Christian Maurer - license see murus.go

import (
  "murus/ker"; "murus/col"; "murus/errh"; "murus/day"
)
const ( // v.
  yy = 13
  mm =  5
  dd = 27
)
var
  v *day.Imp = day.New ()


func Colours () (col.Colour, col.Colour, col.Colour) {
//
  return col.Yellow, col.LightYellow, col.DarkGreen
  return col.MurusF, col.Colour3 (0, 16, 128), col.MurusB
}


func String () string {
//
  return v.String ()
}


func Want (y, m, d uint) {
//
  wanted:= day.New ()
  wanted.SetFormat (day.Yymmdd)
  if wanted.Set (d, m, 2000 + y) {
    if v.Less (wanted) {
      errh.Error ("Your murus " + v.String() + " is outdated. You need " + wanted.String () + ".", 0)
      ker.Halt (-1)
    }
  } else {
    ker.Panic ("parameters for v.Want are nonsense")
  }
}


func init () {
//
  v.Set (dd, mm, 2000 + yy)
  v.SetFormat (day.Yymmdd)
}
