package date

// (c) Christian Maurer   v. 130115 - license see murus.go

import (
  . "murus/obj"; "murus/str"
  "murus/font"; "murus/prt"
  "murus/col"; "murus/scr"
  "murus/day"; "murus/clk"
)
type
  Imp struct {
         day *day.Imp
        time *clk.Imp
             }
const
  separator = ','
var
  one, two, three, lastTime *clk.Imp = clk.New (), clk.New (), clk.New (), clk.New ()


func (x *Imp) imp (Y Object) *Imp {
//
  y, ok:= Y.(*Imp)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}


func New () *Imp {
//
  return &Imp { day.New(), clk.New() }
}


func (x *Imp) Set (d *day.Imp, t *clk.Imp) {
//
  x.day = d.Clone().(*day.Imp)
  x.time = t.Clone().(*clk.Imp)
}


func (x *Imp) Day () *day.Imp {
//
  return x.day.Clone().(*day.Imp)
}


func (x *Imp) Time () *clk.Imp {
//
  return x.time.Clone().(*clk.Imp)
}


func (x *Imp) Normal () bool {
//
  if x.day.Normal () {
    if x.day.Normal1 () { // last Sunday in October
      return two.Eq (x.time) || two.Less (x.time)
    }
    return true
  }
  if x.day.Normal1 () { // last Sunday in March
    return x.time.Less (two)
  }
  return false
}


func (x *Imp) Empty () bool {
//
  return x.day.Empty() || x.time.Empty()
}


func (x *Imp) Clr () {
//
  x.day.Clr()
  x.time.Clr()
}


func (x *Imp) Eq (Y Object) bool {
//
  y:= x.imp (Y)
  return x.day.Eq (y.day) && x.time.Eq (y.time)
}


func (x *Imp) Less (Y Object) bool {
//
  return false
}


func (x *Imp) SetColours (f, b col.Colour) {
//
  x.day.SetColours (f, b)
  x.time.SetColours (f, b)
}


func (x *Imp) SetFormat (d, c Format) {
//
  if d < day.NFormats && c < clk.NFormats {
    x.day.SetFormat (d)
    x.time.SetFormat (c)
  }
}


func (x *Imp) writeMask (l, c uint) {
//
  scr.Write1 (separator, l, c + 10) // TODO depends on Format
}


func (x *Imp) Write (l, c uint) {
//
  x.writeMask (l, c)
  x.day.Write (l, c)
  x.time.Write (l, c + 10 + 1 + 1) // TODO depends on Format
}


func (x *Imp) Edit (l, c uint) {
//
  x.day.Edit (l, c)
  x.time.Write (l, c + 10 + 1 + 1) // TODO depends on Format
}


func (x *Imp) SetFont (f font.Font) {
//
  x.day.SetFont (f)
  x.time.SetFont (f)
}


func (x *Imp) printMask (l, c uint) {
//
  prt.Print1 (separator, l, c + 10, font.Roman) // TODO depends on Format
}


func (x *Imp) Print (l, c uint) {
//
  x.printMask (l, c)
  x.day.Print (l, c)
  x.time.Print (l, c + 12) // TODO depends on Format
}


func (x *Imp) String () string {
//
  return x.day.String() + string(separator) + " " + x.time.String()
}


func (x *Imp) Defined (s string) bool {
//
  var p uint
  if str.Contains (s, separator, &p) {
    return x.day.Defined (s[:p]) && x.time.Defined (s[p+1:])
  }
  x.Clr()
  return false
}


func (x *Imp) Copy (Y Object) {
//
  y:= x.imp (Y)
  x.day.Copy (y.day)
  x.time.Copy (y.time)
}


func (x *Imp) Clone () Object {
//
  y:= New ()
  y.Copy (x)
  return y
}


func (x *Imp) Actualize (x1 *Imp) {
//
  if x.Empty() || x1.Empty() { return }
  x1.Copy (x)
  if x.Normal () {
    return
  }
  x1.time.Inc (one)
  if x1.time.Less (one) {
    x1.day.Inc (day.Daily)
  }
}


func (x *Imp) Normalize () {
//
  if x.Empty() { return }
  if x.day.Normal () {
    if x.day.Normal1 () { // TODO get rid of Leq
/*
      if two.Leq (x.time) && x.time.Less (three) {
      // kritischer Fall 2 <= Zeit < 3 (Stunden 2A und 2B): wenn
      // wenn Zeit Leq Zeit beim vorigen Aufruf war, ist es 2B,
      // d.h. es wird nicht mehr eine Stunde zurückgestellt
        if x.time.Leq (lastTime) {
          return
        }
        lastTime.Copy (x.time)
      } else {
     // für eventuelle weitere Aufrufe im gleichen Programmlauf:
        lastTime.Set (0, 0, 0)
        if three.Leq (x.time) {
          return
        }
      }
*/
    } else {
      return
    }
  } else if x.day.Normal1() && x.time.Less (three) {
    return
  }
  if x.time.Less (one) {
    x.day.Dec (day.Daily)
  }
  x.time.Dec (one)
}


func (x *Imp) Inc (dt *clk.Imp) {
//
  x.time.Inc (dt)
  if x.time.Less (dt) {
    x.day.Inc (day.Daily)
  }
}


func (x *Imp) Dec (dt *clk.Imp) {
//
  if x.time.Less (dt) {
    x.day.Dec (day.Daily)
  }
  x.time.Dec (dt)
}


func (x *Imp) Codelen () uint {
//
  return x.day.Codelen() + x.time.Codelen()
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, x.Codelen())
  a:= x.day.Codelen()
  copy (b[:a], x.day.Encode())
  copy (b[a:], x.time.Encode())
  return b
}


func (x *Imp) Decode (b []byte) {
//
  a:= x.day.Codelen()
  x.day.Decode (b[:a])
  x.time.Decode (b[a:])
}


func init () {
//
  one.Set (1, 0, 0)
  two.Set (2, 0, 0)
  three.Set (3, 0, 0)
  lastTime.Set (0, 0, 0)
  var _ Object = New()
  var _ Editor = New()
}
