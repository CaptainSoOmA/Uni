package enumbase

// (c) Christian Maurer   v. 130115 - license see murus.go

import (
  . "murus/obj"; "murus/str"; "murus/kbd"
  "murus/col"; "murus/box"; "murus/errh"; "murus/sel"
  "murus/font"; "murus/pbox"
)
type
  Imp struct {
         typ,
           b byte
           s [NFormats][]string
         num uint
          wd [NFormats]uint
      cF, cB col.Colour
           f Format
             }
var (
  bx *box.Imp = box.New()
  pbx *pbox.Imp = pbox.New()
)


func imp (X Object) *Imp {
//
  x, ok:= X.(*Imp)
  if ! ok { NotCompatiblePanic() }
  return x
}


func New (t byte, s [NFormats][]string) *Imp {
//
  x:= new (Imp)
  x.typ, x.s = t, s
  x.num = uint(len (s[Short]))
  m:= [NFormats]uint { uint(0), uint(0) }
  for f:= Short; f < NFormats; f++ {
    for i, t:= range (s[f]) {
      str.Set (&s[f][i], t)
      w:= uint(len (s[f][i]))
      if m[f] < w { m[f] = w }
    }
    for i, _:= range (s[f]) { str.Norm (&s[f][i], m[f]) } // TODO gefährlich
  }
  x.wd = m
  x.cF, x.cB = col.ScreenF, col.ScreenB
  x.f = Short
  return x
}


func (x *Imp) SetFormat (f Format) {
//
  x.f = f
}


func (x *Imp) Typ () byte {
//
  return x.typ
}


func (x *Imp) Empty () bool {
//
  return x.b == 0
}


func (x *Imp) Clr () {
//
  x.b = 0
}


func (x *Imp) Eq (Y Object) bool {
//
  return x.b == imp (Y).b
}


func (x *Imp) Copy (Y Object) {
//
  y:= imp (Y)
  x.b = y.b
  x.cF, x.cB = y.cF, y.cB
  x.f = y.f
}


func (x *Imp) Clone () Object {
//
  y:= New (x.typ, x.s)
  y.Copy (x)
  return x
}


func (x *Imp) Codelen () uint {
//
  return 1
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, 1)
  b[0] = byte (x.b)
  return b
}


func (x *Imp) Decode (b []byte) {
//
  if uint(b[0]) < x.num {
    x.b = b[0]
  } else {
    x.b = 0
  }
}


func (x *Imp) Less (Y Object) bool {
//
  y:= imp (Y)
  return str.Less (x.String (), y.String ())
}


func (x *Imp) String () string {
//
  s:= x.s[x.f][x.b]
  str.RemSpaces (&s)
  return s
}


func (x *Imp) Defined (s string) bool {
//
  for b:= uint(0); b < x.num; b++ {
    if str.IsEquivPart0 (s, x.s[x.f][b]) {
      x.b = byte(b)
      return true
    }
  }
  return false
}


func (x *Imp) SetColours (f, b col.Colour) {
//
  x.cF, x.cB = f, b
}


func (x *Imp) Write (l, c uint) {
//
  bx.Colours (x.cF, x.cB)
  bx.Wd (x.wd[x.f])
  bx.Write (x.String (), l, c)
}


func (x *Imp) selected (l, c uint) bool {
//
  if x.num == 0 { return false }
  if x.num == 1 { return true }
  x.Write (l, c)
  i:= uint(0)
  h:= x.num / 2
  if h < 5 { h = 5 }
  if h > x.num { h = x.num }
  errh.Hint (errh.ToSelect)
//  f, b:= x.cF, x.cB
//  f, b = col.Pink, col.Darkmagenta
  sel.Select (func (p, l, c uint, f, b col.Colour) {
                bx.Colours (f, b); bx.Write (x.s[x.f][p], l, c)
              }, x.num, h, x.wd[x.f], &i, l, c, x.cB, x.cF)
  errh.DelHint()
  if i < x.num { x.b = byte(i) }
  return i < x.num
}


func (x *Imp) Edit (l, c uint) {
//
  x.Write (l, c)
  s:= x.String ()
  for {
    bx.Edit (&s, l, c)
    var d uint; cmd:= kbd.LastCommand (&d)
    if cmd == kbd.LookFor {
      if x.selected (l, c) {
        break
      }
    }
    if x.Defined (s) {
      break
    } else {
      errh.Error ("geht nicht", 0)
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
  pbx.Print (x.s[x.f][x.b], l, c)
}


func (x *Imp) Ord () uint {
//
  return uint(x.b)
}


func (x *Imp) Num () uint {
//
  return x.num
}


func (x *Imp) Wd () uint {
//
  return x.wd[x.f]
}


func (x *Imp) Set (n uint) bool {
//
  if n < x.num {
    x.b = byte (n)
    return true
  }
  x.b = 0
  return false
}
