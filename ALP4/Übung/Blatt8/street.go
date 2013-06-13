/*


*/
package main

import (
	"fmt"
)

type Strasse interface {
	auffahren() bool
	ausfahren() bool
}

type ImplStrasse struct {
}

func (this ImplStrasse) auffahren(richtung, groese string) chan bool {
	retChan := make(chan bool)
	return retChan
}

func main() {
	fmt.Println("main")
}
