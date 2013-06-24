package pbar

// (c) Christian Maurer   v. 121215 - license see murus.go

import (
  "murus/col"
  "murus/scr"
)
const
  min = 8
type
  Imp struct {
         max uint
  horizontal bool
       value,
      x0, y0, // left top corner
       width,
      height uint
      cF, cB col.Colour
            }


func New (h bool) *Imp {
//
  B:= new (Imp)
  l:= scr.NLines()
  c:= scr.NX()
  B.max = 100
  B.horizontal = h
  B.value = 0
  if h {
    B.width = c
    B.height = scr.NY1()
    B.Locate (0, (l - 1) * B.height, B.width, B.height)
  } else {
    B.width = scr.NX1()
    B.height = scr.NY()
    B.Locate (c - B.width, 0, B.width, B.height)
  }
  B.cF, B.cB = col.HintF, col.HintB
  B.Locate (B.x0, B.y0, B.width, B.height)
  return B
}


func (B *Imp) terminieren () {
//
  scr.RestoreGr (B.x0, B.y0, B.x0 + B.width, B.y0 + B.height)
}


func (B *Imp) Locate (x, y, w, h uint) {
//
  if x + min > scr.NX() {
    x = scr.NX() - min
  }
  B.x0 = x
  if y + min > scr.NY() {
    y = scr.NY() - min
  }
  B.y0 = y
  if w < min { w = min }
  if B.x0 + w > scr.NX() {
    w = scr.NX() - B.x0
  }
  B.width = w
  if h < min { h = min }
  if B.y0 + h > scr.NY() {
    h = scr.NY() - B.y0
  }
  B.height = h
  scr.SaveGr (B.x0, B.y0, B.x0 + B.width, B.y0 + B.height)
}


func (B *Imp) Def (c uint) {
//
  B.max = c
}


func (B *Imp) Fill (i uint) {
//
  if i > B.max { i = B.max }
  B.value = i
}


func (B *Imp) Filldegree () uint {
//
  return B.value
}


func (B *Imp) SetColours (f, b col.Colour) {
//
  B.cF = f
  B.cB = b
}


func (B *Imp) Write () {
//
  scr.Colour (B.cF)
  var d uint
  if B.horizontal {
    scr.Rectangle (int(B.x0), int (B.y0), int(B.x0 + B.width), int(B.y0 + B.height - 1))
    d = ((B.width - 1) * B.value) / B.max
    scr.RectangleFull (int(B.x0), int(B.y0 + 1), int(B.x0 + d), int(B.y0 + B.height - 2))
    scr.Colour (B.cB)
    if d < B.width - 1 {
      scr.RectangleFull (int(B.x0 + d + 1), int(B.y0 + 1), int(B.x0 + B.width - 1), int(B.y0 + B.height - 2))
    }
  } else {
    scr.Rectangle (int(B.x0), int(B.y0), int(B.x0 + B.width - 1), int(B.y0 + B.height))
    d = ((B.height - 1) * B.value) / B.max
    scr.RectangleFull (int(B.x0 + 1), int(B.y0 + B.height - d), int(B.x0 + B.width - 2), int(B.y0 + B.height))
    scr.Colour (B.cB)
    if d < B.height - 1 {
      scr.RectangleFull (int(B.x0 + 1), int(B.y0 + B.height - 1 - d), int(B.x0 + B.width - 2), int(B.y0 + 1))
    }
  }
}


func (B *Imp) Edit (i *uint) {
//
  xi, yi:= scr.MousePosGr ()
  x, y:= xi, yi
  if scr.UnderMouseGr (int(B.x0), int(B.y0), int(B.x0 + B.width - 1), int(B.y0 + B.height - 1), 0) {
    if B.horizontal {
      B.value = (uint(int(x) - int(B.x0)) * B.max) / (B.width - 1)
    } else {
      B.value = B.max - (uint(int(y) - int(B.y0)) * B.max) / (B.height - 1)
    }
  }
  *i = B.value
  B.Write ()
}
