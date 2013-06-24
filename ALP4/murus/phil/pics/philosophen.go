package main

import ("murus/ker"; "murus/scr"; "murus/errh"; "murus/img")

func main() {
  const a = 100
  var p []string = []string { "Anaxagoras", "Aristoteles", "Cicero", "Jemand", "Demokrit", "Diogenes", "Epikur", "Heraklit", "Platon", "Protagoras", "Pythagoras", "Sokrates", "Thales", "Niemand" }
  var x, y uint
  for _, s:= range p {
    img.Get (s, x, y)
    errh.Error (s, 0)
    if x + 2 * a < scr.NX() {
      x += a
    } else {
      x = 0; y = a
    }
  }
  errh.Error2 ("", uint(len(p)), "Philosophen", 0)
  ker.Terminate()
}
