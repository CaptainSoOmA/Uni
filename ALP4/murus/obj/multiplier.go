package obj

// (c) Christian Maurer   v. 130115 - license see murus.go

type
  Multiplier interface {

// Returns true, iff x is neutral w.r.t. all multiplication.
  One () bool

// x = y * z.
  Mul (y, z Multiplier)

// x = x0 * y, where x0 denotes x before.
  Times (y Multiplier)

// x = x0 * x0, where x0 denotes x before.
  Sqr ()

// x = y / z.
  Div (y, z Multiplier)

// x = x0 / y, where x0 denotes x before.
  DivBy (y Multiplier)
}
