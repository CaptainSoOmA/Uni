package perm

// (c) Christian Maurer   v. 121114 - license see murus.go

import
  "murus/rand"
type
  Imp struct {
 size uint
    p []uint
      }


func New (n uint) *Imp {
//
  if n == 0 { return nil }
  x:= new (Imp)
  x.size = n
  if x.size > 1 {
    x.p = make ([]uint, x.size)
    for i:= uint(0); i < x.size; i++ {
      x.p[i] = i
    }
  }
  x.Permute ()
  return x
}


func (x *Imp) Permute () {
//
  switch x.size { case 1:
    return
  case 2:
    if rand.Natural (rand.Natural (1000)) % 2 == 1 {
      x.p[0], x.p[1] = x.p[1], x.p[0]
    }
  default:
    for i:= uint(0); i < 3 * x.size + rand.Natural (x.size); i++ {
      j, k:= rand.Natural (x.size), rand.Natural (x.size)
      if j != k {
        x.p[j], x.p[k] = x.p[k], x.p[j]
      }
    }
  }
}


func (x *Imp) F (i uint) uint {
//
  if x.size == 1 || x.size <= i {
    return 0
  }
  return x.p[i]
}
