package scr

// (c) Christian Maurer   v. 120909 - license see murus.go

import
  "sync"
var
  writing sync.Mutex


func lock () {
//
  writing.Lock ()
}


func unlock () {
//
  writing.Unlock ()
}
