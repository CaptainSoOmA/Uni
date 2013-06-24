package rob

// (c) Christian Maurer   v. 121204 - license see murus.go

import (
  . "murus/obj"; . "murus/kbd"
)
type
  Richtung byte; const (
  Nord = Richtung(iota)
  West
  Sued
  Ost
  nRichtungen
)
type Aktion byte; const (
  Links = iota; Rechts; Weiter; Zurueck;
  KlotzHin; KlotzWeg; KlotzWeiter
  MarkeHin; MarkeWeg; MauerHin; MauerWeg
  nAktionen )
var
  Aktionstext [nAktionen]string
type
  Klotzzahl uint16 // <= 2 * maxK + 2
const
  MaxK = Klotzzahl(999)
type
  Roboter interface {

// Wenn schon maximal viele Roboter initialisiert sind, ist nichts verändert.
// Andernfalls gilt: Der neue Roboter hat die kleinste verfügbare Nummer;
// er steht in Richtung Süd auf dem ersten freien Platz von der nordwestlichen Ecke aus in Schreibrichtung gesehen;
// er hat 999 Klötze - ggf. abzgl. seiner Klötze in der Roboterwelt - in seiner Tasche.
//  NeuerRoboter ()

// Unter R wird im folgenden immer der aufrufende Roboter verstanden.

// R ist nicht initialisiert.
  Terminieren ()

// Liefert die Nummer von R.
  Nummer () uint

// Liefert genau dann true, wenn R in der nordwestlichen Ecke steht.
  InLinkerObererEcke () bool

// Rs Richtung ist um 90 Grad nach links gedreht.
  LinksDrehen ()

// Rs Richtung ist um 90 Grad nach rechts gedreht.
  RechtsDrehen ()

// Liefert genau dann true, wenn Rs Platz in Rs Richtung keinen Nachbarplatz hat.
  AmRand () bool

// Liefert genau dann true, wenn Rs vorheriger Platz in Rs Richtung einen Nachbarplatz hat,
// der nicht zugemauert ist und auf dem vorher kein Robo stand.
// In diesem Fall steht R jetzt in der gleichen Richtung wie vorher auf diesem Platz; andernfalls ist nichts verändert.
  Gelaufen1 () bool

// Liefert genau dann true, wenn Rs vorheriger Platz entgegen Rs Richtung einen Nachbarplatz hat,
// der nicht zugemauert ist und auf dem kein Robo steht.
// In diesem Fall steht R jetzt in der gleichen Richtung wie vorher auf diesem Platz; andernfalls ist nichts verändert.
  Zurueckgelaufen1 () bool

// Liefert genau dann true, wenn auf Rs Platz kein Klotz liegt.
  Leer () bool

// Wenn auf Rs Platz nicht mindestens ein Klotz liegt, den R dort hingelegt hat, ist nichts verändert.
// Andernfalls liegt auf Rs Platz jetzt ein Klotz weniger als vorher und in Rs Tasche ist jetzt einer mehr als vorher.
  Leeren1 ()

// Liefert genau dann true, wenn Rs Platz in Rs Richtung einen Nachbarplatz hat
// und auf diesem Platz weder ein Robo steht noch ein Klotz liegt.
  NachbarLeer () bool

// Liefert genau dann true, wenn Rs Tasche nicht leer ist.
  HatKloetze () bool

// R hat genau n Klötze in der Tasche.
  KloetzeGeben (n uint)

// Liefert die Anzahl der Klötze in Rs Tasche.
  AnzahlKloetze () uint

// Liefert genau dann true, wenn auf Rs Platz keine Klötze liegen, die ein anderer Robo dort hingelegt hat,
// und dieser Platz nicht von einem anderen Robo markiert ist.
  DarfLegen () bool

// Wenn Rs Tasche leer ist oder R auf den Platz, auf dem R steht, nicht legen darf, ist nichts verändert.
// Andernfalls liegt auf Rs Platz jetzt ein Klotz mehr als vorher und in Rs Tasche ist einer weniger.
  Legen1 ()

// Liefert genau dann true, wenn Rs vorheriger Platz in Rs Richtung einen Nachbarplatz hat,
// auf dem mindestens ein Klotz liegt, den R dort hingelegt hat,
// und dieser Platz wiederum in Rs Richtung einen Nachbarplatz hat, der nicht zugemauert ist,
// auf dem kein anderer Robo steht und der leer und nicht von einem anderen Robo markiert ist.

// In diesem Fall steht R jetzt in der gleichen Richtung wie vorher auf seinem vorigen Nachbarplatz in Rs Richtung
// und die Klötze, die vorher auf diesem Platz lagen, liegen jetzt auf dem jetzigen Nachbarplatz in Rs Richtung;
// andernfalls ist nichts verändert.
  Geschoben1 () bool

// Wenn auf Rs Platz keine Klötze liegen, ist nichts verändert.
// Andernfalls liegen alle Klötze, die dort gelegen haben, jetzt auf dem letzten leeren Platz in Rs Richtung,
// der weder markiert noch zugemauert ist, und auf dem kein anderer Robo steht.
  Schiessen ()

// Liefert genau dann true, wenn Rs Platz markiert ist.
  Markiert () bool

// Liefert genau dann true, wenn Rs Platz nicht von einem anderen Robo markiert ist
// und auf ihm keine Klötze liegen, die ein anderer Robo dort hingelegt hat.
  DarfMarkieren () bool

// Wenn R markieren darf, ist Rs Platz jetzt von ihm markiert; andernfalls ist nichts verändert.
  Markieren ()

// Wenn R markieren darf, ist Rs Platz jetzt nicht mehr markiert; andernfalls ist nichts verändert.
  Entmarkieren ()

// Liefert genau dann true, wenn Rs Platz in Rs Richtung einen Nachbarplatz hat, der markiert ist.
  NachbarMarkiert () bool

// Liefert genau dann true, wenn Rs Platz in Rs Richtung einen Nachbarplatz hat, der zugemauert ist.
// Der Wert kann sich unmittelbar nach dem Aufruf geändert haben, wenn außer R noch weitere Robos initialisiert sind.
  VorMauer () bool

// Liefert genau dann true, wenn Rs vorheriger Platz in Rs Richtung einen Nachbarplatz hat,
// der nicht zugemauert ist und auf dem kein anderer Robo steht. In diesem Fall gilt:
// Rs Platz ist jetzt zugemauert.
// R steht jetzt in der gleichen Richtung wie vorher auf dem vorherigen Nachbarplatz in Rs Richtung;
// wenn auf dem Platz, auf dem R vorher gestanden hat, vorher Klötze gelegen haben, liegen sie jetzt dort nicht mehr,
// sondern sind in der Tasche des Robos, der sie dort hingelegt hat.
// Eine vorher dort etwa vorhandene Markierung ist jetzt entfernt.
// Andernfalls ist nichts verändert.
  Gemauert1 () bool

// Liefert genau dann true, wenn Rs vorheriger Platz in Rs Richtung einen Nachbarplatz hat, der zugemauert ist.
// In diesem Fall steht R in der gleichen Richtung wie vorher auf diesem Platz und sein vorheriger Platz
// ist jetzt nicht mehr zugemauert; andernfalls ist nichts verändert.
  Entmauert1 () bool

// Der Zustand von R und der Roboterwelt ist entsprechend (K, T) verändert.
// Details TODO
  Manipulieren (c Comm, t uint)

  Coder // s. murus/obj/coder.go

// Die Klötze in Rs Welt und eventuelle Mauersteine liegen auf den von Benutzer/in festgelegten Plätzen.
// Wenn R auf einem Platz steht, auf dem Klötze liegen, wird deren Anzahl angezeigt.
// Rs Platz und Richtung, die Anzahl der Klötze in seiner Tasche und der Klötze auf allen Plätzen
// sind beim nächsten Programmlauf mit dieser Welt die gleichen wie beim Aufruf dieser Prozedur.
  Editieren ()

// Für ein = true ist das Verhalten des Editors gemäß den Anforderungen an das Spiel Sokoban vereinfacht.
  SokobanSchalten (ein bool)

// n ist in der untersten Bildschirmzeile ausgegeben.
// Der aufrufende Prozess war danach solange angehalten, bis Benutzer/in die Ausgabe mit <Esc> quittiert hatte.
  Ausgeben (n uint)

// Liefert die Zahl, die von der benutzenden Person in der untersten Bildschirmzeile eingegeben wurde.
  Eingabe () uint

// T und n sind in einer Zeile am unteren Bildschirmrand ausgegeben.
// Der aufrufende Prozess war danach solange angehalten, bis Benutzer/in die Ausgabe mit <Esc> quittiert hatte.
// Jetzt ist die Meldung wieder vom Bildschirm entfernt.
  FehlerMelden (T string, n uint)
}
