package page

// (c) Christian Maurer   v. 130510 - license see murus.go

import (
  . "murus/obj"
  "murus/col"; "murus/font"
  "murus/day"
  "murus/termin/dayattr"; "murus/termin/seq"
)
type
  Imp struct {
             *day.Imp
         fmt day.Period
        list seq.Sequence // *seq.Imp
             }


func (x *Imp) imp (Y Object) *Imp {
//
  y, ok:= Y.(*Imp)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}


func New () *Imp {
//
  x:= new (Imp)
  x.Imp = day.New ()
  x.list = seq.New ()
  return x
}


func (x *Imp) Empty () bool {
//
  return x.list.Empty ()
}


func (x *Imp) Clr () {
//
  x.Imp.Clr ()
  x.list.Clr ()
}


func (x *Imp) Eq (Y Object) bool {
//
  y:= x.imp (Y)
  return x.Imp.Eq (y.Imp) &&
         x.list.Eq (y.list)
}


func (x *Imp) Copy (Y Object) {
//
  y:= x.imp (Y)
  x.Imp.Copy (y.Imp)
  x.list.Copy (y.list)
}


func (x *Imp) Clone () Object {
//
  y:= New ()
  y.Copy (x)
  return y
}


func (x *Imp) Set (d *day.Imp) {
//
  x.Imp = d.Clone ().(*day.Imp)
  x.list.Clr ()
}


func (x *Imp) Less (Y Object) bool {
//
  return x.Imp.Less (x.imp (Y).Imp)
}


func (x *Imp) HasWord () bool {
//
  return x.list.HasWord ()
}


func (x *Imp) SetFormat (p day.Period) {
//
  x.fmt = p
  x.list.SetFormat (p)
}


func (x *Imp) SetColours (f, b col.Colour) {
//
// dummy
}


func (x *Imp) Write (Z, S uint) {
//
  if x.Imp.IsHoliday () {
    x.Imp.SetColours (day.HolidayF, day.HolidayB)
  } else {
    x.Imp.SetColours (day.WeekdayF, day.WeekdayB)
  }
  switch x.fmt { case day.Daily:
    x.Imp.SetFormat (day.WD)
    x.Imp.Write (Z, S)
    x.Imp.SetFormat (day.Dd_mm_yyyy)
    x.Imp.Write (Z, S + 11)
    dayattr.WriteAll (x.Imp, Z, S + 22)
    x.list.Write (Z + 2, S)
  case day.Weekly:
    x.list.Write (Z + 1, S)
  case day.Monthly:
    x.list.Write (Z, S)
  }
}


func (x *Imp) Edit (Z, S uint) {
//
  x.Write (Z, S)
  x.list.Edit (Z + 2, S)
}


func (x *Imp) SetFont (f font.Font) {
//
// dummy
}


func (x *Imp) Print (Z, S uint) {
//
  if x.fmt == day.Daily {
    x.Imp.SetFormat (day.WD)
    x.Imp.Print (Z, S)
    x.Imp.SetFormat (day.Dd_mm_yyyy)
    x.Imp.Print (Z, S + 11)
    x.list.Print (Z + 2, S)
  }
}


func (x *Imp) Codelen () uint {
//
  return x.Imp.Codelen () +
         x.list.Codelen ()
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, x.Codelen())
  a:= x.Imp.Codelen()
  copy (b[:a], x.Imp.Encode ())
  copy (b[a:], x.list.Encode ())
  return b
}


func (x *Imp) Decode (b []byte) {
//
  a:= x.Imp.Codelen()
  x.Imp.Decode (b[:a])
  x.list.Decode (b[a:])
}


func (x *Imp) Day () Any {
//
  return x.Imp.Clone ().(*day.Imp)
}


func (x *Imp) Terminate () {
//
  dayattr.Terminate ()
}


func init () { var _ Page = New() }
