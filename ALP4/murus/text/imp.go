package text

// (c) Christian Maurer   v. 130303 - license see murus.go

import (
  "murus/rand"; . "murus/obj"; "murus/z"; "murus/str"
  "murus/col"; "murus/box"
  "murus/font"; "murus/pbox"
)
type
  Imp struct {
      length uint
     content string
      cF, cB col.Colour
        font font.Font
             }
var (
  bx *box.Imp = box.New ()
  pbx *pbox.Imp = pbox.New()
  Vokale, Konsonanten string = "aeiouy", "bcdfghjklmnpqrstvwxz"
  upper, lower string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "abcdefghijklmnopqrstuvwxyz"
)


func (x *Imp) imp (Y Object) *Imp {
//
  y, ok:= Y.(*Imp)
  if ! ok || x.length != y.length { TypeNotEqPanic (x, Y) }
  return y
}


func New (n uint) *Imp {
//
  if n == 0 { return nil }
  x:= new (Imp)
  x.length = n
  x.content = str.Clr (n)
  x.cF, x.cB = col.ScreenF, col.ScreenB
  x.font = font.Normal
  return x
}


func randomvokal () byte {
//
  return Vokale [rand.Natural (6)]
}


func randomkonsonant () byte {
//
  return Konsonanten [rand.Natural (20)]
}


func (x *Imp) Generate () {
//
  b:= make ([]byte, x.length)
  b[0] = upper [rand.Natural (uint(len (upper)))]
  for i:= uint(1); i < x.length; i++ {
    b[i] = lower [rand.Natural (uint(len (lower)))]
  }
  x.content = string (b)
  z.ToHellWithUTF8 (&x.content)
/*
  n:= 3 + rand.Natural (x.length - 2)
  if n >= x.length {
    n = x.length - 1
  }
  b:= rand.Natural (2) % 2 == 1
  s:= x.length
  for i:= 0; i < int(n); i++ {
    if (i % 2 == 1) == b {
      B[i] = randomkonsonant ()
      if B[i] == 's' && i + 2 < int(n) {
        s = uint(i)
      }
    } else {
      B[i] = randomvokal ()
    }
    if i == 0 {
//      B[i] = CAP (B[i])
    }
  }
  if s < x.length {
    B[s + 1] = 'c'
    B[s + 2] = 'h'
  }
  for i:= n; i <= x.length; i++ {
//    B[i] = 0C
  }
*/
}


func (x *Imp) Empty () bool {
//
  return str.Empty (x.content)
}


func (x *Imp) Clr () {
//
  x.content = str.Clr (x.length)
}


func (x *Imp) Copy (Y Object) {
//
  y:= x.imp (Y)
  x.content = y.content
  x.cF, x.cB = y.cF, y.cB
}


func (x *Imp) Clone () Object {
//
  y:= New (x.length)
  y.Copy (x)
  return y
}


func (x *Imp) Eq (Y Object) bool {
//
  return x.content == x.imp (Y).content
}


func (x *Imp) Less (Y Object) bool {
//
  return str.Less (x.content, x.imp (Y).content)
}


func (x *Imp) Equiv (Y Object) bool {
//
  return str.Equiv (x.content, x.imp (Y).content)
}


func (x *Imp) IsPart (Y Object) bool {
//
  var p uint
  return str.IsPart (x.content, x.imp (Y).content, &p)
}


func (x *Imp) IsEquivalentPart (Y Object) bool {
//
  return str.IsEquivPart0 (x.content, x.imp (Y).content)
}


func (x *Imp) Defined (s string) bool {
//
  if uint(len (s)) > x.length { return false }
  str.Set (&(x.content), s)
  str.Norm (&(x.content), x.length)
  return true
}


func (x *Imp) String () string {
//
  return x.content
}


func (x *Imp) SetColours (f, b col.Colour) {
//
  x.cF, x.cB = f, b
}


func (x *Imp) Write (l, c uint) {
//
  bx.Wd (x.length)
  bx.Colours (x.cF, x.cB)
  bx.Write (x.content, l, c)
}


func (x *Imp) Edit (l, c uint) {
//
  bx.Wd (x.length)
  bx.Colours (x.cF, x.cB)
  bx.Edit (&x.content, l, c)
}


func (x *Imp) SetFont (f font.Font) {
//
  x.font = f
}


func (x *Imp) Print (l, c uint) {
//
  pbx.SetFont (x.font)
  pbx.Print (x.content, l, c)
}


func (x *Imp) Codelen () uint {
//
  return x.length
}


func (x *Imp) Encode () []byte {
//
  return ([]byte)(x.content)
}


func (x *Imp) Decode (b []byte) {
//
  if uint(len (b)) == x.length {
    x.content = string(b)
//    str.Lat1 (&x.content)
  } else {
    x.content = str.Clr (x.length)
  }
}


func (x *Imp) Len () uint {
//
  return x.length
}

/////////////////////////////////////////////////////

func (x *Imp) ProperLen () uint {
//
  return str.ProperLen (x.content) // return str.ProperLen (x.String ())
}


func (x *Imp) IsCap () bool {
//
  s:= x.String ()
  return z.IsCap (s[0])
}


func (x *Imp) ToUpper () {
//
  s:= x.String ()
  str.ToUpper (&s)
  x.Defined (s)
}


func (x *Imp) ToLower () {
//
  s:= x.String ()
  str.ToLower (&s)
  x.Defined (s)
}


// func init () { var _ Text = New (16) }
