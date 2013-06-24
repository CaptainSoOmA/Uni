package brat

// (c) Christian Maurer   v. 130526 - license see murus.go

import (
  "math"
  . "murus/obj"; "murus/str"
  "murus/col"; "murus/box"
  "murus/font"; "murus/pbox"
  "murus/nat"
  "murus/errh"
)
const
  max = 1e9 // numerator and denominator with at most 9 digits
type
  Imp struct {
         num,
       denom uint
        geq0 bool
      cF, cB col.Colour
         fnt font.Font
            }
var (
  Gegenzahl, reciprocal *Imp = New (), New ()
  bx *box.Imp = box.New ()
  pbx *pbox.Imp = pbox.New ()
)


func (x *Imp) imp (Y Any) *Imp {
//
  y, ok:= Y.(*Imp)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}


func New () *Imp {
//
  x:= new (Imp)
  x.cF, x.cB = col.ScreenF, col.ScreenB
  if col.Eq (x.cF, col.White) && col.Eq (x.cB, col.Black) { x.cF = col.LightWhite } // Firlefanz
  x.geq0 = true
  return x
}


func (x *Imp) Empty () bool {
//
  return x.denom == 0
}


func (x *Imp) Numbers () (uint, uint) {
//
  return x.num, x.denom
}


func (x *Imp) Clr () {
//
  x.num, x.denom = 0, 0
  x.geq0 = true
}


func (x *Imp) Copy (Y Object) {
//
  y:= x.imp (Y)
  x.num, x.denom = y.num, y.denom
  x.geq0 = y.geq0
}


func (x *Imp) Clone () Object {
//
  y:= New ()
  y.Copy (x)
  return y
}


func (x *Imp) Eq (Y Object) bool {
//
  y:= x.imp (Y)
  if y.denom == 0 {
    return x.denom == 0
  } else if x.denom == 0 {
    return false
  }
  if x.num == 0 { return y.num == 0 }
  if y.num == 0 { return x.num == 0 }
  if x.geq0 != y.geq0 { return false }
  return uint64 (x.num) * uint64 (y.denom) == uint64 (y.num) * uint64 (x.denom)
}


func (x *Imp) Less (Y Object) bool {
//
  y:= x.imp (Y)
  if y.denom == 0 {
    return x.denom == 0
  } else if x.denom == 0 {
    return false
  }
  if x.num == 0 && y.num == 0 { return false }
  if x.num == 0 { return y.geq0 }
  if y.num == 0 { return ! x.geq0 }
  if x.geq0 && ! y.geq0 { return false }
  if ! x.geq0 && y.geq0 { return true }
  p, q:= uint64 (x.num) * uint64 (y.denom), uint64 (y.num) * uint64 (x.denom)
  if x.geq0 { return p < q } // y.geq0
  return p > q // ! x.geq0, ! y.geq0
}


func (x *Imp) Set (n, d int) bool {
//
  if d == 0 {
    x.num, x.denom = 1, 0
    return false
  }
  if n == 0 {
    x.num, x.denom = 0, 1
    return true
  }
  x.geq0 = n > 0
  if ! x.geq0 {
    n = -n
  }
  if d < 0 {
    d = -d
    x.geq0 = ! x.geq0
  }
  x.num, x.denom = uint(n), uint(d)
  x.reduce ()
  return true
}


func (x *Imp) SetNat (n, d uint, geq0 bool) {
//
  x.num, x.denom = n, d
  x.reduce ()
  x.geq0 = geq0
  if x.num == 0 { x.geq0 = true }
}


func (x *Imp) RealVal () float64 {
//
  if x.denom == 0 {
    return math.Inf(1)
  }
  r:= float64 (x.num) / float64 (x.denom)
  if ! x.geq0 {
    r = - r
  }
  return r
}


func (x *Imp) Defined (s string) bool {
//
  x.Clr ()
  if str.Empty (s) {
    return true
  }
  str.Move (&s, true)
  x.geq0 = s[0] != '-'
  switch s[0] { case '+', '-':
    str.Rem (&s, 0, 1)
  }
  n:= str.ProperLen (s)
  var p uint
  if str.Contains (s, '/', &p) {
    s1:= str.Part (s, p + 1, n - p - 1)
    if nat.Defined (&x.denom, s1) {
      if x.denom >= max {
        return false
      }
    } else {
      return false
    }
  } else {
    p = n
    x.denom = 1
  }
  s1:= str.Part (s, 0, p)
  if nat.Defined (&x.num, s1) {
    if x.num >= max {
      return false
    }
  } else {
    return false
  }
  return true
}


func (x *Imp) String () string {
//
  if x.denom == 0 {
    return "Zähler/Nenner"
  }
  s:= nat.StringFmt (x.num, 9, false)
  str.Move (&s, true)
  str.RemSpaces (&s)
  if ! x.geq0 {
    s = "-" + s
  }
  if x.denom == 1 {
    return s
  }
  s += "/"
  t:= nat.StringFmt (x.denom, 9, false)
  str.Move (&t, true)
  str.RemSpaces (&t)
  return s + t
}


func (x *Imp) SetColours (f, b col.Colour) {
//
  x.cF, x.cB = f, b
}


func (x *Imp) Write (l, c uint) {
//
  bx.Colours (x.cF, x.cB)
  bx.Write (x.String (), l, c)
}


func (x *Imp) Edit (l, c uint) {
//
  s:= x.String ()
  bx.Colours (x.cF, x.cB)
  for {
    bx.Edit (&s, l, c)
    if x.Defined (s) {
      break
    } else {
      errh.Error ("Format: Z = \"0\"|\"1\"|...|\"9\", [\"+\"|\"-\"] Z{Z}[/Z{Z}]; jeweils maximal 9 Ziffern", 0)
    }
  }
  x.reduce ()
  x.Write (l, c)
}


func (x *Imp) SetFont (f font.Font) {
//
  x.fnt = f
}


func (x *Imp) Print (l, c uint) {
//
  pbx.SetFont (x.fnt)
  pbx.Print (x.String (), l, c)
}


func (x *Imp) Codelen () uint {
//
  return 2 * 4 + 1
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, x.Codelen())
  copy (b[:4], Encode (x.num))
  copy (b[4:8], Encode (x.denom))
  b[9] = 0
  if x.geq0 { b[9] = 1 }
  return b
}


func (x *Imp) Decode (b []byte) {
//
  x.num = Decode (uint(0), b[:4]).(uint)
  x.denom = Decode (uint(0), b[4:8]).(uint)
  x.geq0 = b[9] == 1
}


func (x *Imp) Integer () bool {
//
  return x.denom == 1
}


func (x *Imp) Null () bool {
//
  if x.denom == 0 { return false }
  return x.num == 0
}


func (x *Imp) One () bool {
//
  if x.denom == 0 { return false }
  return x.num == x.denom
}


func (x *Imp) GeqNull () bool {
//
  if x.denom == 0 { return false }
  return x.geq0
}


func (x *Imp) reduce () {
//
  if x.denom == 0 {
    x.num = 0
    return
  }
  if x.num == 0 {
    x.denom = 1
    return
  }
  g:= nat.Gcd (x.num, x.denom)
  x.num, x.denom = x.num / g, x.denom / g
}


func gcd (a, b uint64) uint64 {
//
  if a < b { a, b = b, a }
  if b == 0 { return a }
  return gcd (a % b, b)
}


const
  max64 = uint64(1<<64 - 1)


func (x *Imp) Add (Y, Z Adder) {
//
  y, z:= x.imp (Y), x.imp (Z)
  if y.denom == 0 || z.denom == 0 {
    x.Clr ()
    return
  }
  if y.num == 0 {
    x.Copy (z)
    return
  }
  if z.num == 0 {
    x.Copy (y)
    return
  }
  a, b:= uint64 (y.num) * uint64 (z.denom), uint64 (z.num) * uint64 (y.denom)
  var n uint64
  if y.geq0 {
    if z.geq0 {
      if a <= max64 - b {
        n = a + b
        x.geq0 = true
      } else {
        x.Clr ()
        return
      }
    } else {
      if a >= b {
        n = a - b
        x.geq0 = true
      } else {
        n = b - a
        x.geq0 = false
      }
    }
  } else { // ! y.geq0
    if z.geq0 {
      if a < b {
        n = b - a
        x.geq0 = true
      } else {
        n = a - b
        x.geq0 = false
      }
    } else {
      if a < max64 - b {
        n = a + b
        x.geq0 = false
      } else {
        x.Clr ()
        return
      }
    }
  }
  d:= uint64 (y.denom) * uint64 (z.denom)
  g:= gcd (n, d)
  n, d = n / g, d / g
  if n > uint64(max) || d > uint64(max) {
    x.Clr ()
    return
  }
  x.num, x.denom = uint(n), uint(d)
  x.reduce ()
}


func (x *Imp) Plus (Y Adder) {
//
  x.Add (x, Y)
}


func (x *Imp) changeSign () {
//
  x.geq0 = ! x.geq0
  if x.num == 0 { x.geq0 = true }
}


func (x *Imp) Sub (Y, Z Adder) {
//
  y, z:= x.imp (Y), x.imp (Z)
  zg:= z.geq0
  z.geq0 = ! z.geq0
  x.Add (y, z)
  z.geq0 = zg
}


func (x *Imp) Minus (Y Adder) {
//
  x.Sub (x, Y)
}


func (x *Imp) Mul (Y, Z Multiplier) {
//
  y, z:= x.imp (Y), x.imp (Z)
  if y.denom == 0 || z.denom == 0 {
    x.Clr ()
    return
  }
  if y.num == 0 || z.num == 0 {
    x.num, x.denom = 0, 1
    x.geq0 = true
    return
  }
  if y.num == y.denom {
    x.num, x.denom = z.num, z.denom
    x.geq0 = z.geq0
    return
  }
  if z.num == z.denom {
    x.num, x.denom = y.num, y.denom
    x.geq0 = y.geq0
    return
  }
  n, d:= uint64 (y.num) * uint64 (z.num), uint64 (y.denom) * uint64 (z.denom)
  g:= gcd (n, d)
  n, d = n / g, d / g
  if n > uint64(max) || d > uint64(max) {
    x.Clr ()
    return
  }
  x.num, x.denom = uint(n), uint(d)
  x.geq0 = y.geq0 == z.geq0
  x.reduce ()
}


func (x *Imp) Times (Y Multiplier) {
//
  x.Mul (x, Y)
}


func (x *Imp) Sqr () {
//
  x.Mul (x, x)
}


func (x *Imp) Invert () {
//
  if x.denom == 0 {
    x.num = 0
  } else {
    x.num, x.denom = x.denom, x.num
  }
}


func (x *Imp) Div (Y, Z Multiplier) {
//
  y, z:= x.imp (Y), x.imp (Z)
  if y.denom == 0 || z.num == 0 || z.denom == 0 {
    x.Clr ()
    return
  }
  inv:= z.Clone ().(*Imp)
  inv.Invert ()
  x.Mul (y, inv)
}


func (x *Imp) DivBy (Y Multiplier) {
//
  x.Div (x, Y)
}

/*
type Operation byte; const (ADD = iota; SUB; MUL; DIV ) // TODO interface in murus/obj

func (x *Imp) Operate (Y, Z Rational, op Operation) { // TODO Rational -> ...
//
  switch op {
  case ADD:
    x.Sum (Y, Z)
  case SUB:
    x.Diff (Y, Z)
  case MUL:
    x.Prod (Y, Z)
  case DIV:
    x.Quot (Y, Z)
  }
}
*/

func init () {
//
  bx.Wd (1 + 9 + 1 + 9) // sign, numerator, fraction bar, denominator
  var _ Rational = New ()
  var _ Adder = New ()
}
