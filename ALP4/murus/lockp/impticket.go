package lockp

// (c) Christian Maurer   v. 111126 - license see murus.go

// >>> Ticket-Algorithm using FetchAndAddUint32

import
  "murus/lock"
type
  ImpTicket struct {
            ticket,
              turn uint32
                   }


func NewTicket (n uint) *ImpTicket {
//
  return new (ImpTicket)
}


func (L *ImpTicket) Lock (p uint) {
//
  t:= lock.FetchAndAddUint32 (&L.ticket, uint32(1))
  for t != L.ticket { /* do nothing */ }
}


func (L *ImpTicket) Unlock (p uint) {
//
  L.turn++
}
