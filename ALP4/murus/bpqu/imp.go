package bpqu

// (c) Christian Maurer   v. 120909 - license see murus.go

import (
  "sync"
  . "murus/obj"
)
type
  Imp struct {
        heap []Any // heap [0] to save the type
    cap, num int
             }
var
  mutex sync.Mutex


func New (a Any, m uint) *Imp {
//
  Q:= new (Imp)
  Q.heap = make ([]Any, m)
  Q.heap[0] = Clone (a)
  Q.num = 0
  return Q
}


func (Q *Imp) Num () uint {
//
  return uint(Q.num)
}


func (Q *Imp) Full () bool {
//
  return Q.num == Q.cap
}


// lift heap [i] as far as necessary to restore the heap invariant heap [i] <= heap [j]
// for all i:= 1, ..., (Q.num - 1) / 2, j == 2 * i and j == 2 * i + 1
func (Q *Imp) lift () {
//
  mutex.Lock()
  i:= Q.num
  for {
    if i == 1 {
      break
    }
    j:= i / 2 // index above i
    if Less (Q.heap [j], Q.heap [i]) {
      break // i < Q.num, above i heap invariant is ok
    } else {
      Q.heap[i], Q.heap[j] = Q.heap[j], Q.heap[i]
    }
    i = j // continue above
  }
  mutex.Unlock()
}


func (Q *Imp) Ins (a Any) {
//
  if Q.num == Q.cap { return } // Q full
  if ! TypeEq (a, Q.heap[0]) { return } // a has wrong type
  Q.num++
  Q.heap [Q.num] = Clone (a)
  go Q.lift ()
}


// sift heap [1] as far as necessary to restore the heap invariant
func (Q *Imp) sift () {
//
  mutex.Lock()
  i:= 1
  for {
    if i > Q.num / 2 {
      break // nothing more under i
    }
    j:= 2 * i // left under i
    if j < Q.num && ! Less (Q.heap [j], Q.heap [j + 1]) {
      j++ // right under i
    }
    if Less (Q.heap [i], Q.heap [j]) {
      break
    } else {
      Q.heap[i], Q.heap[j] = Q.heap[j], Q.heap[i]
      i = j
    }
  }
  mutex.Unlock()
}


func (Q *Imp) Get () Any {
//
  if Q.num == 0 { return nil }
  return Q.heap [1]
}


func (Q *Imp) Del () Any {
//
  if Q.num == 0 { return nil }
  a:= Q.heap [1]
  Q.heap [1] = Q.heap[Q.num]
  Q.num--
  go Q.sift ()
  return Clone (a)
}
