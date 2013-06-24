package set

// (c) Christian Maurer   v. 130115 - license see murus.go

import
  . "murus/obj"
type
  Set interface {

  Equaler
  Coder
  Sorter
  Iterator

  Write (x0, x1, y, dy uint, f func (Any) string)
}
