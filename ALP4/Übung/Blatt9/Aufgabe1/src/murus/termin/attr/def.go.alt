package attr

// (c) Christian Maurer   v. 130511 - license see murus.go

import
  . "murus/obj"
type
  Attr byte; const (
  undef = Attr(iota)
  priv
  Erl
  Tel
  Brf
  Fin
  Ren
  Hob
  Prog
  Uni
  Arzt
  Geb
  NAttrs
)
const
  Wd = 5
const ( // Format
  Kurz = Format(iota) // charakterisierendes erstes Zeichen
  Mittel              // 4 Zeichen (ohne char. erstes Zeichen)
  Lang                // Kurz + Mittel, d.h. Wd Zeichen
)
type
  Attribute interface {

  Editor
  Printer
  SetFormat (f Format)
  Mark (m bool)
}
type
  AttrSet interface {
  //  NewSet () *Set
  Clearer
  Copy (Y AttrSet)
  Ins (a Attribute)
  Write (l, c uint)
}
