package attr

// (c) Christian Maurer   v. 130526 - license see murus.go

import (
  . "murus/obj"; "murus/str"
  "murus/col"; "murus/font"; "murus/box"; "murus/pbox"; "murus/errh"
)
//const
//  LaengeHilfetext = 2 + NAttrs * (Wd + 1)
type
  Imp struct {
           a Attr
             }
var (
  txt [NAttrs]string
  bx, setbx *box.Imp = box.New(), box.New()
  pbx *pbox.Imp = pbox.New()
  Hilfetext string
  cF, cB, cMF, cMB col.Colour = col.LightWhite, col.Green, col.LightWhite, col.Green
  fmt Format
)


func New () *Imp {
//
  x:= new (Imp)
  x.a = 0
  return x
}


func (x *Imp) Empty () bool {
//
  return x.a == 0
}


func (x *Imp) Clr () {
//
  x.a = 0
}


func (x *Imp) Copy (X Object) {
//
  x1, ok:= X.(*Imp)
  if ! ok { return }
  x.a = x1.a
}


func (x *Imp) Clone () Object {
//
  x1:= New ()
  x1.a = x.a
  return x1
}


func (x *Imp) Eq (X Object) bool {
//
  x1, ok:= X.(*Imp)
  if ! ok { return false }
  return x.a == x1.a
}


func (x *Imp) Less (X Object) bool {
//
  x1, ok:= X.(*Imp)
  if ! ok { return false }
  return x.a < x1.a
}


func (x *Imp) Mark (ja bool) {
//
  if ja {
    bx.Colours (cMF, cMB)
  } else {
    bx.Colours (cF, cB)
  }
}


func (x *Imp) SetFormat (f Format) {
//
  fmt = f
  switch fmt {
  case Kurz:
    bx.Wd (1)
  case Mittel:
    bx.Wd (Wd - 1)
  case Lang:
    bx.Wd (Wd)
  }
}


func (x *Imp) SetColours (f, b col.Colour) {
//
// dummy
}


func (x *Imp) Write (Z, S uint) {
//
  s:= txt[x.a]
  switch fmt {
  case Kurz:
    s = string(s[0])
  case Mittel:
    s = string(s[1:])
  }
  bx.Write (s, Z, S)
}


func (x *Imp) SetFont (f font.Font) {
//
// dummy
}


func (x *Imp) Print (Z, S uint) {
//
  pbx.Print (txt[x.a], Z, S)
}


func (x *Imp) Edit (Z, S uint) {
//
  T:= txt[x.a]
  loop:
  for {
    bx.Edit (&T, Z, S)
    for a:= undef; a < NAttrs; a++ {
      if T[0] == txt[a][0] {
        x.a = a
        x.Write (Z, S)
        break loop
      }
    }
    errh.Error (Hilfetext, 0)
  }
}


func (x *Imp) Codelen () uint {
//
  return 1
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, 1)
  b[0] = byte(x.a)
  return b
}


func (x *Imp) Decode (b []byte) {
//
  if b[0] < byte(NAttrs) {
    x.a = Attr(b[0])
  } else {
    x.a = undef
  }
}


func init () {
//
  txt[undef]= "     "
  txt[priv] = ".priv"
  txt[Erl]  = ">Erl."
  txt[Tel]  = "#Tel."
  txt[Brf]  = "/Brf."
  txt[Fin]  = "$Fin."
  txt[Ren]  = "^Ren."
  txt[Hob]  = "&Hob."
  txt[Prog] = ";Prog"
  txt[Uni]  = "!Uni "
  txt[Arzt] = "@Arzt"
  txt[Geb]  = "*Geb."
  Hilfetext = str.Clr (2)
  for a:= priv; a < NAttrs; a++ { Hilfetext += txt[a] + "  " }
  setbx.Wd (uint(NAttrs))
  setbx.Colours (cF, cB);
  var _ Attribute = New ()
}
