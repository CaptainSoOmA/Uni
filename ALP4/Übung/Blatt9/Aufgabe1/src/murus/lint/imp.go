package lint

// (c) Christian Maurer   v. 130115 - license see murus.go

import (
  "math"; . "math/big"; "strconv"
  . "murus/obj"; "murus/str"
  "murus/col"; "murus/scr"; "murus/box"; "murus/errh"
  "murus/font"; "murus/prt" // ; "murus/pbox";
)
type (
  Imp struct {
           n *Int
        f, b col.Colour
           t font.Font
//         nan bool
             }
)
var (
  zero, one, max, max32 *Imp = New(0), New(0), New(0), New(0)
  tmp, tmp1 *Imp = New(0), New(0)
  bx *box.Imp = box.New ()
//  px *pbox.Imp = pbox.New ()
)


func (x *Imp) impn (Y Any) *Int {
//
  x, ok:= Y.(*Imp)
  if ! ok { TypeNotEqPanic (x, Y) }
  return x.n
}


func New (n int) *Imp {
//
  return New64 (int64(n))
}


func New64 (n int64) *Imp {
//
  x:= new (Imp)
  x.n = NewInt (n)
  x.f, x.b = col.ScreenF, col.ScreenB
  return x
}


func NewNat (n uint) *Imp {
//
  return New64 (int64(n))
}


func NewReal (r float64) *Imp {
//
  x:= New (0)
  x.SetReal (r)
  return x
}


func (x *Imp) Empty () bool {
//
  return x.n.Cmp (zero.n) == 0
//  return x.nan
}


func (x *Imp) Clr () {
//
  x.n.SetInt64 (0)
//  x.nan = true
}


func (x *Imp) Eq (Y Object) bool {
//
  return x.n.Cmp (x.impn (Y)) == 0
}


func (x *Imp) Clone () Object {
//
  y:= New (0)
  y.Copy (x)
  return y
}


func (x *Imp) Copy (Y Object) {
//
  x.n.Set (x.impn (Y))
}


func (x *Imp) Less (Y Object) bool {
//
  return x.n.Cmp (x.impn (Y)) == -1
}


func (x *Imp) Geq0 () bool {
//
  return x.n.Sign () > 0
}


func (x *Imp) ChSign () {
//
  x.n.Neg (x.n)
}


func (x *Imp) Codelen () uint {
//
  return uint(len (x.n.Bytes()))
}


func (x *Imp) Encode () []byte {
//
  return x.n.Bytes ()
}


func (x *Imp) Decode (b []byte) {
//
  x.n.SetBytes (b)
}


func (x *Imp) Set (n int) {
//
  x.n.SetInt64 (int64(n))
}


func (x *Imp) Set32 (n int32) {
//
  x.n.SetInt64 (int64(n))
}


func (x *Imp) Set64 (n int64) {
//
  x.n.SetInt64 (int64(n))
}


func (x *Imp) SetColours (f, b col.Colour) {
//
  x.f, x.b = f, b
}


func (x *Imp) Write (l, c uint) {
//
  s:= x.String ()
  c0:= c
  for n:= 0; n < len (s); n++ {
    scr.Write1 (s[n], l, c)
    if c + 1 < scr.NColumns () {
      c ++
    } else if l + 2 < scr.NLines () {
      l ++
      c = c0
    } else {
      break
    }
  }
}


func (x *Imp) Edit (l, c uint) {
//
  s:= x.String()
  w:= uint(len (s))
  N:= scr.NColumns()
  if c >= N - w {
    x.Write (l, c)
    errh.Error ("zu wenig Platz auf dem Bildschirm", 0) // TODO
    return
  }
  bx.Wd (N - 1 - c)
  bx.Edit (&s, l, c)
  for {
    if x.Defined (s) {
      break
    } else {
      errh.Error ("keine Zahl", 0)
    }
  }
  x.Write (l, c)
}


func (x *Imp) SetFont (f font.Font) {
//
  x.t = f
}


func (x *Imp) Print (l, c uint) {
//
  s:= x.String ()
  c0:= c
  for i:= 0; i < len (s); i++ {
    prt.Print1 (s[i], l, c, x.t)
    if c + 1 < prt.NColumns () {
      c ++
    } else if l + 2 < prt.NLines () {
      l ++
      c = c0
    } else {
      break
    }
  }
  prt.GoPrint ()
}


func (x *Imp) Len () uint {
//
  return uint(len (x.String()))
}


func (x *Imp) Odd () bool {
//
  return x.n.Bit (0) == 1
}


func (x *Imp) Val () int {
//
  n:= x.n.Int64()
  if n < 1<<32 - 1 {
    return int(n)
  }
  return 0
}


func (x *Imp) Val64 () int64 {
//
  if x.Less (max) {
    return x.n.Int64()
  }
  return 0 // ? TODO
}


func (x *Imp) RealVal () float64 {
//
  r, err:= strconv.ParseFloat (x.n.String(), 64)
  if err != nil { return math.NaN() }
  if x.n.Sign () < 0 { r = -r }
  return r
}


func (x *Imp) SetReal (r float64) {
//
  i, _:= math.Modf (r + 0.5)
  s:= strconv.FormatFloat (i, 'f', -1, 64)
  p:= uint(0)
  if str.Contains (s, '.', &p) {
    s = str.Part (s, 0, p)
  }
  if _, ok:= x.n.SetString (s, 10); ! ok {
    x.n.SetInt64 (0)
  }
}


func (x *Imp) Defined (s string) bool {
//
  _, ok:= x.n.SetString (s, 10)
  if ok {
    return true
  }
  x.n.SetInt64 (0) // nan
  return false
}


func (x *Imp) String () string {
//
  return x.n.String ()
}


func (x *Imp) SumDigits () uint {
//
  tmp.n.Abs (x.n)
  a:= uint(0)
  for _, c:= range (tmp.n.String()) {
    a += uint(c)
  }
  return a
}


func (x *Imp) Null () bool {
//
  return x.n.Sign () == 0
}


func (x *Imp) Add (Y, Z Adder) {
//
  x.n.Add (x.impn (Y), x.impn (Z))
}


func (x *Imp) Plus (Y Adder) {
//
  x.n.Add (x.n, x.impn (Y))
}


func (x *Imp) Inc () {
//
  x.n.Add (x.n, one.n)
}


func (x *Imp) Sub (Y, Z Adder) {
//
  x.n.Sub (x.impn (Y), x.impn (Z))
}


func (x *Imp) Minus (Y Adder) {
//
  x.n.Sub (x.n, x.impn (Y))
}


func (x *Imp) Dec () {
//
  x.n.Sub (x.n, one.n)
}


func (x *Imp) One () bool {
//
  return x.Eq (one)
}


func (x *Imp) Mul (Y, Z Multiplier) {
//
  x.n.Mul (x.impn (Y), x.impn (Z))
}


func (x *Imp) Times (Y Multiplier) {
//
  x.n.Mul (x.n, x.impn (Y))
}


func (x *Imp) Sqr () {
//
  x.n.Mul (x.n, x.n)
}


func (x *Imp) Div (Y, Z Multiplier) {
//
  zn:= x.impn (Z)
  if zn.Cmp (zero.n) == 0 { DivBy0Panic() }
  x.n.Quo (x.impn (Y), zn)
}


func (x *Imp) DivBy (Y Multiplier) {
//
  x.Div (x, Y)
}


func (x *Imp) MulMod (Y, M LongInteger) {
//
  x.n.Mul (x.n, x.impn (Y)) // not efficient
  x.n.Mod (x.n, x.impn (M))
}


func (x *Imp) Div2 (Y, R LongInteger) {
//
  yn:= x.impn (Y)
  r, ok:= R.(*Imp)
  if ! ok { NotCompatiblePanic() }
  if yn.Cmp (zero.n) == 0 { DivBy0Panic() }
  _, r.n = x.n.QuoRem (x.n, yn, one.n)
}


func (x *Imp) Gcd (Y LongInteger) {
//
  yn:= x.impn (Y)
  if x.n.Sign () <= 0 || yn.Sign () <= 0 {
    return
  }
  x.n.GCD (tmp.n, tmp1.n, x.n, yn)
}


func (x *Imp) Lcm (Y LongInteger) {
//
  yn:= x.impn (Y)
  if x.n.Sign () <= 0 || yn.Sign () <= 0 {
    return
  }
  x.n.Mul (x.n, yn)
  tmp.n.Set (yn)
  tmp.Gcd (x)
  x.n.Quo (x.n, tmp.n)
}


func (x *Imp) Pow (Y LongInteger) {
//
  x.n.Exp (x.n, x.impn (Y), nil)
}


func (x *Imp) PowMod (Y, M LongInteger) {
//
  x.n.Exp (x.n, x.impn (Y), x.impn (M))
}


func (x *Imp) Fak (n uint) {
//
  x.n.MulRange (1, int64(n))
}


func (x *Imp) ProbabylPrime (n int) bool {
//
  return x.n.ProbablyPrime (n)
}


func (x *Imp) Binom (n, k uint) {
//
  x.n.Binomial (int64(n), int64(k))
}


func (x *Imp) LowFak (n, k uint) {
//
  if n < k {
    x.n.SetInt64 (0)
    return
  }
  if k == 0 {
    x.n.SetInt64 (1)
    return
  }
  x.n.MulRange (int64(n - k + 1), int64(n))
}


func (x *Imp) Stirl2 (n, k uint) {
//
  x.n.SetInt64 (0)
  if n < k {
    return
  }
  if k == 0 {
    if n == 0 {
      x.n.SetInt64 (1)
    }
    return
  }
  tmp.n.SetInt64 (1)
  e:= k % 2 != 0
  nn, ii:= New64 (int64(n)), New (1)
  for i:= uint(1); i <= k; i++ {
    tmp.n.Mul (tmp.n, tmp1.n.SetInt64 (int64(k - i + 1)))
    tmp.n.Div (tmp.n, ii.n)
    tmp1.n.Mul (tmp1.n.Exp (ii.n, nn.n, nil), tmp.n)
    if e {
      x.n.Add (x.n, tmp1.n)
    } else {
      x.n.Sub (x.n, tmp1.n)
    }
    e = ! e
    ii.Inc ()
  }
  x.n.Div (x.n, tmp.n.MulRange (1, int64(k)))
}


func (x *Imp) Bitlen () uint {
//
  return uint(x.n.BitLen ())
}


func (x *Imp) Bit (i int) uint {
//
  return x.n.Bit (i)
}


func (x *Imp) SetBit (i int, b bool) {
//
  u:= uint(0)
  if b { u++ }
  x.n.SetBit (x.n, i, u)
}


func init () {
//
  bx.Wd (64)
  zero.n.SetInt64 (0)
  one.n.SetInt64 (1)
  max.n.SetInt64 (1 << 63 - 1)
  max32.n.SetInt64 (1 << 32 - 1)
  var _ LongInteger = New (0)
}
