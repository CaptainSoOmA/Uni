package cond

// (c) Christian Maurer   v. 130526 - license see murus.go

type
  Condition interface {

  Wait (b bool)
  Signal () bool
  SignalAll ()
//  Awaited() bool
}
