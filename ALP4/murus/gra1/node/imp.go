package node

// (c) Christian Maurer   v. 130526 - license see murus.go

import (
  . "murus/obj"; "murus/str"
  "murus/col"; "murus/scr"; "murus/box"
)
const (
  pack = "gra1/node"
  max = 22
  R0 = 3
)
type
  label byte; const (
  zentral = iota
  rechts
  oben
  links
  unten
)
type
  Imp struct {
          name string
           lab label
        length,
          x, y uint
               }
var (
  Farbe = [2]col.Colour { col.ScreenF, col.SignalRed }
  bx *box.Imp = box.New()
)


func New (n uint) *Imp {
//
  if n >= max { n = max }
  x:= new (Imp)
  x.name = ""
  x.lab = zentral
  x.length = n
  return x
}


func (x *Imp) Empty () bool {
//
  return str.Empty (x.name)
}


func (x *Imp) Clr () {
//
  x.name = ""
}


func (x *Imp) Eq (Y Object) bool {
//
  y, ok:= Y.(*Imp)
  if ! ok { NotCompatiblePanic() }
  return x.name == y.name &&
         x.x == y.x && x.y == y.y
}


func (x *Imp) Copy (Y Object) {
//
  y, ok:= Y.(*Imp)
  if ! ok { NotCompatiblePanic() }
  x.name = y.name
  x.length = y.length
  x.lab = y.lab
  x.x, x.y = y.x, y.y
}


func (x *Imp) Less (Y Object) bool {
//
  return false
}


func (x *Imp) Clone () Object {
//
  y:= New (x.length)
  y.Copy (x)
  return y
}


func (x *Imp) Def (s string, c byte, X, Y uint) {
//
  x.name = s
  switch c { case '<', 'l', '-':
    x.lab = links
  case '>', 'r', '+':
    x.lab = rechts
  case '^', 'o':
    x.lab = oben
  case '_', 'u':
    x.lab = unten
  default:
    x.lab = zentral
  }
  x.x, x.y = X, Y
}


func (x *Imp) Locate () {
//
  r:= x.Radius ()
  xm, ym:= scr.MousePosGr ()
  x.x, x.y = uint(xm), uint(ym)
  if x.x < r { x.x = r }
  if x.x + r >= scr.NY () {
    x.x = scr.NX () - 1 - r
  }
  if x.y < r { x.y = r }
  if x.y + r >= scr.NY() {
    x.y = scr.NY() - 1 - r
  }
}


func (x *Imp) String () string {
//
  return x.name
}


func (x *Imp) Pos () (uint, uint) {
//
  return x.x, x.y
}


func (x *Imp) Radius () uint {
//
  if x.length > 2 {
    return R0
  }
  return ((x.length + 1) * scr.NX1 ()) / 2
}


func UnderMouse (a Any) bool {
//
  x:= a.(*Imp)
  return scr.UnderMouseGr (int(x.x), int(x.y), int(x.x), int(x.y), x.Radius ())
}


func (x *Imp) SetColours (f, a col.Colour) {
//
  Farbe [0], Farbe [1] = f, a
}


func (x *Imp) writeEdge (x1 *Imp, u bool) {
//
  n:= 0; if u { n = 1 }
  scr.Colour (Farbe [n])
/*
  if u {
    scr.LinienbreiteSetzen (scr.dicker)
  } else {
    scr.LinienbreiteSetzen (scr.duenn)
  }
*/
  scr.InfLine (int(x.x), int(x.y), int(x1.x), int(x1.y))
}


func (x *Imp) write () {
//
  r:= x.Radius ()
/*
  if r <= R0 {
    scr.CircleFull (int(x.x), int(x.y), R0)
  } else {
    for n:= uint(0); n <= 1; n++ {
      scr.Circle (int(x.x), int(x.y), r + n)
    }
  }
*/
  if x.length > 0 {
    n:= str.ProperLen (x.name)
    xx:= x.x - (n * scr.NX1()) / 2 + 1
    yy:= x.y - scr.NY1() / 2 + 1
    switch x.lab { case zentral:
      ;
    case rechts:
      xx = x.x + scr.NX1() + 1
    case oben:
      yy -= 5 * scr.NY1() / 6
    case links:
      xx = x.x - n * scr.NX1() - scr.NX1() + 1
    case unten:
      yy += 5 * scr.NY1() / 6
    }
    bx.SetTransparent (transparent)
    bx.Wd (x.length)
    bx.WriteGr (x.name, int(xx), int(yy))
  }
  if r <= R0 {
    scr.CircleFull (int(x.x), int(x.y), R0)
  } else {
    for n:= uint(0); n <= 1; n++ {
      scr.Circle (int(x.x), int(x.y), r + n)
    }
  }
}


func (x *Imp) Write (vis, inv bool) {
//
  if inv {
    scr.CircleInv (int(x.x), int(x.y), x.Radius ())
    return
  }
  b:= col.ScreenB
  if vis {
    b = Farbe [0]
  }
  scr.Colour (b)
  bx.ColourF (b)
  x.write ()
}


func (x *Imp) WriteCond (u bool) {
//
  n:= 0; if u { n = 1 }
  scr.Colour (Farbe[n])
  bx.ColourF (Farbe[n])
  x.write ()
}


func (x *Imp) Edit () {
//
  if x.length > 0 {
    B:= (x.length * scr.NX1 ()) / 2
    H:= scr.NY1 () / 2
    bx.Wd (x.length)
    bx.ColourF (Farbe [0])
    bx.EditGr (&x.name, x.x - B + 1, x.y - H + 1)
    N:= str.Clr (x.length)
    bx.WriteGr (N, int(x.x - B + 1), int(x.y - H + 1))
  }
  x.write ()
}


const
  cluint = uint(4)


func (x *Imp) Codelen () uint {
//
  return x.length + 1 + 3 * cluint
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, x.Codelen())
  i, a:= uint(0), cluint
  copy (b[i:i+a], Encode (x.length))
  i += a
  if x.length > 0 {
    a = x.length
    copy (b[i:i+a], x.name[:])
    i += a
  }
  a = 1
  b[i] = byte(x.lab)
  i += a
  a = cluint
  copy (b[i:i+a], Encode (x.x))
  i += a
  copy (b[i:i+a], Encode (x.y))
  return b
}


func (x *Imp) Decode (b []byte) {
//
  i, a:= uint(0), cluint
  x.length = Decode (uint(0), b[i:i+a]).(uint)
  i += a
  if x.length > 0 {
    a = x.length
    x.name = string (b[i:i+a])
    i += a
  }
  a = 1
  x.lab = label (b[i])
  i += a
  a = cluint
  x.x = Decode (uint(0), b[i:i+a]).(uint)
  i += a
  x.y = Decode (uint(0), b[i:i+a]).(uint)
  var node Node = New (0); if node == nil {}
}
