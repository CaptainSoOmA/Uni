package rob

// (c) Christian Maurer   v. 130115 - license see murus.go

import
  . "murus/obj"
type
  Nummer uint8
const (
  r0 = Nummer(0)
  MaxRobo = Nummer(nX / 2)
  niemand = Nummer(MaxRobo)
)
type
  dieRobos struct {
          derRobo [MaxRobo](*Imp)
                  }
var
  nRobo,
  nInitialisiert Nummer


func (x *dieRobos) impr (Y Any) *dieRobos {
//
  y, ok:= Y.(*dieRobos)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}


func NeueRobos () *dieRobos {
//
  x:= new (dieRobos)
  for r:= r0; r < MaxRobo; r++ {
    x.derRobo[r] = new (Imp)
    x.derRobo[r].nr = niemand
//    x.derRobo[r].richtung = Richtung(Nord)
//    x.derRobo[r].Y = nY
//    x.derRobo[r].X = nX
  }
  return x
}


// Liefert genau dann eine Nummer != niemand,
// wenn noch ein Roboter initialisiert werden kann.
// In diesem Fall ist archiviert genau dann true, wenn
// ein Roboter mit dieser Nummer schon in der Roboterwelt vorhanden war.
func (x *dieRobos) freieNummer (archiviert *bool) Nummer {
//
  n, i:= r0, r0
  for {
    if x.derRobo [n].nr < niemand {
// println ("nr ", n, ": i == ", i)
      if nInitialisiert < nRobo {
        if i == nInitialisiert {
          *archiviert = true
          break
        } else {
          i ++
        }
      }
    } else {
      if nInitialisiert == nRobo {
// println ("init, i == ", i)
        nRobo ++
        *archiviert = false
        break
      }
    }
    if n + 1 < niemand {
      n ++
    } else {
// println ("max erreicht")
      return niemand
    }
  }
  nInitialisiert ++
// println ("ok, freie nr == ", n)
//  x.derRobo [r].nr = r
  return n
}


func (x *dieRobos) Empty () bool {
//
  for r:= r0; r < MaxRobo; r++ {
    if x.derRobo [r].nr != niemand {
      return false
    }
  }
  return true
}


func (x *dieRobos) Clr () {
//
  for r:= r0; r < MaxRobo; r++ {
    x.derRobo [r].leeren ()
  }
}


func (x *dieRobos) Copy (Y Object) {
//
  y:= x.impr (Y)
  for r:= r0; r < MaxRobo; r++ {
    x.derRobo[r].nr = y.derRobo[r].nr
    x.derRobo[r].Y = y.derRobo[r].Y
    x.derRobo[r].X = y.derRobo[r].X
    x.derRobo[r].richtung = y.derRobo[r].richtung
    x.derRobo[r].tasche = y.derRobo[r].tasche
  }
}


func (x *dieRobos) Clone () Object {
//
  y:= NeueRobos ()
  y.Copy (x)
  return y
}


func (x *dieRobos) Eq (Y Object) bool {
//
  y:= x.impr (Y)
  for r:= r0; r < MaxRobo; r++ {
    if x.derRobo[r].nr != y.derRobo[r].nr { return false }
//    x.derRobo[r].Y = Zeilen(0)
//    x.derRobo[r].X = Spalten(0)
//    x.derRobo[r].richtung = Richtung(Nord)
//    x.derRobo[r].tasche = Klotzzahl(0)
  }
  return true
}


func (x *dieRobos) Less (Y Object) bool {
//
  return false
}


func (x *dieRobos) Terminieren (r Nummer) {
//
//  x.derRobo [r].nr = niemand
  nInitialisiert --
  nRobo --
}


func (x *dieRobos) inkrementieren (r Nummer, k Klotzzahl) {
// Vor.: Der Roboter mit der Nummer r ist initialisiert.
// Der r-te archivierte Roboter hat k Klötze mehr in der Tasche.
//
  if x.derRobo [r].tasche + k >= MaxK { /* stop */ }
  x.derRobo [r].tasche += k
}


// Eff.: Liefert genau dann true, wenn auf Platz [y, x] ein Roboter steht.
func (xr *dieRobos) einRoboterDa (y, x uint) (bool, Nummer) {
//
  for r:= r0; r < MaxRobo; r++ {
    if xr.derRobo[r] == nil {
      break
    }
    if xr.derRobo[r].nr < MaxRobo {
      if y == xr.derRobo[r].Y && x == xr.derRobo[r].X {
// print ("einRoboterDa nr ", r, " mit der nr ", xr.derRobo[r].nr, " auf Platz Y == ", xr.derRobo[r].Y, "/X == ", xr.derRobo[r].X)
        return true, xr.derRobo[r].nr
      }
    }
  }
  return false, MaxRobo
}


//func mehrereRobo () bool {
////
//  return nRobo > 1
//}


// Eff.: r ist archiviert.
func (x *dieRobos) archivieren (r *Imp) {
//
// println ("archiviert wird Robo nr ", r.nr)
  if r.nr == MaxRobo { return }
//  dR.derRobo [r.nr] = r
  x.derRobo [r.nr].kopieren (r)
// println ("Roboter archiviert mit nr ", dR.derRobo[r.nr].nr, " auf Platz Y == ", r.Y, "/X == ", r.X)
}


// r stimmt mit dem n-ten archivierten Roboter überein.
func (x *dieRobos) restaurieren (r *Imp, n Nummer) {
//
  r.kopieren (x.derRobo [n])
// println ("Roboter restauriert mit nr ", dR.derRobo[n].nr," auf Platz Y == ", r, dR.derRobo[n].Y, "/X == ", dR.derRobo[n].X)
}


func (x *dieRobos) ausgeben () {
//
  for r:= r0; r < MaxRobo; r++ {
    if x.derRobo [r] != nil {
      if x.derRobo [r].nr < MaxRobo {
        x.derRobo [r].ausgebenR ()
      }
    }
  }
}


func (x *dieRobos) Codelen () uint {
//
  return uint(MaxRobo) * x.derRobo[r0].Codelen ()
}


func (x *dieRobos) Encode () []byte {
//
  i, a:= uint(0), x.derRobo[r0].Codelen()
  b:= make ([]byte, uint(MaxRobo) * a)
  for r:= r0; r < MaxRobo; r++ {
//  if dR.derRobo[r].nr != r && r == 0 {
//    println ("Encode derRobo nr ", r, " mit der Nummer ", x.derRobo [r].nr, " auf Platz ", x.derRobo[r].Y, "/", x.derRobo[r].X)
//    x.derRobo [r].nr = r
//    println ("geändert in ", x.derRobo [r].nr)
//  }
    copy (b[i:i+a], x.derRobo [r].Encode())
    i += a
  }
  return b
}


func (x *dieRobos) Decode (b []byte) {
//
  nRobo = 0
  i, a:= uint(0), x.derRobo[r0].Codelen()
  for r:= r0; r < MaxRobo; r++ {
    x.derRobo [r].Decode (b[i:i+a])
//    println ("Decode derRobo nr ", r, " mit der nr ", dR.derRobo[r].nr, " auf Platz ", dR.derRobo[r].Y, "/", dR.derRobo[r].X)
    if x.derRobo [r].nr < niemand {
      nRobo ++
    } else {
      x.derRobo [r].nr = niemand
    }
    i += a
  }
}
