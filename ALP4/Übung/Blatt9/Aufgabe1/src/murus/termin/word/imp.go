package word

// (c) Christian Maurer   v. 130526 - license see murus.go

import (
  . "murus/obj"; "murus/str"
  "murus/col"; "murus/box"; "murus/font"; "murus/pbox"
  "murus/pseq"; "murus/files"
)
const
  length = 12
type
  Imp struct {
        text string
      marked bool
             }
var (
  cF, cB, cMF, cMB, cFF, cFB, cFMF, cFMB col.Colour =
    col.Black, col.White, col.Black, col.LightWhite,
    col.LightWhite, col.Red, col.Red, col.LightWhite
  bx *box.Imp = box.New()
  pbx *pbox.Imp = pbox.New()
  actual *Imp = New ()
  file *pseq.Imp
)


func New () * Imp {
//
  x:= new (Imp)
  x.Clr ()
  return x
}


func (x *Imp) Empty () bool {
//
  return str.Empty (x.text)
}


func (x *Imp) Clr () {
//
  x.text = str.Clr (length)
  x.marked = false
}


func (x *Imp) Copy (X Object) {
//
  x1, ok:= X.(*Imp)
  if ! ok { return }
  x.text = x1.text
  x.marked = x1.marked
}


func (x *Imp) Clone () Object {
//
  x1:= New ()
  x1.Copy (x)
  return x1
}


func (x *Imp) Eq (X Object) bool {
//
  x1, ok:= X.(*Imp)
  if ! ok { return false }
  return x.text == x1.text
}


func (x *Imp) Ok () bool {
//
  var i uint
  return str.IsEquivPart (actual.text, x.text, &i) &&
         ! str.Empty (actual.text)
}


func (x *Imp) Less (X Object) bool {
//
  x1, ok:= X.(*Imp)
  if ! ok { return false }
  return str.Less (x.text, x1.text)
}


func (x *Imp) String () string {
//
  return x.text
}


func (x *Imp) Defined (s string) bool {
//
  x.text = s
  str.Norm (&x.text, length)
  x.marked = false
  return true
}


func (x *Imp) Mark (m bool) {
//
  x.marked = m
}


func (x *Imp) Marked () bool {
//
  return x.marked
}


func (x *Imp) SetColours (f, b col.Colour) {
//
}


func (x *Imp) Write (Z, S uint) {
//
  if x.Ok () {
    if x.marked {
      bx.Colours (cFMF, cFMB)
    } else {
      bx.Colours (cFF, cFB)
    }
  } else {
    if x.marked {
      bx.Colours (cMF, cMB)
    } else {
      bx.Colours (cF, cB)
    }
  }
  bx.Write (x.text, Z, S)
}


func (x *Imp) Edit (Z, S uint) {
//
  x.Write (Z, S)
  bx.Edit (&x.text, Z, S)
  str.Move (&x.text, true)
  x.Write (Z, S)
}


func EditActual (Z, S uint) {
//
  actual.Write (Z, S)
  bx.Edit (&actual.text, Z, S)
  str.Move (&actual.text, true)
  str.RemSpaces (&actual.text)
  str.Norm (&actual.text, length)
  actual.Write (Z, S)
  file.Seek (0)
  file.Put (actual.text)
}


func (x *Imp) SetFont (f font.Font) {
//
  pbx.SetFont (f)
}


func (x *Imp) Print (Z, S uint) {
//
  pbx.Print (x.text, Z, S)
}


func (x *Imp) Codelen () uint {
//
  return length
}


func (x *Imp) Encode () []byte {
//
  return []byte(x.text)[0:length]
}


func (x *Imp) Decode (b []byte) {
//
  x.text = string (b)
  str.Norm (&x.text, length)
  x.marked = false
}


func init () {
//
  var _ Word = New ()
  files.Cd0 ()
  bx.Wd (length)
  file = pseq.New (actual.text)
  file.Name ("Suchwort.dat")
  if file.Empty () {
    actual.Clr ()
    file.Put (actual.text)
  } else {
    actual.text = file.Get ().(string)
  }
}
