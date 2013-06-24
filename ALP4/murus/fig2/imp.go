package fig2

// (c) Christian Maurer   v. 130302 - license see murus.go

import (
  . "murus/obj"; "murus/str"; "murus/kbd"; "murus/font"
  "murus/col"; "murus/scr"; "murus/box"; "murus/errh"; "murus/sel"
  "murus/img"; "murus/psp"
)
const (
  lenText = 40 // maximal len of text
  BB = 10 // length of names
)
type (
  Imp struct {
        sort Sort
      colour col.Colour
        x, y []int
      marked,
//      bewegt,
      filled bool
          tx string
             }
)
var (
  xx, yy int
  name []string
  bx *box.Imp = box.New ()
)


func New () *Imp {
//
  xx, yy = int(scr.NX()), int(scr.NY())
  f:= new (Imp)
  f.Clr ()
//  f.sort = Pointset
  f.sort = Segments
  f.colour = col.ScreenF
  return f
}


func (f *Imp) Empty () bool {
//
  return len (f.x) == 0
}


func (f *Imp) Clr () {
//
  f.x, f.y = nil, nil
  f.marked, f.filled = false, false
  f.tx = ""
}


func (f *Imp) Def (s Sort) {
//
  f.Clr ()
  f.sort = s
}


func (f *Imp) Select () {
//
  f.Clr ()
  Acolour:= f.colour
  Hcolour:= Acolour
  col.Contrast (&Hcolour)
  scr.SwitchFontsize (font.Normal)
  n:= uint(Rectangle)
  Z, S:= scr.MousePos()
  sel.Select1 (name, NSorts, BB, &n, Z, S, Acolour, Hcolour)
  if n < NSorts {
    f.sort = Sort(n)
  }
}


func (f *Imp) Eq (X Object) bool {
//
  f1, ok:= X.(*Imp)
  if ! ok { return false }
  n, n1:= uint(len (f.x)), uint(len (f1.x))
  if f.sort != f1.sort || n != n1 || f.filled != f1.filled {
    return false
  }
  if n == 0 { return true } // ?
  if f.x[0] != f1.x[0] || f.y[0] != f1.y[0] {
    return false
  }
  switch f.sort { case Text:
    if f.tx != f1.tx {
      return false
    }
  case Image:
    if f.x[1] != f1.x[1] || f.y[1] != f1.y[1] {
      return false
    } else {
      // Vergleich der Images fehlt
      return false
    }
  default:
    for i:= uint(1); i < n; i++ {
      if f.x[i] != f1.x[i] || f.y[i] != f1.y[i] {
        return false
      }
    }
  }
  return true
}


func (x *Imp) Less (Y Object) bool {
//
  return false
}


func (f *Imp) Copy (X Object) {
//
  f1, ok:= X.(*Imp)
  if ! ok { return }
  f.sort = f1.sort
  f.colour = f1.colour
  n1:= uint(len (f1.x))
  f.x, f.y = make ([]int, n1), make ([]int, n1)
  for i:= uint(0); i < n1; i++ {
    f.x[i] = f1.x[i]
    f.y[i] = f1.y[i]
  }
  f.filled = f1.filled
  f.tx = f1.tx
  if f.sort == Image {
    // Kopieren des Image fehlt
  }
}


func (f *Imp) Clone () Object {
//
  f1:= New ()
  f1.Copy (f)
  return f1
}


func (f *Imp) Pos () (int, int) {
//
  return f.x[0], f.y[0]
}


func (f *Imp) On (a, b int, t uint) bool {
//
  if ! f.Empty () {
    switch f.sort {
    case Pointset, Segments:
      return scr.OnSegments (f.x, f.y, a, b, t)
    case Polygon:
      return scr.OnPolygon (f.x, f.y, a, b, t)
    case Curve:
      return scr.OnCurve (f.x, f.y, a, b, t)
    case InfLine:
      return scr.OnInfLine (f.x[0], f.y[0], f.x[1], f.y[1], a, b, t)
    case Rectangle:
      return scr.OnRectangle (f.x[0], f.y[0], f.x[1], f.y[1], a, b, t)
    case Circle:
      return scr.OnCircle (f.x[0], f.y[0], uint(f.x[1]), a, b, t)
    case Ellipse:
      return scr.OnEllipse (f.x[0], f.y[0], uint(f.x[1]), uint(f.y[1]), a, b, t)
    case Text:
      if len (f.x) != 2 { errh.Error ("Incident case Text: len (f.x) ==", uint(len(f.x))) }
      return scr.OnRectangle (f.x[0], f.y[0], f.x[1], f.y[1], a, b, t) // crash: TODO
    case Image:
      return scr.InRectangle (f.x[0], f.y[0], f.x[1], f.y[1], a, b) // , t
    }
  }
  return false
}


func (f *Imp) convex () bool {
//
  n:= uint(len (f.x))
  switch f.sort {
  case Rectangle, Circle, Ellipse, Image:
    return true
  case Polygon:
    switch n { case 0, 1:
      return false
    case 2:
      return true
    }
  default:
    return false
  }
 // polygon with 3 or more nodes
/*
 // TODO
  dxi:= f.x[0] - f.x[n - 1]
  dxk:= f.x[1] - f.x[0]
  dyi:= f.y[0] - f.y[n - 1]
  dyk:= f.y[1] - f.y[0]
  z:= uint(0)
  if dxi * dxk + dyi * dyk < 0 { z = 1 }
  a:= dxi * dyk
  b:= dxk * dyi
  if a == b { // polygon reduced by a node
    return true
    // for n > 3 we are going to roasted in devils oven ...
  }
  gr:= a > b
  var k uint
  for i:= uint(1); i < n; i++ {
    if i < n { k = i + 1 } else { k = 0 }
    dxi = f.x[i] - f.x[i - 1]
    dyi = f.y[i] - f.y[i - 1]
    dxk = f.x[k] - f.x[i]
    dyk = f.y[k] - f.y[i]
    if dxi * dxk + dyi * dyk < 0 { // Winkel < 90 Grad
      z++
      if z > 3 {  // if more than 3 angles are < 90°, then
        return false // the angle sum is < (n - 1) * 180° !
      }
    }
    a = dxi * dyk
    b = dxk * dyi
    if a != b {
      if (a > b) != gr { return false }
    }
  }
*/
  return true
}


func (f *Imp) rectangular () bool {
//
  switch f.sort { case Rectangle, Image:
    return true
  }
  if f.sort != Polygon { return false }
  if len (f.x) != 4 { return false }
  return f.x[1] + f.x[3] == f.x[0] + f.x[2] && f.y[1] + f.y[3] == f.y[0] + f.y[2] &&
         f.x[1] * f.x[1] + f.x[0] * f.x[2] + f.y[1] * f.y[1] + f.y[0] * f.y[2] ==
           f.x[1] * (f.x[0] + f.x[2]) + f.y[1] * (f.y[0] * f.y[2])
}


func (f *Imp) UnderMouse (t uint) bool {
//
  a, b:= scr.MousePosGr ()
  return f.On (a, b, t)
}


// Locate (a, b) = Relocate (a - x[0], b - y[0])
func (f *Imp) Move (a, b int) {
//
  var n uint
  switch f.sort { case Pointset, Segments, Polygon, Curve, InfLine, Rectangle:
    n = uint(len (f.x))
  case Circle, Ellipse:
    n = 1
  case Text, Image:
    n = 2
  }
  for i:= uint(0); i < n; i++ {
    f.x[i] += a
    f.y[i] += b
  }
}


func (f *Imp) Marked () bool {
//
  return f.marked
}


func (f *Imp) Mark (m bool) {
//
  f.marked = m
}


func (f *Imp) SetColour (c col.Colour) {
//
  f.colour = c
  bx.ColourF (f.colour)
  if f.sort == Image {
    // what ?
  }
}


func (f *Imp) Colour () col.Colour {
//
  return f.colour
}


func (f *Imp) Erase () {
//
  switch f.sort { case Image:
    scr.ClearGr (uint(f.x[0]), uint(f.y[0]), uint(f.x[1]), uint(f.y[1]))
  default:
    c:= f.colour
    f.SetColour (col.ScreenB)
    f.Write ()
    f.SetColour (c)
  }
}


func (f *Imp) Write () {
//
  if f.Empty () { return }
  scr.Colour (f.colour)
  switch f.sort {
  case Pointset:
    scr.Pointset (f.x, f.y)
  case Segments:
    scr.Segments (f.x, f.y)
  case Polygon:
    scr.Polygon (f.x, f.y)
    if f.filled {
//      scr.PolygonFull (f.x, f.y) // not yet implemented
    }
  case Curve:
    scr.Curve (f.x, f.y)
    if f.filled {
      n:= len (f.x) - 1
      scr.CircleFull (f.x[n], f.y[n], 4) // ?
    }
  case InfLine:
    scr.InfLine (f.x[0], f.y[0], f.x[1], f.y[1])
  case Rectangle:
    if f.filled {
      scr.RectangleFull (f.x[0], f.y[0], f.x[1], f.y[1])
    } else {
      scr.Rectangle (f.x[0], f.y[0], f.x[1], f.y[1])
    }
  case Circle:
    if f.filled {
      scr.CircleFull (f.x[0], f.y[0], uint(f.x[1]))
    } else {
      scr.Circle (f.x[0], f.y[0], uint(f.x[1]))
    }
  case Ellipse:
    if f.filled {
      scr.EllipseFull (f.x[0], f.y[0], uint(f.x[1]), uint(f.y[1]))
    } else {
      scr.Ellipse (f.x[0], f.y[0], uint(f.x[1]), uint(f.y[1]))
    }
  case Text:
    bx.Wd (str.ProperLen (f.tx))
    bx.ColourF (f.colour)
    bx.WriteGr (f.tx, f.x[0], f.y[0])
  case Image:
//    if bewegt {
//      scr.RectangleFullInv (...)
//    } else {
//      copy from Imageptr in Framebuffer
//    }
    img.Get (f.tx, uint(f.x[0]), uint(f.y[0]))
  }
}


func (f *Imp) Print (p *psp.Imp) {
//
  if f.Empty () { return }
  n:= uint(len (f.x))
  p.SetColour (f.colour)
  switch f.sort {
  case Pointset:
    x, y:= make ([]float64, n), make ([]float64, n)
    for i:= uint(0); i < n; i++ {
      x[i], y[i] = p.S(f.x[i]), p.Sy(f.y[i])
    }
    p.Points (x, y)
  case Segments:
    x, y:= make ([]float64, n), make ([]float64, n)
    for i:= uint(0); i < n; i++ {
      x[i], y[i] = p.S(f.x[i]), p.Sy(f.y[i])
    }
    p.Segments (x, y)
  case Polygon:
    x, y:= make ([]float64, n), make ([]float64, n)
    for i:= uint(0); i < n; i++ {
      x[i], y[i] = p.S(f.x[i]), p.Sy(f.y[i])
    }
    p.Polygon (x, y, f.filled)
  case Curve:
    x, y:= make ([]float64, n), make ([]float64, n)
    for i:= uint(0); i < n; i++ {
      x[i], y[i] = p.S(f.x[i]), p.Sy(f.y[i])
    }
    p.Curve (x, y)
  case InfLine:
    x, y, x1, y1:= p.S(f.x[0]), p.Sy(f.y[0]), p.S(f.x[1]), p.Sy(f.y[1])
    p.Line (x, y, x1, y1)
  case Rectangle:
    x, y, x1, y1:= p.S(f.x[0]), p.Sy(f.y[0]), p.S(f.x[1]), p.Sy(f.y[1])
    p.Rectangle (x, y, x1 - x, y1 - y, f.filled)
  case Circle:
    x, y, r:= p.S(f.x[0]), p.Sy(f.y[0]), p.S(f.x[1])
    p.Circle (x, y, r, f.filled)
  case Ellipse:
    x, y, a, b:= p.S(f.x[0]), p.Sy(f.y[0]), p.S(f.x[1]), p.S(f.y[1])
    p.Ellipse (x, y, a, b, f.filled)
  case Text:
    x, y:= p.S(f.x[0]), p.Sy(f.y[0])
    p.Write (f.tx, x, y)
  case Image:
//i TODO
  }
}


func (f *Imp) Invert () {
//
  if f.Empty () { return }
  switch f.sort { case Pointset:
    scr.PointsetInv (f.x, f.y)
  case Segments:
    scr.SegmentsInv (f.x, f.y)
  case Polygon:
    if f.filled {
      scr.PolygonInv /* TODO Full */ (f.x, f.y)
    } else {
      scr.PolygonInv (f.x, f.y)
    }
  case Curve:
    scr.CurveInv (f.x, f.y)
    if f.filled {
      n:= len (f.x) - 1
      scr.CircleInv (f.x[n], f.y[n], 4)
    }
  case InfLine:
    scr.InfLineInv (f.x[0], f.y[0], f.x[1], f.y[1])
  case Rectangle:
    if f.filled {
      scr.RectangleFullInv (f.x[0], f.y[0], f.x[1], f.y[1])
    } else {
      scr.RectangleInv (f.x[0], f.y[0], f.x[1], f.y[1])
    }
  case Circle:
    if f.filled {
      scr.CircleFullInv (f.x[0], f.y[0], uint(f.x[1]))
    } else {
      scr.CircleInv (f.x[0], f.y[0], uint(f.x[1]))
    }
  case Ellipse:
    if f.filled {
      scr.EllipseFullInv (f.x[0], f.y[0], uint(f.x[1]), uint(f.y[1]))
    } else {
      scr.EllipseInv (f.x[0], f.y[0], uint(f.x[1]), uint(f.y[1]))
    }
  case Text:
// >>>  sollte in bx integriert werden:
//  bx.WriteInvGr (tx, x[0], y[0])
    scr.SwitchTransparence (true)
    scr.WriteInvGr (f.tx, f.x[0], f.y[0])
  case Image:
    scr.RectangleInv (f.x[0], f.y[0], f.x[1], f.y[1])
  }
}


func (f *Imp) invertN () {
//
  switch f.sort { case Pointset:
    scr.PointsetInv (f.x, f.y)
  case Segments:
    scr.SegmentsInv (f.x, f.y)
  case Polygon:
    scr.PolygonInv (f.x, f.y)
  case Curve:
    scr.CurveInv (f.x, f.y)
    if f.filled {
      n:= len (f.x) - 1
      scr.CircleInv (f.x[n], f.y[n], 4)
    }
  }
}


func (f *Imp) editN () {
//
  switch f.sort { case Pointset, Segments, Polygon, Curve: default: return }
  x0:= make ([]int, 2); x0[0] = f.x[0]; f.x = x0
  y0:= make ([]int, 2); y0[0] = f.y[0]; f.y = y0
  f.x[1], f.y[1] = scr.MousePosGr ()
  f.invertN ()
  var ( K kbd.Comm; T uint )
  loop: for {
    K, T = kbd.Command ()
    scr.MouseCursor (true)
    n:= uint(len (f.x))
    switch K { case kbd.Esc:
      break loop
    case kbd.Go,
         kbd.Here, kbd.Pull, kbd.Hither,
         kbd.There, kbd.Push, kbd.Thither,
         kbd.This: // kbd.ToThis:
      f.invertN ()
//      if f.sort == Curve {
//        if n == scr.MaxBezierdegree { break loop }
//      }
      if f.sort == Pointset {
        if K != kbd.Go {
          n++
        }
      } else {
        if K == kbd.Here { // TODO Curve: missing
          n++
        }
      }
      if K == kbd.This {
        n:= len (f.x)
        if n == 0 {
          break loop
        } else { // TODO
          n--
          if n == 0 {
            break loop
//          } else {
//            x0 = make ([]int, n); copy (x0, f.x[:n]); f.x = x0
//            y0 = make ([]int, n); copy (y0, f.y[:n]); f.y = y0
            }
        }
      }
      if n > uint(len (f.x)) {
        x0 = make ([]int, n); copy (x0, f.x); f.x = x0
        y0 = make ([]int, n); copy (y0, f.y); f.y = y0
      }
      f.x[n-1], f.y[n-1] = scr.MousePosGr ()
      f.invertN ()
      if f.sort == Pointset {
        if K == kbd.Hither { break loop }
      } else {
        if K == kbd.Thither { break loop }
      }
    }
  }
  if f.x == nil {
    f.Clr ()
    return
  }
  scr.Colour (f.colour)
  switch f.sort { case Pointset:
    scr.Pointset (f.x, f.y)
  case Segments:
    scr.Segments (f.x, f.y)
  case Polygon:
    scr.Polygon (f.x, f.y)
    f.filled = T > 0 && f.convex ()
    if f.filled {
//      scr.PolygonFull (f.x, f.y) // not yet implemented
    }
  case Curve:
    scr.Curve (f.x, f.y)
    f.filled = T > 0
    if f.filled {
      n:= len (f.x) - 1
      scr.CircleFull (f.x[n], f.y[n], 4)
    }
  }
}


func (f *Imp) invert1 () {
//
  switch f.sort { case InfLine:
    scr.InfLineInv (f.x[0], f.y[0], f.x[1], f.y[1])
  case Rectangle:
    scr.RectangleInv (f.x[0], f.y[0], f.x[1], f.y[1])
  default:
    scr.EllipseInv (f.x[0], f.y[0], uint(f.x[1]), uint(f.y[1]))
  }
}


func (f *Imp) edit1 () {
//
  x0:= make ([]int, 2); x0[0] = f.x[0]; f.x = x0
  y0:= make ([]int, 2); y0[0] = f.y[0]; f.y = y0
  switch f.sort { case InfLine:
    if f.x[0] == 0 {
      f.x[1] = 1
    } else {
      f.x[1] = f.x[0] - 1
    }
    f.y[1] = f.y[0]
  case Rectangle:
    f.x[1] = f.x[0]
    f.y[1] = f.y[0]
  case Circle, Ellipse:
    f.x[1] = 0
    f.y[1] = 0
  default:
    return
  }
//    scr.PointInv (f.x[0], f.y[0])
  f.invert1 ()
  loop: for {
    K, T:= kbd.Command ()
    switch K { case kbd.Pull, kbd.Hither:
      f.invert1 ()
      f.x[1], f.y[1] = scr.MousePosGr ()
      switch f.sort { case InfLine:
        if f.x[1] == f.x[0] && f.y[1] == f.y[0] {
          if f.x[0] == 0 {
            f.x[1] = 1
          } else {
            f.x[1] = f.x[0] - 1
          }
        }
      case Rectangle:
        ;
      case Circle, Ellipse:
        if f.x[1] > f.x[0] {
          f.x[1] -= f.x[0]
        } else {
          f.x[1] = f.x[0] - f.x[1]
        }
        if f.y[1] > f.y[0] {
          f.y[1] -= f.y[0]
        } else {
          f.y[1] = f.y[0] - f.y[1]
        }
        if f.sort == Circle {
          if f.x[1] > f.y[1] {
            f.y[1] = f.x[1]
          } else {
            f.x[1] = f.y[1]
          }
        }
      default:
        // stop (Modul, 1)
      }
      f.invert1 ()
      if K == kbd.Hither {
        f.filled = T > 0
        break loop
      }
    }
  }
  switch f.sort { case InfLine:
    scr.InfLine (f.x[0], f.y[0], f.x[1], f.y[1])
  case Rectangle:
    if f.filled {
      scr.RectangleFull (f.x[0], f.y[0], f.x[1], f.y[1])
    } else {
      scr.Rectangle (f.x[0], f.y[0], f.x[1], f.y[1])
    }
  case Circle, Ellipse:
    if f.filled {
      scr.EllipseFull (f.x[0], f.y[0], uint(f.x[1]), uint(f.y[1]))
    } else {
      scr.Ellipse (f.x[0], f.y[0], uint(f.x[1]), uint(f.y[1]))
    }
  }
}


func (f *Imp) editText () {
//
  if f.sort != Text { return }
  scr.MouseCursor (false)
  bx.Wd (lenText)
  bx.ColourF (f.colour)
  x1:= f.x[0] + int(lenText * scr.NX1()) - 1
  if x1 >= xx { x1 = xx - 1 }
  y1:= f.y[0] + int(scr.NY1()) - 1
  if y1 >= yy { y1 = yy - 1 }
  scr.SaveGr (uint(f.x[0]), uint(f.y[0]), uint(x1), uint(y1))
  bx.SetTransparent (false)
  f.tx = str.Clr (lenText) // wörkeraunt
  bx.EditGr (&f.tx, uint(f.x[0]), uint(f.y[0]))
  bx.SetTransparent (true)
  scr.RestoreGr (uint(f.x[0]), uint(f.y[0]), uint(x1), uint(y1))
  var T uint
  if kbd.LastCommand (&T) == kbd.Enter {
    bx.SetTransparent (true)
//    scr.RestoreGr (f.x[0], f.y[0], x1, y1)
    bx.WriteGr (f.tx, f.x[0], f.y[0])
    k:= str.ProperLen (f.tx)
    x0:= make ([]int, 2); x0[0] = f.x[0]; f.x = x0
    y0:= make ([]int, 2); y0[0] = f.y[0]; f.y = y0
    f.x[1] = f.x[0] + int(scr.NX1() * k) - 1
    f.y[1] = f.y[0] + int(scr.NY1()) - 1
    scr.WarpMouseGr (f.x[0], f.y[1])
  } else {
//    f.tx = str.Clr (lenText)
//    bx.WriteGr (f.tx, f.x[0], f.y[0])
//    f.tx = ""
//    f.x, f.y = nil, nil
  }
  scr.MouseCursor (true)
}


func (f *Imp) editImage () {
//
  if f.sort != Image { return }
  scr.MouseCursor (false)
  errh.Hint ("Name des Bildes eingeben")
  bx.Wd (32) // reine Willkür
  Hf:= col.ScreenB
  bx.Colours (f.colour, Hf)
  f.tx = str.Clr (BB)
  bx.EditGr (&f.tx, uint(f.x[0]), uint(f.y[0]))
  str.RemSpaces (&f.tx)
  W, H:= img.Size (f.tx)
  w, h:= int(W), int(H)
  if w <= xx && h <= yy {
    x0:= make ([]int, 2); x0[0] = f.x[0]; f.x = x0
    y0:= make ([]int, 2); y0[0] = f.y[0]; f.y = y0
    f.x[1] = f.x[0] + w - 1
    f.y[1] = f.y[0] + h - 1
    if f.x[1] >= xx {
      f.x[0] = xx - w
      f.x[1] = xx - 1
    }
    if f.y[1] >= yy {
      f.y[0] = yy - h
      f.y[1] = yy - 1
    }
    errh.DelHint()
//  besser:
//    img.Get ...
//    NEW (Imagespeicher)
//    img.Get ( ... dort rein ...)
//    img.Get (tx, x[0], y[0])
  } else {
    errh.DelHint()
  }
  scr.MouseCursor (true)
}


func (f *Imp) uM () uint {
//
  const ( r = 4; t = 4 )
  a, b:= scr.MousePosGr ()
  n:= uint(len (f.x))
  for i:= uint(0); i < n; i++ {
    if scr.OnCircle (f.x[i], f.y[i], r, a, b, t) {
      return uint(i)
    }
  }
  return n + 1 // ?
}


func (f *Imp) mark (i uint) {
//
//  if f.sort != Curve { return }
  for r:= uint(3); r <= 4; r++ {
    scr.CircleInv (f.x[i], f.y[i], r)
  }
}


func (f *Imp) Edit () {
//
  if f.Empty () {
    scr.Colour (f.colour)
    f.x, f.y = make ([]int, 1), make ([]int, 1)
    f.x[0], f.y[0] = scr.MousePosGr ()
    switch f.sort { case Pointset, Segments, Polygon, Curve:
      f.editN ()
    case InfLine, Rectangle, Circle, Ellipse:
      f.edit1 ()
    case Text:
      f.editText ()
    case Image:
//      ALLOCATE (Imageptr, Groesse())
//      img.Get (tx [...], Imageptr)
      f.editImage ()
    }
    if f.x == nil {
      f.Clr ()
    }
  } else {
    n:= uint(len (f.x))
errh.Error ("Figur hat Länge", n)
    switch f.sort { case Text:
      f.editText ()
    case Image:
      f.editImage ()
    default:
      f.Erase ()
      f.Invert ()
      if true { // f.sort == Curve {
        for i:= uint(0); i < n; i++ { f.mark (i) }
      }
      i:= f.uM ()
      f.x[i], f.y[i] = scr.MousePosGr ()
      loop: for {
        scr.MouseCursor (true)
        c, _:= kbd.Command ()
        switch c { case kbd.Esc:
          break loop
        case kbd.Enter, kbd.Tab, kbd.LookFor:
          f.colour = sel.Colour ()
        case kbd.Here:
          break loop
        case kbd.There:
          i = f.uM ()
        case kbd.Push, kbd.Thither:
          if i < n {
            f.Invert ()
            f.mark (i)
            f.x[i], f.y[i] = scr.MousePosGr ()
            f.mark (i)
            f.Invert ()
            if c == kbd.Thither { i = n } // ? ? ?
          }
        case kbd.This:
          switch f.sort { case Pointset, Segments, Polygon, Curve:
            if f.x == nil {
              f.Clr ()
            } else {
              for i:= uint(0); i < n; i++ { f.mark (i) }
              f.Erase ()
              n-- // ? ? ?
              f.Invert ()
              for i:= uint(0); i < n; i++ { f.mark (i) }
            }
          }
        }
        errh.Hint (c.String ())
      }
      f.Invert ()
      if true { // sort != Text {
        for i:= uint(0); i < n; i++ { f.mark (i) }
      }
      f.Write ()
    }
  }
}


var
  clz = Codelen (uint(0))


func (f *Imp) Codelen () uint {
//
  n:= 1 + col.Codelen () + clz
  switch f.sort { case Text:
    n += 2 * clz + 1 + uint(len (f.tx))
  case Image:
    n += 4 * clz + 1 + uint(len (f.tx))
  default:
    n += 2 * uint(len (f.x)) * clz
  }
  n += 2 * clz // Reserve
  return n
}


func (f *Imp) Encode () []byte {
//
  B:= make ([]byte, f.Codelen ())
  a:= uint(0)
  B[a] = byte(f.sort)
  a++
  copy (B[a:a+3], col.Encode (f.colour))
  a += 3
  var n uint
  if f.sort < Text {
    n = uint(len (f.x))
  } else {
    n = uint(len (f.tx))
  }
  copy (B[a:a+clz], Encode (n))
  a += clz
  if f.sort < Text {
    for i:= uint(0); i < n; i++ {
      copy (B[a:a+clz], Encode (f.x[i]))
      a += clz
      copy (B[a:a+clz], Encode (f.y[i]))
      a += clz
    }
  } else { // Text, Image
    copy (B[a:a+clz], Encode (f.x[0]))
    a += clz
    copy (B[a:a+clz], Encode (f.y[0]))
    a += clz
    if f.sort == Image {
      copy (B[a:a+clz], Encode (f.x[1]))
      a += clz
      copy (B[a:a+clz], Encode (f.y[1]))
      a += clz
    }
    copy (B[a:a+n], []byte(f.tx))
    a += n
  }
  B[a] = 0
  if f.filled { B[a]++ }
  if f.marked { B[a] += 2 }
  return B
}


func (f *Imp) Decode (B []byte) {
//
  a:= uint(0)
  f.sort = Sort(B[a])
  a ++
  col.Decode (&f.colour, B[a:a+3])
  a += 3
  n:= uint(0)
  n = Decode (uint(0), B[a:a+clz]).(uint)
  a += clz
  if f.sort < Text {
    f.x, f.y = make ([]int, n), make ([]int, n)
    for i:= uint(0); i < n; i++ {
      f.x[i] = Decode (f.x[i], B[a:a+clz]).(int)
      a += clz
      f.y[i] = Decode (f.y[i], B[a:a+clz]).(int)
      a += clz
    }
  } else { // sort == Text, Image
    f.x, f.y = make ([]int, 2), make ([]int, 2)
    f.x[0] = Decode (f.x[0], B[a:a+clz]).(int)
    a += clz
    f.y[0] = Decode (f.y[0], B[a:a+clz]).(int)
    a += clz
    if f.sort == Image {
      f.x[1] = Decode (f.x[1], B[a:a+clz]).(int)
      a += clz
      f.y[1] = Decode (f.y[1], B[a:a+clz]).(int)
      a += clz
    }
    f.tx = string(B[a:a+n])
    a += n
    if f.sort == Text {
      f.x[1] = f.x[0] + int(scr.NX1() * n) - 1
      f.y[1] = f.y[0] + int(scr.NY1()) - 1
    }
  }
  f.filled = B[a] % 2 == 1
  f.marked = (B[a] / 2) % 2 == 1
}


func init () {
//
  name = make ([]string, NSorts)
  name[Pointset]    = "Punktfolge"
  name[Segments]    = "Strecke(n)"
  name[Polygon]     = "Polygon   "
  name[Curve]       = "Kurve     "
  name[InfLine]     = "Gerade    "
  name[Rectangle]   = "Rechteck  "
  name[Circle]      = "Kreis     "
  name[Ellipse]     = "Ellipse   "
  name[Text]        = "Text      "
  name[Image]       = "Bild      "
  bx.SetTransparent (true)
  bx.Wd (lenText)
//  var _ Figure2 = New() {}
}
