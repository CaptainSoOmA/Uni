package pseq

// (c) Christian Maurer   v. 130118 - license see murus.go

// >>> still lots of things TODO

import (
  . "murus/ker"; . "murus/obj"
  "murus/str"
//  "murus/seq"
)

const (
  pack = "pseq"
  null = uint64(0)
  one  = uint64(1)
)
type
  Imp struct {
        name,
    tmpName string
 emptyObject,
      object Any
        file *files
       owner,
       group uint
        size,
         pos,
         num uint64
   buf, buf1 []byte
     ordered bool
             }
//var
//  filenames *seq.Imp
/*
  Among others the following problems are not yet solved:
  1. Not more than 1 psequence must be (re-)named with the same name.
     Help: Put names at (re-)defining into the sequence "filenames" and
           remove them at terminating.
  2. Access to psequences is only possible, if the rights are named correspondingly.
     At the moment clients are not protected from trying to access psequences
     without having the rights to.
  3. The following trivial handling of read/write-errors should be replaced by a better concept.
*/

func (x *Imp) check (a Any) {
//
  if ! TypeEq (x.emptyObject, a) { NotCompatiblePanic () }
}


func (x *Imp) imp (a Any) *Imp {
//
  y, ok:= a.(*Imp)
  x.check (y.emptyObject)
  if ! ok || x.size != y.size {
    TypeNotEqPanic (x, a)
  }
  if x.file == nil || y.file == nil { Panic ("pseq-error: file = nil") }
  if x == y { Panic ("pseq error: x == y") }
  return y
}


var
  wasRead, wasWritten uint


func (x *Imp) read (b []byte) {
//
  r:= x.file.read (b[0:x.size])
  wasRead = r
}


func (x *Imp) write (b []byte) {
//
  w:= x.file.write (b[0:x.size])
  if uint64(w) < x.size {
    wasWritten = w
  }
}


func New (a Any) *Imp {
//
  PanicIfNotOk (a)
  x:= new (Imp)
  x.emptyObject = Clone (a)
  x.object = Clone (a)
  x.num = null
  x.size = uint64(Codelen (a))
  x.file = newFile()
  x.buf = make ([]byte, x.size)
  x.buf1 = make ([]byte, x.size)
  x.ordered = false
  return x
}


func (x *Imp) Terminate () { // not in Def !
//
///*
//  n:= str.Length (x.name)
//  if filenames.Ex (x.name, n) {
//    filenames.Del ()
//  } else {
//    Fehler
//  }
///*
  x.file.terminate ()
}


func Length (N string) uint { // < -- uint64 !
//
  return uint(directLength (N))
}


func Erase (N string) {
//
  erase (N)
}

/*
func accessible (Name string, Zugriff Zugriffe) bool {
//
  return file.accessible (Name, VAL (Zugriffe, ORD (Zugriff)))
}
*/

func (x *Imp) Name (N string) {
//
//  if ! files.Defined (N) { Fehler }
  x.name = N
//  str.DelSpaces (&x.name)
//  n:= str.Length (x.name)
//  if filenames.Ex (x.name, n) {
//    // Fehlersituation, siehe oben Bemerkung 1.
//    Fehler
//  }
  x.file.name (x.name)
//    $USER 
  x.pos = 0
  x.num = x.file.length () / x.size
  x.tmpName = x.name + "-tmp"
//  tmpName.Temporieren ()
}


func (x *Imp) Rename (n string) {
//
  if str.Empty (n) || n == x.name {
    return
  }
  x.name = n
  x.file.rename (x.name)
}


func (x *Imp) Empty () bool {
//
  if str.Empty (x.name) { return true }
  return x.file.empty()
}


func (x *Imp) Clr () {
//
  if x.file == nil { Panic ("pseq.Clr: file = nil") }
  x.file.clear ()
  x.pos = 0
  x.num = 0
}


func equal (as, bs []byte) bool {
//
  if len (as) != len (bs) { return false }
  for i, a:= range (as) {
    if a != bs[i] { return false }
  }
  return true
}


func (x *Imp) e (y *Imp, r Rel) bool {
//
  if y.name == x.name { return true }
  if x.num != y.num { return false }
  for i:= null; i < x.num; i ++ {
    x.read (x.buf)
    y.read (y.buf)
    if ! r (x.buf, y.buf) {
      return false
    }
  }
  return true
}


func (x *Imp) Eq (Y Object) bool {
//
  return x.e (x.imp (Y), Eq)
}


func (x *Imp) leq (Y Object) bool { // TODO
//
  y:= x.imp (Y)
  if y.name == x.name { return true }
  if x.num != y.num { return false }
  for i:= null; i < x.num; i ++ {
    x.read (x.buf)
/*
    for x.pos < x.num {
      for {
        x1.read (x1.buf)
        if equal (x.buf, x1. buf)
          continue
        } else {
        }
      }
    if ! equal (x.buf, x1.buf) {
      return false
    }
*/
  }
  return true
}


func (x *Imp) Less (Y Object) bool { // TODO
//
  y:= x.imp (Y)
  if y.name == x.name { return false }
  if x.num == y.num { return false }
  return x.leq (Y)
}


func (x *Imp) Equiv (Y Iterator, r Rel) bool {
//
  return x.e (x.imp (Y), r)
}


func (x *Imp) Num () uint {
//
  return uint(x.file.length() / x.size)
//  return uint(x.num)
}


func (x *Imp) NumPred (p Pred) uint {
//
  n:= Zero
  if x.num == 0 { return 0 }
  x.file.seek (0)
  for i:= null; i < x.num; i ++ {
    x.read (x.buf)
    if p (x.buf) {
      n ++
    }
  }
  return n
}


func (x *Imp) Ex (a Any) bool {
//
  x.check (a)
  if x.num == 0 { return false }
  x.file.seek (0)
  for i:= null; i < x.num; i ++ {
    x.read (x.buf)
    if equal (x.buf, Encode (a)) {
      x.pos = i
      return true
    }
  }
  return false
}


func (x *Imp) Step (forward bool) {
//
  if forward {
    if x.pos * x.size < x.file.length () {
      x.pos ++
    }
  } else if x.pos > 0 {
    x.pos --
  }
}


func (x *Imp) Seek (n uint) {
//
  x.pos = uint64(n)
}


func (x *Imp) Jump (forward bool) {
//
  if forward {
    x.Seek (uint(x.num))
  } else {
    x.Seek (0)
  }
}


func (x *Imp) Offc () bool {
//
  return x.pos * x.size == x.file.length ()
}


func (x *Imp) Eoc (forward bool) bool {
//
  if forward {
    return (x.pos + 1) * x.size == x.file.length ()
  }
  return x.pos == 0
}


func (x *Imp) Pos () uint {
//
  return uint(x.pos)
}


func (x *Imp) Get () Any {
//
  x.file.seek (x.pos * x.size)
  if x.file.position() != x.pos * x.size { Stop (pack, 10000000 + uint(x.pos)) }
  x.read (x.buf)
  xx:= Decode (x.object, x.buf)
  return Clone (xx)
}


func (x *Imp) Put (a Any) {
//
  x.check (a)
  x.file.seek (x.pos * x.size)
  x.write (Encode (a))
}


func (x *Imp) insert (a Any) {
//
  if x.pos >= x.num {
    x.pos = x.num
    x.file.seek (x.file.length ())
    x.write (Encode (a))
    x.pos ++
    x.num ++
    return
  }
// x.pos < x.num:
  x1:= New (x.emptyObject)
  x1.Name (x.tmpName)
  x1.Clr()
  x.file.seek (0)
  if x.pos > 0 {
    for i:= null; i < x.pos; i++ {
      x.read (x.buf)
      x1.write (x.buf)
    }
  }
  x1.write (Encode (a))
  if x.pos < x.num {
    for i:= x.pos; i < x.num; i++ {
      x.read (x.buf)
      x1.write (x.buf)
    }
  }
  x.pos ++
  x.num ++
  n:= x.num
  p:= x.pos
  x.file.clear ()
  x1.file.rename (x.name)
  x1.file.flush ()
  x1.Terminate ()
  x.file.name (x.name)
  x.pos = p
  x.num = n // == x.file.length() / x.size
}


func (x *Imp) insertOrd (a Any) {
//
  ps:= New (x.emptyObject)
  ps.Name (x.tmpName)
  ps.Clr()
  x.file.seek (0)
  i:= null
  n:= x.num
  inserted:= false
  p:= null
  code:= Encode (a)
  for {
    if i == x.num {
      if ! inserted {
        p = i
        ps.write (code)
      }
      break
    }
    x.read (x.buf)
    if ! inserted {
      if Less (code, x.buf) {
        p = i
        ps.write (code)
        inserted = true
      }
      if ! inserted {
        if ! Less (x.buf, code) {
          inserted = true
        }
      }
    }
    ps.write (x.buf)
    i ++
  }
  x.file.clear()
  ps.file.rename (x.name)
  ps.file.flush ()
  ps.Terminate ()
  x.file.name (x.name)
  x.num = x.file.length() / x.size
  if x.num != n + 1 {
    // noch untersuchen
  }
  x.pos = p + 1
}


func (x *Imp) Ins (a Any) {
//
  x.check (a)
  if x.ordered {
    x.insertOrd (a)
  } else {
    x.insert (a)
  }
}


func (x *Imp) Del () Any {
//
  if x.num == 0 || x.pos >= x.num {
    return nil
  }
  n:= x.num
  x1:= New (x.emptyObject)
  x1.Name (x.tmpName)
  x1.Clr()
  x.file.seek (0)
  var a Any
  for i:= null; i < x.num; i++ {
    x.read (x.buf)
    if i == x.pos {
      a = Decode (x.object, x.buf)
    } else {
      x1.write (x.buf)
    }
  }
  if x.pos == x.num - 1 && x.pos > 0 {
    x.pos --
  }
  p:= x.pos
  x.file.clear ()
  x1.file.rename (x.name)
  x1.file.flush ()
  x1.Terminate()
  x.file.name (x.name)
  x.pos = p
  x.num = x.file.length() / x.size // x.num --
  if x.num + 1 != n {
// errh.Error2 ("what to devil", uint(x.num + 1), "is here loose", uint(n))
  }
  return a
}


func (x *Imp) ExPred (p Pred, f bool) bool {
//
  if x.file.empty () { return false }
  n:= x.file.length () / x.size
  if n == 0 { return false }
  i:= null
  if f {
    i = 0
  } else {
    i = n - 1
  }
  x.file.seek (i * x.size)
  for {
    x.read (x.buf)
    if p (Decode (x.emptyObject, x.buf)) {
      x.pos = i
      return true
    }
    if f {
      if i == n - 1 {
        break
      } else {
        i ++
      }
    } else if i == 0 {
      break
    } else {
      i --
    }
  }
  return false
}


func (x *Imp) StepPred (p Pred, f bool) bool {
//
  n:= x.file.length () / x.size
  if n <= 1 { return false }
  if f && x.pos == n - 1 { return false }
  if ! f && x.pos == 0 { return false }
  i:= null
  if x.pos == n {
    if f {
      i = 0
    } else {
      i = n - 1
    }
  } else {
    i = x.pos
    if f {
      i ++
    } else {
      i --
    }
  }
  for {
    x.file.seek (i * x.size)
    x.read (x.buf)
    if p (Decode (x.emptyObject, x.buf)) {
      x.pos = i
      break
    }
    if f {
      if i == n - 1 {
        break
      } else {
        i ++
      }
    } else {
      if i == 0 {
        break
      } else {
        i --
      }
    }
  }
  return false
}


func (x *Imp) All (p Pred) bool {
//
  if x.num == 0 { return true }
  x.file.seek (0)
  for i:= null; i < x.num; i++ {
    x.read (x.buf)
    if ! p (Decode (x.emptyObject, x.buf)) {
      return false
    }
  }
  return true
}


func (x *Imp) Ordered () bool {
//
  if x.num <= 1 { return true }
  x.file.seek (0)
  x.read (x.buf)
  for i:= one; i < x.num; i++ {
    x.read (x.buf1)
    if ! Less (x.buf, x.buf1) && ! Eq (x.buf, x.buf1) {
      return false
    }
    copy (x.buf, x.buf1)
    i ++
  }
  return true
}


func (x *Imp) Sort () {
//
// TODO
}


func (x *Imp) Trav (op Op) {
//
  b:= x.file.length() == 0
  if b {
    if x.num != 0 || ! x.Empty () { println ("Trav oops") }
  }
  x.file.seek (0)
  for i:= null; i < x.num; i++ {
    x.read (x.buf)
    if uint64(wasRead) < x.size {
      copy (x.buf, Encode (x.emptyObject)) // provisorisch
    }
    x.object = Decode (x.emptyObject, x.buf)
    op (x.object)
    if ! equal (x.buf, Encode (x.object)) {
      copy (x.buf, Encode (x.object))
      x.file.seek (i * x.size)
      x.write (x.buf)
      x.file.seek (i * x.size)
    }
  }
  x.file.flush ()
}


func (x *Imp) TravCond (p Pred, op CondOp) {
//
  if x.num == 0 { return }
  if x.Empty () { return }
  x.file.seek (0)
  for i:= null; i < x.num; i++ {
    x.read (x.buf)
    a:= Decode (Clone (x.emptyObject), x.buf)
    op (a, p (a))
    if ! equal (x.buf, Encode (a)) {
      copy (x.buf, Encode (a))
      x.file.seek (i * x.size)
      x.write (x.buf)
      x.file.seek (i * x.size)
    }
  }
  x.file.flush ()
}


func (x *Imp) TravPred (p Pred, op Op) {
//
  if x.num == 0 { return }
  if x.Empty () { return }
  x.file.seek (0)
  for i:= null; i < x.num; i++ {
    x.read (x.buf)
    a:= Decode (x.emptyObject, x.buf)
    if p (a) {
      op (a)
      if ! equal (x.buf, Encode (a)) {
        copy (x.buf, Encode (a))
        x.file.seek (i * x.size)
        x.write (x.buf)
        x.file.seek (i * x.size)
      }
    }
  }
  x.file.flush ()
}


func (x *Imp) Filter (C Iterator, p Pred) {
//
  x1:= x.imp (C)
  if x1 == nil { return }
  if x.num == 0 { return }
  x.file.seek (0)
  x1.Clr()
  x1.pos = 0
  for i:= null; i < x.num; i++ {
    x.read (x.buf)
    if p (Decode (Clone (x.emptyObject), x.buf)) {
      x1.write (x.buf)
      x1.pos ++
    }
  }
  x1.file.flush ()
}


func (x *Imp) Cut (c Iterator, p Pred) {
//
  x1:= x.imp (c)
  if x1 == nil { return }
  x1.Clr()
  if x.name == x1.name { return }
  x2:= New (x.emptyObject)
  x2.Name (x.tmpName)
  x2.Clr()
  x.file.seek (0)
  x.pos = 0
  for i:= null; i < x.num; i++ {
    x.read (x.buf)
    if p (Decode (Clone (x.emptyObject), x.buf)) {
      x1.write (x.buf)
      x1.pos ++
    } else {
      x2.write (x.buf)
      x.pos ++
    }
  }
  x.file.clear ()
  x2.file.rename (x.name)
  x2.file.flush ()
  x2.Terminate()
  x.file.name (x.name)
  x1.file.flush ()
}


func (x *Imp) ClrPred (p Pred) {
//
  x1:= New (x.emptyObject)
  if x1 == nil { return }
  if x.num == 0 { return }
  x.file.seek (0)
  n:= x.pos
  x1.Clr()
  for i:= null; i < x.num; i++ {
    x.read (x.buf)
    if p (Decode (Clone (x.emptyObject), x.buf)) {
      if n == i {
        n ++
      }
    } else {
      x1.write (x.buf)
      x1.num ++
    }
  }
  x1.file.flush ()
  x1.Terminate ()
  x1.pos = n
  x.file.name (x.name)
}


func (x *Imp) Split (C Iterator) {
//
  x1:= x.imp (C)
  if x1 == nil { return }
  if x.num == 0 { return }
  x1.Clr()
  ps:= New (x.emptyObject)
  ps.Name (x.tmpName)
  ps.Clr()
  x.file.seek (0)
  if x.pos == 0 {
//    errh.ReportError ("pseq: Split not yet completely implemented") // >>>> alles nach S1
  } else {
    for i:= null; i < x.pos; i++ {
      x.read (x.buf)
      ps.write (x.buf)
    }
    if x.pos < x.num {
      for i:= one; i <= x.num - x.pos; i++ {
        x.read (x.buf)
        x1.write (x.buf)
      }
    }
    x1.pos = 0
  }
  x.file.clear()
  ps.file.rename (x.name)
  ps.Terminate()
  x.file.name (x.name)
  x.pos = x.num - x.pos - 1
  x.num = x.file.length() / x.size
  x1.num = x1.file.length() / x1.size
  x1.file.flush ()
}


func (x *Imp) concatenate (x1 *Imp) {
//
  if x1.num == 0 { return }
/*
  if x.num == 0 {
    should be more effective: // TODO
    rename ...
    x1.Name -> x.Name
  }
*/
  x.file.seek (x.num * x.size)
  x1.file.seek (0)
  for i:= null; i < x1.num; i++ {
    x1.read (x.buf)
    x.write (x.buf)
  }
  x.file.flush ()
  x.num = x.file.length() / x.size
  x1.Clr()
}


func (S *Imp) join (S1 *Imp) {
//
  if S1.num == 0 { return }
/*
  if S.num == 0 {
    more effective: see concatenate
  }
*/
  ps:= New (S.emptyObject)
  ps.Name (S.tmpName)
  ps.Clr()
  S.file.seek (0)
  S1.file.seek (0)
  S1.read (S1.buf)
  i:= null
  i1:= null
  if S.num > 0 {
    S.read (S.buf)
    for {
      if Less (S.buf, S1.buf) {
        ps.write (S.buf)
        i ++
        if i < S.num {
          S.read (S.buf)
        } else {
          break
        }
      } else {
        if Less (S1.buf, S.buf) {
          ps.write (S1.buf)
          i1 ++
          if i1 < S1.num {
            S1.read (S1.buf)
          } else {
            break
          }
        } else {
          ps.write (S1.buf)
          i ++
          if i < S.num {
            S.read (S.buf)
          }
          i1 ++
          if i1 < S1.num {
            S1.read (S1. buf)
          }
          if (i == S.num) || (i1 == S1.num) {
            break
          }
        }
      }
    }
  }
  for {
    if i == S.num { break }
    ps.write (S.buf)
    i ++
    if i < S.num {
      S.read (S.buf)
    }
  }
  for {
    if i1 == S1.num { break }
    ps.write (S1.buf)
    i1 ++
    if i1 < S1.num {
      S1.read (S1.buf)
    }
  }
  S.file.clear ()
  S.num = S.file.length() / S.size
  S1.Clr()
  ps.file.rename (S.name)
  ps.Terminate ()
}


func (S *Imp) Join (C Iterator) {
//
  S1:= S.imp (C)
  if S1 == nil { return }
  if S.ordered {
    S.join (S1)
  } else {
    S.concatenate (S1)
  }
}


/* func init () {
//
  filenames = seq.New (string)
} */
