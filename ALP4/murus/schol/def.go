package schol

// (c) Christian Maurer   v. 130115 - license see murus.go

import
  . "murus/obj"
const ( // Format
  Minimal = iota //  1 Zeile,  52 Spalten
  VeryShort      //  1 Zeile,  80 Spalten
  Short          //  2 Zeilen, 80 Spalten
  Long           // 21 Zeilen, 80 Spalten
  NFormats
)
const ( // Order
  NameOrder = iota
  AgeOrder
  NOrders
)
var
  ActualOrder Order
type
  Scholar interface {

  Editor
  Printer

// some func's
}
