package perm

// (c) Christian Maurer   v. 111114 - license see murus.go

type
  Permutation interface { // Permutations of the first n natural numbers.

// Pre: n > 1.
// x has size n, i.e. it is a random permutation of the natural numbers < n.
// New (n uint) *Imp

// x is a randomly permuted.
  Permute ()

// Returns 0, if k >= size of x,
// returns otherwise the k-th number of x.
  F (k uint) uint
}
