package ker

// (c) Christian Maurer   v. 130127 - license see murus.go

import (
  "os"
  "strconv"
)
var (
  handler = []func(){}
  terminated bool
)


func Terminate () {
//
  if terminated { return }
  for _, h:= range (handler) {
    h ()
  }
  terminated = true
}


func Panic (s string) {
//
  Terminate ()
  panic (s)
}


func Stop (p string, n uint) {
//
  Terminate ()
  panic ("Programm wegen Fehler Nr. " + strconv.FormatUint (uint64(n), 10) + " im Paket " + p + " abgebrochen")
}


func Halt (s int) {
//
  Terminate ()
  os.Exit (s)
}


func InstallTerm (h func()) {
//
  handler = append (handler, h)
}


// func init () { installTerm (Terminate) } // does not work: attempt to link returns "atexit not defined"
