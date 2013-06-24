package fuday

// (c) Christian Maurer   v. 130309 - license see murus.go

import (
  "murus/nat"; "murus/day"
)
type
  Imp struct {
             *day.Imp
             }
var (
  tmp, tmp1, tmp2 *day.Imp = day.New(), day.New(), day.New()
)


func New () *Imp {
//
  x:= new (Imp)
  x.Imp = day.New ()
  x.Actualize ()
  return x
}


func (x *Imp) set (d *day.Imp) {
//
  x.Imp.Copy (d)
}


func (x *Imp) summer () bool {
//
  y:= x.Year ()
  tmp.Set (1, 4, y)
  if x.Imp.Less (tmp) {
    return false
  }
  tmp.Set (1, 10, y)
  if x.Imp.Less (tmp) {
    return true
  }
  return false
}


func (x *Imp) semester (b, e *day.Imp) {
//
  y:= x.Year () % 100
  tmp.Set (1, 4, y)
  if x.Imp.Less (tmp) {
    b.Set (1, 10, y); b.Dec (day.Yearly)
    e.Set (1,  4, y)
  } else {
    tmp.Set (1, 10, y)
    if x.Imp.Less (tmp) {
      b.Set (1,  4, y)
      e.Set (1, 10, y)
    } else {
      b.Set (1, 10, y)
      e.Set (1,  4, y); e.Inc (day.Yearly)
    }
  }
  e.Dec (day.Daily)
}


func (x *Imp) lectures (b, e *day.Imp) {
//
  x.semester (b, e)
  y:= x.Year () % 100
  w:= uint(14)
  if x.summer () {
    b.Set (14, 4, y)
  } else {
    b.Set (18, 10, y)
    w += 2 + 2 // Weihnachtsferien
  }
  for ! b.IsBeginning (day.Weekly) {
    b.Dec (day.Daily)
  }
  e.Copy (b)
  for i:= uint(0); i < w; i++ {
    e.Inc (day.Weekly)
  }
  e.Dec (day.Daily)
  e.Dec (day.Daily) // Saturday
}


func (x *Imp) string_ () string {
//
  y:= x.Year () % 100
  tmp.Set (1, 4, y)
  s:= nat.StringFmt (y, 2, true)
  if x.Imp.Less (tmp) {
    tmp.Dec (day.Yearly)
    return "WS " + nat.StringFmt (tmp.Year() % 100, 2, true) + "/" + s
  }
  tmp.Set (1, 10, y)
  if x.Imp.Less (tmp) {
    return "SS " + s
  }
  tmp.Inc (day.Yearly)
  return "WS" + s + "/" + nat.StringFmt (tmp.Year() % 100, 2, true)
}


func (x *Imp) lectureDay (d *day.Imp) bool {
//
  summer:= x.summer ()
  x.lectures (tmp1, tmp2)
  if d.Less (tmp1) || tmp2.Less (d) || d.IsHoliday () {
    return false
  }
  if ! summer {
    tmp.Copy (tmp1) // Beginn Akademische Ferien:
    for i:= uint(0); i < 10; i++ { tmp.Inc (day.Weekly) }
// tmp.Write (10, 0)
    tmp2.Copy (tmp) // Vorlesungsbeginn Januar:
    for i:= uint(0); i < 2; i++ { tmp2.Inc (day.Weekly) }
// tmp2.Write (10, 0)
    if tmp.Eq(d) || tmp.Less (d) && d.Less (tmp2) {
      return false
    }
  }
  return true
}


func (x *Imp) numWeeks () uint {
//
  if x.summer () {
    return 14
  }
  return 16
}


func (x *Imp) monday (d *day.Imp, n uint) {
//
  w:= x.numWeeks ()
  if n == 0 || n > w {
    d.Clr()
    return
  }
  x.lectures (d, tmp)
  for i:= uint(0); i + 1 < n; i++ {
    d.Inc (day.Weekly)
  }
  if ! x.summer () && n > 10 { // Akademische Ferien
    d.Inc (day.Weekly)
    d.Inc (day.Weekly)
  }
}
