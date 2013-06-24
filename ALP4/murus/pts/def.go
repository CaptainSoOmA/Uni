package pts

// (c) Christian Maurer   v. 121215 - license see murus.go

/*
import (
  . "murus/obj"
  "murus/col"
  "murus/vect"
  "murus/pt"
)
type
  Points interface {

  Clearer
  Persistent

// Die Folge ist die von Benutzer/in per Parameter beim Programmaufruf oder interaktiv ausgewählte, andernfalls leer.
  Select ()

// Die Folge ist mit dem Parameter beim Programmaufruf als Dateiname definiert.
  DefCall ()

// Pre: n > 0; pt.Punkt <= c <= pt.Polygon. Die Folge enthält noch keinen Punkt der Klasse Start.
// Die durch c, a, v[i] und f definierten Punkte mit der Normale (0, 0, 1) sind für 0 <= i < n an die Folge angehängt.
// Der i-te Punkte hat die Nummer a - 1 - i.
  Insert1 (c pt.Class, a uint, v []*vect.Imp, f col.Colour)

// Pre: a > 0. Die Folge enthält noch keinen Punkt der Klasse Start.
// Die durch c, a, v[i], n[i] und f definierten Punkte sind für 0 <= i < a an die Folge angehängt.
// Der i-te Punkt hat die Nummer a - 1 - i.
  Insert (c pt.Class, a uint, v, n []*vect.Imp, f col.Colour)

// Pre: (x, y, z) != (x1, y1, z1). Die Folge enthält noch keinen Punkt der Klasse Start.
// Der Punkt der Klasse pt.Start mit der Nummer 1, den Koordinaten (x, y, z), der Normale (x1, y1, z1)
// in der Farbe schwarz ist an die Folge angehängt;
// entspricht "Insert (pt.Start, 1, v, n, col.Black" mit v = (x, y, z) und n = (x1, y1, z1).
  Start (x, y, z, x1, y1, z1 float64)

// Pre: Der letzte Punkt der Folge gehört zur Klasse pt.Start.
// Liefert die Koordinaten dieses Punktes und die seiner Normale, wenn der letzte Punkt der Folge
// die Klasse Punkt.Start hat, andernfalls (0, 0, 1) und (0, 0, 0).
  StartCoord () (float64, float64, float64, float64, float64, float64)

// Pre: Die Folge ist definiert und ihr Start ist gelesen. Auge ist definiert.
// Alle Punkte der Folge sind der Reihe nach gl zur Ausgabe übergeben. TODO: Details hierzu.
  Write ()
}
*/
