package bpqu

// (c) Christian Maurer   v. 120909 - license see murus.go */

import
  "murus/pqu"
type
  BoundedPrioQueue interface {

// A bounded priority queue is a
  pqu.PrioQueue // with bounded capacity.

// Returns true, iff x is filled up to its capacity.
  Full () bool
}
