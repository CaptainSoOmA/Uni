package bahnhof

// (c) Christian Maurer   v. 130308 - license see murus.go

import (
  . "murus/obj"; "murus/str"
  "murus/col"; "murus/scr"
  "murus/scale"
  "murus/sub/linie"
)
const (
  dB =  67.62 // km pro Breitengrad bei 52.5° Länge
  dL = 111.13 // km pro Längengrad
)
type
  Imp struct {
       linie linie.Linie
      nummer uint
     umstieg bool
 name, name1 string
beschriftung byte
      breite,        // x
       länge float64 // y
             }
const
  max = 32 // maximale Länge der Namen der Bahnhöfe


func (x *Imp) imp (Y Object) *Imp {
//
  y, ok:= Y.(*Imp)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}


func New () *Imp {
//
  x:= new (Imp)
  x.beschriftung = '.'
  x.breite, x.länge = 52.5, 13.5
  return x
}


func (x *Imp) Clr () {
//
  x.linie, x.nummer = linie.Fußweg, 0
  x.umstieg = false
  x.name, x.name1 = "", ""
  x.beschriftung = '.'
  x.breite, x.länge = 52.5, 13.5
}


func (x *Imp) Empty () bool {
//
  return x.linie == linie.Fußweg &&
         x.nummer == 0
}


func (x *Imp) Def (l linie.Linie, nr uint, n, n1 string, b byte, yy, xx float64) {
//
  x.linie, x.nummer = l, nr
  x.name, x.name1 = n, n1
  switch b { case 'o', 'l', 'u', 'r':
    x.beschriftung = b
  default:
    x.beschriftung = '.'
  }
  x.breite, x.länge = xx * dB, yy * dL
}


func (x *Imp) Linie () linie.Linie {
//
  return x.linie
}


func (x *Imp) Nummer () uint {
//
  return x.nummer
}


func (x *Imp) Umstieg () {
//
//  x.umstieg = true
}


func (x *Imp) Equiv (Y Object) bool {
//
  y, ok:= Y.(*Imp)
  if ! ok { return false }
  return x.nummer != y.nummer &&
         x.breite == y.breite &&
         x.länge == y.länge
}


func (x *Imp) Rescale (xx, yy uint) {
//
  x.breite, x.länge = scale.Rescale (int(xx), int(yy))
}


func (x *Imp) Numerieren (l linie.Linie, nr uint) {
//
  x.linie, x.nummer = l, nr
  x.umstieg = false
}


func (x *Imp) SkalaEditieren () {
//
  scale.Edit ()
}


func (x *Imp) UnterMaus () bool {
//
  xm, ym:= scr.MousePosGr ()
  xx, yy:= scale.Scale (x.breite, x.länge)
  dx, dy:= xm - xx, ym - yy
  if dx < 0 { dx = -dx }
  if dy < 0 { dy = -dy }
  const d = 8 // pixel
  return dx <= d && dy <= d
}


func (x *Imp) Eq (Y Object) bool {
//
  y:= x.imp (Y)
  return x.linie == y.linie &&
         x.nummer == y.nummer
}


func (x *Imp) Less (Y Object) bool {
//
  return false
}


func (x *Imp) Copy (Y Object) {
//
  y:= x.imp (Y)
  x.linie, x.nummer = y.linie, y.nummer
  x.umstieg = y.umstieg
  x.name, x.name1 = y.name, y.name1
  x.beschriftung = y.beschriftung
  x.länge, x.breite = y.länge, y.breite
}


func (x *Imp) Clone () Object {
//
  y:= New ()
  y.Copy (x)
  return y
}


func (x *Imp) Write1 (Y Object) {
//
  y:= x.imp (Y)
  x1, y1:= scale.Scale (x.breite, x.länge)
  x2, y2:= scale.Scale (y.breite, y.länge)
  scr.Line (x1, y1, x2, y2)
}


func (x *Imp) Write (aktuell bool) {
//
  xx, yy:= scale.Scale (x.breite, x.länge)
  lw:= scr.ActLinewidth ()
  scr.SetLinewidth (scr.Thin)
  if aktuell {
    scr.Colours (linie.Farbe [x.linie], col.ScreenB)
  } else { // umstieg
    scr.Colours (col.Black, col.ScreenB)
  }
  const r = 2
  if xx >= r && yy >= r {
    scr.Circle (xx, yy, r)
//    scr.Colour (col.ScreenB)
    scr.CircleFull (xx, yy, r - 1)
  }
  scr.SetLinewidth (lw)
  n:= int(str.ProperLen (x.name))
  if n <= 2 { return }
  n1:= int(str.ProperLen (x.name1))
  if n1 > n {
    n = n1
  }
  xn, yn:= 0, 0
  w, h:= int(scr.NX1()), int(scr.NY1())
  switch x.beschriftung { case 'r':
    xn = xx + w + 1
    if n1 == 0 {
      yn = yy - h / 2
    } else {
      yn = yy - h
    }
  case 'o':
    xn = xx - (n * w) / 2 + 1
    if n1 == 0 {
      yn = yy - h - 1
    } else {
      yn = yy - 2 * h - 1
    }
  case 'l':
    xn = xx - n * w - w + 1
    if n1 == 0 {
      yn = yy - h / 2
    } else {
      yn = yy - h
    }
  case 'u':
    xn = xx - (n * w) / 2 + 1
    yn = yy + h / 2 - 2
  default:
    xn = xx - (n * w) / 2 + 1
    if n1 == 0 {
      yn = yy - h / 2
    } else {
      yn = yy - h
    }
  }
  xxx:= x.name
  for i:= uint(0); i < uint(len (xxx)); i++ {
    if xxx[i] == '_' { str.Replace (&xxx, i, ' ') }
  }
  xxx1:= x.name1
  for i:= uint(0); i < uint(len (xxx1)); i++ {
    if xxx1[i] == '_' { str.Replace (&xxx1, i, ' ') }
  }
  if aktuell {
    scr.Colours (linie.Farbe [x.linie], col.ScreenB)
  } else { // umstieg
    scr.Colours (col.Black, col.ScreenB)
    scr.Colours (col.Black, col.Pink)
  }
  scr.WriteGr (xxx, xn, yn)
  scr.WriteGr (xxx1, xn, yn + h + 0)
}


func (x *Imp) Codelen () uint {
//
  return 4 +
         1 +
         2 * max +
         1 +
         2 * Codelen (0.0)
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, x.Codelen())
  a, i:= uint(4), uint(0)
  n:= 100 * uint(x.linie) + x.nummer
  copy (b[:a], Encode (n))
  i += a
  b[i] = 0; if x.umstieg { b[i] = 1 }
  i ++
  a = max
  copy (b[i:i+a], x.name)
  i += a
  copy (b[i:i+a], x.name1)
  i += a
  b[i] = x.beschriftung
  i ++
  a = Codelen (0.0)
  copy (b[i:i+a], Encode (x.breite))
  i += a
  copy (b[i:i+a], Encode (x.länge))
  return b
}


func (x *Imp) Decode (b []byte) {
//
  a, i:= uint(4), uint(0)
  n:= Decode (uint(0), b[i:i+a]).(uint)
  x.linie, x.nummer = linie.Linie(n / 100), n % 100
  i += a
  x.umstieg = b[i] == 1
  x.umstieg = false
  i ++
  a = max
  x.name = string (b[i:i+a])
  i += a
  x.name1 = string (b[i:i+a])
  i += a
  x.beschriftung = b[i]
  i ++
  a = Codelen (0.0)
  x.breite = Decode (0.0, b[i:i+a]).(float64)
  i += a
  x.länge = Decode (0.0, b[i:i+a]).(float64)
}


func init () {
//
  bMin, lMin:= 12.8800, 52.2850 // x, y
  bMax, hMax:=  1.0647,  0.4859 // 72, 54 km
  if scr.Proportion() > 4 / 3 {
    bMax = bMax / 4 * 3 * scr.Proportion()
  }
  b0, l0:= 13.2610, 52.4550
  scale.Def (b0 * dB, l0 * dL, 16)
  scale.Lim (bMin * dB, lMin * dL, bMax * dB, hMax * dL, 6)
  scr.SwitchTransparence (true)
  scr.MouseCursor (true)
  var bhf Bahnhof = New(); if bhf == nil {}
}
