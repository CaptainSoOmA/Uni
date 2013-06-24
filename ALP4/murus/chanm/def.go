package chanm

// (c) Christian Maurer   v. 120909 - license see murus.go

//     Nichtsequentielle Programmierung mit Go 1 kompakt, 7.7, S. 185 ff.

import
  . "murus/obj"
type
  ChannelModel interface { // "models" of channels (i.e. working only within one process);

// Returns an new Channelmodel, that has the type a,
// i.e. on which messages of that type can be sent and received.
// New (a) *Imp

// Pre: a is of the type of x.
// a is contained in x.
  Send (a Any)

// Returns true, iff there are no messages in x.
  Empty () bool

// Pre: a is the address of an object of the type of x.
// The calling process was blocked, until x contained a message.
// *a is now that message and is is removed from x.
  Recv (a Any) // a == &object to receive !
}
