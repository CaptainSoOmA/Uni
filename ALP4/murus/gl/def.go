package gl

// (c) Christian Maurer   v. 120501 - license see murus.go

type
  Figure byte; const (
  POINTS = iota;
  LINES
  LINE_LOOP
  LINE_STRIP
  TRIANGLES
  TRIANGLE_STRIP
  TRIANGLE_FAN
  QUADS
  QUAD_STRIP
  POLYGON
  UNDEF
  LIGHT
)
const
  MaxL = 16 // <= GL.GL_MAX_LIGHTS
