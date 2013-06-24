package main

// (c) Christian Maurer   v. 130526 - license see murus.go

/* Simple example of the "MVC"-paradigm, the separation of
   - data and algorithms of a problem ("M(odel"),
   - its representation on the screen ("V(iew") and
   - controlling the program ("C(ontroller").
   This package is the "C(ontroller".

   >>> s. Murus-Dokumentation, Seite 23 ff.

   Default device is "cdrom". If you use another device as CD-player,
   you have to call the program with its name (in the directory /dev)
   as parameter (e.g. "cdplay dvd"). */

import (
  "murus/ker"; "murus/rand" // ; "murus/str"
//  "murus/space"; "murus/mouse3d"
  "murus/kbd"; "murus/col"; "murus/scr"; "murus/errh"
  model "murus/cdker"; view "murus/cdio"
)
const (
  program = "cdplay"
  version = "130224"
  author = "Christian Maurer"
)

func main() {
//
/*
  h:= [...]string {
    "     <              Rücktaste <--                      ",
    " <<    >>           Pfeiltaste auf-/abwärts            ",
    "|<<    >>|          Pos1-Taste/Endetaste               ",
    "    2 .. 12         F2 .. F12-Taste                    ",
    "   13 .. 24         Umschalt- + F1 .. F12-Taste        ",
    "               oder Mausklick auf entsprechende Zeile  ",
    " Zufallsstück       Rollen-Taste                       ",
    "                                                       ",
    "  <  1 sek >        Bildtaste auf-/abwärts             ",
    " << 10 sek >>       Umschalt- + Bildtaste auf-/abwärts ",
    "<<<  1 min >>>      Alt- + Bildtaste auf-/abwärts      ",
    "                                                       ",
    "  Lautstärke        Pfeiltaste Links/Rechts und        ",
    "                    Umschalt- + Pfeiltaste Links/Rechts",
    "               oder mit Maus auf entsprechende Anzeige ",
    "   Balance          Alt- + Pfeiltaste Links/Rechts     ",
    "               oder mit Maus auf entsprechende Anzeige ",
    "                                                       ",
    " Pause/weiter       Pause-Taste                        ",
    " Ende/CD raus       Abbruch-/Entfernungstaste          " }
  Help:= make ([]string, len(h))
  for i:= 0; i < len(h); i++ { str.Set (&Help[i], h[i]) }
*/
  errh.WriteHeadline (program, version, author, col.LightWhite, col.DarkBlue)
  f:= model.Soundfile ()
  if f == nil {
    model.Term1 ()
    errh.Error ("kein Zugriff auf CD", 0)
    ker.Terminate(); return
  }
  view.WriteMask ()
  var (help, stop bool)
  quit:= make (chan bool)
  go func () { for ! stop { ker.Sleep (1); if ! help { view.Write() } }; quit <- true } ()
  model.PlayTrack (0)
  loop: for {
    c, t:= kbd.Command ()
    help = false
    scr.MouseCursor (true)
    switch c { case kbd.Esc:
      if t == 0 {
        model.Term()
      } else {
        model.Term1()
      }
      stop = true
      break loop
    case kbd.Enter:
      if t == 0 {
        model.PlayTrack1 (true)
      } else {
        model.PlayTrack (model.NTracks() - 1)
      }
    case kbd.Back:
      switch t { case 0:
        model.PlayTrack0 ()
      case 1:
        model.PlayTrack1 (false)
      default:
        model.PlayTrack (0)
      }
    case kbd.Left, kbd.Right:
      switch t { case 0:
        model.Ctrl1 (model.All, c == kbd.Right)
      case 1:
        _ = uint(model.Volume (model.All))
        for j:= 0; j < 8; j++ { model.Ctrl1 (model.All, c == kbd.Right) }
      default:
        model.Ctrl1 (model.Balance, c == kbd.Left)
      }
    case kbd.Pos1:
      model.PlayTrack (0)
    case kbd.End:
      model.PlayTrack (model.NTracks() - 1)
    case kbd.Up, kbd.Down:
      switch t { case 0, 1:
        model.PlayTrack1 (c == kbd.Down)
      case 2:
        model.PosTime1 (c == kbd.Down, 1)
      case 3:
        model.PosTime1 (c == kbd.Down, 10)
      default:
        model.PosTime1 (c == kbd.Down, 60)
      }
    case kbd.Tab:
      model.PlayTrack1 (t == 0)
    case kbd.Del:
      model.Term1()
      return
    case kbd.Help, kbd.LookFor, kbd.Act, kbd.Cfg,
         kbd.Mark, kbd.Demark, kbd.Deposit, kbd.Paste,
         kbd.Black, kbd.Red, kbd.Green, kbd.Blue:
      if t == 0 {
        if c == kbd.Help {
//          help = true
//          errh.WriteHelp (Help)
        } else {
          model.PlayTrack (uint8(c - kbd.Help))
        }
      } else {
        model.PlayTrack (10 + uint8(c - kbd.Help))
      }
    case kbd.Roll:
      model.PlayTrack (uint8(rand.Natural (uint(model.NTracks()))))
    case kbd.Pause:
      model.Switch()
    case kbd.Here, kbd.Pull:
      var ctrl model.Controller
      if track, b:= view.TrackUnderMouse (); b {
        if c == kbd.Here {
          model.PlayTrack (track)
        }
      } else if ls, c:= view.ControlUnderMouse (&ctrl); c {
        model.Ctrl (ctrl, uint8(ls))
      } else if b, sek, c:= view.TimeUnderMouse (); c {
        model.PosTime (b, sek)
      }
    case kbd.Navigate:
/*
      var mov, rot space.Gridcoordinate
      &mov, &rot = mouse3d.Read()
      model.Ctrl1 (model.All, mov [space.Top] < 0)
      model.Ctrl1 (model.All, rot [space.Right] < 0)
      model.Ctrl1 (model.Balance, rot [space.Front] < 0)
*/
    }
  }
  <-quit
  _ = f.Close()
  ker.Terminate()
}
