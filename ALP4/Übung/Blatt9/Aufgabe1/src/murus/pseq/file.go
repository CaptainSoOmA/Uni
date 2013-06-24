package pseq

// (c) Christian Maurer   v. 121105 - license see murus.go

import (
  . "os"
//  . "murus/ker"
  "murus/str" // ; "murus/env"
  "murus/errh"
)
const
  rights = 0644
/*
type
  clients byte; const (user = iota; group; world)
type
  accesses byte; const (readable = iota; writable; executable)
*/
type
  files struct {
          file *File
       isThere bool
        offset,
     endoffset uint64
        isOpen bool
           err error
               }
/*
var (
  statuscode [3][3]uint // accesses, clients
  caller, callerGroup uint // uid, gid of the calling process
)
*/
/*
func Owner (N string) uint {
//
  fi, err:= Stat (N)
  if err != nil {
    return MaxCard
  }
  return uint(fi.Gid)
}
*/
/*
func FileGroup (N string) uint {
//
  fi, err:= Stat (N)
  if err == nil {
    return uint(fi.Uid)
  }
  return MaxCard
}
*/
/*
func accessible (N string, a accesses) bool {
//
  fi, err:= Stat (N)
  if err == nil {
    return caller == Owner (N)       && statuscode [a, user]  IN fi.Mode ||
           callerGroup == Gruppe (N) && statuscode [a, group] IN fi.Mode
  }
  return false
}
*/

func directLength (N string) uint64 {
//
  fi, err:= Stat (N)
  if err == nil {
    return uint64(fi.Size())
  }
  return 0
}


func erase (N string) {
//
  _, err:= Stat (N)
  if err != nil {
    Remove (N)
  }
}


func newFile () *files {
//
  F:= new (files)
  F.file = nil
  F.isThere = false
  F.offset = 0
  F.endoffset = 0
  F.isOpen = false
  return F
//  return &F // prüfen, ob file = nil
}


func (F *files) terminate () {
//
  F.flush ()
}

/*
func (F *files) report (a, b string, n uint) {
//
  if F.err == nil {
//    errh.Error2 (a + b, n, "ok", 0)
  } else {
    errh.Error2 (a + " Error " + b, n, F.err.String(), 0)
  }
}
*/

func (F *files) name (N string) {
//
  if str.Empty (N) { return }
  F.flush ()
  var fi FileInfo
  fi, F.err = Stat (N)
  F.isThere = F.err == nil
  if F.isThere { // is there a file with name N (?)
//    if ! fi.IsRegular() { errh.Error (N + " is no regular file", 0); Terminate(); Exit (1) } // nothing goes
//    if fi.Permission() != rights { errh.Error (N + " has no rights", 0); Terminate(); Exit (1) } // nothing goes
    F.file, F.err = OpenFile (N, O_RDWR, rights) // ; F.report ("define", "OpenFile", 0)
    if F.err == nil {
      F.endoffset = uint64(fi.Size())
      _ = F.file.Close () // ; F.report ("define", "Close", 0)
    } else {
      F.file = nil
      println (&PathError { "define", N, F.err })
    }
  } else { // there is no file with name N (?)
    if IsPermission (F.err) { println ("no permission ") }
    F.file, F.err = Create (N) // ; F.report ("define", "Create", 0)
    if F.err == nil {
      F.endoffset = 0
      _ = F.file.Close() // ; F.report ("define", "Close", 0)
      F.isThere = true
    } else {
      F.file = nil
    }
  }
  F.offset = 0
  F.isOpen = false
}


func (F *files) rename (s string) {
//
  F.flush ()
  if F.isOpen {
    _ = F.file.Sync () // ; F.report ("redefine", "Sync", 0)
    _ = F.file.Close () // ; F.report ("redefine", "Close", 0)
    F.isOpen = false
  }
  if F.file.Name() == s { return }
  _ = Rename (F.file.Name(), s) // ; F.report ("redefine", "Rename", 0)
  F.isThere = true // = Stat (&name, F.status) == 0
  if ! F.isThere { /* F.report ("redefine", "Stat", 0) */ }
  F.offset = 0
//  F.endoffset = status.Byteanzahl // !!!
  F.isOpen = false
}


func (F *files) empty () bool {
//
  if F.isThere {
    return F.endoffset == 0
  }
  return true
}


func (F *files) clear () {
//
  F.open()
  F.file.Truncate (int64(0))
  F.flush()
  F.offset = 0
  F.endoffset = 0
  F.open()
}


func (F *files) length () uint64 {
//
  return F.endoffset
}


func (F *files) seek (P uint64) {
//
  F.offset = P
}


func (F *files) position () uint64 {
//
  return F.offset
}


func (F *files) open () {
//
  if F.isOpen { return }
//  f, err:= OpenFile (F.file.Name(), /* O_APPEND */ O_RDWR, rights) // ; F.report ("open", F.file.Name(), 0)
  F.file, F.err = OpenFile (F.file.Name(), /* O_APPEND */ O_RDWR, rights) // ; F.report ("open", F.file.Name(), 0)
//  if err == nil {
  if F.err == nil {
//    F.file = f
  } else {
    F.file = nil
  }
  F.isOpen = F.file != nil
}


func (F *files) read (B []byte) uint {
//
  F.open ()
  if ! F.isOpen { errh.Error (F.file.Name(), uint(F.offset)) }
  r:= len (B)
//  r, F.err = F.file.ReadAt (B, int64(F.offset)) // macht Zicken
  _ /* off */, _ = F.file.Seek (int64(F.offset), 0) // ; F.report ("read", "Seek at offset", uint(off))
  r, _ = F.file.Read (B) // ; F.report ("Read", "at offset", uint(off))
  F.offset += uint64(r)
  return uint(r)
}


func (F *files) write (B []byte) uint {
//
  F.open ()
  w:= len (B)
//  w, F.err = F.file.WriteAt (B, int64(F.offset)) // ; F.report ("WriteAt", "", 1000000 + uint(w)
//  var off int64
  /* off */ _, _ = F.file.Seek (int64(F.offset), 0) // ; F.report ("write", "Seek at offset", uint (off))
  w, _ = F.file.Write (B) // ; F.report ("Write", "at offset", uint (off))
  F.offset += uint64(w)
  if F.offset > F.endoffset { F.endoffset = F.offset }
  return uint(w)
}


func (F *files) flush () {
//
  if F.isOpen {
    _ = F.file.Sync () // ; F.report ("flush", "Sync", 0)
    _ = F.file.Close () // ; F.report ("flush", "Close", 0)
    F.isOpen = false
  }
}

/*
func init () {
// nonsense - only some ideas
  for access:= readable; access <= executable; access++ {
    for client:= user; client <= world; client++ {
      statuscode [access, client] = 3 * (2 - client) + (2 - access)
    }
  }
  callingProgram:= env.Parameter (0)
  fi, err:= Stat (callingProgram)
  if err == nil {
    caller = MaxCard
    callerGroup = MaxCard
  } else {
    caller = uint(fi.Uid)
    callerGroup = uint(fi.Gid)
  }
}
*/
