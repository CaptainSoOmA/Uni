package char

// (c) Christian Maurer   v. 130419 - license see murus.go

import (
  . "murus/obj"
  "murus/z"
  "murus/kbd"
  "murus/col"; "murus/scr"
  "murus/font"; "murus/pbox"
)
type
  Imp struct {
      symbol byte
      cF, cB col.Colour
        font font.Font
             }
var
  pbx *pbox.Imp = pbox.New ()


func (x *Imp) imps (Y Object) byte {
//
  y, ok:= Y.(*Imp)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y.symbol
}


func New () *Imp {
//
  x:= new (Imp)
  x.symbol = ' '
  x.cF, x.cB = col.ScreenF, col.ScreenB
  return x
}


func (x *Imp) Def (b byte) { // TODO UTF8
//
  if b < ' ' { b = ' ' }
  x.symbol = b
}


func (x *Imp) Empty () bool {
//
  return x.symbol == ' '
}


func (x *Imp) Clr () {
//
  x.symbol = ' '
}


func (x *Imp) Eq (Y Object) bool {
//
  return x.symbol == x.imps (Y)
}


func (x *Imp) Equiv (Y Object) bool {
//
  return x.symbol % 32 == x.imps (Y) % 32
}


func (x *Imp) Copy (Y Object) {
//
  x.symbol = x.imps (Y)
}


func (x *Imp) Clone () Object {
//
  y:= new (Imp)
  y.Copy (x)
  return y
}


func (x *Imp) Val () uint {
//
  return uint(x.symbol)
}


func (x *Imp) Set (n uint) bool {
//
  if n >= 1<<8 {
    return false
  }
  b:= byte (n)
  if b < ' ' { b = ' ' } // TODO ausschalten unnützer Werte
  x.symbol = b
  return true
}


func (x *Imp) ByteVal () byte {
//
  return x.symbol
}


func (x *Imp) Less (Y Object) bool {
//
  return x.symbol < x.imps (Y)
}


func (x *Imp) String () string {
//
  return string(x.symbol)
}


func (x *Imp) Defined (s string) bool {
//
  if len (s) == 0 { return false }
  if len (s) > 1 { return false } // TODO UTF8
  x.symbol = byte(s[0])
  return true
}


func (x *Imp) SetColours (f, b col.Colour) {
//
  x.cF, x.cB = f, b
}


func (x *Imp) Write (l, c uint) {
//
  scr.Colours (x.cF, x.cB)
  scr.Write1 (x.symbol, l, c)
}


func (x *Imp) Edit (l, c uint) {
//
  b:= x.symbol
  scr.Colours (x.cF, x.cB)
  scr.Write1 (b, l, c)
  loop: for {
    scr.Warp (l, c, scr.Understroke)
    b = kbd.Byte ()
    switch { case ' ' <= b, b < 128, z.IsLatin1(b):
      break loop
    }
  }
  scr.Write1 (b, l, c)
  x.symbol = b
}


func (x *Imp) SetFont (f font.Font) {
//
  x.font = f
}


func (x *Imp) Print (l, c uint) {
//
  pbx.SetFont (x.font)
  pbx.Print (string(x.symbol), l, c)
}


func (x *Imp) Codelen () uint {
//
  return 1
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, 1)
  b[0] = x.symbol
  return b
}


func (x *Imp) Decode (b []byte) {
//
  x.symbol = b[0]
}


func init () { var _ Character = New() }
