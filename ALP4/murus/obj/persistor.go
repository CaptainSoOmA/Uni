package obj

// (c) Christian Maurer   v. 120909 - license see murus.go

type
  Persistor interface {
// In all specifications x denotes the calling persistent object.

// x is defined with the name n, i.e. it is that object, that was last time
// defined with that name, if there exists such; otherwise it is empty.
  Name (n string)

// If there is no other persistent object with the name n,
// x is defined with that name; otherwise nothing has happened.
  Rename (n string)
}
