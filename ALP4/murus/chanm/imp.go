package chanm

// (c) Christian Maurer   v. 120909 - license see murus.go

import (
  "sync"
  . "murus/obj"
)
type
  Imp struct {
     content Any
       first bool
       mutex,
  rendezvous sync.Mutex
             }


func New (a Any) *Imp {
  C:= new (Imp)
  C.content = nil
  C.first = true
  C.rendezvous.Lock()
  return C
}


func (C *Imp) Send (a Any) {
//
  C.mutex.Lock()
  C.content = Clone(a)
  if C.first { // sender first at rendezvous
    C.first = false
    C.mutex.Unlock()
    C.rendezvous.Lock()
    C.mutex.Unlock()
  } else { // receiver first at rendezvous
    C.first = true
    C.rendezvous.Unlock()
  }
}


func (C *Imp) Empty () bool {
//
  return C.content == nil
}


func (C *Imp) Recv (a Any) {
//
  C.mutex.Lock()
  if ! C.first { // sender first at rendezvous
    C.first = true
    b:= Clone(C.content); a = &b
    C.content = nil
    C.rendezvous.Unlock()
  } else { // receiver first at rendezvous
    C.first = false
    C.mutex.Unlock()
    C.rendezvous.Lock()
    b:= Clone(C.content); a = &b
    C.content = nil
    C.mutex.Unlock()
  }
}
