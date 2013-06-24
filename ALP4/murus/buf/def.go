package buf

// (c) Christian Maurer   v. 120910 - license see murus.go

import
  "murus/qu"
type
  Buffer interface { // Queues of bounded capacity;
                     // could also be named "bqu"
                     // The exported functions cannot be interrupted
                     // by calls of these functions of other goroutines.

  qu.Queue

// The default constructor returns an empty buffer of capacity n:
//   func New (a Any, n uint) *Imp

// Num() and Empty():
// For [M|F]Buffer: Should not be used, because the number
// of objects could have been changed immediately after the call.

// For Buffer: Returns true, iff x is filled up to its capacity.
// For [M|F]Buffer: Should not be used, because the number
// of objects could have been changed immediately after the call.
  Full() bool

// For Buffer: Pre: ! x.Full().
// For [M|F]Buffer: The calling process was blocked, until x was not full.
// a is inserted as last object in X.
//   Ins (a Any)

// For Buffer: Pre: ! x.Empty().
// For [M|F]Buffer: The calling process was blocked, until x was not empty.
// Returns the first object of x; that object is removed from x.
//   Get() Any
}
