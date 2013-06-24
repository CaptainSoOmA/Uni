package obj

// (c) Christian Maurer   v. 130115 - license see murus.go

type
  Object interface {

// Most objects in computer science can be compared with others,
// whether they are equal, and can be copied, so they have the type
  Equaler

// Furthermore, usually we can order objects; so they have the type
  Comparer

// Moreover they can be empty and may be cleared with the effect
// of being empty, hence they have the type
  Clearer

// and can be serialized into connected byte sequences,
// e.g. to be written to a storage device or transmitted
// over communication channels, so they have the type
  Coder
}


// Returns true, iff the type of a is bool, string, []byte,
// [u]int{8|16|32}, float[32|64] or complex[64|128].
func Atomic (a Any) bool {
//
  switch a.(type) {
  case bool, string, []byte,
       int8, int16, int, int32, int64,
       uint8, uint16, uint, uint32, uint64,
       float32, float64, complex64, complex128:
    return true
  }
  return false
}


// Returns true, iff the type of a implements Object.
func IsObject (a Any) bool {
//
  _, o:= a.(Object)
  _, e:= a.(Editor)
  return o || e
}


// Returns true, iff a is Atomic or implements Object.
func AtomicOrObject (a Any) bool {
//
  return Atomic (a) || IsObject (a)
}
