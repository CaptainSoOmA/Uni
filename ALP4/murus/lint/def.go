package lint

// (c) Christian Maurer   v. 130115 - license see murus.go

import
  . "murus/obj"
type
  LongInteger interface {

// New (n int) *Imp
// New64 (n int64) *Imp
// NewNat (n uint) *Imp
// NewReal (r float64) *Imp

  Editor
  Geq0 () bool
  Stringer
  Printer
  Adder
  Multiplier

  Set (n int)
  Set32 (n int32)
  Set64 (n int64)
  SetReal (n float64)

  Len () uint
  Odd () bool
  ChSign ()
  Val () int
  Val64 () int64
  RealVal () float64
  SumDigits () uint

// Specs see murus.nat
  Inc ()
  Dec ()
  MulMod (y, m LongInteger)
  Div2 (y, r LongInteger)
  Gcd (y LongInteger)
  Lcm (y LongInteger)
  Pow (y LongInteger)
  PowMod (y, m LongInteger)
  Fak (n uint)
  Binom (n, k uint)
  LowFak (n, k uint)
  Stirl2 (n, k uint)
  ProbabylPrime (n int) bool

  Bitlen () uint
  Bit (i int) uint
  SetBit (i int, b bool)
}
