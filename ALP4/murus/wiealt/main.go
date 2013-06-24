package main

// (c) Christian Maurer   v. 130526 - license see murus.go

import (
  "murus/ker"
  "murus/kbd"
  "murus/col"; "murus/scr"; "murus/errh"
  "murus/prt"
  . "murus/day"
)

func main () {
//
  scr.Switch (scr.TXT)
  var today, birthday *Imp
  today = New()
  today.Actualize ()
  birthday = New()
  birthday.SetColours (col.LightWhite, col.Blue)
  scr.Colour (col.Yellow)
  scr.Write ("Ihr Geburtsdatum:", 12, 0)
  birthday.Edit (12, 18)
  if birthday.Empty() {
    birthday.Actualize()
  } else {
    scr.Write (" war ein ", 12, 26)
    birthday.SetFormat (WD)
    birthday.Write (12, 35)
    errh.Error2 ("Sie sind heute", birthday.Distance (today), "Tage alt.", 0)
  }
  scr.Colours (col.ScreenF, col.ScreenB)
  scr.Cls()
  errh.Hint (" vor-/rückwärts: Pfeiltasten               fertig: Esc ")
  var ( c kbd.Comm; t uint )
  neu:= true
  loop: for {
    if neu {
      birthday.WriteYear (0, 0)
      neu = false
    }
    switch c, t = kbd.Command (); c { case kbd.Esc:
      break loop
    case kbd.Down:
      if t == 0 {
        birthday.Inc (Yearly)
      } else {
        birthday.Inc (Decadic)
      }
      neu = true
    case kbd.Up:
      if t == 0 {
        birthday.Dec (Yearly)
      } else {
        birthday.Dec (Decadic)
      }
      neu = true
    case kbd.PrintScr:
      birthday.PrintYear (0, 0)
      prt.GoPrint()
    }
  }
  ker.Terminate ()
}
