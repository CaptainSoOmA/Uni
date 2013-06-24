package rob1

// (c) Christian Maurer   v. 130526 - license see murus.go

import (
  "murus/ker"; "murus/env"; "murus/obj";
  "murus/col"; "murus/scr"; "murus/errh"
  "murus/pseq"; "murus/day"
  "murus/rob"
)
type
  Imp struct {
             *rob.Imp
             }
var (
  Robo *Imp
  amEditieren bool
)


func stop (a rob.Aktion) {
//
  if ! amEditieren {
    errh.Error2 ("Programmstop: Verletzung der Vor. in Methode", 0, "rob." + rob.Aktionstext [a] + "()", 0)
    ker.Terminate ()
  }
}


func LinksDrehen () {
//
  Robo.LinksDrehen ()
}


func RechtsDrehen () {
//
  Robo.RechtsDrehen ()
}


func AmRand () bool {
//
  return Robo.AmRand ()
}


func Laufen1 () {
//
  if Robo.Gelaufen1 () {
  } else {
    stop (rob.Weiter)
  }
}


func Zuruecklaufen1 () {
//
  if Robo.Zurueckgelaufen1 () {
  } else {
    stop (rob.Zurueck)
  }
}


func Leer () bool {
//
  return Robo.Leer ()
}


func Leeren1 () {
//
  Robo.Leeren1 ()
}


func NachbarLeer () bool {
//
  return Robo.NachbarLeer ()
}


func HatKloetze () bool {
//
  return Robo.HatKloetze ()
}


func KloetzeGeben (n uint) {
//
  Robo.KloetzeGeben (n)
}


func AnzahlKloetze () uint {
//
  return Robo.AnzahlKloetze ()
}


func Legen1 () {
//
  Robo.Legen1 ()
}


func Schieben1 () {
//
  if Robo.Geschoben1 () {
  } else {
    stop (rob.KlotzWeiter)
  }
}


func Schiessen () {
//
  Robo.Schiessen ()
}


func Markiert () bool {
//
  return Robo.Markiert ()
}


func Markieren () {
//
  Robo.Markieren ()
}


func Entmarkieren () {
//
  Robo.Entmarkieren ()
}


func NachbarMarkiert () bool {
//
  return Robo.NachbarMarkiert ()
}


func VorMauer () bool {
//
  return Robo.VorMauer ()
}


func Mauern1 () {
//
  if Robo.Gemauert1 () {
  } else {
    stop (rob.MauerHin)
  }
}


func Entmauern1 () {
//
  if Robo.Entmauert1 () {

  } else {
    stop (rob.MauerWeg)
  }
}


var
  programmdatei *pseq.Imp


func prog (Zeile string) {
//
  for i:= 0; i < len (Zeile); i++ {
    programmdatei.Ins (byte(Zeile[i]))
  }
  programmdatei.Ins (byte(10))
}


const
  LF = string(byte(10))


func (R *Imp) programmErzeugen () {
//
  programmdatei = pseq.New (byte (0))
  heute:= day.New()
  heute.SetFormat (day.Dd)
  heute.Actualize()
  programmdatei.Name (env.User () + "-" + heute.String() + ".go") // TODO aktuelle Zeit ?
  programmdatei.Clr()
  scr.Cls()
  scr.Colours (col.Yellow, col.DarkBlue)
  prog ("package main" + LF + LF + "import . \"murus/robo\"" + LF)
  prog ("func main () {" + LF + "//")
  Robo.Trav (func (a obj.Any) { prog ("  " + rob.Aktionstext[a.(rob.Aktion)] + "()") })
  prog ("  Fertig()" + LF + "}")
  programmdatei.Terminate ()
}


func InLinkerObererEcke () bool {
//
  return Robo.InLinkerObererEcke ()
}


var
  sollProtokolliertWerden bool


func Editieren () {
//
//  rob.RoboterweltDefinieren()
  amEditieren = true
//  errh.WriteHelp1 ()
  Robo.Editieren ()
//  errh.DelHint ()
  amEditieren = false
  if sollProtokolliertWerden {
    Robo.programmErzeugen ()
//    rob.Roboterwelt.Clr ()
  }
  if sollProtokolliertWerden {
//    ZurueckInLinkeObereEcke() // TODO
  }
  Robo.Terminieren ()
  ker.Terminate ()
}


func ProtokollSchalten (ein bool) {
//
  sollProtokolliertWerden = ein
  scr.Colours (col.ErrorF, col.ErrorB)
  if sollProtokolliertWerden {
//    rob.aktionsfolge = rob.aktionsfolge [0:0]
    scr.Write ("Protokoll eingeschaltet", 0, scr.NColumns () - 23)
  } else {
    scr.Write ("                       ", 0, scr.NColumns () - 23)
  }
}


func SokobanSchalten (ein bool) {
//
  Robo.SokobanSchalten (ein)
}


func Ausgeben (n uint) {
//
  Robo.Ausgeben (n)
}


func Eingabe () uint {
//
  return Robo.Eingabe ()
}


func FehlerMelden (s string, n uint) {
//
  Robo.FehlerMelden (s, n)
}


func Fertig () {
//
  Robo.FehlerMelden ("Programm beendet", 0)
  ker.Terminate()
}


func init () {
//
  Robo = &Imp { rob.NeuerRoboter() }
//  Roboterwelt = NeueWelt()
  rob.WeltDefinieren()
  Robo.Aktualisieren()
}
