package obj

// (c) Christian Maurer   v. 130526 - license see murus.go

type (
  Any interface{} // hides anything

// statements
  Stmt func ()
  StmtSpectrum func (uint)
// conditions
  Cond func () bool
  CondSpectrum func (uint) bool
// operations
  Op func (Any)
  Op3 func (Any, Any, Any)
  OpSpectrum func (Any, uint)
// functions
  Func func (Any) Any
  FuncSpectrum func (Any, uint) Any
  ObjectFunc func (Object) Object
  ObjectFuncSpectrum func (Object, uint) Object
// predicates
  Pred func (Any) bool
  PredSpectrum func (Any, uint) bool
// conditioned operations
  CondOp func (Any, bool)
  CondOp3 func (Any, Any, Any, bool)
// relations
  Rel func (Any, Any) bool
// writings
  Writing func (Any, uint, uint)
  Writing2 func (Any, uint, uint, uint, uint)
)

// Stmt
func Null () { }

// Op[3]
func Null1 (a Any) { }
func Null3 (a, a1, a2 Any) { }

// Func[Spectrum]
func Nil (a Any) Any { return nil }
func NilSp (a Any, i uint) Any { return nil }

// Pred[Spectrum]
func True (a Any) bool { return true }
func TrueSp (a Any, i uint) bool { return true }

// CondOp[3]
func CondNull1 (a Any, b bool) { }
func CondNull3 (a, a1, a2 Any, b bool) { }
