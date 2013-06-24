package world

// (c) Christian Maurer   v. 130308 - license see murus.go

import (
  . "murus/obj"
  "murus/life/species"
)

func Sys (s species.System) { sys(s) }

type World interface {

  Editor
  Persistor
}
