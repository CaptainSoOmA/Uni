package stk

// (c) Christian Maurer   v. 120909 - license see murus.go

import
  . "murus/obj"
type
  Stack interface {

// x.Top() equals a and the result of x.Push(); x.Pop() is x; i.e.
// a is the object on top of x, the stack below a is x before.
  Push (a Any)

// Returns true, iff there is no object on x.
// In the concurrent case, this value is not reliable,
// as another process could have pushed on object on the stack
// immediately after the call.
// Senseless to be called for mstacks or fstacks !
//  Empty () bool

// If x was empty, nothing has happened.
// In the concurrent case, the calling process has been blocked,
// until x was not empty; otherwise, the object on top of x is
// removed; i.e. x now equals the stack below that object before.
  Pop ()

// Returns nil, if x is empty, otherwise a copy of the object
// on top of x.
// x is not changed.
  Top () Any
}
