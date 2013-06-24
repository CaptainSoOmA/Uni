package clk

// (c) Christian Maurer   v. 130309 - license see murus.go

import (
  . "murus/ker"; . "murus/obj"; "murus/str"
  "murus/col"; "murus/scr"; "murus/box"; "murus/font"; "murus/pbox"
  "murus/errh"; "murus/nat"
)
const (
  pack = "clk"
  maxlength = 8 // maximal Formatlength for "Hh_mm_ss"
  _MS = 60 // m / h = s / m
  _H = 24 // hours per day
  maxTimeCode = _H * _MS * _MS
)
type
  Imp struct {
        hour, // <= _H // 24 for the empty clocktime
      minute, // < _MS
      second uint // < _MS
         fmt Format
      cF, cB col.Colour
        font font.Font
             }
var (
  textlength [NFormats]uint
  currentTime, clock *Imp = New(), New()
  bx, clockbx *box.Imp = box.New(), box.New()
  line, column uint
  pbx *pbox.Imp = pbox.New()
)


func (x *Imp) imp (Y Object) *Imp {
//
  y, ok:= Y.(*Imp)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}


func New () *Imp {
//
  x:= new (Imp)
  x.hour = _H
  x.cF, x.cB = col.ScreenF, col.ScreenB
  x.fmt = Hh_mm
  return x
}


func (x *Imp) Actualize () {
//
  x.hour, x.minute, x.second = ActualizeTime ()
}


func (x *Imp) Empty () bool {
//
  return x.hour == _H
}


func (x *Imp) Clr () {
//
  x.hour = _H
  x.minute, x.second = 0, 0
}


func (x *Imp) Copy (Y Object) {
//
  y:= x.imp (Y)
  x.hour, x.minute, x.second = y.hour, y.minute, y.second
}


func (x *Imp) Clone () Object {
//
  y:= New ()
  y.Copy (x)
  return y
}


func (x *Imp) internalCode () uint {
//
  c:= (_MS * x.hour + x.minute) * _MS + x.second
  if c > maxTimeCode { Panic ("jaul") }
  return c
}


func (x *Imp) Eq (Y Object) bool {
//
  return x.internalCode () == x.imp (Y).internalCode ()
}


func less (C, C1 uint) bool {
//
  if C == maxTimeCode {
    return C1 != maxTimeCode
  } else if C1 == maxTimeCode {
    return false
  }
  return C < C1
}


func (x *Imp) Less (Y Object) bool {
//
  return x.internalCode () < x.imp (Y).internalCode ()
}


func (x *Imp) Elapsed () bool {
//
  currentTime.Actualize ()
  return x.internalCode () < currentTime.internalCode()
}


func (x *Imp) Distance (Y Object) uint {
//
  y:= x.imp (Y)
  if x.Empty () || y.Empty () { return MaxNat }
  c, d:= x.internalCode (), y.internalCode ()
  if d > c {
    return d - c
  }
  return c - d
}


func (x *Imp) NSeconds () uint {
//
  if x.Empty () { return MaxNat }
  return x.internalCode ()
}


func (x *Imp) Hours () uint {
//
  if x.Empty () { return _H }
  return x.hour
}


func (x *Imp) Minutes () uint {
//
  if x.Empty () { return _MS }
  return x.minute
}


func (x *Imp) Seconds () uint {
//
  if x.Empty () { return _MS }
  return x.second
}


func (x *Imp) Inc (Y Object) {
//
  y:= x.imp (Y)
  if x.Empty () || y.Empty () { return }
  c:= x.internalCode ()
  c+= y.internalCode ()
  c = c % maxTimeCode
  x.second = c % _MS
  c = c / _MS
  x.minute, x.hour = c % _MS, c / _MS
}


func (x *Imp) Dec (Y Object) {
//
  y:= x.imp (Y)
  if x.Empty () || y.Empty () { return }
  c:= x.internalCode () + maxTimeCode
  c-= y.internalCode ()
  c = c % maxTimeCode
  x.second = c % _MS
  c = c / _MS
  x.minute, x.hour = c % _MS, c / _MS
}


func (x *Imp) SetFormat (f Format) {
//
  if f < NFormats {
    x.fmt = f
  }
}


func (x *Imp) SetColours (f, b col.Colour) {
//
  x.cF, x.cB = f, b
}


func (x *Imp) String() string {
//
  if x.Empty () {
    return str.Clr (textlength[x.fmt])
  }
  s:= ""
  if x.fmt <= Hh_mm_ss { // Hh_mm
    s = nat.StringFmt (x.hour, 2, true) + "."
  }
  s += nat.StringFmt (x.minute, 2, true)
  if x.fmt >= Hh_mm_ss {
    s += ":" + nat.StringFmt (x.second, 2, true)
  }
  return s
}


func (x *Imp) defined (h, m, s uint) bool {
//
  if h < _H { x.hour = h } else { return false }
  if m < _MS { x.minute = m } else { return false }
  if s < _MS { x.second = s } else { return false }
  x.hour, x.minute, x.second = h, m, s
  return true
}


func (x *Imp) Defined (t string) bool {
//
  x.Clr ()
  if str.Empty (t) { return true }
  n, ss, P, L:= nat.DigitSequences (t)
  if n == 0 || n > 3 { return false }
  if n == 3 {
    if x.fmt == Hh_mm { return false }
  }
  if L[0] >= textlength[x.fmt] { return false }
  h, m, s:= uint(0), uint(0), uint(0)
  if n == 1 {
    if ! nat.Defined (&h, str.Part (t, P[0], 2)) { return false }
    if L[0] > 2 {
      if ! nat.Defined (&m, str.Part (t, P[0] + 2, 2)) { return false }
      if L[0] > 4 {
        if ! nat.Defined (&s, str.Part (t, P[0] + 4, 2)) { return false }
      }
    }
  } else {
    if ! nat.Defined (&h, ss[0]) { return false }
    if ! nat.Defined (&m, ss[1]) { return false }
    if n == 2 && x.fmt == Mm_ss {
      s, m, h = m, h, 0
    }
    if n == 3 {
      if ! nat.Defined (&s, ss[2]) { return false }
    }
  }
  return x.defined (h, m, s)
}


func (x *Imp) Set (h, m, s uint) {
//
  x.Clr ()
  if h < _H && m < _MS && s < _MS {
    x.hour, x.minute, x.second = h, m, s
  }
}


func (x *Imp) Write (l, c uint) {
//
  bx.Wd (textlength[x.fmt])
  bx.Colours (x.cF, x.cB)
  bx.Write (x.String(), l, c)
}


func (x *Imp) Edit (l, c uint) {
//
  x.Write (l, c)
  s:= x.String()
  err:= uint(0)
  for {
    bx.Edit (&s, l, c)
    if x.Defined (s) {
      s = x.String ()
      bx.Write (s, l, c)
      return
    } else {
      err ++
      switch err { case 1: // --> errh.WriteError
        errh.Error ("Die Uhrzeitangabe ist unverständlich!", 0) // , l + 1, c)
      case 2:
        errh.Error ("Bitte wiederholen Sie die Uhrzeitangabe, sie ist immer noch unverständlich !", 0) // , l + 1, c)
      case 3:
        errh.Error ("Jetzt passen Sie doch gefälligst auf !", 0) // l + 1, c)
      case 4:
        errh.Error ("Was soll der Quatsch? Ist das eine Uhrzeit ?", 0) // l + 1, c)
      case 5:
        errh.Error ("Schaffen Sie es nicht, eine Uhrzeit richtig einzugeben ?", 0) // l + 1, c)
      default:
        errh.Error ("Vergessen Sie's ...", 0) // , l + 1, c)
        x.Actualize ()
        return
      }
    }
  }
}


func (x *Imp) SetFont (f font.Font) {
//
  x.font = f
}


func (x *Imp) Print (l, c uint) {
//
  pbx.SetFont (x.font)
  pbx.Print (x.String(), l, c)
}


func (x *Imp) Encode () []byte {
//
  bs:= make ([]byte, x.Codelen())
  copy (bs, Encode (x.internalCode()))
  return bs
}


func (x *Imp) Codelen () uint {
//
  return Codelen (uint(0))
}


func (x *Imp) Decode (bs []byte) {
//
  n:= Decode (uint(0), bs).(uint)
  x.second = n % _MS
  n /= _MS
  x.minute = n % _MS
  n /= _MS
  if n > _H {
    x.hour = _H
  } else {
    x.hour = n
  }
}


func SetAttributes (l, c uint, f, b col.Colour) {
//
  line, column = l, c
  clockbx.Colours (f, b)
}


func Show () {
//
  if line >= scr.NLines() {
    SetAttributes (Zero, scr.NColumns () - textlength[clock.fmt], col.HintF, col.HintB)
  }
  for {
    clock.Actualize ()
    clockbx.Write (clock.String(), line, column)
    Sleep (1) // not precise, but good enough for practical purposes
              // more precise would be: sleep until AlarmClock rings
  }
}


func init () {
//
  textlength[Hh_mm] =    5
  textlength[Hh_mm_ss] = maxlength
  textlength[Mm_ss] =    5
  clock.fmt = Hh_mm_ss
  line = 1000
//  SetAttributes (Zero, scr.NColumns () - textlength[clock.fmt], col.HintF, col.HintB)
}
