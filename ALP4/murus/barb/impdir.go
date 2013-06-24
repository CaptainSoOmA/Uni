package barb

// (c) Christian Maurer   v. 130424 - license see murus.go

import
  "sync"
type
  ImpDir struct {
        waiting,
          mutex sync.Mutex
              n int
                }


func NewDir () *ImpDir {
//
  x:= new (ImpDir)
  x.waiting.Lock()
  return x
}


func (x *ImpDir) Customer () {
//
  x.mutex.Lock()
  x.n ++
  if x.n == 0 {
//    x.mutex.Unlock() // *
    x.waiting.Unlock()
  } else {
    x.mutex.Unlock()
  }
}


func (x *ImpDir) Barber () {
//
  x.mutex.Lock()
  x.n --
  if x.n == -1 {
    x.mutex.Unlock()
    x.waiting.Lock()
//    x.mutex.Lock() // *
  }
  x.mutex.Unlock()
}
