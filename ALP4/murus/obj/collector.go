package obj

// (c) Christian Maurer   v. 130115 - license see murus.go

// Collections of Objects or of variables of a concrete atomic type
// ([u]int.., float.., string, etc.). Every collector has either
// exactly one actual object or its actual object is undefined.
//
// In all specifications x denotes the calling collector.

// The constructor function "func New (a Any) *Imp"
// returns a new collector for objects of the type of a
// that is empty, so its actual object is undefined.

type
  Collector interface {

// Empty:   Returns true, iff x does not contain any object.
// Clr:     x is empty; its actual object is undefined.
  Clearer

// Returns true, iff the actual object of x is not defined.
  Offc () bool

// Returns the nunber of objects of x.
  Num () uint

// Pre: a has the type of the objects in x. 
// If x does not carry any order:
//   If the actual object of x was undefined, x is appended behind
//   the end of x, otherwise x is inserted before the actual object.
//   All other objects and their order in x and the actual object of x
//   are not influenced.
// Otherwise, i.e. if x is ordered w.r.t. to an
// order relation // (reflexive, transitive and antisymmetric)
// or a strict order relation (irreflexive and transitive):
//   x is inserted behind the last object b in x, for which r(b,a) == true,
//   i.e. that under r "is smaller" than a.
//   If r is a strict order ("<"), nothing has changed.
//   If r is a order ("<="), then a is contained once more in x,
//   also if it was already contained in x before.
//   If x was ordered w.r.t. r before, x is still ordered now.
// In both cases the actual object of x is now the object behind
// the inserted one, if there exists one, otherwise it is undefined.
  Ins (a Any)

// If x is not empty, then
// if f and the actual object of x was defined, then
// the actual object now is the object after the former actual object,
// if there exists one, otherwise it not is not defined.
// If !f and the actual object of x was defined and was not the first
// object of x, then it is now the object before the former one;
// if it was not defined, then it is now the last object of x.
// In all other cases, nothing has happened.
  Step (f bool)

// If f is empty, the actual object is undefined; otherwise for
// f/!f the actual object of x now is the last/first object of x.
  Jump (f bool)

// Returns true, iff for f/!f the last/first object of x is its actual object.
  Eoc (f bool) bool

// Returns a copy of the actual object of x, if that is defined; nil otherwise.
  Get () Any

// Pre: a has the type of the objects in x. 
// If the actual object of x was undefined, a copy of a is appended
// behind the end of x and is now the actual object of x.
// Otherwise the actual object of x is replaced by a.
  Put (a Any)

// Returns nil, if the actual object of x is not undefined,
// otherwise, the actual object and that was removed from x,
// and the actual object is now the object after it,
// if the former actual object was not the last object of x.
// In that case the actual object of x is now undefined.
  Del () Any

// Returns true, iff a is contained in x. In that case
// case the first such object is the actual object of x;
// otherwise, the actual object is the same as before.
  Ex (a Any) bool
}
