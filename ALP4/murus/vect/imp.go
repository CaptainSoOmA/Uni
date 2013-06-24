package vect

// (c) Christian Maurer   v. 130115 - license see murus.go

import (
  "math"; "strconv"
  . "murus/spc"; . "murus/obj"; "murus/str"
  "murus/col"; "murus/box"; "murus/errh"
  "murus/font"; "murus/pbox"
)
const (
  um = math.Pi / 180.0
  null = 0.0
)
type
  Imp struct {
           x [NDirs]float64
             }
var (
  temp, temp1 *Imp = New(), New()
  bx *box.Imp = box.New()
  pbx *pbox.Imp = pbox.New()
)


func (x *Imp) imp (Y Any) *Imp {
//
  y, ok:= Y.(*Imp)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}


func New () *Imp {
//
  return new (Imp)
}


func (x *Imp) Set3 (x0, x1, x2 float64) {
//
  x.x[Right], x.x[Front], x.x[Top] = x0, x1, x2
}


func (x *Imp) Set (c Coord) {
//
  for d:= D0; d < NDirs; d++ {
    x.x[d] = c[d]
  }
}

func (x *Imp) Coord3() (float64, float64, float64) {
//
  return x.x[Right], x.x[Front], x.x[Top]
}


func (x *Imp) Coord (d Direction) float64 {
//
  return x.x[d]
}


func (v *Imp) DefPolar (x, y, z, r, phi, theta float64) {
//
  v.x[Right] = x + r * math.Cos (phi * um) * math.Sin (theta * um)
  v.x[Front] = y + r * math.Sin (phi * um) * math.Sin (theta * um)
  v.x[Top]   = z + r                       * math.Cos (theta * um)
}


func (x *Imp) Project (A, B, C Vector) {
//
  a, b, c:= x.imp (A), x.imp (B), x.imp (C)
  for d:= D0; d < NDirs; d++ {
    a.x[d], b.x[d], c.x[d] = null, null, null
  }
  a.x[Right], b.x[Front], c.x[Top] = x.x[Right], x.x[Front], x.x[Top]
}


func (x *Imp) Empty () bool {
//
  a:= null
  for d:= D0; d < NDirs; d++ {
    a += math.Abs (x.x[d])
  }
  return a < epsilon
}


func (x *Imp) Clr () {
//
  x.Set3 (null, null, null)
}


func (x *Imp) Copy (Y Object) {
//
  y:= x.imp (Y)
  for d:= D0; d < NDirs; d++ {
    x.x[d] = y.x[d]
  }
}


func (x *Imp) Clone () Object {
//
  y:= New ()
  y.Copy (x)
  return y
}


func (x *Imp) Eq (Y Object) bool {
//
  y, a:= x.imp (Y), null
  for d:= D0; d < NDirs; d++ {
    a += math.Abs (x.x[d] - y.x[d])
  }
  return a < epsilon
}


func (x *Imp) Less (Y Object) bool {
//
  return false
}


func (x *Imp) Int (Y Vector) float64 {
//
  y, a:= x.imp (Y), null
  for d:= D0; d < NDirs; d++ {
    a += x.x[d] * y.x[d]
  }
  return a
}


func (x *Imp) Cross (Y Vector) {
//
  y:= x.imp (Y)
  var a [NDirs]float64
  for d:= D0; d < NDirs; d++ {
    a[d] = x.x[d]
  }
  for d:= D0; d < NDirs; d++ {
    d1, d2:= Next (d), Prev (d)
    x.x[d] = a[d1] * y.x[d2] - a[d2] * y.x[d1]
  }
}


func (x *Imp) Ext (Y, Z Vector) {
//
  y, z:= x.imp (Y), x.imp (Z)
  for d:= D0; d < NDirs; d++ {
    d1, d2:= Next (d), Prev (d)
    x.x[d] = y.x[d1] * z.x[d2] - z.x[d2] * y.x[d1]
  }
}


func (x *Imp) Collinear (Y Vector) bool {
//
  y:= x.imp (Y)
  if x.Empty () || y.Empty () {
    return true
  }
  temp.Copy (x)
  temp.Cross (y)
  return temp.Empty ()
}


// >>> deprecated !!!
func (x *Imp) Scale (a float64, Y Vector) {
//
  y:= Y.(*Imp)
  for d:= D0; d < NDirs; d++ {
    x.x[d] = a * y.x[d]
  }
}


func (x *Imp) Dilate (a float64) { // TODO name ?
//
  for d:= D0; d < NDirs; d++ {
    x.x[d] *= a
  }
}


func (x *Imp) Null () bool {
//
  return x.Empty ()
}


func (x *Imp) Add (Y, Z Adder) {
//
  y, z:= x.imp (Y), x.imp (Z)
  for d:= D0; d < NDirs; d++ {
    x.x[d] = y.x[d] + z.x[d]
  }
}


func (x *Imp) Plus (Y Adder) {
//
  y:= Y.(*Imp)
  for d:= D0; d < NDirs; d++ {
    x.x[d] += y.x[d]
  }
}


func (x *Imp) Sub (Y, Z Adder) {
//
  y, z:= x.imp (Y), x.imp (Z)
  for d:= D0; d < NDirs; d++ {
    x.x[d] = y.x[d] - z.x[d]
  }
}


func (x *Imp) Minus (Y Adder) {
//
  y:= Y.(*Imp)
  for d:= D0; d < NDirs; d++ {
    x.x[d] -= y.x[d]
  }
}


func (x *Imp) Parametrize (Y, Z Vector, t float64) {
//
  y, z:= x.imp (Y), x.imp (Z)
  for d:= D0; d < NDirs; d++ {
    x.x[d] = y.x[d] + t * (z.x[d] - y.x[d])
  }
}


func (x *Imp) Len () float64 {
//
  return math.Sqrt (x.Int (x))
}


func (x *Imp) Distance (Y Vector) float64 {
//
  y:= Y.(*Imp)
  a, s:= null, null
  for d:= D0; d < NDirs; d++ {
    s = x.x[d] - y.x[d]
    a += s * s
  }
  return math.Sqrt (a)
}


func (x *Imp) Centre (Y, Z Vector) float64 {
//
  y, z:= x.imp (Y), x.imp (Z)
  a, s:= null, null
  for d:= D0; d < NDirs; d++ {
    x.x[d] = (y.x[d] + z.x[d]) / 2.0
    s = y.x[d] - z.x[d]
    a += s * s
  }
  return math.Sqrt (a) / 2.0
}


func (x *Imp) Flat (Y Vector) bool {
//
  y:= Y.(*Imp)
  return math.Abs (x.x[Top] - y.x[Top]) < epsilon
}


func (x *Imp) Norm () {
//
  a:= math.Sqrt (x.Int (x))
  for d:= D0; d < NDirs; d++ {
    x.x[d] /= a
  }
}


func (x *Imp) Normed () bool {
//
  return math.Abs (x.Len () - 1.0) < epsilon
}


func (x *Imp) Rot (Y Vector, a float64) {
//
  y:= Y.(*Imp)
  for a <= -180. { a += 360. }
  for a > 180. { a -= 360. }
  if x.Collinear (y) { return } // error
//  d.Norm () // avoid rounding errors
  c:= math.Cos (a * um)
// x = cos(a) * x0 + <x0, y> * (1 - cos(a)) * y + sin(a) * [y, x0]
//  temp.Scale ((1. - c) * x.Int (y), y)
  temp.Copy (y)
  temp.Dilate ((1. - c) * x.Int (y))
  temp1.Copy (y)
  temp1.Cross (x)
  temp1.Dilate (math.Sin (a * um))
  x.Dilate (c)
  x.Plus (temp)
  x.Plus (temp1)
}


func (x *Imp) Defined (s string) bool {
//
  x.Clr()
  n:= uint(len (s))
  if n < 7 { return false }
  if s[0] != '(' || s[n - 1] != ')' { return false }
  t:= str.Part (s, 1, n - 2) + ","
  var p uint
  for d:= D0; d < NDirs; d++ {
    if ! str.Contains (t, ',', &p) {
      return false
    }
    r, err:= strconv.ParseFloat (t[:p], 64)
    if err == nil {
      x.x[d] = r
    } else {
      return false
    }
    str.Rem (&t, 0, p+1)
  }
  return true
}


func (x *Imp) String () string {
//
  s:= "("
  for d:= D0; d < NDirs; d++ {
    s += strconv.FormatFloat (x.x[d], 'f', 2, 64)
    if d == NDirs - 1 {
      s += ") "
    } else {
      s += ", "
    }
  }
  return s
}


func (x *Imp) SetColours (f, b col.Colour) {
//
  bx.Colours (f, b)
}


func (x *Imp) Edit (l, c uint) {
// func (v *Imp) Edit (p ... uint) {
//
  s:= x.String()
  m:= uint(len (s))
  bx.Wd (m)
  for {
    bx.Edit (&s, l, c)
//    bx.Edit (&s, p[0], p[1])
    if x.Defined (s) {
      break
    } else {
      errh.Error ("kein Vektor", 0)
    }
  }
}


func (x *Imp) Write (l, c uint) {
// func (v *Imp) Write (p ... uint) {
//
  bx.Wd (uint (len(x.String())))
  bx.Write (x.String(), l, c)
//  bx.Write (v.String(), p[0], p[1])
}


func (x *Imp) SetFont (f font.Font) {
//
  pbx.SetFont (f)
}


func (x *Imp) Print (l, c uint) {
//
  pbx.Print (x.String(), l, c)
}


var
  clfloat = Codelen(null)


func (x *Imp) Codelen () uint {
//
  return uint(NDirs) * clfloat
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, x.Codelen())
  i, a:= uint(0), clfloat
  for d:= D0; d < NDirs; d++ {
    copy (b[i:i+a], Encode (x.x[d]))
    i += a
  }
  return b
}


func (x *Imp) Decode (b []byte) {
//
  i, a:= uint(0), clfloat
  for d:= D0; d < NDirs; d++ {
    x.x[d] = Decode (null, b[i:i+a]).(float64)
    i += a
  }
}


func (V *Imp) Minimax (N, X Vector) {
//
  Min, n:= N.(*Imp)
  Max, x:= X.(*Imp)
  if ! n || ! x { return }
  for d:= D0; d < NDirs; d++ {
    if V.x[d] < Min.x[d] {
      Min.x[d] = V.x[d]
    }
    if V.x[d] > Max.x[d] {
      Max.x[d] = V.x[d]
    }
  }
}


func init () { var _ Adder = New (); var _ Vector = New () }
