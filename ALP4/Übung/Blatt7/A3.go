package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Counter int
	done    chan bool
)

//ÄNDERUNG 3: Sleeptime erhöht (von 1e5);
func v() { time.Sleep(time.Duration(rand.Int63n(1e8))) }

//Simuliert atomares ++
//ÄNDERUNGEN 4:
//v() vor Beginn hinzugefügt (Unterbrechung direkt nach Aufruf möglich)
func inc(n *int) {
	v()
	Accu := *n // "LDA n"
	v()
	Accu++ // "INA"
	v()
	*n = Accu // "STA n"
	v()
	done <- true
}

//ÄNDERUNG 2: N um Zehnerpotenz erhöht (von 5)
func count(p int) {
	const N = 50
	for n := 0; n < N; n++ {
		inc(&Counter)
	}
	done <- true
}

func main() {
	var (
		low, high int
		av        float32
	)

	for i := 0; i < 10000; i++ {
		Counter = 0
		done = make(chan bool)
		go count(0)
		go count(1)
		//ÄNDERUNG 1: Anz. Prozesse erhöht
		go count(2)
		go count(3)
		go count(4)
		go count(5)
		go count(6)
		go count(7)
		<-done
		<-done
		<-done
		<-done
		<-done
		<-done
		<-done
		<-done

		if i == 0 {
			low, high = Counter, Counter
		}

		if Counter < low {
			low = Counter
		}

		if Counter > high {
			high = Counter
		}

		av += float32(Counter)
	}

	fmt.Printf("Kleinster Zählerstand = %d\nHöchster Zählerstand = %d\nDurchschnittlicher Zählerstand: %v\n", low, high, av/10000)

}

/*
Ausgabe Ohne Änderungen:

Kleinster Zählerstand = 0
Höchster Zählerstand = 6
Durchschnittlicher Zählerstand: 1.2954

Ausgabe nach Änderung 1:

Kleinster Zählerstand = 0
Höchster Zählerstand = 5
Durchschnittlicher Zählerstand: 1.1207

Ausgabe nach Änderung 1 u. 2:

Kleinster Zählerstand = 0
Höchster Zählerstand = 5
Durchschnittlicher Zählerstand: 1.1317

Ausgabe nach Änderung 1, 2 u. 3:

Kleinster Zählerstand = 0
Höchster Zählerstand = 14
Durchschnittlicher Zählerstand: 1.2031

Ausgabe nach Änderung 1, 2, 3 u. 4:

Kleinster Zählerstand = 0
Höchster Zählerstand = 55
Durchschnittlicher Zählerstand: 12.2102

*/