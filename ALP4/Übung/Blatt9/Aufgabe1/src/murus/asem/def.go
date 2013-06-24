package asem

// (c) Christian Maurer   v. 120909 - license see murus.go

type
  AddSemaphore interface { // Spec see my book, p. 99

  P (n uint)

  V (n uint)
}
