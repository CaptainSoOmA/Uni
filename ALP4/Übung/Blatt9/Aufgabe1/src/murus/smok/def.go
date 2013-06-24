package smok

// (c) Christian Maurer   v. 130424 - license see murus.go

type
  Smokers interface {

// Pre: u < 3.
// Der aufrufende Prozess ist als Wirtin zur Ablage der
// zu u komplementären Utensilien im kritischen Abschnitt.
// Diese Utensilien sind verfügbar.
  AgentIn (u uint)

// Pre: u < 3.
// Der aufrufende Prozess war ggf. solange blockiert, bis keiner raucht.
  AgentOut ()

// Pre: u < 3.
// Die zu u komplementären Utensilien sind nicht mehr verfügbar,
// sondern in exklusivem Besitz des aufrufenden Rauchers.
// Er war ggf. solange blockiert, bis das möglich war.
  SmokerIn (u uint)

// Der aufrufende Raucher raucht nicht mehr.
  SmokerOut ()
}
