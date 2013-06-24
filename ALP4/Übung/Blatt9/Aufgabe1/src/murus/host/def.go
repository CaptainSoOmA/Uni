package host

// (c) Christian Maurer   v. 121120 - license see murus.go

import
  . "murus/obj"
const ( // Format
  Hostname = iota
  IPnumber
  NFormats
)

// Returns true, iff h is the name of the computer, that runs the calling process.
func Local (h string) bool { return local(h) }

// Returns the local host
func LocalHost () Host { return localHost }

type
  Host interface { // Hostnames and their IP-numbers

  Editor
  Formatter
  Stringer
//  Printer TODO
  Marker

// Returns the IP4-number of x converted into a uint32.
// >>>  deprecated  !!!
  Number () uint32

// Returns the IP-number of x as byte sequence.
  IP () []byte

// Returns true, if x has the name or the IP-number s.
  Equiv (s string) bool

// Returns true, iff the IP4-number of x is less than the IP4-number of the local host.
  Sonstewas () bool // TODO name of function
}
