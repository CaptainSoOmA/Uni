package gl

// (c) Christian Maurer   v. 120909 - license see murus.go

// #cgo LDFLAGS: -lGL
// #include <GL/gl.h> 
import
  "C"
import (
  "math"
  "murus/ker"; "murus/spc"
  "murus/col"; "murus/scr"
  "murus/vect"
)
const (
  pack = "gl"
  nLamp = 12
)
var (
  lmAmb, mAmbi, mDiff [4]C.GLfloat
  sin, cos [nLamp + 2]C.GLdouble
  lightSource [MaxL]*vect.Imp
  lightColour /* diffus */ [MaxL]col.Colour
  lightInitialized [MaxL]bool
  lightVis bool
  initialized bool
  aa, dd [MaxL][4]C.GLfloat
  matrix [4][4]C.GLdouble
  right, front, top, eye [3]float64
)


func Cls (c col.Colour) {
//
  r, g, b:= col.Float (c) // 0.0, 0.0, 0.0
  C.glClearColor (C.GLclampf(r), C.GLclampf(g), C.GLclampf(b), C.GLclampf(0.0))
}


func Init (fern float64) { // called by points.Start ()
//
  const (
    D = 2.0 // -fache Bildschirmbreite
    nah = C.GLdouble(0.2)
  )
  if ! initialized {
    initialize ()
  }
  C.glMatrixMode (C.GL_PROJECTION)
  C.glLoadIdentity ()
  deg:= D * math.Atan ((0.5 / D) / scr.Proportion())
  deg /= 0.9 // experimentelle Weitwinkelkorrektur
  var m [4][4]C.GLdouble
  m[1][1] = 1.0 / C.GLdouble(math.Tan (deg)) // Cot
  m[0][0] = m[1][1] / C.GLdouble(scr.Proportion())
//  delta:= C.GLdouble(fern) - nah
//  m[2][2] = - (C.GLdouble(fern) + nah) / delta
//  m[2][3] = GLdouble(-1.0)
//  m[3][2] = -2. * nah * C.GLdouble(fern) / delta
  m[2][2] = C.GLdouble(-1.0)
  m[2][3] = C.GLdouble(-1.0)
  m[3][2] = C.GLdouble(-1.0) * nah
  C.glMultMatrixd (&m[0][0])
//  q:= C.GLdouble(0.75)
//  GLFrustum (-1.0 * nah, 1.0 * nah, -q * nah, q * nah, 1.0 * nah, C.GLdouble(fern))
  C.glMatrixMode (C.GL_MODELVIEW)
}


// Pre: n < MaxL, 0 <= h[i] <= 1 für i = 0, 1.
// Wenn Licht n schon eingeschaltet war, ist nichts verändert; andernfalls ist es
// an der Position v in Farbe f mit der Ambienz h[0] und der Diffusität h[1] eingeschaltet.
func InitLight (n uint, v, h *vect.Imp, c col.Colour) {
//
//
  if lightInitialized[n] { return }
  var a [4]float64
  a[0], a[1], a[2] = h.Coord3 ()
  // Arbeitsdrumrum, weil die Punkte bisher nur eine Farbe transportieren, hier die diffuse.
  // In L wird die ambiente Farbe geliefert.
  for i:= 0; i < 3; i++ { aa[n][i] = C.GLfloat(a[i]) }; aa[n][3] = C.GLfloat(1.0)
  lightColour[n] = c
  d0, d1, d2:= col.Float (c)
  dd[n][0], dd[n][1], dd[n][2] = C.GLfloat(d0), C.GLfloat(d1), C.GLfloat(d2)
  dd[n][3] = C.GLfloat(1.0)
  lightSource[n].Copy (v)
  ActualizeLight (n)
  lightInitialized[n] = true
}


// Pre: Licht n ist eingeschaltet.
// Licht n hat die Position v.
func PosLight (n uint, v *vect.Imp) {
//
  if ! lightInitialized[n] { ker.Stop (pack, 3) }
  lightSource[n].Copy (v)
}


func ActualizeLight (n uint) { // n < MaxL
//
  var L [4]float64
  L[0], L[1], L[2] = lightSource[n].Coord3 ()
  var l [4]C.GLfloat
  for i:= 0; i < 3; i++ { l[i] = C.GLfloat(L[i]) }; l[3] = C.GLfloat(1.0)
  C.glLightfv (C.GL_LIGHT0 + C.GLenum(n), C.GL_POSITION, &l[0])
  C.glLightfv (C.GL_LIGHT0 + C.GLenum(n), C.GL_AMBIENT, &aa[n][0])
  C.glLightfv (C.GL_LIGHT0 + C.GLenum(n), C.GL_DIFFUSE, &dd[n][0])
  C.glEnable (C.GL_LIGHT0 + C.GLenum(n))
}


var (
  yyy [3]C.GLdouble
  nn uint
  fig Figure = POINTS
)

func vector2yyy (v *vect.Imp) {
//
  for i:= 0; i < 3; i++ {
    yyy[i] = C.GLdouble(v.Coord (spc.Direction(i)))
  }
}


func Write0 () {
//
  if ! initialized {
    initialize ()
  }
  if ! scr.UnderX () { ker.Stop (pack, 1) }
  C.glMatrixMode (C.GL_MODELVIEW)
  C.glLoadIdentity ()
  for i:= 0; i < 3; i++ {
    matrix[i][0] = C.GLdouble(right[i])
    matrix[i][1] = C.GLdouble(top[i])
    matrix[i][2] = C.GLdouble(-front[i])
  }
  C.glMultMatrixd (&matrix[0][0])
  C.glTranslated (C.GLdouble(-eye[0]), C.GLdouble(-eye[1]), C.GLdouble(-eye[2]))
  C.glClear (C.GL_COLOR_BUFFER_BIT + C.GL_DEPTH_BUFFER_BIT)
  for n:= uint(0); n < MaxL; n++ {
    if lightInitialized[n] {
      ActualizeLight (n)
    }
  }
  C.glBegin (POINTS)
  nn = 0
}


func Write (f Figure, a uint, V, N []*vect.Imp, c col.Colour) {
//
  switch f { case UNDEF:
    nn = 0 // forces glEnd / glBegin
println ("gl.Write UNDEF")
    return
  case LIGHT:
    lightVis = true
    if a >= MaxL { ker.Stop (pack, 2) }
    InitLight (a, V[0], N[0], c)
    nn = 0
println ("gl.Write LIGHT")
    return
  }
  if f != fig || a != nn || nn == 0 {
    fig = f
    nn = a
    C.glEnd ()
    C.glBegin (C.GLenum(f))
  }
  C.glColor3ub (C.GLubyte(c.R), C.GLubyte(c.G), C.GLubyte(c.B))
  for i:= uint(0); i < a; i++ {
println ("gl.Write", i)
    vector2yyy (V[i]); C.glVertex3dv (&yyy[0])
    vector2yyy (N[i]); C.glNormal3dv (&yyy[0])
  }
/*
  tmp:= vect.New ()
  C.glEnd ()
  for i:= uint(0); i < a; i++ {
    C.glBegin (LINES)
    C.glColor3ub (C.GLubyte(0), C.GLubyte(255), C.GLubyte(0))
    vector2yyy (V[i]); C.glVertex3dv (&yyy[0])
    tmp.Copy (V[i])
    tmp.Inc (N[i])
    vector2yyy (tmp); C.glVertex3dv (&yyy[0])
    C.glEnd ()
  }
  nn = 0
  C.glBegin (POINTS)
*/
}


// func Write1 (d chan bool) {
func Write1 () {
//
  C.glEnd()
  if lightVis {
    for n:= uint(0); n < MaxL; n++ { lamp (n) }
  }
  scr.WriteGlx ()
//  d <- true
}


func Actualize (R, V, O, A *vect.Imp) {
//
  right[0], right[1], right[2] = R.Coord3 ()
  front[0], front[1], front[2] = V.Coord3 ()
  top[0], top[1], top[2] = O.Coord3 ()
  eye[0], eye[1], eye[2] = A.Coord3 ()
}

/*
func Hold () {
//
  C.glPushMatrix ()
}


func Continue () {
//
  C.glPopMatrix ()
}
*/

func lamp (n uint) {
//
  if ! lightInitialized[n] {
    return
  }
  xx, yy, zz:= lightSource[n].Coord3 ()
  x, y, z:= C.GLdouble(xx), C.GLdouble(yy), C.GLdouble(zz)
  r:= C.GLdouble(0.1)
  C.glBegin (TRIANGLE_FAN)
  C.glColor3ub (C.GLubyte(lightColour[n].R), C.GLubyte(lightColour[n].G), C.GLubyte(lightColour[n].B))
  C.glNormal3d (C.GLdouble(0.0), C.GLdouble(0.0), C.GLdouble(-1.0))
  C.glVertex3d (C.GLdouble(x), C.GLdouble(y), C.GLdouble(z + r))
  r0:= r * sin[1]
  z0:= z + r * cos[1]
  for l:= 0; l <= nLamp; l++ {
    C.glNormal3d (-sin[1] * cos[l], -sin[1] * sin[l], -cos[1])
    C.glVertex3d (x + r0 * cos[l],   y + r0 * sin[l],   z0)
  }
  C.glEnd ()
  C.glBegin (QUAD_STRIP)
  var r1, z1 C.GLdouble
  for b:= 1; b <= nLamp / 2 - 2; b++ {
    r0, z0 = r * sin[b],   z + r * cos[b]
    r1, z1 = r * sin[b+1], z + r * cos[b+1]
    for l:= 0; l <= nLamp; l++ {
      C.glNormal3d (-sin[b+1] * cos[l], -sin[b+1] * sin[l], -cos[b+1])
      C.glVertex3d (x + r1 * cos[l], y + r1 * sin[l], z1)
      C.glNormal3d (-sin[b] * cos[l], -sin[b] * sin[l], -cos[b])
      C.glVertex3d (x + r0 * cos[l], y + r0 * sin[l], z0)
    }
  }
  C.glEnd ()
  C.glBegin (TRIANGLE_FAN)
  C.glNormal3d (0., 0., 1.)
  C.glVertex3d (x, y, z - r)
  r0, z0 = r * sin[1], z - r * cos[1]
  b:= nLamp / 2 - 1
  for l:= 0; l <= nLamp; l++ {
    C.glNormal3d (-sin[b] * cos[l], -sin[b] * sin[l], -cos[b])
    C.glVertex3d (x + r0 * cos[l], y + r0 * sin[l], z0)
  }
  C.glEnd ()
}


func initialize () {
//
  if ! initialized {
    initialized = true
    scr.Switch (scr.XGA) // (scr.MaxMode())
    C.glViewport (0, 0, C.GLsizei(scr.NX()), C.GLsizei(scr.NY()))
  }
}


func init () {
//
  right[0], front[1], top[2] = 1.0, 1.0, 1.0
  matrix[3][3] = 1.
  for l:= 0; l < MaxL; l++ {
    lightSource[l] = vect.New()
  }
  w:= 2.0 * math.Pi / float64 (nLamp)
  sin[0], cos[0] = C.GLdouble(0.0), C.GLdouble(1.0)
  sin[nLamp], cos[nLamp] = sin[0], cos[0]
  for g:= 1; g < nLamp; g++ {
    sin[g] = C.GLdouble(math.Sin (float64 (g) * w))
    cos[g] = C.GLdouble(math.Cos (float64 (g) * w))
  }
  sin[nLamp+1], sin[1] = cos[nLamp+1], cos[1]
//  C.glDepthFunc (C.GL_LESS) // default
  C.glEnable (C.GL_DEPTH_TEST)
  C.glShadeModel (C.GL_SMOOTH)
  for i:= 0; i < 3; i++ { lmAmb[i] = C.GLfloat(0.2) } // default: 0.2
  lmAmb[3] = C.GLfloat(1.0) // default: 1.0
  C.glLightModelfv (C.GL_LIGHT_MODEL_AMBIENT, &lmAmb[0])
  for i:= 0; i < 3; i++ { mAmbi[i] = C.GLfloat(0.2) } // default: 0.2
  mAmbi[3] = C.GLfloat(1.0) // default: 1.0
//  C.glLightModelfv (C.GL_LIGHT_MODEL_TWO_SIDE, 1)
  C.glMaterialfv (C.GL_FRONT_AND_BACK, C.GL_AMBIENT_AND_DIFFUSE, &mAmbi[0])
  for i:= 0; i < 3; i++ { mDiff[i] = C.GLfloat(0.8) } // default: 0.8
  mDiff[3] = C.GLfloat(1.0) // default: 1.0
  w = 1.
  C.glClearDepth (C.GLclampd(w))
//  C.glMaterialfv (C.GL_FRONT_AND_BACK, C.GL_DIFFUSE, mDiff)
//  C.glColorMaterial (C.GL_FRONT_AND_BACK, C.GL_DIFFUSE)
//  C.glColorMaterial (C.GL_FRONT, C.GL_AMBIENT)
  C.glColorMaterial (C.GL_FRONT_AND_BACK, C.GL_AMBIENT_AND_DIFFUSE)
  C.glEnable (C.GL_COLOR_MATERIAL)
  C.glEnable (C.GL_LIGHTING)
  initialize ()
}
