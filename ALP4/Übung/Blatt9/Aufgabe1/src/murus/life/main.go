package main

// (c) Christian Maurer   v. 130526 - license see murus.go

//     cp *.life ~/.life

// >>> highly recommendable:
//     http://www.lcc.uma.es/~fjv/UMA/LCC/web/Teaching/brainfood/life_lexicon/lex.htm

import (
  "murus/str"
  "murus/col"; "murus/box"; "murus/errh"
  . "murus/menue"
  "murus/life/species"; "murus/life/world"
)


func defined () (string, bool) {
//
  bx:= box.New()
  bx.Wd (6)
  bx.Colours (col.LightCyan, col.Black)
  bx.Write ("Welt:", 1, 0)
  const n = 12
  bx.Wd (n)
  name:= str.Clr (n)
  name = ""
  errh.Hint ("Namen der Welt eingeben                Programmende: leere Eingabe")
  bx.Edit (&name, 1, 6)
  str.RemSpaces (&name)
  errh.DelHint ()
  return name, ! str.Empty (name)
}


func sim () {
//
  w:= world.New ()
  for {
    if name, ok:= defined (); ok {
      w.Name (name)
      w.Write (0, 0)
      w.Edit (0, 0)
      w.Clr ()
    } else {
      break
    }
  }
}


func main () {
//
  x:= New ("Spiel des Lebens                v. 120722")

  game:= New ("Game of Life (John Conway)")
  game.Leaf (func() { world.Sys (species.Life); sim() }, true)
  x.Ins (game)

  ecosys:= New ("Einfaches Modell eines Ökosystems aus Füchsen, Hasen und Pflanzen")
  ecosys.Leaf (func() { world.Sys (species.Eco); sim() }, true)
  x.Ins (ecosys)

  x.Exec ()
}
