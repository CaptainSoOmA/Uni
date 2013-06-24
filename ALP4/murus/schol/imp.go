package schol

// (c) Christian Maurer   v. 130526 - license see murus.go

import (
  . "murus/obj"
  "murus/kbd"
  "murus/col"; "murus/box"
//  "murus/errh"
  "murus/font"; "murus/pbox"
  "murus/text"; "murus/pers"; "murus/cntry"; "murus/addr"
  "murus/langs"; "murus/enum"
)
const
  pack = "schol"
type
  Imp struct {
      person *pers.Imp
placeOfBirth *text.Imp
 nationality *cntry.Imp
     address *addr.Imp
legalGuardian *pers.Imp
   addressLG *addr.Imp
     langSeq *langs.Imp
    religion *enum.Imp
     notUsed byte
      format Format
             }
const
  lenPlace = 22
var (
  bx *box.Imp = box.New ()
  cF, cB col.Colour = col.White, col.Black
  pbx *pbox.Imp = pbox.New ()
  temp, temp1 *Imp = New (), New ()
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
  x.person = pers.New ()
  x.placeOfBirth = text.New (lenPlace)
  x.nationality = cntry.New ()
  x.address = addr.New ()
  x.legalGuardian = pers.New ()
  x.legalGuardian.SetFormat (pers.LongT)
  x.addressLG = addr.New ()
  x.langSeq = langs.New ()
  x.religion = enum.New (enum.Religion)
  x.format = Short
  x.SetColours (col.LightCyan, col.Black)
  return x
}


func (x *Imp) Empty () bool {
//
  return x.person.Empty ()
}


func (x *Imp) Clr () {
//
  x.person.Clr ()
  x.placeOfBirth.Clr ()
  x.nationality.Clr ()
  x.address.Clr ()
  x.legalGuardian.Clr ()
  x.addressLG.Clr ()
  x.langSeq.Clr ()
  x.religion.Clr ()
  x.notUsed = 0
}


func (x *Imp) Copy (Y Object) {
//
  y:= x.imp (Y)
  y.person.Copy (x.person)
  y.placeOfBirth.Copy (x.placeOfBirth)
  y.nationality.Copy (x.nationality)
  y.address.Copy (x.address)
  y.legalGuardian.Copy (x.legalGuardian)
  y.addressLG.Copy (x.addressLG)
  y.langSeq.Copy (x.langSeq)
  y.religion.Copy (x.religion)
  y.notUsed = x.notUsed
}


func (x *Imp) Clone () Object {
//
  y:= New ()
  y.Copy (x)
  return x
}


func (x *Imp) Eq (Y Object) bool {
//
  y:= x.imp (Y)
  switch x.format { case Minimal, VeryShort:
    return x.person.Eq (y.person)
  case Short:
    return x.person.Eq (y.person) &&
           x.langSeq.Eq (y.langSeq)
  } // Long:
  return x.person.Eq (y.person) &&
         x.langSeq.Eq (y.langSeq) &&
         x.placeOfBirth.Eq (y.placeOfBirth) &&
         x.nationality.Eq (y.nationality) &&
         x.address.Eq (y.address) &&
         x.legalGuardian.Eq (y.legalGuardian) &&
         x.addressLG.Eq (y.addressLG) &&
         x.religion.Eq (y.religion)
}


func (x *Imp) Equiv (Y Object) bool {
//
  pers.ActualOrder = ActualOrder
  return x.person.Equiv (x.imp (Y).person)
}


func (x *Imp) Less (Y Object) bool {
//
  pers.ActualOrder = ActualOrder
  return x.person.Less (x.imp (Y).person)
}


func (x *Imp) String () string {
//
  return x.person.String ()
}

/*
func (x *Imp) Num (l []*enum.Imp, v, b []uint) uint {
//
  return x.langSeq.Num (l, v, b)
}
*/

func (x *Imp) FullAged () bool {
//
  return x.person.FullAged ()
}


func (x *Imp) SetFormat (f Format) {
//
  if f < NFormats {
    x.format = f
    switch f { case Minimal:
      x.person.SetFormat (pers.ShortB)
    case Short:
      x.langSeq.SetFormat (langs.Short)
    case Long:
      x.langSeq.SetFormat (langs.Long)
    }
  }
}


func (x *Imp) SetColours (f, b col.Colour) {
//
  x.person.SetColours (f, b)
  x.placeOfBirth.SetColours (f, b)
  x.nationality.SetColours (f, b)
  x.address.SetColours (f, b)
  x.legalGuardian.SetColours (cF, cB)
  x.addressLG.SetColours (cF, cB)
  x.langSeq.SetColours (f, b)
  x.religion.SetColours (f, b)
}


var
  lLs, lPb, lNa, lAd, lLg, lAg, lRe,
  cLs, cPb, cNa, cAd, cLg, cAg, cRe uint


func (x *Imp) writeMask (l, c uint) {
//
  switch x.format { case Minimal:
    lLs = 0; cLs = 0
  case VeryShort:
    lLs = 0; cLs = 0
  case Short:
    lLs = 1; cLs = 16
  case Long:
/*        1         2         3         4         5         6         7
0123456789012345678901234567890123456789012345678901234567890123456789012345
Geburtsort: ______________________ Staatsangehörigkeit: ____________________

gesetzl. Vertreter(in):
Person, Anschrift
Sprachenfolge: Sprachenfolge
 
Religionszugehörigkeit: ______________________
*/
    lPb = 1; cPb = 12; lNa = 1; cNa = 56
    lAd = 3; cAd = 0; lLg = 7; cLg = 0; lAg = 9; cAg = 0
    lLs = 12; cLs = 15; lRe = 17; cRe = 25
  }
  bx.ColoursScreen ()
  switch x.format {
    case Minimal, VeryShort:
  default:
    bx.Wd (14)
    bx.Write ("Sprachenfolge:", l + lLs, c + cLs - 15)
  }
  if x.format == Long {
    bx.Wd (11)
    bx.Write ("Geburtsort:", l + lPb, c + cPb - 12)
    bx.Wd (20)
    bx.Write ("Staatsangehörigkeit:", l + lNa, c + cNa - 21)
    bx.Wd (24)
    bx.Write ("gesetzl. Vertreter(in)", l + lLg - 1, c + 1)
    bx.Wd (23)
    bx.Write ("Religionszugehörigkeit:", l + lRe, c + cRe - 24)
  }
}


func (x *Imp) Write (l, c uint) {
//
  x.writeMask (l, c)
  x.person.Write (l, c)
  switch x.format { case Minimal, VeryShort:
  default:
    x.langSeq.Write (l + lLs, c + cLs)
  }
  if x.format == Long {
    x.placeOfBirth.Write (l + lPb, c + cPb)
    x.nationality.Write (l + lNa, c + cNa)
    x.address.Write (l + lAd, c + cAd)
    x.legalGuardian.Write (l + lLg, c + cLg)
    x.addressLG.Write (l + lAg, c + cAg)
    x.religion.Write (l + lRe, c + cRe)
  }
}


func (x *Imp) Edit0 (l, c uint) {
//
  x.person.Edit (l, c)
}


func (x *Imp) Edit (l, c uint) {
//
  const nKomponenten = 8
  x.Write (l, c)
  i:= 1
  loop: for {
    switch i { case 1:
    x.person.Edit (l, c)
    case 2:
      if x.format == Long {
        x.placeOfBirth.Edit (l + lPb, c + cPb)
      }
    case 3:
      if x.format == Long {
        x.nationality.Edit (l + lNa, c + cNa)
      }
    case 4:
      if x.format == Long {
        x.address.Edit (l + lAd, c + cAd)
//        if ! x.person.FullAged () {
//          x.address.Copy (x.addressLG)
//        }
      }
    case 5:
      if x.format == Long {
        if x.person.FullAged () {
          // x.legalGuardian.Clr ()
        } else {
          x.legalGuardian.Edit (l + lLg, c + cLg)
        }
      }
    case 6:
      if x.format == Long {
        if x.person.FullAged () {
          // x.legalGuardian.Clr ()
        } else {
          x.addressLG.Edit (l + lAg, c + cAg)
        }
      }
    case 7:
      if x.format != Minimal && x.format != VeryShort {
        x.langSeq.Edit (l + lLs, c + cLs)
      }
    case nKomponenten:
      if x.format == Long {
        x.religion.Edit (l + lRe, c + cRe)
      }
    }
    var d uint
    switch kbd.LastCommand (&d) { case kbd.Enter:
      if d == 0 /* aufpassen bei i == 0 ! */ {
        if i < nKomponenten { i++ } else { break loop }
      } else {
        break loop
      }
    case kbd.Esc:
      break loop
    case kbd.Down:
      if i < nKomponenten { i++ } else { break loop }
    case kbd.Up:
      if i > 1 { i -- }
    }
//    if ! x.person.Identifiable () {
//      errh.Error ("Name, Vorname, Geb.-Datum ?", 0)
//    }
  }
}


func (x *Imp) SetFont (f font.Font) {
//
  pbx.SetFont (f)
}


func (x *Imp) printMask (l, c uint) {
//
  switch x.format {
  case Minimal:
    lLs = 0; cLs = 0
  case VeryShort:
    lLs = 0; cLs = 0
  case Short:
    lLs = 1; cLs = 16
  case Long:
/*        1         2         3         4         5         6         7
0123456789012345678901234567890123456789012345678901234567890123456789012345
Person
Geburtsort: ______________________ Staatsangehörigkeit: ____________________

Anschrift
gesetzl. Vertreter(in):
Person, Anschrift

Sprachenfolge: ___________ von Klasse __ bis Klasse __
               ___________ von Klasse __ bis Klasse __
               ___________ von Klasse __ bis Klasse __
               ___________ von Klasse __ bis Klasse __

Religionszugehörigkeit: ______________________
*/
    lPb = 1; cPb = 12; lNa = 1; cNa = 56
    lAd = 3; cAd = 0; lLg = 7; cLg = 0; lAg = 9; cAg = 0
    lLs = 12; cLs = 15; lRe = 17; cRe = 25
  }
  switch x.format { case Minimal, VeryShort:
  default:
    pbx.Print ("langSeq:", l + lLs, c + cLs - 15)
  }
  if x.format == Long {
    pbx.Print ("Geburtsort:", l + lPb, c + cPb - 12)
    pbx.Print ("Staatsangehörigkeit:", l + lNa, c + cNa - 21)
    pbx.Print ("gesetzl. Vertreter(in):", l + lLg - 1, c + 1)
    pbx.Print ("Religionszugehörigkeit:", l + lRe, c + cRe - 24)
  }
}


func (x *Imp) Print (l, c uint) {
//
  x.printMask (l, c)
  x.person.Print (l, c)
  switch x.format { case Minimal, VeryShort:
  default:
    x.langSeq.Print (l + lLs, c + cLs)
  }
  if x.format == Long {
    x.placeOfBirth.Print (l + lPb, c + cPb)
    x.nationality.Print (l + lNa, c + cNa)
    x.address.Print (l + lAd, c + cAd)
    x.legalGuardian.Print (l + lLg, c + cLg)
    x.addressLG.Print (l + lAg, c + cAg)
    x.religion.Print (l + lRe, c + cRe)
  }
}


func (x *Imp) Codelen () uint {
//
  return 45 + // x.person
       + 22 + // lenPlace
       +  2 + // x.nationality.Codelen ()
       + 66 + // x.address.Codelen ()
       + 45 + // x.legalGuardian.Codelen ()
       + 66 + // x.addressLG.Codelen ()
       +  8 + // x.langSeq.Codelen ()
       +  1 + // x.religion.Codelen ()
       +  1   // x.notUsed
} //    256


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, x.Codelen())
  i, a:= uint (0), x.person.Codelen()
  copy (b[i:i+a], x.person.Encode ())
  i += a
  a = lenPlace
  copy (b[i:i+a], x.placeOfBirth.Encode ())
  i += a
  a = x.nationality.Codelen()
  copy (b[i:i+a], x.nationality.Encode ())
  i += a
  a = x.address.Codelen()
  copy (b[i:i+a], x.address.Encode ())
  i += a
  a = x.legalGuardian.Codelen()
  copy (b[i:i+a], x.legalGuardian.Encode ())
  i += a
  a = x.addressLG.Codelen()
  copy (b[i:i+a], x.addressLG.Encode ())
  i += a
  a = x.langSeq.Codelen()
  copy (b[i:i+a], x.langSeq.Encode ())
  i += a
  a = x.religion.Codelen()
  copy (b[i:i+a], x.religion.Encode ())
  i += a
  b[i] = x.notUsed
  return b
}


func (x *Imp) Decode (b []byte) {
//
  i, a:= uint (0), x.person.Codelen()
  x.person.Decode (b[i:i+a])
  i += a
  a = lenPlace
  x.placeOfBirth.Decode (b[i:i+a])
  i += a
  a = x.nationality.Codelen()
  x.nationality.Decode (b[i:i+a])
  i += a
  a = x.address.Codelen()
  x.address.Decode (b[i:i+a])
  i += a
  a = x.legalGuardian.Codelen()
  x.legalGuardian.Decode (b[i:i+a])
  i += a
  a = x.addressLG.Codelen()
  x.addressLG.Decode (b[i:i+a])
  i += a
  a = x.langSeq.Codelen()
  x.langSeq.Decode (b[i:i+a])
  i += a
  a = x.religion.Codelen()
  x.religion.Decode (b[i:i+a])
  i += a
  x.notUsed = b[i]
}


func init () { var _ Scholar = New () }
