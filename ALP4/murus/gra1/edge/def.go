package edge

// (c) Christian Maurer   v. 121215 - license see murus.go

// >>> Nur zur Implementierung von graph1, soll weiter oben nicht verwendet werden !

import (
  . "murus/obj"; "murus/col"
  "murus/gra1/node"
)
var
  WithValues bool = true
type
// Edges with natural numbers < 10^... as values, represented as lines on the screen.
// The endpoints of the edges are defined by their nodes.
  Edge interface {

  Object // Empty edges have value 1.
//  Valuator // missing Ok TODO

// x has the value n.
  Def (n uint) // n < 10^...

// f is the normal and a the actual colour of all edges.
  SetColours (f, a col.Colour)

// x is written on the screen between n and n1.
  Write (n, n1 node.Node, d, vis, inv bool)

// x is written at its position, for a = true ...
  WriteCond (n, n1 node.Node, g, a bool)

// x has the name and the value edited by the user.
  Edit (n, n1 node.Node, g bool)

  Val () uint // Valuator ? - missing Ok TODO
}
