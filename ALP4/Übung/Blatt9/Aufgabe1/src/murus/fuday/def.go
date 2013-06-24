package fuday

// (c) Christian Maurer   v. 130306 - license see murus.go

import
  "murus/day"
// Vorlesungszeiten an der Freien Universität Berlin
// Das Semester eines per func New() *Imp erzeugten Objekts
// ist das Systemdatum des aufrufenden Rechners.

// Das Semester von x ist jetzt dasjenige, in dem der Tag x liegt.
func (x *Imp) Set (d *day.Imp) { x.set(d) }

// b und e sind der erste und letzte Vorlesungstag des Semesters von x.
func (x *Imp) Lectures (b, e *day.Imp) { x.lectures(b,e) }

// Liefert die Kurzbezeichnung des Semesters von x (z.B. "SS 13" oder "WS 13/14").
func (x *Imp) String () string { return x.string_() }

// Liefert genau dann true, wenn d in der Vorlesungszeit
// des Semesters von x liegt und kein Feiertag ist.
func (x *Imp) LectureDay (d *day.Imp) bool { return x.lectureDay(d) }

// Liefert die Anzahl der Vorlesungswochen im Semster von x.
func (x *Imp) NumWeeks () uint { return x.numWeeks() }

// Wenn die n-te Woche (beginnend mit n == 1) ab Beginn der Vorlesungszeit
// in der Vorlesungszeit liegt, ist d der Montag dieser Woche.
// Andernfalls ist d leer.
func (x *Imp) Monday (d *day.Imp, n uint) { x.monday(d,n) }
