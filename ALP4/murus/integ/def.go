package integ

// (c) Christian Maurer   v. 130312 - license see murus.go

import
  "murus/col"

// Specifications analogously to those in murus/nat.

func Wd (z int) uint { return wd(z) }

func Defined (z *int, s string) bool { return defined(z,s) }

func String (z int) string { return string_(z) }

func StringFmt (z int, w uint) string { return stringFmt(z,w) }

func SetColours (f, b col.Colour) { setColours(f,b) }

func Write (z int, l, c uint) { }

func SetWd (w uint) { setWd(w) }

func Edit (z *int, l, c uint) { edit(z,l,c) }
