package rob

// (c) Christian Maurer   v. 130526 - license see murus.go

import (
  . "murus/obj"; "murus/str"
  "murus/col"
  "murus/scr"
  "murus/sel"
  "murus/pseq"
)
const (
  nX = 32; nY = 16 // passt zu WVGApp = 1024 x 576, denn:
                   // 576 - 16 * imp.zellengroesse = 4 * scr.NY1, d.h. wir haben genug Platz
                   // für je 2 Kopf- und Fußzeilen (Weltname usw. und Fehlermeldungen/Hinweise)
  null = uint(0)
)
type (
  platz struct {
       kloetze Klotzzahl
      besitzer Nummer // genau dann == niemand, wenn der platz von keinem Robo
      markiert,       // markiert oder mit Klötzen belegt
    zugemauert bool
               }
  Welt struct {
            p [nY][nX]platz
     schatten *dieRobos
              }
)
var (
  Roboterwelt *Welt = neueWelt ()
  datei *pseq.Imp
)


func neueWelt () *Welt {
//
  W:= new (Welt)
  W.schatten = NeueRobos ()
  W.Clr ()
  return W
}


func (W *Welt) Definieren () { //
//
  _, dateiname:= sel.Names ("Robos Welt:", suffix, 8, 0, 0, col.LightWhite, col.Blue)
  datei = pseq.New (Roboterwelt)
  datei.Name (dateiname)
  if ! datei.Empty() {
    W.Copy (datei.Get().(*Welt))
  }
  W.Ausgeben()
}


func (W *Welt) Terminieren () {
//
  if datei == nil { println ("Welt.go/Terminieren: datei == nil") }
  datei.Seek (0)
  datei.Put (Roboterwelt)
  datei.Terminate ()
//  W.schatten.Terminieren (R.nr)
//  W.Ausgeben()
}


func (W *Welt) Empty () bool {
//
  for y:= null; y < nY; y++ {
    for x:= null; x < nX; x++ {
      if W.p[y][x].kloetze > 0 { return false }
      if W.p[y][x].besitzer != niemand { return false }
    }
  }
  if ! W.schatten.Empty () {
    return false
  }
  return true
}


func (W *Welt) Clr () {
//
  for y:= null; y < nY; y++ {
    for x:= null; x < nX; x++ {
      W.p[y][x].kloetze = 0
      W.p[y][x].besitzer = niemand
      W.p[y][x].markiert = false
      W.p[y][x].zugemauert = false
    }
  }
  W.schatten.Clr ()
}


func (W *Welt) Copy (X Object) {
//
  W1, ok:= X.(*Welt)
  if ! ok { return }
  for y:= null; y < nY; y++ {
    for x:= null; x < nX; x++ {
      W.p[y][x].kloetze = W1.p[y][x].kloetze
      W.p[y][x].besitzer = W1.p[y][x].besitzer
      W.p[y][x].markiert = W1.p[y][x].markiert
      W.p[y][x].zugemauert = W1.p[y][x].zugemauert
    }
  }
  W.schatten.Copy (W1.schatten)
}


func (W *Welt) Clone () Object {
//
  W1:= neueWelt ()
  W1.Copy (W)
  return W1
}


func (W *Welt) Eq (X Object) bool {
//
  W1, ok:= X.(*Welt)
  if ! ok { return false }
  for y:= null; y < nY; y++ {
    for x:= null; x < nX; x++ {
      if W.p[y][x].kloetze != W1.p[y][x].kloetze { return false }
      if W.p[y][x].besitzer != W1.p[y][x].besitzer { return false }
      if W.p[y][x].markiert != W1.p[y][x].markiert { return false }
      if W.p[y][x].zugemauert != W1.p[y][x].zugemauert { return false }
    }
  }
//  schatten ?
  return true
}


func (W *Welt) Less (Y Object) bool {
//
  return false
}


func (W *Welt) Codelen () uint {
//
  return nY * nX * (Codelen (uint16(0)) + 1) + W.schatten.Codelen()
}


func (W *Welt) Encode () []byte {
//
  B:= make ([]byte, W.Codelen())
  var k Klotzzahl
  a:= 0
  for y:= null; y < nY; y++ {
    for x:= null; x < nX; x++ {
      if W.p[y][x].zugemauert {
        k = 2 * MaxK + 2
      } else {
        k = 2 * W.p[y][x].kloetze
        if W.p[y][x].markiert {
          k++
        }
      }
      copy (B[a:a+2], Encode(uint16(k)))
      a += int(Codelen(uint16(0)))
      B[a] = byte(W.p[y][x].besitzer)
      a += 1
    }
  }
  copy (B[a:a+int(W.schatten.Codelen())], W.schatten.Encode ())
  return B
}


func (W *Welt) Decode (B []byte) {
//
  var k Klotzzahl
  a:= 0
  for y:= null; y < nY; y++ {
    for x:= null; x < nX; x++ {
      k = Klotzzahl(Decode (uint16(0), B[a:a+2]).(uint16))
      W.p[y][x].zugemauert = k > 2 * MaxK + 1
      if W.p[y][x].zugemauert {
        W.p[y][x].kloetze = 0
        W.p[y][x].markiert = false
      } else {
        W.p[y][x].kloetze = k / 2
        W.p[y][x].markiert = k % 2 == 1
      }
      a += int(Codelen(uint16(0)))
      W.p[y][x].besitzer = Nummer (B[a])
      if W.p[y][x].zugemauert {
        W.p[y][x].besitzer = niemand
      }
      a += 1
    }
  }
  W.schatten.Decode (B[a:])
}


func (W *Welt) ausgeben1 (Y, X uint) {
//
  platzAusgeben (Y, X)
  if Y > 0 { platzAusgeben (Y - 1, X) }
  if Y + 1 < nY { platzAusgeben (Y + 1, X) }
  if X > 0 { platzAusgeben (Y, X - 1) }
  if X + 1 < nX { platzAusgeben (Y, X + 1) }
  W.schatten.ausgeben()
}


func (W *Welt) ausgeben2 (Y, X uint) {
//
  for y:= null; y < nY; y++ {
    platzAusgeben (y, X)
  }
  for x:= null; x < nX; x++ {
    platzAusgeben (Y, x)
  }
  W.schatten.ausgeben()
}


func (W *Welt) Ausgeben () {
//
  l:= str.Clr (scr.NColumns())
  scr.Colours (farbeV, farbeH)
  scr.Write (l, 0, 0) // kollidiert mit geplanter Ausgabe des Weltnamens
  scr.Write (l, 1, 0)
  scr.Write (l, scr.NLines() - 2, 0)
//  scr.Write (l, scr.NLines() - 3, 0)
  for y:= null; y < nY; y++ {
    for x:= null; x < nX; x++ {
      platzAusgeben (y, x)
    }
  }
  W.schatten.ausgeben()
}


func (W *Welt) besitzer (y, x uint) Nummer {
//
  return W.p[y][x].besitzer
}


func (W *Welt) besetzen (y, x uint, n Nummer) {
//
  W.p[y][x].besitzer = n
}


func (W *Welt) kloetze (y, x uint) Klotzzahl {
//
  return W.p[y][x].kloetze
}


func (W *Welt) klotzen (y, x uint, k Klotzzahl) {
//
  W.p[y][x].kloetze = k
}


func (W *Welt) kloetzeInkrementieren (y, x uint) {
//
  W.p[y][x].kloetze ++
}


func (W *Welt) kloetzeDekrementieren (y, x uint) {
//
  W.p[y][x].kloetze --
}


func (W *Welt) markiert (y, x uint) bool {
//
  return W.p[y][x].markiert
}


func (W *Welt) markieren (y, x uint, m bool) {
//
  W.p[y][x].markiert = m
}


func (W *Welt) zugemauert (y, x uint) bool {
//
  return W.p[y][x].zugemauert
}


func (W *Welt) zumauern (y, x uint, m bool) {
//
  W.p[y][x].zugemauert = m
}


func init () {
//
  Roboterwelt = neueWelt ()
}
