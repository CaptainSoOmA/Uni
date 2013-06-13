// Definition des Buches fuer die FetchAndAdd Funktion
// inertn realisiert mit XADD
FetchAndAdd(k,n) = AddUint32(k,n) - n 

var interested uint 	// Wert fuer die aktuelle Warteschlange
var turn uint 			// wert in der Wartschleife des Aktuellen Prozesses

func Lock() {
	// "zieht" die naechst hoehere Wartenummer und wartet bis er an der Tuer ist
	myTurn = FetchAndAdd(*interested,1)
	for myTurn != turn {Null()}
}

func Unlock(){
	// beim Austritt die den aktuellen Prozess erhoehen
	FetchAndAdd(*turn,1)
}
