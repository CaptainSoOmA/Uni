/*Die Richtung des Aufzugs wird zu (0 = hoch, 2runter) abstrahiert.

Die eigentliche Implementation des universellen kritischen Abschnitts ist 
aus den Packeten von Maurer genommen. */

package main

import . "murus/cs"
import "fmt"
import "time"
import . "murus/obj"

import "math/rand"

var AutoOben uint // Wieviele Autos sind oben

var nO, nU uint   // wievlie Autos warten 
var aUnten = true // Wo ist der Aufzug
var aZug = false  // ist der Aufzug in benutzung
var x *Imp

// Diese Funktion representiert den Aufzug. Sie flippt, das Aufzug
// Bit und gibt die Arbeit des Aufzugs an.
func call_elevator() {
	if aUnten {
		aUnten = false
		time.Sleep(100 * time.Millisecond)

		fmt.Println(" Aufzug fährt hoch")
	} else {
		aUnten = true
		time.Sleep(100 * time.Millisecond)

		fmt.Println(" Aufzug fährt runter")

	}
}

// --------------------------------------------------------------------
// Universeller Kritischer Abschnitt
// --------------------------------------------------------------------

// Bedingungs Funktion für den Universellen kritischen Abschnitt
func c(k uint) bool {
	var ret bool
	if k == 0 {
		// für k = 0 hochfahren
		ret = !aZug && (aUnten || !x.Blocked(1))
	} else {
		// für k = 1 runterfahren
		ret = !aZug && (!aUnten || !x.Blocked(0))
	}

	return ret
}

// eingangsfunktion
func i(x Any, k uint) {
	aZug = true
	if k == 0 {
		if !aUnten {
			call_elevator()
		}
		AutoOben += 1
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("es sind jetzt %d Autos oben\n", AutoOben)
	} else {
		if aUnten {
			call_elevator()
		}
		AutoOben -= 1
		time.Sleep(100 * time.Millisecond)

		fmt.Printf("es sind jetzt %d Autos oben\n", AutoOben)
	}
}

// Ausgangsfunktion
func o(x Any, k uint) {
	aZug = false
}

// -----------------------------------------------------------------------

func reinfahren() {
	time.Sleep(100 * time.Millisecond)

	fmt.Println("ein Auto will einfahren -------------------------------------")
	x.Enter(0, nil)
	call_elevator()
	x.Leave(0, nil)

}

func rausfahren() {
	time.Sleep(100 * time.Millisecond)

	fmt.Println("ein Auto will ausfahren -------------------------------------")
	x.Enter(1, nil)
	call_elevator()
	x.Leave(1, nil)
}

func Car(t time.Duration) {

	time.Sleep(t)
	reinfahren()
	time.Sleep(t)
	rausfahren()
}

func main() {
	x = New(2, c, i, o)

	rand.Seed(time.Now().Unix())

	// 100 Autos mit zufälliger verweildauer auf dem Parkdach fahren ein

	for i := 0; i <= 10; i++ {
		t := rand.Intn(1000)
		go Car(time.Duration(t))
	}

	time.Sleep(10000)

	Car(time.Duration(rand.Intn(1000)))

}
