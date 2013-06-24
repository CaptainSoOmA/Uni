package atom

// (c) Christian Maurer   v. 130115 - license see murus.go

import (
  . "murus/obj"; "murus/ker"; "murus/font"; "murus/col"
  "murus/enum"; "murus/tval"; "murus/char"; "murus/text"; "murus/bnat"; "murus/breal"
  "murus/clk"; "murus/day"; "murus/euro"; "murus/cntry";
  "murus/pers"; "murus/phone"; "murus/addr"
)
const
  pack = "atom"
type
  Imp struct {
             Object
         typ Atomtype
             }


func (x *Imp) imp (Y Object) *Imp {
//
  y, ok:= Y.(*Imp)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}


func New (o Object) *Imp {
//
  x:= new (Imp)
  switch o.(type) {
  case *enum.Imp:
    x.Object, x.typ = enum.New (enum.Enum(o.(*enum.Imp).Typ ())), Enumerator
  case *tval.Imp:
    x.Object, x.typ = tval.New (), TruthValue
  case *char.Imp:
    x.Object, x.typ = char.New (), Character
  case *text.Imp:
    x.Object, x.typ = text.New (o.(*text.Imp).Len ()), Text
  case *bnat.Imp:
    x.Object, x.typ = bnat.New (o.(*bnat.Imp).Startval ()), Natural
  case *breal.Imp:
    x.Object, x.typ = breal.New (4), Real
  case *clk.Imp:
    x.Object, x.typ = clk.New (), Clocktime
  case *day.Imp:
    x.Object, x.typ = day.New (), Calendarday
  case *euro.Imp:
    x.Object, x.typ = euro.New (), Euro
  case *cntry.Imp:
    x.Object, x.typ = cntry.New (), Country
  case *pers.Imp:
    x.Object, x.typ = pers.New (), Person
  case *phone.Imp:
    x.Object, x.typ = phone.New (), PhoneNumber
  case *addr.Imp:
    x.Object, x.typ = addr.New (), Address
  default:
    ker.Panic ("atom.New: parameter does not characterize an atom.Atomtype")
  }
  return x
}


func (x *Imp) Type () Atomtype {
//
  return x.typ
}


func (x *Imp) Equiv (Y Object) bool {
//
  return x.typ == x.imp (Y).typ
}


func (x *Imp) Empty () bool {
//
  return x.Object.Empty ()
}


func (x *Imp) Clr () {
//
  x.Object.Clr ()
}


func (x *Imp) Eq (Y Object) bool {
//
  return x.Object.Eq (x.imp (Y).Object)
}


func (x *Imp) Copy (Y Object) {
//
  x.Object.Copy (x.imp (Y).Object)
}


func (x *Imp) Clone () Object {
//
  y:= New (x.Object)
  y.Copy (x)
  return y
}


func (x *Imp) Less (Y Object) bool {
//
  return x.Object.Less (x.imp (Y).Object)
}


func (x *Imp) SetFormat (f Format) {
//
  if y, ok:= x.Object.(Formatter); ok {
    y.SetFormat (f)
  }
}


func (x *Imp) SetColours (f, b col.Colour) {
//
  if e, ok:= x.Object.(Editor); ok {
    e.SetColours (f, b)
  }
}


func (x *Imp) Write (l, c uint) {
//
  if e, ok:= x.Object.(Editor); ok {
    e.Write (l, c)
  }
}


func (x *Imp) Edit (l, c uint) {
//
  if e, ok:= x.Object.(Editor); ok {
    e.Edit (l, c)
  }
}


func (x *Imp) SetFont (f font.Font) {
//
  if p, ok:= x.Object.(Printer); ok {
    p.SetFont (f)
  }
}


func (x *Imp) Print (l, c uint) {
//
  if p, ok:= x.Object.(Printer); ok {
    p.Print (l, c)
  }
}


func (x *Imp) Codelen () uint {
//
  return x.Object.Codelen ()
}


func (x *Imp) Encode () []byte {
//
  return x.Object.Encode ()
}


func (x *Imp) Decode (b []byte) {
//
  x.Object.Decode (b)
}


func init () { var _ Atom = New (day.New()) }
