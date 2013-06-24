package nchan

// (c) Christian Maurer   v. 121029 - license see murus.go

//     Nichtsequentielle Programmierung mit Go 1 kompakt, S. 194

import
  . "murus/obj"
const
  Port0 = 49152 // first private port
type
// Channels for passing objects between two processes over the net.
// x always means the calling channel.
  NetChannel interface {

// Pre: c is the name of a computer, either in /etc/hosts or resolvable per DNS.
//      p0+p < 2^16;
//      the port p0+n is not used by a network service on the local host or on c.
// Returns x with the type of a and the port p0+p.
// If o == true, x is a 1:1-channel with c as communication partner of the calling process of x,
// otherwise, x is a multiplex channel with c as server.
// The port of x is now used by a network service on the local host and c.
// x is not yet activated.
//  New0 (a Any, c string, p uint16, o bool) *Imp

// see New0. x is activated.
//  New (a Any, c string, p uint16, o bool) *Imp

// x is activated.
  Go ()

// Returns true, iff the calling process is the server process.
  IsServer () bool

// Spec TODO
  Chan () (chan Any, chan Any)

// Pre: a is of the type of x.
// The object a is sent on x (resp. if x is a multiplex channel
// and the calling process is a server, on the actual subchannel of x)
// to the communication partner of the calling process.
// The calling process was blocked until the object was received.
  Send (a Any)

// Returns the object of the type of x, that was received on x (resp. if x as a multiplex channel
// and the calling process is a server, on the actual subchannel of x)
// from the communication partner, if that was completely received, nil /* an empty object or a zero value ? */ otherwise.
// The calling process was blocked, until an object was received.
  Recv () Any

// Spec TODO
  Terminate ()
}
