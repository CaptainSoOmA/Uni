package fmon

// (c) Christian Maurer   v. 121030 - license see murus.go

//     Nichtsequentielle Programmierung mit Go 1 kompakt,
//     Kapitel 8, insbesondere Abschnitt 8.3

import
  . "murus/obj"
type
  FarMonitor interface {
// Specifications: Buy my book and read chapter 8.

// x is not yet activated.
//  New0 (a Any, n uint, f FuncSpectrum, p PredSpectrum, s string, p uint) *Imp

// x is activated.
//  New (a Any, n uint, f FuncSpectrum, p PredSpectrum, h string, p uint) *Imp

// x is activated.
  Go ()

  Prepare (s Stmt)

//  Func (i uint, a Any) Any // deprecated, replaced by F
  F (a Any, i uint) Any
  S (a Any, i uint, c chan Any) // experimental

  Terminate ()
}
