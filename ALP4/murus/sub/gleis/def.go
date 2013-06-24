package gleis

// (c) Christian Maurer   v. 130308 - license see murus.go

import (
  . "murus/obj"
  "murus/sub/linie"
)
type
  Gleis interface { // Strecke mit Linie und natürlicher Zahl < 10 als Wert

  Object
  Valuator

// Wenn l und n zulässige Werte haben, gehört x zur Linie l und hat den Wert n.
  Def (l linie.Linie, n uint)

// In scr ist die Farbe von x gesetzt.
  Write (b bool)
}
