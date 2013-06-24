package dlock

// (c) Christian Maurer   v. 121109 - license see murus.go

import
  . "murus/lock"
type
  DistributedLock interface {

// h is the slice of the names of all hosts involved.
// New (h []string) *Imp

// h is the slice of all hosts involved,
// p the slice of different ports (with len(p) == len(h)).
//  NewR (h []*host.Imp, p []uint) *ImpR

  Locker
}
