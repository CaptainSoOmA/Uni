package pt

// (c) Christian Maurer   v. 121216 - license see murus.go

import (
  . "murus/obj"
  "murus/col"
  "murus/vect"
)
type
  Class byte; const (
  None = iota; Start; Light
  Points; Lines; LineStrip; LineLoop
  Triangles; TriangleStrip; TriangleFan
  Quads; QuadStrip
  Polygon
  nClasses
)
type // Coloured points in 3-space with a class, a current number and a normal vector.
  Point interface {

// Returns an point with NoClass, number 1, coordinates (0, 0, 0), normal (0, 0, 0) and foregroundcolour of the screen.
//  New () *Imp

  Object

//  Terminate ()

// x is the endpoint of v with class c, number a, colour f and normal n.
  Set (c Class, a uint, f col.Colour, v, n *vect.Imp)

// Returns the class of x.
  ClassOf () Class

// Returns the current number of x.
  Number () uint

// Returns the colour of x.
  Colour () col.Colour

// Returns the vector with the endpoint x.
  Read () *vect.Imp

// Returns the vector with the endpoint x and the normal of x.
  Read2 () (*vect.Imp, *vect.Imp)
}
