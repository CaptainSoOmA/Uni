package edge

// (c) Christian Maurer   v. 130526 - license see murus.go

import (
  "math"
  . "murus/obj" // ; "murus/kbd"
  "murus/col"; "murus/scr"; "murus/box"; "murus/nat"
  "murus/gra1/node"
)
const (
  pack = "edge"
  zk = 2
  r0 = 4
)
type
  Imp struct {
         val uint
        x, y,
      x1, y1 uint
             }
var (
  Farbe = [2]col.Colour { col.ScreenF, col.SignalRed }
  bx *box.Imp = box.New()
)


func New () *Imp {
//
  return &Imp { val: uint(1) }
}


func (x *Imp) Def (n uint) {
//
  x.val = n
}


func (x *Imp) Empty () bool {
//
  return x.val == 1
}


func (x *Imp) Clr () {
//
  x.val = uint(1)
}


func (x *Imp) Eq (Y Object) bool {
//
  y, ok:= Y.(*Imp)
  if ! ok { NotCompatiblePanic() }
  return x.val == y.val &&
         x.x == y.x && x.y == y.y && x.x1 == y.x1 && x.y1 == y.y1
}


func (x *Imp) Less (Y Object) bool {
//
  return false
}


func (x *Imp) Copy (Y Object) {
//
  y, ok:= Y.(*Imp)
  if ! ok { NotCompatiblePanic() }
  x.val = y.val
  x.x, x.y = y.x, y.y
  x.x1, x.y1 = y.x1, y.y1
}


func (x *Imp) Clone () Object {
//
  y:= New()
  y.Copy (x)
  return y
}


func (x *Imp) Val () uint {
//
  return x.val
}

/*
func (x *Imp) Ok (v uint) bool {
//
  x.val = v
  return WHAT ?
}
*/

func (x *Imp) SetColours (f, a col.Colour) {
//
  Farbe [0], Farbe [1] = f, a
}

/*
func (x *Imp) UnderMouse (e Edge, n, n1 Node) bool {
//
  x, y, _, _, x0, y0, _, _, _:= e.Pos (e, n, n1)
  return scr.UnderMouseGr (x0 - 5, y0 - 5, x0 + 5, y0 + 5, 10)
}
*/


func (x *Imp) pos (n, n1 *node.Imp, d bool) (uint, uint, uint, uint, bool) {
//
  x.x, x.y = n.Pos ()
  x.x1, x.y1 = n1.Pos ()
  if x.x == x.x1 && x.y == x.y1 {
    return 0, 0, 0, 0, false
  }
  dx:= math.Abs (float64 (x.x) - float64 (x.x1))
  dy:= math.Abs (float64 (x.y) - float64 (x.y1))
  if dx * dx + dy * dy < 0.001 {
    return 0, 0, 0, 0, false
  }
  d0:= 1.0 / math.Sqrt (dx * dx + dy * dy)
  dx, dy = d0 * dx, d0 * dy
  r:= n.Radius ()
  h:= uint (dx * float64 (r + r0 + 1) + 0.5) + 1
  if x.x < x.x1 {
    x.x += h
    x.x1 -= h
  } else {
    x.x -= h
    x.x1 += h
  }
  h = uint (dy * float64 (r + r0 + 1) + 0.5) + 1
  if x.y < x.y1 {
    x.y += h
    x.y1 -= h
  } else {
    x.y -= h
    x.y1 += h
  }
  x0:= (x.x + x.x1) / 2 - (zk * scr.NX1()) / 2 + 1
  y0:= (x.y + x.y1) / 2 - scr.NY1() / 2 + 1
  var x1, y1 uint
//  x1, y1 = (x.x + 7 * x.x1) / 8, (x.y + 7 * x.y1) / 8
  if d {
    x1, y1 = x.x1, x.y1
  } else {
    x1, y1 = x.x, x.y
  }
  return x0, y0, x1, y1, true
}


func (x *Imp) aus (n, n1 *node.Imp, directed bool) {
//
  x0, y0, x1, y1, ok:= x.pos (n, n1, directed)
  if ! ok { return }
  scr.InfLine (int(x.x), int(x.y), int(x.x1), int(x.y1))
  if directed {
    scr.CircleFull (int(x1), int(y1), r0)
  }
  if zk > 0 && WithValues {
    T:= nat.StringFmt (x.val, zk, false)
    bx.WriteGr (T, int(x0), int(y0))
  }
}


func (x *Imp) Write (N, N1 node.Node, directed, vis, inv bool) {
//
  n, n1:= N.(*node.Imp), N1.(*node.Imp)
  if inv {
    _, _, x1, y1, ok:= x.pos (n, n1, directed)
    if !ok { return }
    scr.InfLineInv (int(x.x), int(x.y), int(x.x1), int(x.y1))
    if directed {
      scr.CircleInv (int(x1), int(y1), r0)
    }
    return
  }
  b:= col.ScreenB
  if vis {
    b = Farbe [0]
  }
  bx.ColourF (b)
  if vis {
    scr.Colour (col.ScreenB)
    x.aus (n, n, directed)
  }
  scr.Colour (b)
  x.aus (n, n1, directed)
}


func (x *Imp) WriteCond (N, N1 node.Node, directed, aktuell bool) {
//
  n, n1:= N.(*node.Imp), N1.(*node.Imp)
  f:= 0; if aktuell { f = 1 }
  bx.ColourF (Farbe[f])
//  scr.SetLinewidth (scr.Thicker)
  if ! aktuell {
    scr.Colour (col.ScreenB)
    x.aus (n, n1, directed)
//    scr.SetLinewidth (scr.Thin)
  }
  scr.Colour (Farbe [f])
  x.aus (n, n1, directed)
  if aktuell {
//    scr.SetLinewidth (scr.Thin)
  }
}


func (x *Imp) Edit (N, N1 node.Node, directed bool) {
//
  n, n1:= N.(*node.Imp), N1.(*node.Imp)
  x.Write (n, n1, directed, true, false)
  x0, y0, _, _, _:= x.pos (n, n1, directed)
  if zk == 0 || ! WithValues {
    /* accept only Del
       _, c, _:= kbd.Read () // ?
       if c == kbd.Del { x.val = 0 }
    */
  } else {
    T:= nat.StringFmt (x.val, zk, false)
    bx.ColourF (Farbe [0])
    for {
      bx.EditGr (&T, x0, y0)
      if nat.Defined (&x.val, T) { break }
    }
  }
}


const
  cluint = uint(4)


func (x *Imp) Codelen () uint {
//
  return cluint
}


func (x *Imp) Encode () []byte {
//
  return Encode (x.val)
}


func (x *Imp) Decode (b []byte) {
//
  x.val = Decode (uint(0), b).(uint)
}


func init () {
//
//  bx.SetNumerical ()
  bx.Wd (zk)
//  var e Object = New(); if e == nil {}
//  var e1 Valuator = New(); if e1 == nil {}
}
