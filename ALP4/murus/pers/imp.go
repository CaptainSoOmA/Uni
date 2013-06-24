package pers

// (c) Christian Maurer   v. 130305 - license see murus.go

import (
  . "murus/obj"; "murus/str"; "murus/kbd"
  "murus/col"; "murus/box"; "murus/font"; "murus/pbox"
  "murus/text"; "murus/tval"; "murus/enum"; "murus/day"
)
const (
  lenName = uint(26)
  lenFirstName = uint(15)
  lenShort = lenName + lenFirstName + 2 // ", "
)
type (
  Imp struct {
     surname,
   firstName *text.Imp
          fm *tval.Imp // True == f, False == m
          bd *day.Imp
          ti *enum.Imp
         fmt Format
             }
)
var (
  bx, shbx *box.Imp = box.New (), box.New ()
  pbx = pbox.New()
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
  x.surname = text.New (lenName)
  x.firstName = text.New (lenFirstName)
  x.fm = tval.New ()
  x.fm.SetFormat ("m", "w")
  x.bd = day.New ()
  x.ti = enum.New (enum.Title)
  x.fmt = LongB
  return x
}


func (x *Imp) Empty () bool {
//
  if x.surname.Empty () {
    return x.firstName.Empty ()
  }
  return false
}


func (x *Imp) Clr () {
//
  x.surname.Clr ()
  x.firstName.Clr ()
  x.fm.Clr ()
  x.bd.Clr ()
  x.ti.Clr ()
}


func (x *Imp) Identifiable () bool {
//
  return ! x.surname.Empty () &&
         ! x.firstName.Empty () &&
         ! x.bd.Empty()
}


func (x *Imp) FullAged () bool {
//
  if x.bd.Empty () { return false }
  d:= x.bd.Clone ().(*day.Imp)
  for i:= uint(0); i < 18; i++ {
    d.Inc (day.Yearly)
  }
  return d.Elapsed ()
}


func (x *Imp) Copy (Y Object) {
//
  y:= x.imp (Y)
  x.surname.Copy (y.surname)
  x.firstName.Copy (y.firstName)
  x.fm.Copy (y.fm)
  x.bd.Copy (y.bd)
  x.ti.Copy (y.ti)
}


func (x *Imp) Clone () Object {
//
  y:= New ()
  y.Copy (x)
  return y
}


func (x *Imp) Eq (Y Object) bool {
//
  y:= x.imp (Y)
  g:= x.surname.Eq (y.surname) &&
      x.firstName.Eq (y.firstName) &&
      x.bd.Eq (y.bd)
  switch x.fmt { case LongB, LongTB:
    g = g && x.fm.Eq (y.fm)
  }
  switch x.fmt { case LongT, LongTB:
    g = g && x.ti.Eq (y.ti)
  }
  return g
}


func (x *Imp) Equiv (Y Object) bool {
//
  y:= x.imp (Y)
  if ActualOrder == NameOrder {
    return x.surname.Eq (y.surname) &&
           x.firstName.Eq (y.firstName) &&
           x.bd.Eq (y.bd)
  }
  return x.bd.Eq (y.bd)
}


func (x *Imp) Less (Y Object) bool {
//
  y:= x.imp (Y)
  if ActualOrder == NameOrder {
    if x.surname.Eq (y.surname) {
      if x.firstName.Eq (y.firstName) {
        return x.bd.Less (y.bd)
      }
      return x.firstName.Less (y.firstName)
    }
    return x.surname.Less (y.surname)
  } // ActualOrder == AgeOrder
  if x.bd.Eq (x.imp (Y).bd) {
    if x.surname.Eq (y.surname) {
      return x.firstName.Less (y.firstName)
    }
    return x.surname.Less (y.surname)
  }
  return x.bd.Less (x.imp (Y).bd)
}


func (x *Imp) IsPart (Y Object) bool {
//
  y:= x.imp (Y)
  if ! x.surname.IsPart (y.surname) {
    return false
  }
  if ! x.firstName.IsPart (y.firstName) {
    return false
  }
  if ! x.bd.Empty () && y.bd.Empty () {
    return false
  }
  if ! x.bd.Eq (y.bd) {
    return false
  }
  return true
}


func (x *Imp) SetFormat (f Format) {
//
  if f < NFormats {
    x.fmt = f
  }
}


func (x *Imp) SetColours (f, b col.Colour) {
//
  x.surname.SetColours (f, b)
  x.firstName.SetColours (f, b)
  shbx.Colours (f, b)
  x.fm.SetColours (f, b)
  x.bd.SetColours (f, b)
  x.ti.SetColours (f, b)
}


var
  Sna, Svn, Smw, Sgb, San uint
/* without mask:
          1         2         3         4         5         6         7
01234567890123456789012345678901234567890123456789012345678901234567890123456789
VeryShort    Name, Vorname                                   1 Zeile, 44 Spalten
__________________________, _______________

Name: __________________________ Vorname: _______________ m/w: _ geb.: ________
Anr.: _____________ 

Short, ShortB: Name, Vorname (GebDat)                   1 Zeile,  43 (54) Spalten
__________________________, _______________ (________)

ShortT, ShortTB: Name, Vorname, Anrede (GebDat)
__________________________, _______________, _____________ (________)

with mask:
Long        Name, Vorname, m/w     1 Zeile,  64 Spalten
LongB       Lang, GebDat           1 Zeile,  80 Spalten
LongT       Lang, Anrede           2 Zeilen, 64 Spalten

LongT, LongTB: Name, Vorname, m/w, (geb)              2 Zeilen, 64 (79) Spalten
Name: __________________________ Vorname: _______________ m/w: _ geb.: ________
Anr.: _____________
******************************************************************************/

func (x *Imp) writeMask (Z, S uint) {
//
  switch x.fmt { case Short, ShortB:
    Sna = 0; Svn = 28; Sgb = 44
  default:
    Sna = 6; Svn = 42; Smw = 63; Sgb = 71; San = Sna
  }
  bx.Wd (1)
  bx.ColoursScreen ()
  switch x.fmt { case Short:
    bx.Write (",", Z, S + Svn - 2)
    return
  case ShortB:
    bx.Write (",", Z, S + Svn - 2)
    bx.Write ("(", Z, S + Sgb - 1)
    bx.Write (")", Z, S + Sgb + 8)
    return
  default:
    bx.Wd (5)
    bx.Write ("Name:", Z, S + Sna - 6)
    bx.Wd (8)
    bx.Write ("Vorname:", Z, S + Svn - 9)
    bx.Wd (4)
    bx.Write ("m/w:", Z, S + Smw - 5)
    bx.Wd (5)
  }
  switch x.fmt { case LongB, LongTB:
    bx.Wd (5)
    bx.Write ("geb.:", Z, S + Sgb - 6)
  }
  switch x.fmt { case LongT, LongTB:
    bx.Wd (5)
    bx.Write ("Anr.:", Z + 1, S + San - 6)
  }
}


func (x *Imp) String () string {
//
  n, f:= x.surname.String(), x.firstName.String()
  str.RemSpaces (&n); str.RemSpaces (&f)
  nf, fn, b:= n + ", " + f, f + " " + n, x.bd.String()
  switch x.fmt {
  case VeryShort, Short:
    return nf
  case ShortB:
    return nf + " (" + b + ")"
  case ShortT:
    if ! x.ti.Empty() { nf = x.ti.String() + " " + nf }
    return nf
  case ShortTB:
    if ! x.ti.Empty() { fn = x.ti.String() + " " + fn }
    return nf + " (" + b + ")"
/* mit Maske:
  case Long:    // Name, Vorname, m/w     1 Zeile,  64 Spalten
    return ""
  case LongB:   // lang, GebDat           1 Zeile,  80 Spalten
    return ""
  case LongT:   // lang, Anrede           2 Zeilen, 64 Spalten
    return ""
  case LongTB:  // lang, GebDat, Anrede   2 Zeilen, 80 Spalten
    return ""
*/
  }
  return nf
}


func (x *Imp) Defined (s string) bool { // trivial version, bette TODO
//
  if ! x.surname.Defined (s[:26]) { return false }
  if ! x.firstName.Defined (s[26:41]) { return false }
  if ! x.fm.Defined (s[41:42]) { return false }
  if ! x.bd.Defined (s[42:50]) { return false }
//  if ! x.title.Defined (s[49:]) { return false }
  return true
}


func (x *Imp) Write (Z, S uint) {
//
  if x.fmt == VeryShort {
    shbx.Write (x.String(), Z, S)
    return
  }
  x.writeMask (Z, S)
  x.surname.Write (Z, S + Sna)
  x.firstName.Write (Z, S + Svn)
  switch x.fmt { case Short, ShortB:
  default:
    x.fm.Write (Z, S + Smw)
  }
  switch x.fmt { case ShortB, LongB, LongTB:
    x.bd.Write (Z, S + Sgb)
  }
  switch x.fmt { case LongT, LongTB:
    x.ti.Write (Z + 1, S + San)
  }
}


func (x *Imp) Edit (Z, S uint) {
//
  x.Write (Z, S)
  if x.fmt == VeryShort { return }
  var i, d uint
  if kbd.LastCommand (&d) == kbd.Up {
    i = 4
  } else {
    i = 0
  }
  loop: for {
    switch i { case 0:
      x.surname.Edit (Z, S + Sna)
    case 1:
      x.firstName.Edit (Z, S + Svn)
    case 2:
      switch x.fmt { case Short, ShortB:
        ;
      default:
        x.fm.Edit (Z, S + Smw)
      }
    case 3:
      switch x.fmt { case ShortB, LongB, LongTB:
        x.bd.Edit (Z, S + Sgb)
      }
    case 4:
      switch x.fmt { case LongT, LongTB:
        x.ti.Edit (Z + 1, S + San)
      }
    }
    switch kbd.LastCommand (&d) { case kbd.Esc:
      break loop
    case kbd.Enter:
      if d == 0 {
        if i < 4 { i++ } else { break loop }
      } else {
        break loop
      }
    case kbd.Down, kbd.Right:
      if i < 4 { i++ } else { break loop }
    case kbd.Up, kbd.Left:
      if i > 0 { i-- } else { break loop }
    case kbd.LookFor:
      break loop
    }
  }
}


func (x *Imp) SetFont (f font.Font) {
//
  x.surname.SetFont (f)
  x.firstName.SetFont (f)
  x.fm.SetFont (f)
  x.bd.SetFont (f)
  x.ti.SetFont (f)
}


func (x *Imp) printMask (Z, S uint) {
//
  switch x.fmt { case Short, ShortB:
    Sna = 0; Svn = 28; Sgb = 44
  default:
    Sna = 6; Svn = 42; Smw = 63; Sgb = 71; San = Sna
  }
  switch x.fmt { case Short:
    pbx.Print (",", Z, S + Svn - 2)
    return
  case ShortB:
    pbx.Print (",", Z, S + Svn - 2)
    pbx.Print ("(", Z, S + Sgb - 1)
    pbx.Print (")", Z, S + Sgb + 8)
    return
  default:
    pbx.Print ("Name:", Z, S + Sna - 6)
    pbx.Print ("Vorname:", Z, S + Svn - 9)
    pbx.Print ("m/w:", Z, S + Smw - 5)
  }
  switch x.fmt { case LongB, LongTB:
    pbx.Print ("geb.:", Z, S + Sgb - 6)
  }
  switch x.fmt { case LongT, LongTB:
    pbx.Print ("Anr.:", Z + 1, S + San - 6)
  }
}


func (x *Imp) Print (Z, S uint) {
//
  x.printMask (Z, S)
  if x.fmt == VeryShort {
    pbx.Print (x.String(), Z, S)
    return
  }
  x.surname.SetFont (font.Bold)
  x.surname.Print (Z, S + Sna)
  x.firstName.SetFont (font.Bold)
  x.firstName.Print (Z, S + Svn)
  switch x.fmt { case Short, ShortB:
  default:
    x.fm.Print (Z, S + Smw)
  }
  switch x.fmt { case ShortB, LongB, LongTB:
    x.bd.Print (Z, S + Sgb)
  }
  switch x.fmt { case LongT, LongTB:
    x.ti.Print (Z + 1, S + San)
  }
}


func (x *Imp) Codelen () uint {
//
  return lenName +         // 26
         lenFirstName +    // 15
         x.bd.Codelen () + //  2
         x.fm.Codelen () + //  1
         x.ti.Codelen ()   //  1
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, x.Codelen())
  i, a:= uint(0), lenName
  copy (b[i:i+a], x.surname.Encode())
  i += a
  a = lenFirstName
  copy (b[i:i+a], x.firstName.Encode())
  i += a
  a = x.bd.Codelen()
  copy (b[i:i+a], x.bd.Encode())
  i += a
  a = x.fm.Codelen()
  copy (b[i:i+a], x.fm.Encode())
  i += a
  a = x.ti.Codelen()
  copy (b[i:i+a], x.ti.Encode())
  return b
}


func (x *Imp) Decode (b []byte) {
//
  i, a:= uint(0), lenName
  x.surname.Decode (b[i:i+a])
  i += a
  a = lenFirstName
  x.firstName.Decode (b[i:i+a])
  i += a
  a = x.bd.Codelen()
  x.bd.Decode (b[i:i+a])
  i += a
  a = x.fm.Codelen()
  x.fm.Decode (b[i:i+a])
  i += a
  a = x.ti.Codelen()
  x.ti.Decode (b[i:i+a])
}


func RotOrder () {
//
  ActualOrder = Order (1 - ActualOrder)
}


func init () {
//
  shbx.Wd (lenShort)
  var _ Person = New()
//  init1 ()
}
