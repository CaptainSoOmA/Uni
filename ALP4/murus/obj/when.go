package obj

// (c) Christian Maurer   v. 130526 - license see murus.go

// selective communication
func When (b bool, c chan Any) chan Any {
  if b {
    return c
  }
  return nil
}
