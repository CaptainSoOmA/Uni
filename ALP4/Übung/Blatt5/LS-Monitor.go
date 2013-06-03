package ReaderWriterProblem

import . "sync"

//Loesung moeglichst dicht an [MAURER] S.137 f. wobei wie auch auf S.135 von 
//signal-and-continue ausgegangen (also das mutex erst nach dem signal freigegeben) wird

//'go build LS-Monitor.go' ohne Fehler

type Imp struct {
	nR, nW, rCnt, wCnt uint
	mutex              Mutex
	okR, okW           *Cond
}

func New() *Imp {
	x := new(Imp)
	x.okR = NewCond(&x.mutex)
	x.okW = NewCond(&x.mutex)
	return x
}

//Schreiber mit Prioritaet
func (x *Imp) ReaderIn() {
	x.mutex.Lock()
	//Awaited() um auf blockierte Reader/Writer zu pruefen gibts hier nicht, 
	//daher zaehlen rCnt bzw. wCnt mit
	if x.nW > 0 || x.wCnt > 0 {
		x.rCnt++
		x.okR.Wait()
		x.rCnt--
	}
	x.nR++
	x.mutex.Unlock()
}

func (x *Imp) ReaderOut() {
	x.mutex.Lock()
	x.nR--
	if x.nR == 0 {
		x.okW.Signal()
	}
	x.mutex.Unlock()
}

func (x *Imp) WriterIn() {
	x.mutex.Lock()
	if x.nR > 0 || x.nW > 0 {
		x.wCnt++
		x.okW.Wait()
		x.wCnt--
	}
	x.nW = 1
	x.mutex.Unlock()
}

//Priorisiert Reader
func (x *Imp) WriterOut() {
	x.mutex.Lock()
	x.nW = 0
	if x.rCnt > 0 {
		x.okR.Signal()
	} else {
		x.okW.Signal()
	}
	x.mutex.Unlock()
}