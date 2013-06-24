package rob1

// (c) Christian Maurer   v. 120318 - license see murus.go

/* Verwaltet eine rechteckige Welt aus schachbrettförmig angeordneten Plätzen,
   auf denen Klötze liegen oder die zugemauert sein können, und den Roboter.
   Die Welt ist 32 Plätze breit (West-Ost) und 16 Plätze hoch (Nord-Süd).
   Von einem Platz aus in einer der Himmelsrichtungen (Nord, West, Süd, Ost) gesehen
   gibt es entweder genau einen Nachbarplatz oder keinen.

   Die Roboterwelt ist die mit dem bei Programmbeginn eingegebenen Namen.

   Der Roboter steht immer auf einem der Plätze, der als "Roboters Platz" bezeichnet wird.
   Er steht immer in genau einer der vier Himmelsrichtungen, die als "Roboters Richtung" bezeichnet wird.

   Der Roboter hat eine Tasche mit Klötzen und jederzeit Zugriff auf Mauersteine.

   Der Platz, die Richtung und die Anzahl der Klötze in der Tasche des Roboters sind ggf. die gleichen
   wie beim letzten Programmlauf mit dieser Welt; wenn die Welt neu ist, ist sie leer
   und der Roboter steht in Richtung Süd in der nordwestlichen Ecke und hat 999 Klötze
   - ggf. abzgl. derjenigen, die er schon in die Welt gelegt hat - in seiner Tasche.

   Anfangs ist das Protokoll nicht eingeschaltet. */

// Liefert genau dann true, wenn der Roboter in der nordwestlichen Ecke steht.
//  InLinkerObererEcke () bool

// Roboters Richtung ist um 90 Grad nach links gedreht.
//  LinksDrehen ()

// Roboters Richtung ist um 90 Grad nach rechts gedreht.
//  RechtsDrehen ()

// Liefert genau dann true, wenn Roboters Platz in Roboters Richtung keinen Nachbarplatz hat.
//  AmRand () bool

// Vor.: Roboters Platz hat in Roboters Richtung einen Nachbarplatz, der nicht zugemauert ist.
// Der Roboter steht in der gleichen Richtung wie vorher auf diesem Nachbarplatz.
//  Laufen1 ()

// Vor.: Roboters Platz hat entgegen Roboters Richtung einen Nachbarplatz, der nicht zugemauert ist.
// Der Roboter steht in der gleichen Richtung wie vorher auf diesem Nachbarplatz.
//  Zuruecklaufen1 ()

// Liefert genau dann true, wenn auf Roboters Platz kein Klotz liegt.
//  Leer () bool

// Vor.: Auf Roboters Platz liegt mindestens ein Klotz.
// Auf Roboters Platz liegt ein Klotz weniger als vorher, in seiner Tasche ist einer mehr.
//  Leeren1 ()

// Liefert genau dann true, wenn Roboters Platz in Roboters Richtung einen Nachbarplatz hat
// und auf diesem Platz kein Klotz liegt.
//  NachbarLeer () bool

// Liefert genau dann true, wenn die Tasche des Roboters nicht leer ist.
//  HatKloetze () bool

// Vor.: n < 1000.
// Der Roboter hat genau n Klötze in der Tasche.
//  KloetzeGeben (n uint)

// Liefert die Anzahl der Klötze in der Tasche des Roboters.
//  AnzahlKloetze () uint

// Vor.: Die Tasche des Roboters ist nicht leer. 
// Auf Roboters Platz liegt ein Klotz mehr als vorher, in seiner Tasche ist einer weniger.
//  Legen1 ()

// Vor.: Roboters Platz hat in Roboters Richtung einen Nachbarplatz,
//       auf dem mindestens ein Klotz liegt. Dieser Nachbarplatz hat wiederum einen
//       Nachbarplatz in Roboters Richtung, der leer und nicht zugemauert ist.
// Der Roboter steht in der gleichen Richtung wie vorher auf dem vorherigen
// Nachbarplatz in seiner Richtung und die Klötze, die vorher auf ihm
// lagen, liegen jetzt auf dem jetzigen Nachbarplatz in Roboters Richtung.
//  Schieben1 ()

// Wenn auf Roboters Platz keine Klötze liegen, ist nichts verändert.
// Andernfalls liegen alle Klötze, die dort gelegen haben, jetzt
// auf dem letzten leeren Platz in Roboters Richtung, der weder
// markiert noch zugemauert ist, und auf dem kein anderer Roboter steht.
//  Schiessen ()

// Liefert genau dann true, wenn Roboters Platz markiert ist.
//  Markiert () bool

// Roboters Platz ist markiert.
//  Markieren ()

// Roboters Platz ist nicht markiert.
//  Entmarkieren ()

// Liefert genau dann true, wenn Roboters Platz in Roboters Richtung einen Nachbarplatz hat, der markiert ist.
//  NachbarMarkiert () bool

// Liefert genau dann true, wenn Roboters Platz in Roboters Richtung einen Nachbarplatz hat, der zugemauert ist.
//  VorMauer () bool

// Vor.: Roboters Platz hat in Roboters Richtung einen Nachbarplatz, der nicht zugemauert ist.
// Der Roboter steht in der gleichen Richtung wie vorher auf diesem Nachbarplatz.
// Wenn auf dem Platz, auf dem der Roboter vorher gestanden hat, vorher Klötze gelegen haben,
// liegen sie dort jetzt nicht mehr, sondern sind in seiner Tasche; dafür ist dieser Platz jetzt zugemauert.
// Eine vorher dort etwa vorhandene Markierung ist jetzt entfernt.
//  Mauern1 ()

// Vor.: Roboters Platz hat in Roboters Richtung einen Nachbarplatz, der zugemauert ist.
// Der Roboter steht in der gleichen Richtung wie vorher auf diesem Nachbarplatz
// und dieser Platz ist jetzt nicht mehr zugemauert.
//  Entmauern1 ()

// Die Klötze in der Roboterwelt und eventuelle Mauersteine liegen auf den von Benutzer/in festgelegten Plätzen.
// Wenn der Roboter auf einem Platz steht, auf dem Klötze liegen, wird deren Anzahl angezeigt.
// Roboters Platz und Richtung, die Anzahl der Klötze in seiner Tasche und der Klötze auf allen Plätzen
// sind beim nächsten Programmlauf mit dieser Welt die gleichen wie beim Aufruf dieser Methode.
// Wenn das Protokoll eingeschaltet ist, ist der Editiervorgang in einem Go-Quelltext
// (unter dem Namen der Roboterwelt mit dem Suffix ".go") protokolliert.
// Das aus diesem Quelltext durch Übersetzung erzeugte Programm simuliert schrittweise den Editiervorgang.
//  Editieren ()

// Das Protokoll ist genau dann eingeschaltet, wenn ein = true (siehe editieren).
//  ProtokollSchalten (ein bool)

// Für ein = true ist das Verhalten des Editors gemäß den Anforderungen an das Spiel Sokoban vereinfacht.
//  SokobanSchalten (ein bool)

// n ist in der untersten Bildschirmzeile ausgegeben.
// Der aufrufende Prozess war danach solange angehalten, bis Benutzer/in die Ausgabe mit <Esc> quittiert hatte.
//  Ausgeben (n uint)

// Liefert die Zahl, die von der benutzenden Person in der untersten Bildschirmzeile eingegeben wurde.
//  Eingabe() uint

// s und n sind in einer Zeile am unteren Bildschirmrand ausgegeben.
// Der aufrufende Prozess war danach solange angehalten, bis Benutzer/in die Ausgabe mit <Esc> quittiert hatte.
// Jetzt ist die Meldung wieder vom Bildschirm entfernt.
//  FehlerMelden (s string, n uint)

// Das Programm ist mit der Fehlermeldung ("Programm beendet") angehalten.
// Fertig()
