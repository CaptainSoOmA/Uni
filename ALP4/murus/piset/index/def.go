package index

// (c) Christian Maurer   v. 121216 - license see murus.go

import
  . "murus/obj"
type
  IndexObject interface { // TODO detailed explanations

  Object

// x has indexobject a and position n.
  Set (a Any, n uint)

// Returns the position of x.
  Pos () uint
}
