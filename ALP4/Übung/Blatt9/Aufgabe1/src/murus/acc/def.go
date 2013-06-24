package acc

// (c) Christian Maurer   v. 120220 - license see murus.go

import
  "murus/euro"
type
  Account interface {

// TODO Spec
  Deposit (e *euro.Imp) *euro.Imp

// TODO Spec
  Draw (e *euro.Imp) *euro.Imp
}
