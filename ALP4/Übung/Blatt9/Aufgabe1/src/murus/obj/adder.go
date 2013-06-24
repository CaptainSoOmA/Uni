package obj

// (c) Christian Maurer   v. 130115 - license see murus.go

type
  Adder interface {

// Returns true, iff x is neutral w.r.t. addtion.
  Null () bool

// x = y + z.
  Add (y, z Adder)

// x = x0 + y, where x0 denotes x before.
  Plus (y Adder)

// x = y - z.
  Sub (y, z Adder)

// x = x0 - y, where x0 denotes x before.
  Minus (y Adder)
}
