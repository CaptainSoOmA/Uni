package obj

// (c) Christian Maurer   v. 130115 - license see murus.go

type
  Iterator interface {

  Collector

//  ExGeq (a Object) bool // TODO

// Returns the number of those objects in x, for which p returns true.
  NumPred (p Pred) uint

// Returns NumPred(p) == Num(), i.e. true, iff p returns true
// on all objects of x, particularly if x is empty.
  All (p Pred) bool

// Returns true, iff there is an object in x, for which p returns true.
// In that case the actual object of x is for f/!f the last/first such
// object, otherwise the actual object of x is the same as before.
  ExPred (p Pred, f bool) bool

// Returns true, iff there is an object in x in direction f
// from the actual object of x, for which p returns true.
// In that case the actual object of x is for f/!f the next/previous
// such object, otherwise the actual object of x is the same as before.
  StepPred (p Pred, f bool) bool

// Returns true, if x.Mum() == y.Num() and for each i < x.Num()
// the i-th object of x and of y (in this order) are in relation unter r.
  Equiv (y Iterator, r Rel) bool

// Pre: If x is ordered, o is strongly monotone with respect to that order,
//      i.e. x < y implies o(x) < o(y), where < denotes the order.
// o was applied to all objects of x (in their order in x).
// The actual object of x is the same as before.
  Trav (op Op)

// Pre: See Trav.
// If the actual object of x was undefined, nothing has changed.
// op was applied to all objects of x (in their order in x),
// for which p returns true.
// The actual object of x is the same as before.
  TravPred (p Pred, o Op)

// Pre: See Trav.
// o(-, p(-)) was applied to all objects in x (in their order in x), i.e.
// o(-, true) was applied to all objects of x, for which p returns true,
// otherwise p(-, false). The actual object of x is the same as before.
  TravCond (p Pred, o CondOp)

// Pre: y is a collector of objects of the same type as x
// (especially contains objects of the same type as a).
// y consists exactly of those objects of x before
// (in their order in x), for which p returns true.
// The actual object of x is undefined; y is unchanged.
  Filter (y Iterator, p Pred)

// Pre: See Filter.
// y consists exactly of those objects of x (in their order in x),
// for which p returns true, and exactly those objects are removed from x.
// The actual objects of x and y are undefined.
  Cut (y Iterator, p Pred)

// In x all objects, for which p returns true, are removed.
// If the actual object of x was one of them, now it is undefined.
  ClrPred (p Pred)

// Pre: See Filter.
// If the actual object of x was undefined, x is not changed and y is empty.
// Otherwise y contains exactly the objects before the former actual object in x
// and y contains the former actual object of x and exactly all objects behind it,
// both in their former order in x. In this case the actual object of x is undefined
// and the actual object of y is its first object.
  Split (c Iterator)

// Pre: See Filter.
// If x == y or if x and y do not have the same type, nothing has changed. Otherwise:
// If x does not carry any order:
//   x consists of exactly all objects of x before (in their order in x)
//   and behind them all objects of y before (in their order in y).
//   If the actual object of x was undefined, now the former first object of y is
//   the actual object of x, otherwise the actual object of x is the same as before.
//   y is empty and its actual object is undefined.
// Otherwise, i.e. if x is ordered w.r.t. to an order relation,
//   Pre: r is a strict order ("<") and x and y are ordered w.r.t. r,
//        and do not contain any object more than once, or r is a order ("<=").
//   x consists exactly of all objects of x and y before.
//   If r is strict ("<"), then the objects, which are contained in x as well as in y,
//   are contained in x only once, otherwise ("<=") in their multiplicity.
//   x is ordered w.r.t. r and y is empty. The actual objects of x and y are undefined.
  Join (y Iterator)

// If x is empty or x is ordered, nothing has changed.
// Otherwise x contains exactly the same objects as before in their former order,
// with the following exception: for f == true the former last object of x
// is now the first one and for f == false the former first object of x
// is now the last one. The actual object of x is the same as before.
//  Rotate () // ? TODO
}
