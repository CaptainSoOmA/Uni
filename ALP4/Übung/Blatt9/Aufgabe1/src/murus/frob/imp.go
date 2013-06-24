package frob

// (c) Christian Maurer   v. 130131 - license see murus.go

import (
  . "murus/obj"; "murus/kbd"
  "murus/fmon"; "murus/rob"
)
const (
  Anbieter = "venus" // <<< Name des Servers
  clu = 4 // uint
  clb = 1 // bool
)
const (
  start = iota; term
  amanfang; links; rechts; amrand; weiter; zurueck; keinklotz;
  klotzda; klotzhin; klotzweg; klotzweiter; klotzschuss;
  markeda; markehin; markeweg; vormauer; mauerhin; mauerweg; manip
  AnzahlOperationen
)
type
  botschaft struct {
              Robi *rob.Imp
          Robiwelt *rob.Welt
                 c kbd.Comm
                 d uint
              wert bool
                   }
var (
  Monitor *fmon.Imp
  Robi *rob.Imp
  b *botschaft
)


func f (a Any, i uint) Any {
//
  b:= a.(botschaft)
  if i != start {
    Robi = b.Robi
  }
  switch i { case start:
    Robi = rob.NeuerRoboter()
  case term:
    Robi.Terminieren()
  case amanfang:
    b.wert = Robi.InLinkerObererEcke()
  case links:
    Robi.LinksDrehen ()
  case rechts:
    Robi.RechtsDrehen ()
  case amrand:
    b.wert = Robi.AmRand()
  case weiter:
    b.wert = Robi.Gelaufen1()
  case zurueck:
    b.wert = Robi.Zurueckgelaufen1()
  case keinklotz:
    b.wert = Robi.Leer()
  case klotzda:
    b.wert = Robi.HatKloetze()
  case klotzhin:
    Robi.Legen1()
  case klotzweg:
    Robi.Leeren1()
  case klotzweiter:
    b.wert = Robi.Geschoben1()
  case klotzschuss:
    Robi.Schiessen()
  case markeda:
    b.wert = Robi.Markiert()
  case markehin:
    Robi.Markieren()
  case markeweg:
    Robi.Entmarkieren()
  case vormauer:
    b.wert = Robi.VorMauer()
  case mauerhin:
    b.wert = Robi.Gemauert1()
  case mauerweg:
    b.wert = Robi.Entmauert1()
  case manip:
    Robi.Manipulieren (b.c, b.d)
  }
  b.Robiwelt = rob.Roboterwelt
  return b.wert
}


// Vor.: Robi ist auf dem Anbieter und dem lokalen Rechner initialisiert.
// Robi ist vom Anbieter mit Op bearbeitet. Die lokale Robiwelt
// hat den Zustand der Welt auf dem Anbieter unmittelbar nach
// der Bearbeitung von Op und ist auf dem scr ausgegeben.
// Robi ist im aus der Bearbeitung von Op resultierenden Zustand.
func F (i uint) Any {
//
  b.Robi = Robi
  b = Monitor.F (b, i).(*botschaft)
  b.Robiwelt.Ausgeben()
  Robi = b.Robi
  return b.wert
}


func Initialisieren () {
//
  Robi = rob.NeuerRoboter ()
  F (start)
}


func Terminieren () {
//
  F (term)
  Robi.Terminieren ()
}


func InLinkerObererEcke () bool {
//
  return F (amanfang).(bool)
}


func LinksDrehen () {
//
  F (links)
}


func RechtsDrehen () {
//
  F (rechts)
}


func AmRand () bool {
//
  return F (amrand).(bool)
}


func Gelaufen1 () bool {
//
  return F (weiter).(bool)
}


func Zurueckgelaufen1 () bool {
//
  return F (zurueck).(bool)
}


func Leer () bool {
//
  return F (keinklotz).(bool)
}


func HatKloetze () bool {
//
  return F (klotzda).(bool)
}


func Legen1 () {
//
  F (klotzhin)
}


func Leeren1 () {
//
  F (klotzweg)
}


func Geschoben1 () bool {
//
  return F (klotzweiter).(bool)
}


func Schiessen () {
//
  F (klotzschuss)
}


func Markiert () bool {
//
  return F (markeda).(bool)
}


func Markieren () {
//
  F (markehin)
}


func Entmarkieren () {
//
  F (markeweg)
}


func VorMauer () bool {
//
  return F (vormauer).(bool)
}


func Gemauert1 () bool {
//
  return F (mauerhin).(bool)
}


func Entmauert1 () bool {
//
  return F (mauerweg).(bool)
}


func Editieren () {
//
  for {
    b.c, b.d = kbd.Command ()
    switch b.c { case kbd.Esc:
      break
    case kbd.Enter,
         kbd.Left, kbd.Right, kbd.Up, kbd.Down, kbd.Pos1, kbd.End, kbd.Tab, kbd.Del, kbd.Ins, kbd.Help,
         kbd.Mark, kbd.Demark: // kbd.PrintScr:
      F (manip)
    }
  }
  Terminieren()
}


func terminieren () {
//
  Monitor.Terminate ()
}


func init () {
//
  b = new (botschaft)
  Monitor = fmon.New (b, AnzahlOperationen, f, TrueSp, Anbieter, 90)
  Monitor.Prepare (func () { rob.WeltDefinieren() })
//  TerminierungInstallieren (terminieren)
  Initialisieren() // läuft nicht mehr auf dem Anbieter !
}
