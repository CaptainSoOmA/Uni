package tval

// (c) Christian Maurer   v. 130115 - license see murus.go

import (
  . "murus/obj"; "murus/str"
  "murus/col"; "murus/box"; "murus/errh"
  "murus/font"; "murus/pbox"
)
const
  max = 15
type
  value byte; const (
  undecidable = iota
  falseVal
  trueVal
  nValues
)
type (
  representation [nValues]string
  Imp struct {
         val value
      length uint
         rep representation
      cF, cB col.Colour
          fo font.Font
             }
)
var (
  bx *box.Imp = box.New()
  pbx *pbox.Imp = pbox.New()
)


func (x *Imp) imp (Y Any) *Imp {
//
  y, ok:= Y.(*Imp)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}


func New () *Imp {
//
  x:= new (Imp)
  x.val = undecidable
  x.length = 1
  x.rep = [nValues]string { " ", "n", "j" }
  x.cF, x.cB = col.ScreenF, col.ScreenB
  return x
}


func (x *Imp) SetFormat (f, t string) {
//
  x.rep [falseVal] = f
  x.length = str.ProperLen (x.rep [falseVal])
  if x.length == 0 {
    x.rep = [nValues]string { " ", "n", "j" }
    return
  }
  x.rep [trueVal] = t
  n:= str.ProperLen (x.rep [falseVal])
  if n > x.length { x.length = n }
  x.rep [undecidable] = str.Clr (x.length)
}


func (x *Imp) Empty () bool {
//
  return x.val == undecidable
}


func (x *Imp) Clr () {
//
  x.val = undecidable
}


func (x *Imp) Copy (Y Object) {
//
  y:= x.imp (Y)
  x.val = y.val
  x.cF, x.cB = y.cF, y.cB
}


func (x *Imp) Clone () Object {
//
  y:= New ()
  y.Copy (x)
  return y
}


func (x *Imp) Eq (Y Object) bool {
//
  return x.val == x.imp (Y).val
}


func (x *Imp) Less (Y Object) bool {
//
  return x.val < x.imp (Y).val
}


func (x *Imp) Codelen () uint {
//
  return uint(1)
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, 1)
  b[0] = byte(x.val)
  return b
}


func (x *Imp) Decode (b []byte) {
//
  x.val = value(b[0])
}


func (x *Imp) Defined (s string) bool {
//
//  if ! str.Empty (s) { str.RemSpaces (&s) }
//  errh.Error2 (s, 999, "Länge", uint(len(s)))
  switch s[0] { case ' ', '?':
    x.val = undecidable
    return true
  }
  var p uint
  for v:= value(1); v < nValues; v++ {
    if str.IsEquivPart (s, x.rep [v], &p) && p == 0 {
      x.val = v
      return true
    }
  }
  return false
}


func (x *Imp) String () string {
//
  return x.rep [x.val]
}


func (x *Imp) SetColours (f, b col.Colour) {
//
  x.cF, x.cB = f, b
}


func (x *Imp) Write (l, c uint) {
//
  bx.Wd (x.length)
  bx.Colours (x.cF, x.cB)
  bx.Write (x.rep [x.val], l, c)
}


func (x *Imp) Edit (l, c uint) {
//
  bx.Wd (x.length)
  bx.Colours (x.cF, x.cB)
  var input string
  for {
    input = x.rep [x.val]
    bx.Write (input, l, c)
    bx.Edit (&input, l, c)
    if x.Defined (input) {
      break
    } else {
      errh.Error ("Eingabe unverständlich", 0) // , l + 1, c)
    }
  }
  x.Write (l, c)
}


func (x *Imp) SetFont (f font.Font) {
//
  pbx.SetFont (f)
}


func (x *Imp) Print (l, c uint) {
//
  pbx.SetFont (x.fo)
  pbx.Print (x.rep [x.val], l, c)
}


func (x *Imp) Set (b bool) {
//
  x.val = falseVal
  if b {
    x.val = trueVal
  }
}


// func init () { var _ TruthValue = New () }
