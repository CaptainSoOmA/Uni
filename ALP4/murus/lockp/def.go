package lockp

// (c) Christian Maurer   v. 111127 - license see murus.go

// Ensures the access to a critical section.
// The functions Lock and Unlock cannot be interrupted
// by calls of Lock or Unlock of other goroutines.

type
  LockerP interface {

// Pre: The calling goroutine is not in the critical section.
// It is the only one in the critical section.
  Lock (p uint)

// Pre: The calling goroutine is in the critical section.
// It is not in the critical section.
  Unlock (p uint)
}
