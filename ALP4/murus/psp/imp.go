package psp

// (c) Christian Maurer   v. 130519 - license see murus.go

import (
  "os"; "strconv"; "math"
  "murus/ker"; "murus/z"
  "murus/col"; "murus/scr"; "murus/font"
)
type
  Imp struct {
        file *os.File
          lw float64 // linewidth
             }


func (x *Imp) S (pt int) float64 {
//
  return float64(pt) / float64(scr.NX()) * ker.A4wdPt
}


func (x *Imp) Sy (pt int) float64 {
//
  return ker.A4htPt - x.S (pt)
}


func (x *Imp) write (s string) {
//
  x.file.Write ([]byte(s))
}


func (x *Imp) newpath () {
//
  x.write ("newpath\n")
}


func (x *Imp) closepath () {
//
  x.write ("closepath\n")
}


func (x *Imp) fill () {
//
  x.write ("fill\n")
}


func (x *Imp) stroke () {
//
  x.write ("stroke\n")
}


func (x *Imp) f (r float64) string {
//
  return strconv.FormatFloat (r, 'f', 4, 64)
}


func (x *Imp) moveto (a, b float64) {
//
  x.write (x.f(a) + " " + x.f(b) + " moveto\n")
}


func (x *Imp) lineto (a, b float64) {
//
  x.write (x.f(a) + " " + x.f(b) + " lineto\n")
}


func (x *Imp) rmoveto (a, b float64) {
//
  x.write (x.f(a) + " " + x.f(b) + " rmoveto\n")
}


func (x *Imp) rlineto (a, b float64) {
//
  x.write (x.f(a) + " " + x.f(b) + " rlineto\n")
}


func (x *Imp) arc (x0, x1, r, a, b float64) {
//
  x.write (x.f(x0) + " " + x.f(x1) + " " + x.f(r) + " " + x.f(a) + " " + x.f(b) + " arc\n")
}


func (x *Imp) scale (s float64) {
//
  x.write ("1 " + strconv.FormatFloat (s, 'f', 4, 64) + " scale\n")
}


func New () *Imp {
//
  x:= new (Imp)
  x.lw = 0.4
  const ppi = ker.PointsPerInch
  return x
}


func (x *Imp) Name (n string) {
//
  var err error
  x.file, err = os.Create (n + ".ps")
  if err != nil { panic ("open error") }
  x.write ("%!PS-Adobe-2.0\n")
  x.write ("%%Creator murus/psp.go (c) Christian Maurer\n")
//  x.write ("%%BoundingBox: 0 0 596 842 \n") // A4
  x.write ("%%DocumentPaperSize: a4\n")
  x.write ("%%EndComments\n")
  x.write (x.f(x.lw) + " setlinewidth\n")
  x.SetFont (font.Roman)
  x.SwitchFontsize (font.Normal)
  x.write ("72 72 translate\n")
}


func (x *Imp) Terminate () {
//
  x.write ("showpage\n")
  x.file.Close ()
}


func (x *Imp) SetUnit (pt float64) {
//
  x.write (x.f(pt) + " dup scale\n")
}


func (x *Imp) Translate (l, b float64) {
//
  x.write (x.f(l) + " " + x.f(b) + " translate\n")
}


func g (n uint8) string {
//
  return strconv.FormatFloat (float64(n) / 255, 'f', 4, 64)
}


func (x *Imp) SetColour (c col.Colour) {
//
  x.write (g (c.R) + " " + g (c.G) + " " + g (c.B) + " setrgbcolor\n")
}


func (x *Imp) SetFont (f font.Font) { // TODO
//
  var s string
  switch f {
  case font.Roman:
    s = "terminus-normal 16"
  case font.Bold:
    s = "terminus-bold"
  case font.Slanted:
    s = "Times-Roman" // nonsense
  case font.Italic:
    s = "Times-Roman-Italic" // nonsense
  }
  x.write ("/" + s + " findfont\n")
}


func (x *Imp) SwitchFontsize (f font.Size) {
//
  var h int
  switch f {
  case font.Tiny:
    h =  7
  case font.Small:
    h = 10
  case font.Normal:
    h = 16
  case font.Big:
    h = 24
  case font.Huge:
    h = 32
  }
  x.write (strconv.Itoa(h) + " scalefont setfont\n")
}


func (x *Imp) Write (s string, x0, y0 float64) {
//
  x.newpath ()
  x.moveto (x0, y0)
  for i:= 0; i < len (s); i++ {
    if z.IsLatin1 (s[i]) {
      x.write ("/" + z.Postscript (s[i]) + " glyphshow\n")
    } else {
      x.write ("(" + string(s[i]) + ") show\n")
    }
  }
  x.stroke ()
}


func (x *Imp) SetLinewidth (w float64) {
//
  x.lw = w
  x.write (x.f(x.lw) + " setlinewidth\n")
}


func (x *Imp) Point (x1, y1 float64) {
//
  x.newpath ()
  x.arc (x1, y1, x.lw, 0, 360)
  x.fill ()
  x.stroke ()
}


func (x *Imp) Points (xs, ys []float64) {
//
  n:= len(xs)
  if n == 0 || len(ys) != n { return }
  x.newpath ()
  for i:= 0; i < n; i++ {
    x.arc (xs[i], ys[i], 2 * x.lw, 0, 360)
    x.fill ()
  }
  x.stroke ()
}


func (x *Imp) Line (x1, y1, x2, y2 float64) {
//
  x.newpath ()
  x.moveto (x1, y1)
  x.lineto (x2, y2)
  x.stroke ()
}


func (x *Imp) Lines (x0, y0, x1, y1 []float64) {
//
  n:= len(x0)
  if n < 1 || len(y0) != n || len(x1) != n || len(y1) != n { return }
  x.newpath ()
  for i:= 0; i < n; i++ {
    x.moveto (x0[i], y0[i])
    x.lineto (x1[i], y1[i])
  }
  x.closepath ()
  x.stroke ()
}


func (x *Imp) Segments (xs, ys []float64) {
//
  n:= len (xs)
  if n < 1 || len (ys) != n { return }
  if n == 1 {
    x.Point (xs[0], ys[0])
    return
  }
  x.newpath ()
  x.moveto (xs[0], ys[0])
  for i:= 1; i < n; i++ {
    x.lineto (xs[i], ys[i])
  }
  x.stroke ()
}


func (x *Imp) Rectangle (x0, y0, w, h float64, f bool) {
//
  x.newpath ()
  x.moveto (x0, y0)
  x.rlineto (w, 0)
  x.rlineto (0, h)
  x.rlineto (-w, 0)
  x.closepath ()
  if f { x.fill () }
  x.stroke ()
}


func (x *Imp) Polygon (xs, ys []float64, f bool) {
//
  n:= len (xs)
  if n < 1 || len (ys) != n { return }
  if n == 1 {
    x.Point (xs[0], ys[0])
    return
  }
  x.newpath ()
  x.moveto (xs[0], ys[0])
  for i:= 1; i < n; i++ {
    x.lineto (xs[i], ys[i])
  }
  x.closepath ()
  if f { x.fill () }
  x.stroke ()
}


func (x *Imp) Arc (x0, y0, r, a, b float64) {
//
  x.newpath ()
  x.arc (x0, y0, r, a, b)
  x.stroke ()
}


func (x *Imp) Circle (x0, y0, r float64, f bool) {
//
  x.newpath ()
  x.arc (x0, y0, r, 0, 360)
  if f { x.fill () }
  x.stroke ()
}


func (x *Imp) Ellipse (x0, y0, a, b float64, f bool) {
//
  x.write ("/ellipse { 7 dict begin\n")
  x.write ("/" + x.f(b) + " exch def\n")
  x.write ("/" + x.f(a) + " exch def\n")
  x.write ("/" + x.f(y0) + " exch def\n")
  x.write ("/" + x.f(x0) + " exch def\n")
  x.write ("/mat matrix currentmatrix def\n")
  x.write (x.f(x0) + " " + x.f(y0) + " translate\n")
  x.write (x.f(a) + " " + x.f(b) + " scale\n")
  x.write ("0 0 1 0 360 arc\n")
  x.write ("mat setmatrix\n")
  x.write ("end\n")
  x.write ("} def\n")
  x.newpath ()
  x.write (x.f(x0) + " " + x.f(y0) + " " + x.f(a) + " " + x.f(b) + " ellipse\n")
  if f { x.fill () }
  x.stroke ()
}


func p (t, a float64, k uint) float64 {
//
  if k == 0 { return a }
  if k % 2 == 0 {
    return p (t * t, a, k / 2)
  }
  return p (t * t, t * a, k / 2)
}


func (x *Imp) nodes (xs, ys []float64) int {
//
  l:= len (xs)
  if l == 0 || l != len (ys) { return 0 }
  n:= 0
  for i:= 1; i < l; i++ {
    dx, dy:= math.Abs (xs[i] - xs[i-1]), math.Abs (ys[i] - ys[i-1])
    n += int(math.Sqrt (dx * dx + dy * dy + 0.5))
  }
  return n
}


func bezier (t float64, n uint, xs, ys []float64) (float64, float64) {
//
  var x, y float64
  for i:= uint(0); i <= n; i++ {
    a:= float64(ker.Binomial (n, i)) * p (1 - t, 1, n - i) * p (t, 1, i)
    x += a * xs[i]
    y += a * ys[i]
  }
  return x, y
}


func (x *Imp) Curve (xs, ys []float64) {
//
  n:= len (xs)
  if len (ys) != n { return }
  x.newpath ()
  m:= x.nodes (xs, ys)
  if m == 0 { return }
  x.moveto (xs[0], ys[0])
  for i:= 1; i < m; i++ {
    xb, yb:= bezier (float64(i) / float64(m), uint(n - 1), xs, ys)
    x.lineto (xb, yb)
  }
  x.stroke ()
}


// func init () { var _ PostScriptPage = New() }