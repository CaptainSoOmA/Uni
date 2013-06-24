package bahnhof

// (c) Christian Maurer   v. 130308 - license see murus.go

import (
  . "murus/obj"
  . "murus/sub/linie"
)
type
  Bahnhof interface {

  Object

  Def (l Linie, nr uint, n, n1 string, b byte, y, x float64)
  Rescale (x, y uint)
  Linie () Linie
  Nummer () uint
  Umstieg ()
  Numerieren (l Linie, nr uint)
  Equiv (Y Object) bool
  SkalaEditieren ()
  UnterMaus () bool
  Write (b bool)
}
