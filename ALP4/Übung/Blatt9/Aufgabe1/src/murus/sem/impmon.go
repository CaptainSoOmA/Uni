package sem

//     Christian Maurer   v. 130521
// (c) Go-Authors

import
  "sync"
type
  ImpMon struct {
          count int
           lock *sync.Mutex
         wakeup *sync.Cond
                }


func NewMon (n int) *ImpMon {
//
  x:= new (ImpMon)
  x.count = n
  x.lock = new (sync.Mutex)
  x.wakeup = sync.NewCond (x.lock)
  return x
}


func (s *ImpMon) P() {
//
  s.Wait (1)
}


func (s *ImpMon) Wait (units uint) {
//
  s.lock.Lock()
  s.count -= int(units)
  for s.count < 0 {
    s.wakeup.Wait()
  }
  s.lock.Unlock()
}


func (s *ImpMon) V() {
//
  s.Signal (1)
}


func (s *ImpMon) Signal (units uint) {
//
  s.lock.Lock()
  wakeOthers:= s.count < 0
  s.count += int(units)
  if wakeOthers {
    s.wakeup.Signal()
  }
  s.lock.Unlock()
}
