package main

// (c) Christian Maurer   v. 130308 - license see murus.go

import
  "murus/sub/netz"


func main () {
//
  for netz.StartUndZielGewählt () {
    netz.KürzestenWegSuchen ()
  }
}
