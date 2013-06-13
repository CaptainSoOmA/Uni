package main

import "fmt"
import "time"

type Matrix interface {
	mult() *Matrix
	getCol() []uint8
}

type ImpMatrix struct {
	row  [][]uint8
	cols uint8
	rows uint8
}

func New(XSize, YSize uint8) *ImpMatrix {

	// erzeuegen der Arrays
	tmpArray := make([][]uint8, YSize)
	for i := range tmpArray {
		tmpArray[i] = make([]uint8, XSize)
	}

	// füllen mit random werten
	for x := range tmpArray {
		for y := range tmpArray[x] {
			// TODO: random werte einfügen
			tmpArray[x][y] = uint8(x + y*2)
		}
	}

	r := ImpMatrix{tmpArray, XSize, YSize}
	return &r
}

// gibt eine Einzige Spalte als Vektor zurück
func (this ImpMatrix) getCol(c int) *[]uint8 {

	tmpArray := make([]uint8, this.cols)

	for y := range this.row {
		tmpArray[y] = this.row[y][c]
	}

	return &tmpArray
}

func (this ImpMatrix) getRow(c int) *[]uint8 {
	return &this.row[c]
}

func (this ImpMatrix) mult(other ImpMatrix) *ImpMatrix {

	tmp := New(this.cols, other.rows)

	for x := 0; x < int(this.rows); x++ {
		for y := 0; y < int(other.cols); y++ {
			go scalar(*this.getCol(x), *this.getRow(y), tmp, x, y)
		}
	}

	return tmp
}

func scalar(v1, v2 []uint8, mtx *ImpMatrix, posX, posY int) {
	var ret uint8 = 0
	for x := range v1 {
		ret += v1[x] * v2[x]
	}

	mtx.row[posX][posY] = ret
}

func main() {
	m1 := *New(10, 10)
	m2 := *New(10, 10)

	m3 := m1.mult(m2)

	time.Sleep(10000)

	fmt.Println(m3)

}
