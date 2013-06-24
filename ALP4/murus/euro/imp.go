package euro

// (c) Christian Maurer   v. 130526 - license see murus.go

import (
  "math"
  . "murus/obj"; "murus/str"
  "murus/col"; "murus/box"; "murus/errh"
  "murus/nat"
  "murus/font"; "murus/pbox"
)


const (
  hundred = uint(100)
  tenMillions = uint(1e7)
  undefined = uint(tenMillions * hundred)
  nDigits = 7 // höchstens 9.999.999 Euro
  length = nDigits + 1 /* Komma */ + 2
)
type (
  Imp struct {
        cent uint
      cF, cB col.Colour
          fo font.Font
             }
  Texte [length]byte
)
var (
  bx *box.Imp = box.New()
  pbx *pbox.Imp = pbox.New()
)


func (x *Imp) impc (Y Any) uint {
//
  y, ok:= Y.(*Imp)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y.cent
}


func New () *Imp {
//
  x:= new (Imp)
  x.Clr()
  x.SetColours (col.LightGreen, col.Black)
  return x
}


func (x *Imp) Empty() bool {
//
  return x.cent >= undefined
}


func (x *Imp) Clr () {
//
  x.cent = undefined
}


func (x *Imp) Copy (Y Object) {
//
  x.cent = x.impc (Y)
}


func (x *Imp) Clone () Object {
//
  y:= New ()
  y.Copy (x)
  return y
}


func (x *Imp) Eq (Y Object) bool {
//
  return x.cent == x.impc (Y)
}


func (x *Imp) Less (Y Object) bool {
//
  return x.cent < x.impc (Y)
}


func (x *Imp) Val () uint {
//
  return x.cent
}


func (x *Imp) Set (c uint) bool {
//
  x.cent = c
  return x.cent < undefined
}


func (x *Imp) Val2 () (uint, uint) {
//
  return x.cent / hundred, x.cent % hundred
}


func (x *Imp) Set2 (e, c uint) bool {
//
  if e >= tenMillions || c >= hundred {
    x.cent = undefined
  } else {
    x.cent = hundred * e
    x.cent += c
  }
  return x.cent < undefined && c < hundred
}


func (x *Imp) RealVal () float64 {
//
  return float64 (x.cent) / float64 (hundred)
}


func (x *Imp) SetReal (r float64) bool {
//
  if r >= 0. && r < float64(tenMillions) {
    x.cent = uint(math.Trunc (float64 (hundred) * r + 0.5))
  } else {
    x.cent = undefined
  }
  return x.cent < undefined
}


func (x *Imp) Null () bool {
//
  return x.cent == 0
}


func (x *Imp) Add (Y, Z Adder) {
//
  yc, zc:= x.impc (Y), x.impc (Z)
  x.cent = yc + zc
  if x.cent > undefined { x.cent = undefined }
}


func (x *Imp) Plus (Y Adder) {
//
  x.Add (x, Y)
}


func (x *Imp) Sub (Y, Z Adder) {
//
  yc, zc:= x.impc (Y), x.impc (Z)
  if x.cent == undefined || yc == undefined || zc == undefined { return }
  if yc >= zc {
    x.cent = yc - zc
  } else {
    x.cent = undefined
  }
}


func (x *Imp) Minus (Y Adder) {
//
  x.Sub (x, Y)
}


func (x *Imp) Operate (Faktor, Divisor uint) {
//
  if x.cent == undefined { return }
  if Divisor == 0 { x.cent = undefined; return }
  if Faktor == 0 { x.cent = 0; return }
  if x.cent / Divisor < (tenMillions * hundred) / Faktor {
    x.cent *= Faktor
    x.cent += Divisor / 2
    x.cent /= Divisor
  } else {
    x.cent = undefined
  }
}


func toThe (q float64, n uint) float64 {
//
  if n == 0 {
    return 1.
  }
  return q * toThe (q, n - 1)
}


func (x *Imp) ChargeInterest (p, n uint) {
//
  if x.cent == undefined { return }
  f:= toThe (1.0 + float64 (p) / 10000.0, n)
  b:= float64 (x.cent) * f + 0.5
  if b < float64 (tenMillions * hundred) {
    x.cent = uint(math.Trunc (b))
  } else {
    x.cent = undefined
  }
}


func (x *Imp) Round (Y Euro) {
//
  yc:= x.impc (Y)
  if x.cent == undefined || yc == undefined { return }
  x.cent = yc * (x.cent / yc)
}


func (x *Imp) String () string {
//
  if x.Empty () { return str.Clr (length) }
  return nat.StringFmt (x.cent / hundred, nDigits, false) + "," +
         nat.StringFmt (x.cent % hundred, 2, true)
}


func (x *Imp) Defined (s string) bool {
//
  if str.Empty (s) {
    x.cent = undefined
    return true
  }
  a, t, P, L:= nat.DigitSequences (s)
  var k uint
  hatKomma:= str.Contains (s, ',', &k)
  if ! hatKomma {
    hatKomma = str.Contains (s, '.', &k)
  }
  var n uint
  if ! nat.Defined (&n, t[0]) {
    return false
  }
  switch a { case 1:
    if hatKomma && k < P[0] { // Komma vor der Ziffernfolge
      switch L[0] { case 1:
        x.cent = 10 * n
      case 2:
        x.cent = n
      default:
        return false
      }
      return true
    }
    if hatKomma && k >= P[0] + L[0] || ! hatKomma {
      if L[0] <= nDigits {
        x.cent = hundred * n
        return true
      }
    }
  case 2:
    if ! hatKomma { return false }
    if k < P[0] + L[0] || P[1] <= k { return false }
    if L[0] > nDigits {
      return false
    } else {
      x.cent = hundred * n
    }
    if ! nat.Defined (&n, t[1]) { return false }
    switch L [1] { case 1:
      x.cent += 10 * n
    case 2:
      x.cent += n
    default:
      return false
    }
    return true
  }
  return false
}


func (x *Imp) SetColours (f, b col.Colour) {
//
  x.cF, x.cB = f, b
}


func (x *Imp) Write (l, c uint) {
//
  bx.Colours (x.cF, x.cB)
  bx.Write (x.String(), l, c)
}


func (x *Imp) Edit (l, c uint) {
//
  s:= x.String()
  bx.Colours (x.cF, x.cB)
  for {
    bx.Edit (&s, l, c)
    if x.Defined (s) {
      break
    } else {
      errh.Error ("kein Geldbetrag", 0) // l + 1, c)
    }
  }
  x.Write (l, c)
}


func (x *Imp) SetFont (f font.Font) {
//
  x.fo = f
}


func (x *Imp) Print (l, c uint) {
//
  pbx.SetFont (x.fo)
  pbx.Print (x.String(), l, c)
}


func (x *Imp) Codelen () uint {
//
  return Codelen (x.cent)
}


func (x *Imp) Encode () []byte {
//
  bs:= make ([]byte, Codelen (x.cent))
  bs = Encode (x.cent)
  return bs
}


func (x *Imp) Decode (bs []byte) {
//
  x.cent = Decode (x.cent, bs).(uint)
}


func init () {
//
  bx.Wd (length)
//  bx.SetNumerical ()
  var _ Euro = New ()
}
