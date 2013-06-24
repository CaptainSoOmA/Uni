package addr

// (c) Christian Maurer   v. 130526 - license see murus.go

import (
  . "murus/obj"; "murus/kbd"
  "murus/col"; "murus/box"
  "murus/font"; "murus/pbox"; "murus/masks"
  "murus/text"; "murus/bnat"; "murus/phone"
)
const (
  LenStreet = 28
  LenCity   = 22
)
type
  Imp struct {
      street *text.Imp
    postcode *bnat.Imp
        city *text.Imp
 phonenumber,
  cellnumber *phone.Imp
             }
var (
  bx *box.Imp = box.New()
  pbx *pbox.Imp = pbox.New()
  cF, cB col.Colour = col.LightCyan, col.Black
  mask *masks.Imp = masks.New ()
  cst, cpc, cci, cph uint
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
  x.street, x.city = text.New (LenStreet), text.New (LenCity)
  x.postcode = bnat.New (10000)
  x.phonenumber, x.cellnumber = phone.New (), phone.New ()
  x.SetColours (cF, cB)
  return x
}


func (x *Imp) Empty () bool {
//
  return x.street.Empty () &&
         x.postcode.Empty () &&
         x.city.Empty () &&
         x.phonenumber.Empty () &&
         x.cellnumber.Empty ()
}


func (x *Imp) Clr () {
//
  x.street.Clr ()
  x.postcode.Clr ()
  x.city.Clr ()
  x.phonenumber.Clr ()
  x.cellnumber.Clr ()
}


func (x *Imp) Clone () Object {
//
  y:= New ()
  y.Copy (x)
  return y
}


func (x *Imp) Copy (Y Object) {
//
  y:= x.imp (Y)
  x.street.Copy (y.street)
  x.postcode.Copy (y.postcode)
  x.city.Copy (y.city)
  x.phonenumber.Copy (y.phonenumber)
  x.cellnumber.Copy (y.cellnumber)
}


func (x *Imp) Eq (Y Object) bool {
//
  y:= x.imp (Y)
  return x.street.Eq (y.street) &&
         x.postcode.Eq (y.postcode) &&
         x.city.Eq (y.city) &&
         x.phonenumber.Eq (y.phonenumber) &&
         x.cellnumber.Eq (y.cellnumber)
}


func (x *Imp) Equiv (Y Object) bool {
//
  if x.postcode.Eq (x.imp (Y).postcode) {
    return true
  }
  return false
}


func (x *Imp) Less (Y Object) bool {
//
  y:= x.imp (Y)
  if x.postcode.Eq (y.postcode) {
    if x.city.Eq (y.city) {
      return x.street.Less (y.street)
    } else {
      return x.city.Less (y.city)
    }
  }
  return x.postcode.Less (y.postcode)
}


func (x *Imp) SetColours (f, b col.Colour) {
//
  x.street.SetColours (f, b)
  x.postcode.SetColours (f, b)
  x.city.SetColours (f, b)
  x.phonenumber.SetColours (f, b)
  x.cellnumber.SetColours (f, b)
}


func (x *Imp) Write (l, c uint) {
//
  mask.Write (l, c)
  x.street.Write (l, c + cst)
  x.postcode.Write (l + 1, c + cpc)
  x.city.Write (l + 1, c + cci)
  x.phonenumber.Write (l, c + cph)
  x.cellnumber.Write (l + 1, c + cph)
}


func (x *Imp) Edit (l, c uint) {
//
  x.Write (l, c)
  i:= 0
  var d uint; if kbd.LastCommand (&d) == kbd.Up {
    i = 4
  }
  loop: for {
    switch i { case 0:
      x.street.Edit (l, c + cst)
    case 1:
      x.postcode.Edit (l + 1, c + cpc)
    case 2:
      x.city.Edit (l + 1, c + cci)
    case 3:
      x.phonenumber.Edit (l, c + cph)
    case 4:
      x.cellnumber.Edit (l + 1, c + cph)
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
      if i < 4 { i ++ } else { break loop }
    case kbd.Up, kbd.Left:
      if i > 0 { i -- } else { break loop }
    }
  }
}


func (x *Imp) SetFont (f font.Font) {
//
  x.street.SetFont (f)
  x.postcode.SetFont (f)
  x.city.SetFont (f)
  x.phonenumber.SetFont (f)
  x.cellnumber.SetFont (f)
}


func (x *Imp) Print (l, c uint) {
//
  mask.Print (l, c)
  x.street.Print (l, c + cst)
  pbx.Print ("Tel.:", l, c + cph - 6)
  x.phonenumber.Print (l, c + cph)
  x.postcode.Print (l + 1, c + cpc)
  x.city.Print (l + 1, c + cci)
  pbx.Print ("Funk:", l + 1, c + cph - 5)
  x.cellnumber.Print (l + 1, c + cph)
}


func (x *Imp) Codelen () uint {
//
  return LenStreet +                  // 28
         x.postcode.Codelen () +      //  4
         LenCity +                    // 22
         2 * x.phonenumber.Codelen () // 12
}                                     // 66


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, x.Codelen())
  i, a:= uint(0), x.street.Codelen()
  copy (b[i:i+a], x.street.Encode())
  i += a
  a = x.postcode.Codelen()
  copy (b[i:i+a], x.postcode.Encode())
  i += a
  a = x.city.Codelen()
  copy (b[i:i+a], x.city.Encode())
  i += a
  a = x.phonenumber.Codelen()
  copy (b[i:i+a], x.phonenumber.Encode())
  i += a
  copy (b[i:i+a], x.cellnumber.Encode())
  return b
}


func (x *Imp) Decode (b []byte) {
//
  i, a:= uint(0), x.street.Codelen()
  Decode (x.street, b[i:i+a])
  i += a
  a = x.postcode.Codelen()
  Decode (x.postcode, b[i:i+a])
  i += a
  a = x.city.Codelen()
  Decode (x.city, b[i:i+a])
  i += a
  a = x.phonenumber.Codelen()
  Decode (x.phonenumber, b[i:i+a])
  i += a
  Decode (x.cellnumber, b[i:i+a])
}


func init () {
//           1         2         3         4         5         6
// 012345678901234567890123456789012345678901234567890123456789012
// Anschrift: ____________________________  Tel.: ________________
// PLZ: _____  Ort: ______________________  Funk: ________________
  mask.Ins ("Anschrift:", 0,  0)
  mask.Ins ("PLZ:",       1,  0)
  mask.Ins ("Ort:",       1, 12)
  mask.Ins ("Tel.:",      0, 41)
  mask.Ins ("Funk:",      1, 41)
  cst, cpc, cci, cph = 11, 5, 17, 47
}
