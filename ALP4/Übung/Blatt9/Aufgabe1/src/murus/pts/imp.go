package pts

// (c) Christian Maurer   v. 130526 - license see murus.go

import (
  "murus/ker"; "murus/env"
//  . "murus/obj"
  "murus/gl"
  "murus/col"; "murus/errh"
  "murus/vect"
  "murus/sel"
  "murus/pseq"
  "murus/pt"
)
const (
  pack = "pts"
  suffix = "mug"
  null = 0.0
  eins = 1.0
)
/*
type
  Points struct {
           file *pseq.Imp
     eye, focus,
         vv, nn []*vect.Imp
                }
*/
var (
  file *pseq.Imp
  eye, focus []*vect.Imp = make ([]*vect.Imp, 1), make ([]*vect.Imp, 1)
  vv, nn []*vect.Imp
  started bool
)

/*
func New () *Imp {
//
  x:= new (Imp)
  x.file = pseq.New (pt.New ())
  x.eye, x.focus:= make ([]*vect.Imp, 1), make ([]*vect.Imp, 1)
  x.eye[0], x.focus[0] = vect.New (), vect.New ()
  return x
}
*/

func Clr () {
//
  file.Clr ()
}


func Empty () bool {
//
  return file.Empty ()
}


func Name (s string) {
//
  file.Name (s + "." + suffix)
  eye[0].Set3 (null, -eins, null)
  focus[0].Clr ()
  n:= file.Num ()
  if n > 0 {
    vv, nn = make ([]*vect.Imp, n), make ([]*vect.Imp, n)
    for i:= uint(0); i < n; i++ {
      vv[i], nn[i] = vect.New (), vect.New ()
    }
    file.Seek (n - 1)
    p:= pt.New ()
    p = file.Get ().(*pt.Imp)
    if p.ClassOf () == pt.Start {
      eye[0], focus[0] = p.Read2 ()
    } else {
      eye[0].Set3 (null, null, eins)
      focus[0].Clr ()
    }
  }
}


func Rename (s string) {
//
  file.Rename (s + "." + suffix)
}


func DefCall () {
//
  Name (env.Par (1))
}


func Select () {
//
  name, filename:= sel.Names ("Graphik:", suffix, 64, 0, 0, col.ScreenF, col.ScreenB)
  if name == "" {
    errh.Error ("nicht vorhanden", 0)
    Clr ()
  } else {
    Name (filename)
  }
}


func Ins1 (c pt.Class, a uint, v []*vect.Imp, f col.Colour) {
//
//  if started { ker.Stop (pack, 1) }
  if c > pt.Polygon { ker.Stop (pack, 2) }
//  if uint(len (v)) != a { println ("pts.Ins1: len(v) = ", len(v), " != a = ", a) } // ker.Stop (pack, 98) }
  p:= pt.New ()
  n:= vect.New ()
  n.Set3 (null, null, eins)
  for i:= uint(0); i < a; i++ {
    p.Set (c, a - 1 - i, f, v[i], n)
    file.Ins (p)
  }
}


func Ins (c pt.Class, a uint, v, n []*vect.Imp, f col.Colour) {
//
//  if started { ker.Stop (pack, 3) }
  p:= pt.New ()
  if c == pt.Light {
    p.Set (c, a, f, v[0], n[0])
    file.Ins (p)
    return
  }
  if len (v) != len (n) { ker.Stop (pack, 98) }
  if uint(len (v)) != a { println ("pts.Ins: len(v) = ", len(v), " != a = ", a) } // ker.Stop (pack, 98) }
  for i:= uint(0); i < a; i++ {
    p.Set (c, a - 1 - i, f, v[i], n[i])
    file.Ins (p)
  }
}


func Start (x, y, z, x1, y1, z1 float64) {
//
  if x == x1 && y == y1 && z == z1 { ker.Stop (pack, 4) }
  eye[0].Set3 (x, y, z)
  focus[0].Set3 (x1, y1, z1)
  Ins (pt.Start, 1, eye, focus, col.Red)
  started = true
}


func StartCoord () (float64, float64, float64, float64, float64, float64) {
//
  x, y, z:= eye[0].Coord3 ()
  x1, y1, z1:= focus[0].Coord3 ()
  gl.Init (500.0 * eye[0].Distance (focus[0]))
  return x, y, z, x1, y1, z1
}


func Write () {
//
// TODO: pt der Class Start zuerst 
  p1:= pt.New()
  fn:= file.Num ()
  vv, nn:= make ([]*vect.Imp, fn), make ([]*vect.Imp, fn)
  for i:= uint(0); i < fn; i++ {
    vv[i], nn[i] = vect.New (), vect.New ()
  }
//  pts:= make ([]*pt.Imp, fn)
  i:= uint (0)
//  file.Traverse (func (a Any) { pts[i] = pt.New(); pts[i].Copy (a.(*pt.Imp)) })
  file.Seek (0)
  gl.Write0 ()
// println ("pts.Write: gl.Write0 aufgerufen")
// println ("vor for: file.Pos == ", file.Pos (), "/ fn == ", fn)
  for file.Pos () + 1 < fn {
// println ("file.Pos == ", file.Pos ())
    i = uint(0)
    var a uint
    for {
      p1 = file.Get ().(*pt.Imp)
      k:= p1.Number ()
      if i == 0 {
        if p1.ClassOf () == pt.Light {
          a = k
          k = 0
        } else {
          a = k + 1 // !
        }
      }
      vv[i], nn[i] = p1.Read2 ()
      i ++
// println ("pts.Write: i == ", i)
      file.Step (true)
      if k == 0 { break }
    }
    var f gl.Figure
    switch p1.ClassOf () { case pt.None:
      f = gl.UNDEF
    case pt.Start:
      return
    case pt.Light:
      f = gl.LIGHT
    case pt.Points:
      f = gl.POINTS
    case pt.Lines:
      f = gl.LINES
    case pt.LineStrip:
      f = gl.LINE_STRIP
    case pt.LineLoop:
      f = gl.LINE_LOOP
    case pt.Triangles:
      f = gl.TRIANGLES
    case pt.TriangleStrip:
      f = gl.TRIANGLE_STRIP
    case pt.TriangleFan:
      f = gl.TRIANGLE_FAN
    case pt.Quads:
println ("pts.Write found Quads")
      f = gl.QUADS
    case pt.QuadStrip:
      f = gl.QUAD_STRIP
    case pt.Polygon:
      f = gl.POLYGON
    }
// println ("pts.Write: nach switch: f == ", f)
    gl.Write (f, a, vv, nn, p1.Colour())
  }
  gl.Write1 ()
  println ("pts.Write: gl.Write aufgerufen")
//  println ("written")
}


func Terminate () {
//
  file.Terminate ()
}


func init () {
//
  file = pseq.New (pt.New ())
  eye[0], focus[0] = vect.New (), vect.New ()
  gl.Cls (col.LightWhite)
}
