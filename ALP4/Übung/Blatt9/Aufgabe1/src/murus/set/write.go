package set

// (c) Christian Maurer   v. 130115 - license see murus.go

import (
  "murus/obj"
  . "murus/scr"
)


func (x *tree) write (x0, x1, y, dy uint, f func (obj.Any) string) {
//
  if x == nil { return }
  xm:= (x0 + x1) / 2
  y1:= int(y + NY1() / 2) - 1
  if x.left != nil {
    Line (int(xm), y1, int(x0 + xm) / 2, y1 + int (dy))
  }
  if x.right != nil {
    Line (int(xm), y1, int(xm + x1) / 2, y1 + int (dy))
  }
  WriteGr (f (x.root), int(xm - NX1()), int(y))
  x.left.write (x0, xm, y + dy, dy, f)
  x.right.write (xm, x1, y + dy, dy, f)
}


func (x *Imp) Write (x0, x1, y, dy uint, f func (obj.Any) string) {
//
  x.anchor.write (x0, x1, y, dy, f)
}
