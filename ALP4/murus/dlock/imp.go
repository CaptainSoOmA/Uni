package dlock

// (c) Christian Maurer   v. 121030 - license see murus.go

/* >>> Distributed mutual exclusion due to
       Ricart, G., Agrawala, A. K.:
         An Optimal Algorithm for Mutual Exclusion in Computer Networks
         Comm. ACM 24 (1981), 9-17, 581 */

import (
  "sync"
  "murus/host"; "murus/nchan"
)
const (
  zero = uint(0)
  ok = uint(1)
)
type
  Imp struct {
           n uint // number of hosts involved
        host []string // their names
          me uint
     request,
       reply [][]*nchan.Imp
       mutex,
    critSect sync.Mutex
        time,
     ownTime,
   noReplies uint
  requesting,
  terminated bool
    deferred []bool
             }
var (
  terminated bool
  done chan bool
  nFLocks uint
)


func New (H []string) *Imp {
//
  if len (H) < 2 { return nil }
  x:= new (Imp)
  x.n = uint(len (H))
  x.deferred = make ([]bool, x.n)
  x.host = make ([]string, x.n)
  for h:= zero; h < x.n; h++ {
    if h > 0 {
      for i:= zero; i < x.n; i++ {
        if H[h] == x.host[i] { return nil }
      }
    }
    x.host[h] = H[h]
    if host.Local (x.host[h]) {
      x.me = h
    }
  }
  x.critSect.Lock ()
  x.request = make ([][]*nchan.Imp, x.n)
  x.reply = make ([][]*nchan.Imp, x.n)
  for i:= zero; i < x.n; i++ {
    x.request[i] = make ([]*nchan.Imp, x.n)
    x.reply[i] = make ([]*nchan.Imp, x.n)
  }
  partner:= ""
  for h:= zero; h < x.n; h++ {
    for i:= zero; i < x.n; i++ {
      if h != i && (x.me == h || x.me == i) {
        if x.me == h {
          partner = x.host[i]
        } else { // me == i
          partner = x.host[h]
        }
        k:= uint16((nFLocks * x.n * x.n + h) * x.n + i)
        x.request[h][i] = nchan.New (x.ownTime, partner, k, true)
        x.reply[h][i] = nchan.New (x.ownTime, partner, k + uint16(x.n * x.n), true)
      }
    }
  }
  for h:= zero; h < x.n; h++ {
    if h != x.me {
      go func (i uint) { // bookkeeping of request
        for ! terminated {
          otherTime:= x.request[i][x.me].Recv ().(uint)
          x.mutex.Lock ()
          if otherTime > x.time {
            x.time = otherTime
          }
          if x.requesting && (x.ownTime < otherTime || (x.ownTime == otherTime && x.me < i)) {
            x.deferred[i] = true
          } else {
            x.reply[x.me][i].Send (ok)
          }
          x.mutex.Unlock ()
        }
        done <- true
      }(h)
      go func (i uint) { // bookkeeping of ok-replies
        for ! terminated {
          _ = x.reply[i][x.me].Recv ().(uint)
          x.mutex.Lock ()
          x.noReplies++
          if x.noReplies == x.n - 1 {
            x.critSect.Unlock ()
          }
          x.mutex.Unlock ()
        }
        done <- true
      }(h)
    }
  }
  nFLocks++
  return x
}


func (x *Imp) Terminate () {
//
  for h:= zero; h < x.n; h++ {
    if h != x.me {
      terminated = true // stop the bookkeeping goroutines
      for i:= zero; i < x.n * x.n; i++ {
        <-done
      }
    }
  }
  for h:= zero; h < x.n; h++ {
    for i:= zero; i < x.n; i++ {
      if h != i && (x.me == h || x.me == i) {
        x.request[h][i].Terminate()
        x.reply[h][i].Terminate()
      }
    }
  }
}


func (x *Imp) Lock () {
//
  x.mutex.Lock ()
  x.requesting = true
  x.ownTime = x.time + 1
  x.mutex.Unlock ()
  x.noReplies = 0
  for h:= zero; h < x.n; h++ {
    if h != x.me {
      x.request[x.me][h].Send (x.ownTime)
    }
  }
  x.critSect.Lock ()
}


func (x *Imp) Unlock () {
//
  x.mutex.Lock ()
  x.requesting = false
  x.mutex.Unlock ()
  for h:= zero; h < x.n; h++ {
    if x.deferred[h] {
      x.deferred[h] = false
      x.reply[x.me][h].Send (ok)
    }
  }
}
