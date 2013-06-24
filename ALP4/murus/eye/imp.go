package eye

// (c) Christian Maurer   v. 130115 - license see murus.go

import (
  "math"
  . "murus/spc"; "murus/col"
  "murus/errh"
  "murus/vect"
  "murus/gl" // PosLight, Actualize
)
const
  epsilon = 1.0E-6
type
  Imp struct {
      origin,
   originOld,
       focus,
        temp *vect.Imp
         vec[NDirs]*vect.Imp
       delta float64 // Invariant: delta == Distance (Auge, focus)
      colour col.Colour
    flaechig bool
             }
var
  nB, nD, nF uint
//  var e *Imp = New () // -> Anwendung


func New () *Imp {
//
  E:= new (Imp)
  E.origin = vect.New ()
  E.originOld = vect.New ()
  E.focus = vect.New ()
  E.temp = vect.New ()
  for d:= D0; d < NDirs; d++ {
    E.vec[d] = vect.New ()
    E.vec[d].Set (Unit[d])
  }
  E.delta = E.origin.Distance (E.focus)
  E.colour = col.ScreenF
  E.flaechig = false
  return E
}


func (E *Imp) SetLight (n uint) {
//
  gl.PosLight (n, E.origin)
}


func (E *Imp) Actualize () {
//
  gl.Actualize (E.vec[Right], E.vec[Front], E.vec[Top], E.origin)
}


func (E *Imp) DistanceFrom (aim *vect.Imp) float64 {
//
  return E.origin.Distance (aim)
}


func (E *Imp) focusAnpassen () {
//
  E.focus.Scale (E.delta, E.vec[Front])
  E.focus.Plus (E.origin)
  E.Actualize ()
}


func (E *Imp) Distance () float64 {
//
  if math.Abs (E.origin.Distance (E.focus) - E.delta) > epsilon {
    E.focusAnpassen ()
  }
  return E.delta
}


func (E *Imp) Read (A[]vect.Imp) bool {
//
  A[0].Copy (E.originOld)
  if len (A) > 1 {
    A[1].Copy (E.origin)
  }
  return E.flaechig
}


func (E *Imp) Flatten (f bool) {
//
  E.flaechig = f
}


func (E *Imp) Move (d Direction, dist float64) {
//
  nB ++
//  E.vec[Top].Copy (E.vec[Right])
//  E.vec[Top].Cross (E.vec[Front])
  E.vec[Top].Ext (E.vec[Right], E.vec[Front])
  E.vec[Top].Norm ()
  E.originOld.Copy (E.origin)
  E.temp.Scale (dist, E.vec[d])
  E.origin.Plus (E.temp)
  E.focusAnpassen ()
}


func (E *Imp) rotate (d Direction, alpha float64) { // ziemlich abenteuerliche Konstruktion
//
  V1:= E.vec[Next (d)]
  V1.Rot (E.vec[d], alpha)
  V1.Norm ()
  V2:= E.vec[Prev (d)]
//  V2.Copy (E.vec[d])
//  V2.Cross (V1)
  V2.Ext (E.vec[d], V1)
  V2.Norm ()
}


func (E *Imp) Turn (d Direction, alpha float64) {
//
  nD++
  E.rotate (d, alpha)
  E.focusAnpassen ()
}


func (E *Imp) Invert () {
//
  nD++
  E.vec[Right].Dilate (-1.0)
  E.vec[Front].Dilate (-1.0)
  E.focusAnpassen ()
}


func (E *Imp) originAnpassen () {
//
  E.temp.Scale (E.delta, E.vec[Front])
//  E.origin.Copy (E.focus)
//  E.origin.Minus (E.temp)
  E.origin.Sub (E.focus, E.temp)
  E.Actualize ()
}


func (E *Imp) Focus (d float64) {
//
  if d < epsilon { return }
  E.delta = d
  E.originAnpassen ()
}


func (E *Imp) TurnAroundFocus (D Direction, alpha float64) {
//
  if E.delta < epsilon { return }
  nF++
println ("TurnAroundFocus")
  E.rotate (D, - alpha)
println ("rotated")
// Dieser Vorzeichenwechsel ist ein Arbeitsdrumrum um einen mir bisher nicht erklärbaren Fehler.
// Vermutlich liegt das daran, dass ich irgendeine suboptimal dokumentierte Eigenschaft von openGL noch nicht begriffen habe.
  E.originAnpassen ()
}


func (E *Imp) Set (x, y, z, xf, yf, zf float64) {
//
  E.origin.Set3 (x, y, z)
  E.focus.Set3 (xf, yf, zf)
  E.delta = E.origin.Distance (E.focus)
  if E.delta < epsilon { return } // error
  if math.Abs (z - zf) < epsilon { // Blick horizontal
    E.vec[Top].Set (Unit[Top])
//    E.vec[Front].Copy (E.focus)
//    E.vec[Front].Minus (E.origin)
    E.vec[Front].Sub (E.focus, E.origin)
    E.vec[Front].Norm ()
//    E.vec[Right].Copy (E.vec[Front])
//    E.vec[Right].Cross (E.vec[Top])
    E.vec[Right].Ext (E.vec[Front], E.vec[Top])
    E.vec[Right].Norm ()
  } else { // z != zf
    if math.Abs (x - xf) < epsilon && math.Abs (y - yf) < epsilon { // x == xf und y == yf
      E.vec[Right].Set (Unit[Right])
      E.vec[Front].Set (Unit[Top])
      E.vec[Top].Set (Unit[Right])
      if z > zf { // Blick von Top, x -> Right, y -> Top
        E.vec[Front].Dilate (-1.0)
      } else { // z < zf *) // Blick von unten, x -> Right, y -> unten
        E.vec[Top].Dilate (-1.0)
      }
    } else { // x != xf oder y != yf
//      E.vec[Front].Copy (E.focus)
//      E.vec[Front].Minus (E.origin)
      E.vec[Front].Sub (E.focus, E.origin)
      E.vec[Front].Norm ()
      v2:= E.vec[Front].Coord (Top)
      E.vec[Top].Copy (E.vec[Front])
      if z < zf { v2 = -v2 }
      E.temp.Set3 (0., 0., - 1. / v2)
      E.vec[Top].Plus (E.temp)
      E.vec[Top].Norm ()
//      E.vec[Right].Copy (E.vec[Front])
//      E.vec[Right].Cross (E.vec[Top])
      E.vec[Right].Ext (E.vec[Front], E.vec[Top])
      E.vec[Right].Norm ()
    }
  }
  E.Actualize()
}


func Report () {
//
  errh.Error2 ("Bewegungen:", nB, "/ Drehungen:", nD)
}


var (
  stack[]([]byte) = make ([]([]byte), 100)
  v *vect.Imp = vect.New ()
)

// Vielleicht geht das folgende ja noch einfacher ...

func zl() uint {
//
  return 4 * v.Codelen() + col.Codelen()
}


func (E *Imp) Push (c col.Colour) {
//
  B:= make ([]byte, zl())
  a:= 0
  copy (B[a:a+8], E.origin.Encode ())
  a += 8
  for d:= D0; d < NDirs; d++ {
    copy (B[a:a+8], E.vec[d].Encode ())
    a += 8
  }
  copy (B[a:a+97], col.Encode (c))
  stack = append (stack, B)
}


func (E *Imp) Colour () col.Colour {
//
  B:= stack[len(stack) - 1]
  a:= 0
  E.origin.Decode (B[a:a+8])
  a += 8
  for d:= D0; d < NDirs; d++ {
    E.vec[d].Decode (B[a:a+8])
    a += 8
  }
  var c col.Colour
  col.Decode (&c, B[a:a+3])
  stack = stack[0:len(stack) - 2]
  return c
}
