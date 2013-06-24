package gra1

// (c) Christian Maurer   v. 120909 - license see murus.go

// import ("murus/col"; "murus/gra")

// Verwaltet einen Graphen

// Der Graph ist genau dann gerichtet, wenn d == true.
// Seine Ecken können Namen aus bis zu n Zeichen haben.
// Für e == false haben alle Kanten den Wert 0, sonst können seine Kanten natürliche Zahlen < 10^... als Werte haben.
// Er ist der Graph, der in der Datei namens s abgespeichert war.
//  Init (e, d bool, n uint, s string) // n < 22

//  SetColours (n, a col.Colour)

//  Write ()

//  Edit ()

// Im Graphen sind die postaktuelle und die aktuelle Ecke gemäß Benutzereingaben definiert; sie fallen nicht zusammen.
//  NodesSelected () bool

// Im Graphen ist die postaktuelle Ecke gemäß Benutzereingaben definiert; sie fällt mit der aktuellen Ecke zusammen.
//  NodeSelected () bool

// Liefert genau dann TRUE, wenn WHAT ?.
//  Done () bool

//  Set (s graph.Demoset)
