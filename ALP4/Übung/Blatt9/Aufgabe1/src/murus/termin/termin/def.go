package termin

// (c) Christian Maurer   v. 130127 - license see murus.go

import
  . "murus/obj"
const ( // Format
  Lang = Format(iota) // eine vollständige Bildschirmzeile
  Kurz                // eine Zeile, zehn Spalten
  GanzKurz            // ein Zeichen
)
type
  Termin interface {

  Formatter
  Marker
  Editor
  Printer
}
