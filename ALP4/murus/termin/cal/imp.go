package cal

// (c) Christian Maurer   v. 130526 - license see murus.go

import (
  "murus/ker"; . "murus/obj"; "murus/kbd"
  "murus/errh"; "murus/day"; "murus/piset"
  "murus/termin/word"; "murus/termin/page"
)
const
  pack = "cal"
var (
  content *piset.Imp
  globalPage *page.Imp = page.New ()
  globalDay *day.Imp = day.New ()
)


func SetFormat (p day.Period) {
//
  globalPage.SetFormat (p)
}


func Seek (d *day.Imp) {
//
  globalPage.Set (d)
  if content.Ex (globalPage) { // richtige Seite gefunden
    globalPage = content.Get ().(*page.Imp)
  } else {
    globalPage.Clr ()
  }
}


func WriteDay (Z, S uint) {
//
  globalPage.Write (Z, S)
}


func ClearDay (d *day.Imp, Z, S uint) {
//
  globalPage.Set (d)
  globalPage.Write (Z, S)
}


func Edit (d *day.Imp, Z, S uint) {
//
  globalPage.Set (d)
  globalPage.SetFormat (day.Daily)
  exists:= content.Ex (globalPage)
  if exists { // haben wir an diesem Tag Termine
    errh.Hint (errh.ToSelect)
    loop: for {
      globalPage = content.Get ().(*page.Imp)
      globalPage.Write (Z, S)
      K, _:= kbd.Command ()
      switch K { case kbd.Enter:
        break loop
      case kbd.Esc:
        errh.DelHint()
        return
      case kbd.Down, kbd.Up:
        content.Step (K == kbd.Down)
      case kbd.Pos1, kbd.End:
        content.Jump (K == kbd.End)
      }
    }
    errh.DelHint()
    globalDay = globalPage.Day ().(*day.Imp)
  }
  globalPage.Edit (Z, S)
  if globalPage.Empty () {
    if exists {
      content.Del ()
    }
  } else if exists {
    content.Put (globalPage)
  } else {
    content.Ins (globalPage)
  }
}


func EditWord (Z, S uint) {
//
  word.EditActual (Z, S)
}


func Print (Z, S uint) {
//
  globalPage.Print (Z, S)
}


func LookFor (b Op) {
//
  content.Trav (func (a Any) { if a.(*page.Imp).HasWord () { b (globalPage.Day ()) } })
}


func Index (X Object) Object {
//
  x, ok:= X.(*page.Imp)
  if ! ok { ker.Stop (pack, 1) }
  return Clone (x.Day()).(*day.Imp)
}


func Terminate () {
//
  globalPage.Terminate ()
}


func init () {
//
  content = piset.New (globalPage, Index)
  content.Name ("Termine")
}
