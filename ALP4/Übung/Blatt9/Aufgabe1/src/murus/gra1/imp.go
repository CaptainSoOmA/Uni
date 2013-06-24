package gra1

// (c) Christian Maurer   v. 130402 - license see murus.go

import (
//  "murus/ker"
  . "murus/obj"
  "murus/str"
  "murus/kbd"
  "murus/col"; "murus/scr"; "murus/errh"
//  "murus/pseq"
//  "murus/img"
  "murus/gra"
  "murus/gra1/node"; "murus/gra1/edge"
)
const (
  pack = "gra1"
  suffix = "gra"
)
var (
  Graph *gra.Imp
  initialisiert bool
  Ecke, Ecke1, tempEcke, tempEcke1 *node.Imp
  Kante, tempKante *edge.Imp
  help []string
)


func Init (e, d bool, n uint, s string) {
//
  if initialisiert {
    Graph.Terminate ()
  } else {
    scr.Switch (scr.WVGApp)
    initialisiert = true
  }
  Ecke = node.New (n)
  Ecke1 = node.New (n)
  tempEcke = node.New (n)
  tempEcke1 = node.New (n)
  Kante = edge.New ()
  tempKante = edge.New ()
  Graph = gra.New (d, Ecke, Kante)
  edge.WithValues = e
  Dateiname:= ""
  str.RemAllSpaces (&s)
  if str.Empty (s) {
    Dateiname = "temp"
  } else {
    Dateiname = s
  }
//  sel.DateinameEditieren (Dateiname)
  Graph.Name (Dateiname + "." + "gra")
}


func SetColours (n, a col.Colour) {
//
  tempEcke.SetColours (n, a)
  tempKante.SetColours (n, a)
}


func ausgewaehlt (Ecke *node.Imp) bool {
//
  loop: for {
    c, _:= kbd.Command ()
    switch c { case kbd.Here:
      if Graph.ExPred (node.UnderMouse) {
        Ecke = Graph.Get ().(*node.Imp)
        return true
      }
    case kbd.Esc:
      break loop
    }
  }
  return false
}


func NodesSelected () bool {
//
  scr.MouseCursor (true)
  g:= false
  for {
    errh.Hint ("Start auswählen")
    if ausgewaehlt (Ecke) { // Ecke aktuell
      Graph.Position (true) // Ecke postaktuell
      Write()
      errh.Hint ("Ziel auswählen")
      if ausgewaehlt (Ecke1) { // Ecke1 aktuell
        errh.DelHint()
        if Graph.Positioned () {
          errh.Error ("Fehler: Start und Ziel sind gleich !", 0)
        } else {
          g = true
          break
        }
      }
    } else {
      break
    }
  }
  Write()
  errh.DelHint()
  scr.MouseCursor (false)
  return g
}


func NodeSelected () bool {
//
  scr.MouseCursor (true)
  errh.Hint ("Ecke auswählen")
  g:= false
  if ausgewaehlt (Ecke) { // Ecke aktuell
    Graph.Position (true) // Ecke postaktuell
    g = true
    Write()
  }
  errh.DelHint()
  scr.MouseCursor (false)
  return g
}


func write (n Any, a bool) {
//
  n.(*node.Imp).WriteCond (a)
}


func write3 (e, n, n1 Any, a bool) {
//
  e.(*edge.Imp).WriteCond (n.(*node.Imp), n1.(*node.Imp), Graph.Directed(), a)
}


func Write () {
//
  scr.Buf (true)
  Graph.Trav3Cond (write3)
  Graph.TravCond (write)
  scr.Buf (false)
}


func wr (vis, inv bool) {
//
  g:= Graph.Directed ()
  tempEcke = Graph.Get ().(*node.Imp) // tempEcke aktuell
  tempEcke.Write (vis, inv)
  Graph.Position (true) // tempEcke auch postaktuell
  k:= Graph.NumLocal ()
  if k > 0 {
    for i:= uint(0); i < k; i++ {
      Graph.Step (i, true)
      n, n1, e:= Graph.Get3 () // tempEcke aktuell
      tempEcke, tempEcke1, tempKante = n.(*node.Imp), n1.(*node.Imp), e.(*edge.Imp)
      tempKante.Write (tempEcke, tempEcke1, g, vis, inv)
      Graph.Step (i, false)
    }
  }
  if ! g { return }
  Graph.InvAct ()
  k = Graph.NumLocal ()
  if k > 0 {
    for i:= uint(0); i < k - 1; i++ {
      Graph.Step (i, true)
      n, n1, e:= Graph.Get3 () // tempEcke aktuell
      tempEcke, tempEcke1, tempKante = n.(*node.Imp), n1.(*node.Imp), e.(*edge.Imp)
      tempKante.Write (tempEcke, tempEcke1, true, vis, inv) // !
      Graph.Step (i, false)
    }
  }
  Graph.InvAct ()
}


func Edit () {
//
  Write ()
  errh.Hint ("Graph editieren: Hilfe per F1, fertig: Abbruchtaste (Esc)")
  scr.MouseCursor (true)
  loop: for {
    K, i:= kbd.Command ()
//    errh.DelHint()
    switch K { case kbd.Esc:
      break loop
    case kbd.Help:
      errh.WriteHelp (help)
    case kbd.Here: // neue Ecke oder Namen vorhandener Ecke ändern:
      if Graph.ExPred (node.UnderMouse) {
        if i > 0 {
          Ecke = Graph.Get ().(*node.Imp) // aktuell: Ecke
          Ecke.Edit ()
          Graph.Put (Ecke)
        }
      } else {
        Ecke.Clr ()
        Ecke.Locate ()
        Ecke.Write (true, true)
        Ecke.Edit ()
        Graph.Ins (Ecke)
      }
    case kbd.Del: // Ecke entfernen
      if Graph.ExPred (node.UnderMouse) {
        wr (false, false)
        Graph.Del ()
      }
    case kbd.There: // Ecke verschieben
      switch i { case 0:
        if Graph.ExPred (node.UnderMouse) {
          wr (false, false)
          wr (false, true)
          loop1: for {
            geschoben:= false
            kk, _:= kbd.Command ()
            switch kk { case kbd.Push:
              geschoben = true
              wr (false, true)
              Ecke = Graph.Get ().(*node.Imp)
              Ecke.Locate ()
              Graph.Put (Ecke)
              wr (false, true)
            case kbd.Thither:
              wr (false, true)
              if geschoben {
                Graph.Put (Ecke)
              }
              wr (true, false)
              break loop1
            }
          }
        }
      default: // Ecke entfernen
        if Graph.ExPred (node.UnderMouse) {
          wr (false, false)
          Graph.Del ()
        }
      }
    case kbd.This: // Ecken verbinden / Kante entfernen:
      x0, y0:= scr.MousePosGr ()
      x, y:= x0, y0
      if Graph.ExPred (node.UnderMouse) {
        Ecke = Graph.Get ().(*node.Imp) // Ecke aktuell
        Graph.Position (true) // Ecke auch postaktuell
        loop2: for {
          kk, _:= kbd.Command()
          switch kk {
          case kbd.Move:
            scr.LineInv (x0, y0, x, y)
            x, y = scr.MousePosGr ()
            scr.LineInv (x0, y0, x, y)
          case kbd.Thus:
            scr.LineInv (x0, y0, x, y)
            if Graph.ExPred (node.UnderMouse) {
              Ecke1 = Graph.Get ().(*node.Imp) // Ecke1 aktuell
              g:= Graph.Directed ()
              if g {
                Kante.Write (Ecke, Ecke1, g, false, false)
              }
              Kante.Edit (Ecke, Ecke1, g)
              Graph.Edge1 (Kante)
              if Kante.Val () == 0 {
                Kante.Write (Ecke, Ecke1, g, false, false)
                Graph.Del1 ()
              }
            }
            break loop2
          }
        }
      }
    case kbd.PrintScr:
//      img.Write (".tmp.Graph", 0, 0, scr.NX(), scr.NY())
//      img.Print()
    }
  }
  errh.DelHint()
}


func Done () bool {
//
  return Graph.NumAct() == Graph.Num ()
/* >>>>>            falsch ^^^^^^^^^; richtig wäre: Anzahl der
   Ecken der Zusammenhangskomponente, in der die aktuelle Ecke liegt */
}


func Terminate () {
//
  if initialisiert {
    Graph.Terminate ()
  }
}


func Set (d gra.Demo) {
//
  Graph.Install (write, write3)
  Graph.Set (d)
}


func init () {
//
  h:= [...]string { "        neue Ecke: linke Maustaste              ",
                    " Ecke verschieben: rechte Maustaste             ",
                    "  Ecken verbinden: mittlere Maustaste           ",
                    "                                                ",
                    "Eckennamen ändern: Vorwahl- und linke Maustaste ",
                    "     Ecke löschen: Vorwahl- und rechte Maustaste",
                    " Graph ausdrucken: Drucktaste                   ",
                    "                                                ",
                    "(Vorwahltaste = Umschalt-, Strg- oder Alt-Taste)",
                    "                                                ",
                    "             Done: Abbruchtaste (Esc)           " }
  help = make ([]string, len (h))
  for i, l:= range (h) { str.Set (&help[i], l) }
//  scr.Fullscreen ()
//  scr.Switch (scr.WVGApp)
//  ker.InstallTerm (Terminate) // goes not :-(
}
