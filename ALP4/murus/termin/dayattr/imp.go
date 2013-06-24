package dayattr

// (c) Christian Maurer   v. 130526 - license see murus.go

import (
  . "murus/obj"
  "murus/font"; "murus/col"; "murus/scr"; "murus/box"
  "murus/day"
  "murus/termin/pdays"
)
const (
  pack = "dayattr"
  length = 8 // of names
)
type
  attribute byte; const (suchwort = iota; ferien; casetta; nAttrs)
var (
  name [nAttrs]string
  set[nAttrs]*pdays.Imp
  actual attribute
  bx *box.Imp
  workdayAF, workdayAB, holidayAF, holidayAB col.Colour
)


func Normalize () {
//
  actual = attribute (0)
}


func Change (w bool) {
//
  if w {
    if actual + 1 < nAttrs {
      actual++
    } else {
      actual = attribute (0)
    }
  } else if actual > 0 {
    actual--
  } else {
    actual = attribute (nAttrs - 1)
  }
}


func write (a attribute, visible bool, Z, S uint) {
//
  if visible {
    bx.Write (name [a], Z, S)
  } else {
    // bx.Clr (Z, S)
    scr.Clr (Z, S, length, 1)
  }
}


func WriteActual (Z, S uint) {
//
  write (actual, true, Z, S)
}


func hasAttribute (d *day.Imp) bool {
//
  return set[actual].Ex (d)
}


func Actualize (d *day.Imp, gesetzt bool) {
//
  if true /* actual > 0 */ {
    if gesetzt {
      set[actual].Ins (d)
    } else {
      set[actual].Del (d)
    }
  }
}


func WriteAll (d *day.Imp, Z, S uint) {
//
  for a:= attribute (0); a < nAttrs; a++ {
    write (a, set[a].Ex (d), Z, S)
    S += length + 1
  }
}


func Clr () {
//
  set[0].Clr ()
}


func Attrib (a Any) {
//
  d, ok:= a.(*day.Imp)
  if ! ok { return }
  if set[actual].Ex (d) {
    if d.IsHoliday () {
      d.SetColours (holidayAF, holidayAB)
      d.SetFont (font.Bold)
    } else {
      d.SetColours (workdayAF, workdayAB)
      d.SetFont (font.Bold)
    }
  } else if d.IsHoliday () {
    d.SetColours (day.HolidayF, day.HolidayB)
    d.SetFont (font.Bold)
  } else {
    d.SetColours (day.WeekdayF, day.WeekdayB)
    d.SetFont (font.Roman)
  }
}


func Terminate () {
//
  for a:= 0; a < nAttrs; a++ {
    set[a].Terminate ()
  }
}


func init() {
//
  name[suchwort] = "Suchwort"
  name[ferien]   = "Ferien  "
  name[casetta]  = "Casetta "
  bx = box.New()
  bx.Wd (length)
  workdayAF, workdayAB = col.LightWhite, col.Blue
  holidayAF, holidayAB = col.LightWhite, col.Red
  bx.Colours (workdayAF, workdayAB)
//  day.Attribute (Attrib)
  for a:= 0; a < nAttrs; a++ {
    set[a] = pdays.New ()
    set[a].Name (name[a])
  }
}
