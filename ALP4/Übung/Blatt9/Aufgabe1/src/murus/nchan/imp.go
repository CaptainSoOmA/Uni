package nchan

// (c) Christian Maurer   v. 130309 - license see murus.go

import (
  "time"; "sync"; "net"
  . "murus/ker"; . "murus/obj"; "murus/str"
  "murus/nat"; "murus/errh"; "murus/pseq"
  "murus/host"
)
const (
  pack = "nchan"
  tcp = "tcp"
  prefix = "/tmp/inet-"
  tst = true
)
type
  Imp struct {
      object Any
       width uint
      server string
        port uint16
     farHost *host.Imp
      oneOne,
    isServer,
        gone,
        info bool
        conn net.Conn
    nClients uint
        addr *net.TCPAddr
        list net.Listener
         buf []byte
     in, out chan Any
        s, r int
         err error
             }
var
  mutex sync.Mutex


func name (n uint16) string {
//
  const d = 5 // n < 2^16 contains at most 5 digits
  return prefix + nat.StringFmt (uint(n), d, true)
}


func first (n uint16) bool {
//
  file:= pseq.New (byte(0))
// TODO mutual exclusion by file locking
  file.Name (name (n))
  f:= file.Empty ()
  if f {
    file.Put ('#')
  }
  file.Terminate ()
  return f
}


func delete (n uint16) {
//
  N:= name (n)
  if pseq.Length (N) > 0 {
    file:= pseq.New (byte(0))
    file.Name (N)
    file.Clr()
    file.Terminate()
  }
}


func (x *Imp) Chan () (chan Any, chan Any) {
//
  if x.isServer && ! x.oneOne {
    return x.in, x.out
  }
  return nil, nil
}


func (x *Imp) serve (c net.Conn) {
//
  if x.info { println ("new client") }
  for {
    x.r, _ = c.Read (x.buf)
    x.checkR ()
    if x.r == 0 {
      break
    }
    x.in <- Decode (x.object, x.buf)
//  println ("server is blocked")
// at this point the calling process is blocked
// until the function of the client is executed
      a:= <-x.out
    x.s, _ = c.Write (Encode (a))
    x.checkS ()
  }
  x.nClients--
  if x.info {
    if x.nClients > 0 {
      if x.info { println ("one client less") }
    } else {
      if x.info { println ("no clients") }
    }
  }
  c.Close ()
}


func New0 (a Any, h string, p uint16, o bool) *Imp {
//
  if str.Empty (h) || p >= 1 << 16 - Port0 { return nil }
  if a == nil { a = false }
  x:= new (Imp)
  if tst { println ("nchan.New0 started for host", h, "/ port", p) }
  x.object, x.width = Clone (a), Codelen (a)
  str.RemSpaces (&h)
  x.farHost = host.New ()
  if ! x.farHost.Defined (h) {
    errh.Error ("Hostname " + h + " is not resolvable", 0); Stop (pack, 1)
  }
  x.isServer = host.Local (h)
  x.server = h
  x.port = Port0 + p
  x.oneOne = o
  if x.oneOne {
    if x.isServer {
      x.isServer = first (x.port)
    } else {
      x.isServer = x.farHost.Sonstewas ()
    }
  }
  x.buf = make ([]byte, x.width)
//  x.info = true
  if tst { println ("nchan.New0 for host", h, "is done") }
  return x
}


func (x *Imp) Go () {
//
  if x.gone {
    return
  }
  x.gone = true
  portstring:= nat.StringFmt (uint(x.port), nat.Wd (uint(x.port)), true)
  if x.isServer {
    if x.info { println ("nchan.Go started on server", x.server) }
    x.addr, _ = net.ResolveTCPAddr (tcp, ":" + portstring)
//    if x.err != nil { Stop (pack, 2) }
    x.list, _ = net.ListenTCP (tcp, x.addr)
//    if x.err != nil { Stop (pack, 3) }
    if x.oneOne {
      var e error
      x.conn, e = x.list.Accept () // x.conn.RemoteAddr().(*net.TCPAddr).IP == ipFar
      if e != nil || x.conn == nil { Stop (pack, 4) }
    } else {
      if x.info { println ("server", x.server, "is up") }
      x.in, x.out = make (chan Any), make (chan Any)
      go func () {
        for {
          conn, err:= x.list.Accept ()
          if err == nil {
            x.nClients ++
            go x.serve (conn)
          } else {
// if tst { println ("not accepted because " + err.String()) }
          }
        }
      }()
    }
  } else { // client
    if tst { println ("nchan.Go started for", x.server, "as client") }
    x.farHost.SetFormat (host.Hostname)
    dialaddr:= x.farHost.String () + ":" + portstring
    for {
//      println ("waiting for server", x.server)
      x.conn, x.err = net.Dial (tcp, dialaddr)
      if x.err == nil {
        break
      }
      time.Sleep (2e9)
    }
    if x.conn == nil { Stop (pack, 5) }
    if tst { println ("nchan.Go.Dial", x.server, "war erfolgreich") }
//    if x.info { errh.DelHint () }
  }
}


func New (a Any, c string, p uint16, o bool) *Imp {
//
  x:= New0 (a, c, p, o)
  x.Go ()
  return x
}


func (x *Imp) IsServer () bool {
//
  return x.isServer
}


func (x *Imp) checkS () {
//
  if x.s < 0 { Stop (pack, 6) }
  if x.s < int(x.width) {
    println ("sent only ", x.s, " bytes instead of ", x.width)
  }
}


func (x *Imp) Send (a Any) {
//
  if ! x.gone { Stop (pack, 7) }
  if ! TypeEq (x.object, a) { Stop (pack, 8) }
  if x.conn == nil { Stop (pack + " / " + x.server, 9) }
  x.s, _ = x.conn.Write (Encode (a))
  print ("sent a message to "); if x.isServer { println (x.farHost.String()) } else { println (x.server) }
  x.checkS ()
}


func (x *Imp) checkR () {
//
  switch { case x.r < 0:
    Stop (pack, 10)
  case x.r == 0:
    if x.isServer {
      if x.oneOne {
        if x.info { println ("connection down") } // TODO
      } else {
//      TODO // can happen while serving
      }
    } else { // client
      if x.oneOne {
        if x.info { println ("connection down") } // TODO
      } else {
//      TODO
      }
    }
  case x.r < int(x.width):
    println ("received only ", x.r, " bytes instead of ", x.width)
  }
}


func (x *Imp) Recv () Any {
//
//  if x.isServer && ! x.oneOne { Stop (pack, 11) }
  if ! x.gone { Stop (pack, 12) }
  for {
    x.r, x.err = x.conn.Read (x.buf)
    if x.err == nil {
      break
    }
    x.checkR ()
  }
  return Decode (x.object, x.buf)
}


func (x *Imp) Send2 (c net.Conn, a Any) {
//
  if ! TypeEq (x.object, a) { Stop (pack, 13) }
  x.s, _ = c.Write (Encode (a))
  x.checkS ()
}


func (x *Imp) Recv2 () (net.Conn, Any) {
//
  if x.isServer && ! x.oneOne { Stop (pack, 14) }
  for {
    x.r, x.err = x.conn.Read (x.buf)
    if x.err == nil {
      break
    }
    x.checkR ()
  }
  return x.conn, Decode (x.object, x.buf)
}


func (x *Imp) Terminate () {
//
  if x.isServer && x.oneOne { Stop (pack, 15) } // Blödsinn ?
  x.conn.Close()
  if x.isServer {
    delete (x.port)
    if ! x.oneOne {
      close (x.in)
      close (x.out)
    }
  }
}
