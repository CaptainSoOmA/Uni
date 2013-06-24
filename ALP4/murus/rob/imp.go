package rob

// (c) Christian Maurer   v. 130526 - license see murus.go

import (
  "murus/ker"; . "murus/obj"
  "murus/kbd"
  "murus/col"; "murus/scr"; "murus/errh"; "murus/nat"; "murus/img"
)
type
  arten byte; const (
  nichts = iota
  einRoboter
  klotz
  marke
  mauer
  nArten
)
type
  Imp struct {
          nr Nummer
        Y, X uint // muss noch umbenannt werden
    richtung Richtung
      tasche Klotzzahl
    aktionen []Aktion
             }
const (
  suffix = "robo"
  zellengroesse = 32 // 2 * scr.NY1()
)
var (
  farbeV, farbeH, randfarbe, mauerfarbe col.Colour
  schrittweise, amEditieren, Sokoban bool
  bild [nArten][zellengroesse]string
  farbe [MaxRobo+1]col.Colour
  Hilfe [21]string
  Sokobanhilfe [5]string
)


func schreiten () {
//
  if ! schrittweise { return }
  errh.Hint (errh.ToContinueOrNot)
  loop: for {
    c, _:= kbd.Command ()
    switch c { case kbd.Enter:
      break loop
    case kbd.Esc:
      ker.Terminate()
    }
  }
}


func (R *Imp) Farbe () col.Colour {
//
  return farbe [R.nr]
}


func (R *Imp) melden () {
//
  if Roboterwelt.kloetze (R.Y, R.X) > 0 {
    nat.SetColours (farbe [R.nr], farbeH)
    nat.Write (uint(Roboterwelt.kloetze (R.Y, R.X)), 0, 32 + 2)
  } else {
    scr.Colours (farbeH, farbeH)
    scr.Write ("   ", 0, 32)
  }
}


func ausgeben1 (a arten, r Richtung, y uint, x uint, f, b col.Colour) {
//
  y0:= (zellengroesse / int(scr.NY1())) * int(y + 1)
  x0:= (zellengroesse / int(scr.NX1())) * int(x)
  x0 *= int(scr.NX1())
  y0 = int(scr.NY1()) * (y0) // * (y0 - 1)
  if a == nichts { // schneller:
    scr.Colour (farbeH)
    scr.RectangleFull (x0 + 1, y0 + 1, x0 + zellengroesse - 2, y0 + zellengroesse - 1)
    scr.Colour (randfarbe)
    scr.Rectangle (x0, y0, x0 + zellengroesse - 1, y0 + zellengroesse - 1)
  } else {
// Datenstrukturen zur Beschleunigung der Ausgabe unter X um das 171-fache
// durch Reduktion der Aufrufe von X-Routinen von 2 * 32 * 32 = 2048 auf 2 * 6 = 12:
    const (
      n = 6 // Farbzahl
      zz = zellengroesse * zellengroesse
    )
    var (
      nr int
      zelle [zellengroesse][zellengroesse]int
    )
    for dy:= 0; dy < zellengroesse; dy++ {
      for dx:= 0; dx < zellengroesse; dx++ {
        nr = 0 // farbeH
        switch bild [a][dy][dx] { case 'o':
          nr = 1 // F
        case 'x':
          nr = 2 // randfarbe
        case '+':
          nr = 3 // farbeV
        case 'k':
          if Roboterwelt.kloetze (y, x) > 0 {
            nr = 4 // B
          }
        case 'm':
          if Roboterwelt.markiert (y, x) {
            nr = 4 // B
          }
        case ' ':
          if a == mauer {
            nr = 5 // mauerfarbe
          } else {
            nr = 0 // farbeH
          }
        default:
          return // stop
        }
        if a != einRoboter { r = Nord }
        switch r { case Nord:
          zelle[dx][dy] = nr
        case West:
          zelle[dy][dx] = nr
        case Sued:
          zelle[dx][zellengroesse - 1 - dy] = nr
        case Ost:
          zelle[zellengroesse - 1 - dy][dx] = nr
        }
      }
    }
    var (
      anzahl [n]int
      xx, yy [n][zz]int
    )
    for dy:= 0; dy < zellengroesse; dy++ {
      for dx:= 0; dx < zellengroesse; dx++ {
        nr = zelle[dx][dy]
        xx[nr][anzahl[nr]] = x0 + dx
        yy[nr][anzahl[nr]] = y0 + dy
        anzahl[nr] ++
      }
    }
    c:= [n]col.Colour { farbeH, f, randfarbe, farbeV, b, mauerfarbe }
    for i:= 0; i < n; i++ {
      scr.Colour (c[i])
      scr.Pointset (xx[i][:], yy[i][:])
    }
  }
}


func platzAusgeben (y uint, x uint) {
//
  if da, _:= Roboterwelt.schatten.einRoboterDa (y, x); ! da {
    if Roboterwelt.zugemauert (y, x){
      ausgeben1 (mauer, Nord, y, x, farbeH, mauerfarbe)
    } else {
      var art arten
      if Roboterwelt.kloetze (y, x) == 0 {
        if Roboterwelt.markiert (y, x) {
          art = marke
        } else {
          art = nichts
        }
      } else {
        art = klotz
      }
      f:= farbe [Roboterwelt.besitzer (y, x)]
      ausgeben1 (art, Nord, y, x, f, f)
    }
  }
}


func (R *Imp) ausgebenR () {
//
  f:= farbe [R.nr]
  b:= farbe [Roboterwelt.besitzer (R.Y, R.X)]
  ausgeben1 (einRoboter, R.richtung, R.Y, R.X, f, b)
  if amEditieren {
    R.melden ()
  }
}


func platzDa (Y *uint, X *uint) bool {
//
  for {
    if da, _:= Roboterwelt.schatten.einRoboterDa (*Y, *X); da || Roboterwelt.zugemauert (*Y, *X) {
      if *X + 1 < nX {
        *X++
      } else if *Y + 1 < nY {
        *Y++
        *X = 0
      } else {
        break
      }
    }
    return true
  }
  return false
}


func (R *Imp) WeltAusgeben() {
//
  Roboterwelt.Ausgeben()
}


func (R *Imp) Aktualisieren () {
//
  for y:= null; y < nY; y++ {
    for x:= null; x < nX; x++ {
      da, n:= Roboterwelt.schatten.einRoboterDa (y, x)
      if da && n == R.nr {
        Roboterwelt.schatten.restaurieren (R, n)
/*
        R.Y, R.X = y, x
        R.richtung = Roboterwelt.schatten.derRobo[n].richtung
*/
      }
    }
  }
  Roboterwelt.Ausgeben()
}


func NeuerRoboter () *Imp {
//
  var archiviert bool
  r:= Roboterwelt.schatten.freieNummer (&archiviert)
  if r == niemand { // nix geht mehr
//    errh.Error ("Die Roboterwelt ist voll !", 0)
    return nil // besser: Error erzeugen, der abgefragt werden kann
  }
  R:= new (Imp)
  if archiviert {
// println ("NeuerRoboter schon im Archiv enthalten")
    Roboterwelt.schatten.restaurieren (R, r)
  } else {
    R.nr = r
// println ("NeuerRoboter hat die Nummer ", r)
    R.Y = 0
    R.X = uint(R.nr)
    if ! platzDa (&(R.Y), &(R.X)) { // Vorsicht, X jetzt erhöht, weil schon im Schattenarchiv !
// println ("Neuer Roboter: nach ! platzDa, Pos = ", R.Y, "/", R.X)
      errh.Error ("kein Platz mehr frei für Roboter", 0)
      return nil
    }
    R.richtung = Sued
    R.tasche = MaxK
    for y:= null; y < nY; y++ {
      for x:= null; x < nX; x++ {
        if y == R.Y && x == R.X {
          R.tasche -= Roboterwelt.kloetze (y, x)
        }
      }
    }
    Roboterwelt.schatten.archivieren (R)
  }
  R.aktionen = make ([]Aktion, 0)
  return R
}


func (R *Imp) kopieren (R1 *Imp) {
//
  R.nr = R1.nr
  R.Y, R.X = R1.Y, R1.X
  R.richtung = R1.richtung
  R.tasche = R1.tasche
//  R.aktionen
}


func WeltDefinieren () {
//
  Roboterwelt.Definieren ()
}


func (R *Imp) Terminieren () {
//
  Roboterwelt.Terminieren ()
}


func (R *Imp) leeren () {
//
  R.nr = niemand
  R.aktionen = R.aktionen [0:0]
}


func (R *Imp) Nummer () uint {
//
  return uint(R.nr)
}


func (R *Imp) InLinkerObererEcke () bool {
//
  return uint(R.Y) + uint(R.X) == 0
}


func (R *Imp) ablegen (a Aktion) {
//
  R.aktionen = append (R.aktionen, a) // TODO
}


func (R *Imp) letzteAktion () Aktion {
//
  return R.aktionen[len(R.aktionen) - 1]
}


func (R *Imp) remove () {
//
  R.aktionen = R.aktionen[0:len(R.aktionen)-1]
}


func (R *Imp) Trav (op Op) {
//
  if len (R.aktionen) > 0 {
    for i:= len (R.aktionen) - 1; i > 0; i-- {
      op (R.aktionen [i])
    }
  }
  R.aktionen = R.aktionen[0:0]
}


func (R *Imp) links () {
//
  R.richtung = (R.richtung + 1) % nRichtungen
  Roboterwelt.schatten.archivieren (R)
  Roboterwelt.ausgeben1 (R.Y, R.X)
}


func (R *Imp) rechts () {
//
  R.richtung = (R.richtung + nRichtungen - 1) % nRichtungen
  Roboterwelt.schatten.archivieren (R)
  Roboterwelt.ausgeben1 (R.Y, R.X)
}


func (R *Imp) LinksDrehen () {
//
  schreiten()
  R.links ()
  R.ablegen (Links)
}


func (R *Imp) linksDrehenZurueck () {
//
  R.rechts ()
  R.remove ()
}


func (R *Imp) RechtsDrehen () {
//
  schreiten ()
  R.rechts ()
  R.ablegen (Rechts)
}


func (R *Imp) rechtsDrehenZurueck () {
//
  R.links ()
  R.remove ()
}


func (R *Imp) AmRand () bool {
//
  switch R.richtung { case Nord:
    return R.Y == 0
  case West:
    return R.X == 0
  case Sued:
    return R.Y + 1 == nY
  case Ost:
    break
  }
  return R.X + 1 == nX
}


// Liefert genau dann true, wenn es keinen Nachbarplatz entgegen Rs Richtung gibt.
func (R *Imp) randHinten () bool {
//
  switch R.richtung { case Nord:
    return R.Y + 1 == nY
  case West:
    return R.X + 1 == nX
  case Sued:
    return R.Y == 0
  case Ost:
    break
  }
  return R.X == 0
}


// (y, x) ist die Position von Rs Nachbarplatz in Rs Richtung, falls R nicht am Rand steht, andernfalls Rs Position.
func (R *Imp) nachbarPosition () (y uint, x uint) {
//
  y, x = R.Y, R.X
  switch R.richtung { case Nord:
    if y > 0 { y-- }
  case West:
    if x > 0 { x-- }
  case Sued:
    if y + 1 < nY { y++ }
  case Ost:
    if x + 1 < nX { x++ }
  }
  return
}


// (y, x) ist die Positon von Rs Nachbarplatz entgegen Rs Richtung,
// falls nicht hinter R der Rand ist, andernfalls Rs Position.
func (R *Imp) hintereNachbarPosition () (y uint, x uint) {
//
  y, x = R.Y, R.X
  switch R.richtung { case Nord:
    if y + 1 < nY { y++ }
  case West:
    if x + 1 < nX { x++ }
  case Sued:
    if y > 0 { y-- }
  case Ost:
    if x > 0 { x-- }
  }
  return
}


func (R *Imp) weiter () bool {
//
  if R.AmRand () {
    return false
  }
  y, x:= R.nachbarPosition ()
  da, _:= Roboterwelt.schatten.einRoboterDa (y, x)
  if Roboterwelt.zugemauert (y, x) || da {
    return false
  }
  R.Y, R.X = y, x
  Roboterwelt.schatten.archivieren (R)
  Roboterwelt.ausgeben1 (y, x)
  return true
}


func (R *Imp) zurueck () bool {
//
  if R.randHinten () {
    return false
  }
  y, x:= R.hintereNachbarPosition ()
  da, _:= Roboterwelt.schatten.einRoboterDa (y, x)
  if Roboterwelt.zugemauert (y, x) || da {
    return false
  }
  R.Y, R.X = y, x
  Roboterwelt.schatten.archivieren (R)
  Roboterwelt.ausgeben1 (y, x)
  return true
}


func (R *Imp) Gelaufen1 () bool {
//
  schreiten()
  if ! R.weiter () {
    return false
  }
  R.ablegen (Weiter)
  return true
}


func (R *Imp) laufen1zurueck () {
//
  if R.zurueck () {
    R.remove ()
  }
}


func (R *Imp) Zurueckgelaufen1 () bool {
//
  schreiten ()
  if ! R.zurueck () {
    return false
  }
  R.ablegen (Zurueck)
  return true
}


func (R *Imp) zuruecklaufen1zurueck () {
//
  if R.weiter () {
    R.remove ()
  }
}


func (R *Imp) Leer () bool {
//
  return Roboterwelt.kloetze (R.Y, R.X) == 0
//  return R.nr == niemand
}


func (R *Imp) NachbarLeer () bool {
//
  if R.AmRand () {
    return false
  }
  y, x:= R.nachbarPosition ()
  da, _:= Roboterwelt.schatten.einRoboterDa (y, x)
  return Roboterwelt.kloetze (y, x) == 0 &&
         ! Roboterwelt.zugemauert (y, x) &&
         ! da
}


func (R *Imp) DarfLegen () bool {
//
  if Roboterwelt.zugemauert (R.Y, R.X) { /* stop */ }
  if Roboterwelt.markiert (R.Y, R.X) || Roboterwelt.kloetze (R.Y, R.X) > 0 {
    return Roboterwelt.besitzer (R.Y, R.X) == R.nr
  } else { // ! narkiert && Kloetze == 0
    if Roboterwelt.besitzer (R.Y, R.X) >= niemand {
    } else { // ! Markiert && Kloetze == 0 && Besitzer != niemand
//      stop ()
    }
  }
  return true
}


func (R *Imp) klotzweg () bool {
//
  if ! R.DarfLegen () || Roboterwelt.kloetze (R.Y, R.X) == 0 {
    return false
  }
  Roboterwelt.kloetzeDekrementieren (R.Y, R.X)
  if Roboterwelt.kloetze (R.Y, R.X) == 0 {
    if ! Roboterwelt.markiert (R.Y, R.X) {
      Roboterwelt.besetzen (R.Y, R.X, niemand)
    }
  }
  if R.tasche >= MaxK { /* stop */ }
  R.tasche++
  Roboterwelt.schatten.archivieren (R)
  Roboterwelt.ausgeben1 (R.Y, R.X)
  return true
}


func (R *Imp) klotzhin () bool {
  //
  if ! R.DarfLegen () || R.tasche == 0 {
    return false
  }
  if Roboterwelt.kloetze (R.Y, R.X) >= MaxK { return false }
  R.tasche--
  Roboterwelt.kloetzeInkrementieren (R.Y, R.X)
  Roboterwelt.besetzen (R.Y, R.X, R.nr)
  Roboterwelt.schatten.archivieren (R)
  Roboterwelt.ausgeben1 (R.Y, R.X)
  return true
}


func (R *Imp) Leeren1 () {
//
  schreiten()
  if R.klotzweg () {
    R.ablegen (KlotzWeg)
  }
}


func (R *Imp) leeren1zurueck () {
//
  if R.klotzhin () {
    R.remove ()
  }
}


func (R *Imp) HatKloetze () bool {
//
  return R.tasche > 0
}


func (R *Imp) KloetzeGeben (n uint) {
//
  R.tasche = Klotzzahl(n)
}


func (R *Imp) AnzahlKloetze () uint {
//
  return uint(R.tasche)
}


func (R *Imp) Legen1 () {
//
  schreiten()
  if R.klotzhin () {
    R.ablegen (KlotzHin)
  }
}


func (R *Imp) legen1zurueck () {
//
  if R.klotzweg () {
    R.remove ()
  }
}


func (R *Imp) klotzWeiter () bool {
//
  if R.AmRand () || R.VorMauer () {
    return false
  }
  y0, x0:= R.Y, R.X
  R.Y, R.X = R.nachbarPosition ()
  k:= Roboterwelt.kloetze (R.Y, R.X)
  if R.AmRand () ||
     R.VorMauer () ||
     k == 0 ||
     Roboterwelt.besitzer (R.Y, R.X) != R.nr {
    R.Y, R.X = y0, x0
    return false
  }
  y1, x1:= R.Y, R.X
  R.Y, R.X = R.nachbarPosition()
  da, _:= Roboterwelt.schatten.einRoboterDa (R.Y, R.X)
  if da || Roboterwelt.kloetze (R.Y, R.X) > 0 ||
     Roboterwelt.markiert (R.Y, R.X) && Roboterwelt.besitzer (R.Y, R.X) != R.nr {
    R.Y, R.X = y0, x0
    return false
  }
  Roboterwelt.klotzen (R.Y, R.X, k)
  Roboterwelt.besetzen (R.Y, R.X, R.nr)
  R.Y, R.X = y1, x1
  Roboterwelt.klotzen (R.Y, R.X, 0)
  if ! Roboterwelt.markiert (R.Y, R.X) {
    Roboterwelt.besetzen (R.Y, R.X, niemand)
  }
  Roboterwelt.schatten.archivieren (R)
  Roboterwelt.ausgeben1 (R.Y, R.X)
  return true
}


func (R *Imp) Geschoben1 () bool {
//
  schreiten()
  if ! R.klotzWeiter () {
    return false
  }
  R.ablegen (KlotzWeiter)
  return true
}


func (R *Imp) schieben1zurueck () {
//
  if ! R.DarfLegen () {
    return
  }
  y, x:= R.hintereNachbarPosition ()
  da, _:= Roboterwelt.schatten.einRoboterDa (y, x)
  if Roboterwelt.zugemauert (y, x) || da {
    return
  }
  y1, x1:= R.nachbarPosition ()
  k:= Roboterwelt.kloetze (y1, x1)
  Roboterwelt.klotzen (y1, x1, 0)
  if ! Roboterwelt.markiert (y1, x1) {
    Roboterwelt.besetzen (y1, x1, niemand)
  }
  Roboterwelt.klotzen (R.Y, R.X, k)
  Roboterwelt.besetzen (R.Y, R.X, R.nr)
  R.Y, R.X = y, x
  Roboterwelt.schatten.archivieren (R)
  Roboterwelt.ausgeben2 (R.Y, R.X)
  R.remove ()
}


func (R *Imp) virtuellWeiter () bool {
//
  var (y uint; x uint)
  if R.NachbarLeer () {
    y, x = R.nachbarPosition ()
  } else {
    return false
  }
  if Roboterwelt.markiert (y, x) {
    return false
  }
  R.Y, R.X = y, x
  return true
}


func (R *Imp) schiessen () {
//
  k:= Roboterwelt.kloetze (R.Y, R.X)
  Roboterwelt.klotzen (R.Y, R.X, 0)
  r:= Roboterwelt.besitzer (R.Y, R.X)
  Roboterwelt.besetzen (R.Y, R.X, niemand)
  if R.virtuellWeiter () {
    Roboterwelt.klotzen (R.Y, R.X, k)
    Roboterwelt.besetzen (R.Y, R.X, r)
    R.Schiessen ()
    y, x:= R.hintereNachbarPosition ()
    R.Y, R.X = y, x
  } else {
    Roboterwelt.klotzen (R.Y, R.X, k)
    Roboterwelt.besetzen (R.Y, R.X, r)
  }
}


func (R *Imp) Schiessen () {
//
  if ! R.Leer () {
    R.schiessen ()
    Roboterwelt.schatten.archivieren (R)
    Roboterwelt.ausgeben2 (R.Y, R.X)
  }
}


func (R *Imp) Markiert () bool {
//
  return Roboterwelt.markiert (R.Y, R.X)
}


func (R *Imp) NachbarMarkiert () bool {
//
  if R.AmRand () || R.VorMauer () {
    return false
  }
  y, x:= R.nachbarPosition ()
  return Roboterwelt.markiert (y, x)
}


func (R *Imp) DarfMarkieren () bool {
//
  return R.DarfLegen ()
}


func (R *Imp) markeHin () bool {
//
  if ! R.DarfMarkieren () || Roboterwelt.markiert (R.Y, R.X) {
    return false
  }
  Roboterwelt.markieren (R.Y, R.X, true)
  Roboterwelt.besetzen (R.Y, R.X, R.nr)
  Roboterwelt.schatten.archivieren (R)
  Roboterwelt.ausgeben1 (R.Y, R.X)
  return true
}


func (R *Imp) markeweg () bool {
//
  if ! R.DarfMarkieren () {
    return false
  }
  if Roboterwelt.markiert (R.Y, R.X) {
    Roboterwelt.markieren (R.Y, R.X, false)
    if Roboterwelt.kloetze (R.Y, R.X) == 0 {
      Roboterwelt.besetzen (R.Y, R.X, niemand)
    }
  } else {
    return false
  }
  Roboterwelt.schatten.archivieren (R)
  Roboterwelt.ausgeben1 (R.Y, R.X)
  return true
}


func (R *Imp) Markieren () {
//
  schreiten()
  if ! R.markeHin () {
    return
  }
  R.ablegen (MarkeHin)
}


func (R *Imp) markierenZurueck () {
//
  if R.markeweg () {
    R.remove ()
  }
}


func (R *Imp) Entmarkieren () {
//
  schreiten()
  if R.markeweg () {
    R.ablegen (MarkeWeg)
  }
}


func (R *Imp) entmarkierenZurueck () {
//
  if R.markeweg () {
    R.remove ()
  }
}


// Alle Markierungen, die R gesetzt hat, sind entfernt.
func (R *Imp) allesEntmarkieren () {
//
  for y:= null; y < nY - 1; y++ {
    for x:= null; x < nX - 1; x++ {
      if Roboterwelt.markiert (y, x) && Roboterwelt.besitzer (y, x) == R.nr {
        Roboterwelt.markieren (y, x, false)
        if Roboterwelt.kloetze (y, x) == 0 {
          Roboterwelt.besetzen (y, x, niemand)
        }
      }
    }
  }
  Roboterwelt.Ausgeben()
}


func (R *Imp) VorMauer () bool {
//
  if R.AmRand () {
    return false
  }
  y, x:= R.nachbarPosition ()
  return Roboterwelt.zugemauert (y, x)
}


func (R *Imp) Gemauert1 () bool {
//
  if R.AmRand () || R.VorMauer () {
    return false
  }
  y1, x1:= R.nachbarPosition ()
  if da, _:= Roboterwelt.schatten.einRoboterDa (y1, x1); da {
    return false
  }
  schreiten ()
  Roboterwelt.markieren (R.Y, R.X, false)
  if Roboterwelt.kloetze (R.Y, R.X) > 0 {
    if Roboterwelt.besitzer (R.Y, R.X) == R.nr {
      R.tasche += Roboterwelt.kloetze (R.Y, R.X)
    } else {
      Roboterwelt.schatten.inkrementieren (Roboterwelt.besitzer (R.Y, R.X), Roboterwelt.kloetze (R.Y, R.X))
    }
    Roboterwelt.klotzen (R.Y, R.X, 0)
  }
  Roboterwelt.besetzen (R.Y, R.X, niemand)
  Roboterwelt.zumauern (R.Y, R.X, true)
  R.Y, R.X = y1, x1
  R.ablegen (MauerHin)
  Roboterwelt.schatten.archivieren (R)
  Roboterwelt.ausgeben1 (R.Y, R.X)
  return true
}


func (R *Imp) mauern1zurueck () {
//
  y, x:= R.hintereNachbarPosition ()
  if ! Roboterwelt.zugemauert (y, x) {
    return
  }
  Roboterwelt.zumauern (y, x, false)
  R.Y, R.X = y, x
  Roboterwelt.schatten.archivieren (R)
  Roboterwelt.ausgeben1 (R.Y, R.X)
  R.remove ()
}


func (R *Imp) Entmauert1 () bool {
//
  if R.AmRand () || ! R.VorMauer () {
    return false
  }
  schreiten()
  R.Y, R.X = R.nachbarPosition ()
  if da, _:= Roboterwelt.schatten.einRoboterDa (R.Y, R.X); da { /* stop */ }
  if Roboterwelt.besitzer (R.Y, R.X) != R.nr && Roboterwelt.besitzer (R.Y, R.X) != niemand { /* stop */ }
  Roboterwelt.zumauern (R.Y, R.X, false)
  R.ablegen (MauerWeg)
  Roboterwelt.schatten.archivieren (R)
  Roboterwelt.ausgeben1 (R.Y, R.X)
  return true
}


func (R *Imp) entmauern1zurueck () {
//
  y, x:= R.hintereNachbarPosition ()
  da, _:= Roboterwelt.schatten.einRoboterDa (y, x)
  if da || Roboterwelt.zugemauert (y, x) {
    return
  }
  Roboterwelt.zumauern (R.Y, R.X, true)
  R.Y, R.X = y, x
  Roboterwelt.schatten.archivieren (R)
  Roboterwelt.ausgeben1 (R.Y, R.X)
  R.remove ()
}


func (R *Imp) Codelen () uint {
//
  return 1 + // SizeOf (Nummer)
         2 + // SizeOf (uint16)
         1 + // SizeOf (Richtung) // Fehler in der Modula-2-Version !
         2   // SizeOf (Klotzzahl)
}


func (R *Imp) Encode () []byte {
//
//  println ("Encode Robotor nr. , R.nr, " auf Platz ", R.Y, "/", R.X)
  b:= make ([]byte, R.Codelen())
  b[0] = byte(R.nr)
// if R.X != 0 || R.Y != 0 {
// }
  s:= uint16(R.Y) + 256 * uint16(R.X)
  copy (b[1:3], Encode (s))
  b[3] = byte(R.richtung)
  copy (b[4:6], Encode (uint16(R.tasche)))
  return b
}


func (R *Imp) Decode (b []byte) {
//
  R.nr = Nummer (b[0])
  if R.nr > niemand { R.nr = niemand }
  s:= uint(Decode (uint16(0), b[1:3]).(uint16))
  R.Y, R.X = s % 256, s / 256
// println ("Decode Roboter nr. ", R.nr, " auf Platz ", R.Y, "/", R.X)
  if Sokoban {
    R.richtung = Nord
  } else {
    R.richtung = Richtung(b[3])
  }
  R.tasche = Klotzzahl(Decode (uint16(0), b[4:6]).(uint16))
}


func (R *Imp) umkehren () {
//
  schreiten()
  switch R.richtung { case Nord:
    R.richtung = Sued
  case West:
    R.richtung = Ost
  case Sued:
    R.richtung = Nord
  case Ost:
    R.richtung = West
  }
  R.ablegen (Links)
  R.ablegen (Links)
  Roboterwelt.schatten.archivieren (R)
  Roboterwelt.ausgeben1 (R.Y, R.X)
}


func (R *Imp) Manipulieren (K kbd.Comm, T uint) {
//
  s:= schrittweise
  schrittweise = false
  amEditieren = true
  switch K { case kbd.Esc:
    return
  case kbd.Enter:
    if T == 0 {
      if R.Geschoben1 () { }
    } else {
      R.schiessen ()
    }
  case kbd.Left:
    if Sokoban {
      switch R.richtung { case Nord:
        R.LinksDrehen ()
      case West:
        ;
      case Sued:
        R.RechtsDrehen ()
      case Ost:
        R.umkehren ()
      }
      if R.NachbarLeer () {
        if R.Gelaufen1 () { }
      } else {
        if R.Geschoben1 () { }
      }
    } else {
      switch R.richtung { case Nord:
        R.LinksDrehen ()
      case West:
        if T == 0 {
          if R.Gelaufen1 () { }
        } else {
          for R.Gelaufen1 () { }
        }
      case Sued:
        R.RechtsDrehen ()
      case Ost:
        R.umkehren ()
      }
    }
  case kbd.Right:
    if Sokoban {
      switch R.richtung { case Nord:
        R.RechtsDrehen ()
      case West:
        R.umkehren ()
      case Sued:
        R.LinksDrehen ()
      case Ost:
        ;
      }
      if R.NachbarLeer () {
        if R.Gelaufen1 () { }
      } else {
        if R.Geschoben1 () { }
      }
    } else {
      switch R.richtung { case Nord:
        R.RechtsDrehen ()
      case West:
        R.umkehren ()
      case Sued:
        R.LinksDrehen ()
      case Ost:
        if T == 0 {
          if R.Gelaufen1 () { }
        } else {
          for R.Gelaufen1 () { }
        }
      }
    }
  case kbd.Up:
    if Sokoban {
      switch R.richtung { case Nord:
        ;
      case West:
        R.RechtsDrehen ()
      case Sued:
        R.umkehren ()
      case Ost:
        R.LinksDrehen ()
      }
      if R.NachbarLeer () {
        if R.Gelaufen1 () { }
      } else {
        if R.Geschoben1 () { }
      }
    } else {
      if T == 2 {
        R.Entmarkieren ()
      } else {
        switch R.richtung { case Nord:
          if T == 0 {
            if R.Gelaufen1 () { }
          } else {
            for R.Gelaufen1 () { }
          }
        case West:
          R.RechtsDrehen ()
        case Sued:
          R.umkehren ()
        case Ost:
          R.LinksDrehen ()
        }
      }
    }
  case kbd.Down:
    if Sokoban {
      switch R.richtung { case Nord:
        R.umkehren ()
      case West:
        R.LinksDrehen ()
      case Sued:
        ;
      case Ost:
        R.RechtsDrehen ()
      }
      if R.NachbarLeer () {
        if R.Gelaufen1 () { }
      } else {
        if R.Geschoben1 () { }
      }
    } else {
      if T == 2 {
        R.Markieren ()
      } else {
        switch R.richtung { case Nord:
          R.umkehren ()
        case West:
          R.LinksDrehen ()
        case Sued:
          if T == 0 {
            if R.Gelaufen1 () { }
          } else {
            for R.Gelaufen1 () { }
          }
        case Ost:
          R.RechtsDrehen ()
        }
      }
    }
  case kbd.Pos1:
    if ! Sokoban {
      if R.Gemauert1 () { }
    }
  case kbd.End:
    if ! Sokoban {
      if R.Entmauert1 () { }
    }
  case kbd.Tab:
    if ! Sokoban {
      switch T { case 0:
        if Roboterwelt.markiert (R.Y, R.X) {
          R.Entmarkieren ()
        } else {
          R.Markieren ()
        }
/*
        R.Markieren ()
      1: // unter X funktioniert Umschalt + Tab nicht
        R.Entmarkieren ()
      } else {
        R.allesEntmarkieren ()
*/
      }
    }
  case kbd.Ins:
    if ! Sokoban {
      if R.tasche > 0 {
        R.Legen1 ()
      }
    }
  case kbd.Del:
    if ! Sokoban {
      if T == 0 {
        R.Leeren1 ()
      } else {
        Roboterwelt.klotzen (R.Y, R.X, 0)
      }
    }
  case kbd.Back:
    for len (R.aktionen) > 0 {
      switch R.letzteAktion () { case Links:
        R.linksDrehenZurueck ()
      case Rechts:
        R.rechtsDrehenZurueck ()
      case Weiter:
        R.laufen1zurueck ()
      case Zurueck:
        R.zuruecklaufen1zurueck ()
      case KlotzWeg:
        R.leeren1zurueck ()
      case KlotzHin:
        R.legen1zurueck ()
      case KlotzWeiter:
        R.schieben1zurueck ()
      case MarkeHin:
        R.markierenZurueck ()
      case MarkeWeg:
        R.entmarkierenZurueck ()
      case MauerHin:
        R.mauern1zurueck ()
      case MauerWeg:
        R.entmauern1zurueck ()
      }
      if T == 0 {
        break
      }
    }
  case kbd.Help:
    if Sokoban {
      errh.WriteHelp (Sokobanhilfe[:])
    } else {
      errh.WriteHelp (Hilfe[:])
    }
  case kbd.LookFor:
    // Roboter wechseln
  case kbd.Act:
    // neuen Roboter initialisieren
  case kbd.Cfg:
    // terminieren (R)
  case kbd.Mark:
    if ! Sokoban {
      R.Markieren ()
    }
  case kbd.Demark:
    if ! Sokoban {
      if T == 0 {
        R.Entmarkieren ()
      } else {
        R.allesEntmarkieren ()
      }
    }
  case kbd.Paste:
    ;
  case kbd.Deposit:
    ;
  case kbd.Black:
    ;
  case kbd.Red:
    ;
  case kbd.Green:
    ;
  case kbd.Blue:
    ;
  case kbd.PrintScr:
    img.Print (0, 0, scr.NX(), scr.NY() - scr.NY1())
  case kbd.Roll:
    if T == 0 {
      // Roboter wechseln
    } else {
      // terminieren (R)
    }
  case kbd.Pause:
    // Roboter killen
  case kbd.Go:
    ;
/*
  case kbd.Here:
    if scr.UnderMouse (2 * R.Y + 1, 4 * R.X, 4, 2) {
      // ist angeklickt
    }
*/
  case kbd.Pull:
    // wenn Roboter getroffen, vorläufig hierhin
  case kbd.Hither:
    // wenn Roboter getroffen, endgültig hierhin
  case kbd.There:
    ;
  case kbd.Push:
    ;
  case kbd.Thither:
    ;
  case kbd.This:
    ;
  case kbd.Move:
    ;
  case kbd.Thus:
    ;
  case kbd.Navigate:
    ;
  }
  amEditieren = false
  schrittweise = s
}


func (R *Imp) Editieren () {
//
//  errh.WriteHelp1 ()
  R.Aktualisieren()
  errh.Hint ("Hilfe mit F1")
//  R.aktionen = R.aktionen [0:0]
  var (K kbd.Comm; T uint)
  for {
    K, T = kbd.Command ()
    R.Manipulieren (K, T)
    if K == kbd.Esc {
      break
    }
  }
  errh.DelHint()
}


func (R *Imp) SokobanSchalten (ein bool) {
//
  Sokoban = ein
}


func (R *Imp) Ausgeben (n uint) {
//
  nat.SetColours (col.Yellow, col.Blue)
  nat.Write (n, scr.NLines() - 2, 9)
  kbd.Wait (true)
}


func (R *Imp) Eingabe () uint {
//
  var n uint
  nat.SetColours (col.Yellow, col.Blue)
  nat.SetWd (10); nat.Edit (&n, scr.NLines() - 2, 9); nat.SetWd (0)
  return n
}


func (R *Imp) FehlerMelden (Text string, n uint) {
//
  errh.Error (Text, n)
}


func init () {
//
  Aktionstext = [nAktionen]string {
   "LinksDrehen", "RechtsDrehen",
   "Laufen1", "Zuruecklaufen1",
   "Legen1", "Leeren1", "Schieben1",
   "Markieren", "Entmarkieren",
   "Mauern1", "Entmauern1" }
  farbeV, farbeH = col.Black, col.LightWhite
  randfarbe = col.White
  mauerfarbe = col.LightRed // ziegelrot
  schrittweise = true

  Hilfe = [...]string {
    "               Roboter auf der Stelle drehen:  Pfeiltasten              ",
    "         Roboter einen Schritt laufen lassen:  Pfeiltasten              ",
    "                                                                        ",
    "einen Klotz auf der Stelle ablegen/aufnehmen:  Einfüge-/Entfernungstaste",
    " Klotz(-haufen) einen Schritt weiterschieben:  Eingabetaste (Enter)     ",
    "                  Klotz(-haufen) wegschießen:  Umschalt- + Eingabetaste ",
    "                                                                        ",
    "       Mauer setzen und einen Schritt laufen:  Anfangstaste (Pos1)      ",
    "    Mauer vor Roboter abreißen, weiterlaufen:  Endetaste                ",
    "                                                                        ",
    "            Markierung auf der Stelle setzen:  Tabulatortaste, F5-Taste ", // oder Bild abwärts
    "         Markierung auf der Stelle entfernen:  Tabulatortaste, F6-Taste ", // oder Bild aufwärts
    "                 alle Markierungen entfernen:  Umschalt- + F6-Taste     ",
    "                                                                        ",
    "            jeweils letzten Zug zurücknehmen:  Rücktaste (<-)           ",
    "alle Züge zurücknehmen, d.h. ganz zum Anfang:  Umschalt- + Rücktaste    ",
    "                                                                        ",
    "                      Roboterwelt ausdrucken:  Drucktaste               ",
    "                                                                        ",
    "                              Editor beenden:  Schlusstaste (Esc)       ",
    "                                                                        " }
  Sokobanhilfe = [...]string {
    Hilfe [0],
    "",
    "laufen: Pfeiltasten           Zug zurücknehmen: <-           fertig: Esc",
    "",
    "            Bedienungshinweise für Sokoban folgen irgendwann            " }

  for reihe:= 1; reihe < zellengroesse - 1; reihe++ {
    bild [nichts][reihe] = "x                              x"
  }
  bild [nichts][0] = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  bild [nichts][zellengroesse - 1] = bild [nichts][0]
/*
  bild [einRoboter] = [zellengroesse]string {
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    "x             ++++             x",
    "x            ++++++            x",
    "x           +oo++oo+           x",
    "x           +oo++oo+           x",
    "x           ++++++++           x",
    "x            +oooo+            x",
    "x             ++++             x",
    "x             +oo+             x",
    "x           +oooooo+           x",
    "x         +oooooooooo+         x",
    "x        +ooo+oooo+ooo+        x",
    "x       +ooo +oooo+ ooo+       x",
    "x      +ooo  +oooo+  ooo+      x",
    "x      +ooo  +oooo+   ooo+     x",
    "x      +oo   +oooo+    oo+     x",
    "x     +oo    +oooo+    oo+     x",
    "x     +oo   ++++++++   oo+     x",
    "x    o+o    +oooooo+  o+o      x",
    "x    o+o    +oooooo+  o+o      x",
    "x          +ooo  ooo+          x",
    "x         +ooo    ooo+         x",
    "x         +ooo    ooo+         x",
    "xkkkkkkkk +ooo    ooo+   mmmm  x",
    "xkkkkkkkk +ooo    ooo+  mmmmmm x",
    "xkkkkkkkk +oo     oo+  mmm  mmmx",
    "xkkkkkkkk +oo     oo+  mm    mmx",
    "xkkkkkkkk +oo     oo+  mm    mmx",
    "xkkkkkkkk +oo     oo+  mmm  mmmx",
    "xkkkkkkkooooo     ooooo mmmmmm x",
    "xkkkkkkkooooo     ooooo  mmmm  x",
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}
*/
  bild [einRoboter] = [zellengroesse]string {
    "xxxxxxxxxxxooxxxxxxxxxxxxxxxxxxx",
    "x          oo                  x",
    "x          oo                  x",
    "x          oo ++++             x",
    "x          oo++++++            x",
    "x          o++++++++           x",
    "x      oooo++oo++oo++oooo      x",
    "x        oo+ooo++ooo+oo        x",
    "x         o++oo++oo++o         x",
    "x         ++++++++++++         x",
    "x        ++++oo++oo++++        x",
    "x        +++++oooo+++++        x",
    "x             ++++             x",
    "x       +oo++++++++++oo+       x",
    "x      +ooo+oooooooo+ooo+      x",
    "x      +ooo+oooooooo+ooo+      x",
    "x     +ooo++oooooooo++ooo+     x",
    "x     +oo+ +oooooooo+ +oo+     x",
    "x    +ooo+ +oooooooo+ +ooo+    x",
    "x   +ooo+  ++++++++++  +ooo+   x",
    "x   oooo   oooooooooo   oooo   x",
    "xkkko+o+kkkoooooooooo   +o+o   x",
    "xkkko+o+kkk+oooooooo+   +o+o   x",
    "xkkko+o+kkk+oo++++oo+   +o+om  x",
    "xkkkkkkkkkk+oo+  +oo+   mmmmmm x",
    "xkkkkkkkkkk+oo+  +oo+  mmm  mmmx",
    "xkkkkkkkkkk+oo+  +oo+  mm    mmx",
    "xkkkkkkkkkk+oo+  +oo+  mm    mmx",
    "xkkkkkkkkoooooo  oooooommm  mmmx",
    "xkkkkkkkooooooo  ooooooommmmmm x",
    "xkkkkkkoooooooo  oooooooommmm  x",
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}

  bild [klotz] = [zellengroesse]string {
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    "x                              x",
    "x           oo                 x",
    "x          o  oooo             x",
    "x         o       oooo         x",
    "x        o   mmmmmm   oooo     x",
    "x       o  mmmmmmmmmm     ooo  x",
    "x      o  mmmmm  mmmmm     oo  x",
    "x     o   mmmm     mmmm   o o  x",
    "x    o     mmmmm  mmmmm  o  o  x",
    "x   oo      mmmmmmmmmm  o   o  x",
    "x   oooooo    mmmmmm   o  o o  x",
    "x   oo o ooooo        o     o  x",
    "x   o o o o o oooo   o  o   o  x",
    "x   oo o o o o o oooo     o o  x",
    "x   o o o o o o o o o o     o  x",
    "x   oo o o o o o o oo   o   o  x",
    "x   o o o o o o o o o     o o  x",
    "x   oo o o o o o o oo o     o  x",
    "x   o o o o o o o o o   o   o  x",
    "x   oo o o o o o o oo     o o  x",
    "x   o o o o o o o o o o     o  x",
    "x   oo o o o o o o oo   o  o   x",
    "x   o o o o o o o o o     o    x",
    "x   oo o o o o o o oo o  o     x",
    "x   ooo o o o o o o o   o      x",
    "x     oooo o o o o oo  o       x",
    "x         ooooo o o o o        x",
    "x             oooo ooo         x",
    "x                 ooo          x",
    "x                              x",
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}

  bild [marke] = [zellengroesse]string {
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x            oooooo            x",
    "x          oooooooooo          x",
    "x         ooooo  ooooo         x",
    "x         oooo     oooo        x",
    "x          ooooo  ooooo        x",
    "x           oooooooooo         x",
    "x             oooooo           x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}

  bild [mauer] = [zellengroesse]string {
    "oooooooooooooooooooooooooooooooo",
    "           o               o    ",
    "           o               o    ",
    "           o               o    ",
    "oooooooooooooooooooooooooooooooo",
    "   o               o            ",
    "   o               o            ",
    "   o               o            ",
    "oooooooooooooooooooooooooooooooo",
    "           o               o    ",
    "           o               o    ",
    "           o               o    ",
    "oooooooooooooooooooooooooooooooo",
    "   o               o            ",
    "   o               o            ",
    "   o               o            ",
    "oooooooooooooooooooooooooooooooo",
    "           o               o    ",
    "           o               o    ",
    "           o               o    ",
    "oooooooooooooooooooooooooooooooo",
    "   o               o            ",
    "   o               o            ",
    "   o               o            ",
    "oooooooooooooooooooooooooooooooo",
    "           o               o    ",
    "           o               o    ",
    "           o               o    ",
    "oooooooooooooooooooooooooooooooo",
    "   o               o            ",
    "   o               o            ",
    "   o               o            "}

  farbe = [MaxRobo+1]col.Colour {
    col.FlashRed, col.FlashGreen, col.FlashBlue, col.DarkYellow,
    col.LightCyan, col.FlashMagenta, col.Orange, col.Pink,
    col.Red, col.Green, col.Blue, col.Brown,
    col.Cyan, col.Magenta, col.DarkOrange, col.DeepPink,
    col.LightWhite }

  scr.Switch (scr.WVGApp)
}
