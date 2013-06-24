package atom

// (c) Christian Maurer   v. 130116 - license see murus.go

import (
  "murus/col"; "murus/scr"; "murus/sel"
  "murus/enum"; "murus/tval"; "murus/char"; "murus/text"; "murus/bnat"; "murus/breal"
  "murus/clk"; "murus/day"; "murus/euro"; "murus/cntry";
  "murus/pers"; "murus/phone"; "murus/addr"
)
const
  M = 14
var (
  name []string
)


func Selected (l, c uint) *Imp {
//
  cF, cH:= col.ScreenF, col.ScreenB
  col.Contrast (&cH)
  n:= uint(0)
  z, s:= scr.MousePos()
  x:= new (Imp)
  sel.Select1 (name, uint(NAtomtypes), M, &n, z, s, cF, cH)
  if n < uint(NAtomtypes) {
    x.typ = Atomtype(n)
  } else {
    return nil
  }
  switch x.typ { case Enumerator:
    e:= enum.Title // TODO e per select-menue aussuchen
    x.Object = enum.New (e)
  case TruthValue:
    x.Object = tval.New ()
  case Character:
    x.Object = char.New ()
  case Text:
    n:= uint(10) // TODO n editieren
    x.Object = text.New (n)
  case Natural:
    n:= uint(10) // TODO n editieren
    x.Object = bnat.New (n)
  case Real:
    n:= uint(6) // TODO n editieren
    x.Object = breal.New (n)
  case Clocktime:
    x.Object = clk.New ()
  case Calendarday:
    x.Object = day.New ()
  case Euro:
    x.Object = euro.New ()
  case Country:
    x.Object = cntry.New ()
  case Person:
    x.Object = pers.New ()
  case PhoneNumber:
    x.Object = phone.New ()
  case Address:
    x.Object = addr.New ()
  }
  return New (x)
}


func init () {
//
  name = []string { "Enumerator",
                    "TruthValue", "Character", "Text", "Natural", "Real",
                    "Clocktime", "Calendarday", "Euro", "Country",
                    "Person", "PhoneNumber", "Address" }
}
