package errh

// (c) Christian Maurer   v. 130309 - license see murus.go

import (
//  "murus/env"
  "murus/str"; "murus/kbd"
  "murus/col"; . "murus/scr"; "murus/box"; "murus/nat"
)
var (
  hintbox, errorbox, licenseBox, choiceBox *box.Imp = box.New(), box.New(), box.New(), box.New ()
  hintWritten, hintPosWritten /* , DocExists */ bool
  hintwidth uint
  transparent bool
//  actualFontsize FontSizes
  license []string
)


func wait () { // TODO -> kbd, other name
//
  loop: for {
    _, c, _ := kbd.Read ()
    switch c { case kbd.Enter, kbd.Esc, kbd.Back, kbd.Here, kbd.There:
      break loop
    }
  }
}


func pre() {
//
  transparent = TransparenceOn ()
  if transparent { SwitchTransparence (false) }
//  actualFontsize = Fontgroesse ()
//  if actualFontsize # Normal {
//    SwitchFontsize (Normal)
//  }
//  hintbox.Wd (width)
//  errorbox.Wd (width)
}


func post() {
//
  if transparent { SwitchTransparence (true) }
//  if actualFontsize # Normal {
//    SwitchFontsize (Normal)
//  }
}


func hint (s string) {
//
  delHint()
  pre()
  w:= NColumns()
  str.Lat1 (&s)
  str.Center (&s, w)
  l:= NLines() - 1
  Save (l, 0, w, 1)
  hintbox.Wd (w)
  hintbox.Write (s, l, 0)
  hintWritten = true
  post()
}


func delHint () {
//
  pre()
  if hintWritten {
    hintWritten = false
    Restore (NLines() - 1, 0, NColumns(), 1)
  }
  post()
}


func hintPos (s string, l, c uint) {
//
//  delHintPos (s)
  pre()
  if l >= NLines () { l = NLines () - 1 }
  w:= uint(len (s))
  if c + w >= NColumns () { c = NColumns () - w }
  Save (l, c, w, 1)
  hintbox.Wd (w)
  hintbox.Write (s, l, c)
  hintPosWritten = true
  post()
}


func delHintPos (s string, l, c uint) {
//
  if hintPosWritten {
    hintPosWritten = false
    Restore (l, c, uint(len (s)), 1)
  }
}


func error (s string, n uint) {
//
  pre()
  if n > 0 { s += " " + nat.String (n) + " " }
  str.Lat1 (&s)
  str.Center (&s, NColumns())
  l:= NLines() - 1
  Save (l, 0, NColumns(), 1)
  errorbox.Wd (NColumns())
  errorbox.Write (s, l, 0)
  kbd.Wait (false)
  Restore (l, 0, NColumns(), 1)
  post()
}


func error2 (s string, n uint, s1 string, n1 uint) {
//
  str.Lat1 (&s)
  str.Lat1 (&s1)
  if n > 0 { s = s + " " + nat.String (n) }
  if n1 > 0 { s1 = s1 + " " + nat.String (n1) }
  error (s + " " + s1, 0)
}


func errorPos (s string, n, l, c uint) {
//
  pre()
  str.Lat1 (&s)
  if n > 0 { s = s + " " + nat.String (n) }
  if l >= NLines () { l = NLines () - 1 }
  w:= uint(len (s))
  if c + w >= NColumns () { c = NColumns () - w }
  Save (l, c, w, 1)
  errorbox.Wd (w)
  errorbox.Write (s, l, c)
  kbd.Wait (false)
  Restore (l, c, w, 1)
  post()
}


func error2Pos (s string, n uint, s1 string, n1 uint, l, c uint) {
//
  str.Lat1 (&s)
  str.Lat1 (&s1)
  if n > 0 { s = s + " " + nat.String (n) }
  if n1 > 0 { s1 = s1 + " " + nat.String (n1) }
  errorPos (s + s1, 0, l, c)
}


func confirmed () bool {
//
  pre()
  s:= "Sind Sie sicher?  j(a / n(ein"
  w:= NColumns()
  str.Center (&s, w)
  l:= NLines() - 1
  Save (l, 0, w, 1)
  errorbox.Wd (w)
  errorbox.Write (s, l, 0)
  b, _, _:= kbd.Read ()
  a:= b & ' ' == 'J'
  Restore (l, 0, w, 1)
  post()
  return a
}


func writeLicense (project, version, author string, f, cl, b col.Colour, g []string, Text *string) {
//
  pre()
  post()
  w, h:= uint(len (g[0])), uint(25) /* == len (license), see func init */ + 6
  l, c:= (NLines () - h) / 2, (NColumns () - w) / 2
//  l0, c0:= l, c
//  scr.Save (l, c, width, height)
  licenseBox.Wd (w)
  licenseBox.Colours (cl, b)
  emptyLine:= str.Clr (w)
  licenseBox.Write (emptyLine, l, c); l ++
  var s string
  str.Set (&s, project + " v. " + version)
  str.Center (&s, w)
  licenseBox.Write (s, l, c); l ++
  licenseBox.Write (emptyLine, l, c); l ++
  str.Set (&s, "(c) " + author)
  str.Center (&s, w)
  licenseBox.Colours (f, b)
  licenseBox.Write (s, l, c); l ++ // l, c = 30, 52
  licenseBox.Colours (cl, b)
  licenseBox.Write (emptyLine, l, c); l ++
  for i:= 0; i < len (g); i++ {
    licenseBox.Write (g[i], l, c); l ++
  }
  licenseBox.Write (emptyLine, l, c); l ++
  licenseBox.Colours (f, b)
/*
  var line string
  if DocExists {
    str.Set (&line, "ausführliche Bedienungshinweise: siehe Dokumentation")
  } else {
    line = env.Parameter (0)
    if line == "murus" {
      line = str.Clr (w)
    } else {
      str.Set (&line, "kurze Bedienungshinweise: F1-Taste")
    }
  }
  if ! str.Empty (line) { str.Center (&line, w) }
  licenseBox.Write (line, l, c); l ++
  licenseBox.Write (emptyLine, l, c)
*/
//  kbd.Wait (true)
//  scr.Restore (l0, c0, w, height)
}


func murusLicense (project, version, author string, f, l, b col.Colour, t *string) {
//
  writeLicense (project, version, author, f, l, b, license, t)
}


func writeHeadline (project, version, author string, f, b col.Colour) {
//
  pre()
  n:= NColumns ()
  Text:= project + "       (c) " + author + "   v. " + version
  str.Center (&Text, n)
  licenseBox.Wd (n)
  licenseBox.Colours (f, b)
  licenseBox.Write (Text, 0, 0)
  post()
}


func writeHelp (H []string) {
//
  pre()
  h:= uint(len (H))
  var w, l, c uint
  for i:= uint(0); i < h; i++ {
    c = uint(len (H[i]))
    if c > w { w = c }
  }
  if h + 2 > NLines() { h = NLines() - 2 }
  if w + 4 > NColumns() { w = NColumns() - 4 }
  mouseOn:= MouseCursorOn ()
  if false { // mouseOn {
    l, c = MousePos ()
    if l >= NLines() - h - 1 { l = NLines() - h - 2 }
    if c > NColumns() - w - 4 { c = NColumns() - w - 4 }
    MouseCursor (false)
  } else {
    l, c = (NLines() - h - 2) / 2, (NColumns() - w - 4) / 2
  }
  Save (l, c, w + 4, h + 2)
  hintbox.Wd (w + 4)
  T:= str.Clr (w + 4)
  for i:= uint(0); i <= h + 1; i++ {
    hintbox.Write (T, l + i, c)
  }
  hintbox.Wd (w)
  for i:= uint(0); i < h; i++ {
    hintbox.Write (H[i], l + 1 + i, c + 2)
  }
  wait ()
  Restore (l, c, w + 4, h + 2)
  if mouseOn { MouseCursor (true) }
  post()
}


func writeHelp1 () {
//
  pre()
  s:= "kurze Bedienungshinweise: F1-Taste"
  w:= uint(len (s))
//  mouseOn:= MouseCursorOn ()
  var l, c uint
  if false { // mouseOn {
    l, c = MousePos()
    if l >= NLines () - 2 { l = NLines () - 3 }
    if c > NColumns () - w { c = NColumns () - w }
    MouseCursor (false)
  } else {
    l = (NLines () - 3) / 2
    c = (NColumns () - w - 4) / 2
  }
  hintbox.Wd (w + 4)
  t:= str.Clr (w + 4)
  Save (l, c, w + 4, 3)
  for i:= uint(0); i <= 2; i++ { hintbox.Write (t, l + i, c) }
  hintbox.Wd (w)
  hintbox.Write (s, l + 1, c + 2)
  wait ()
  Restore (l, c, w + 4, 3)
//  if mouseOn { MouseCursor (true) }
  post()
}


func init () {
//           1         2         3         4         5         6         7         8         9
// 012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345
  lic:= [...]string {
  " Die Quellen von murus sind lediglich zum Einsatz in der Lehre konstruiert und haben demzufolge ",
  " einen rein akademischen Charakter; sie liefern u.a. eine Reihe von Beispielen für das Lehrbuch ",
  " \"Nichtsequentielle Programmierung mit Go 1 kompakt\" (Springer, 2. Aufl. 2012, 223 S. 14 Abb.). ",
  " Für Lehrzwecke an Universitäten und in Schulen sind die Quelltexte uneingeschränkt verwendbar; ",
  " jegliche Form weitergehender (insbesondere kommerzieller) Nutzung ist jedoch strikt untersagt. ",
  " Davon abweichende Bedingungen sind der schriftlichen Vereinbarung mit dem Urheber vorbehalten. ",
  "                                                                                                ",
  " THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDER AND THE CONTRIBUTORS \"AS IS\" AND ANY EXPRESS ",
  " OR IMPLIED WARRANTIES, INCLUDING BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY ",
  " AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL THE COPYRIGHT OWNER OR ",
  " ANY CONTRIBUTOR BE LIABLE  FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSE- ",
  " QUENTIAL DAMAGES  (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; ",
  " LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)  HOWEVER CAUSED  AND ON ANY THEORY OF ",
  " LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT  (INCLUDING NEGLIGENCE OR OTHERWISE) ",
  " ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH ",
  " DAMAGE. APART FROM THIS THE TEXT IN GERMAN ABOVE AND BELOW IS A MANDATORY PART OF THE LICENSE. ",
  "                                                                                                ",
  " Die Quelltexte von murus sind mit größter Sorgfalt entwickelt und werden laufend überarbeitet. ",
  " ABER: Es gibt keine Software ohne Bugs - auch in DIESEN Quelltexten stecken sicherlich Fehler. ",
  " Ihre Verwendung in Programmen könnte zu SCHÄDEN führen, z. B. zur Inbrandsetzung von Rechnern, ",
  " zur Entgleisung von Eisenbahnzügen, zum GAU in Atomkraftwerken oder zum Absturz des Mondes ... ",
  " Deshalb wird vor der Einbindung irgendwelcher Quelltexte von murus in Programme zu ernsthaften ",
  " Zwecken AUSDRÜCKLICH GEWARNT ! (Ausgenommen sind nur Demo-Programme zum Einsatz in der Lehre.) ",
  "                                                                                                ",
  " Meldungen entdeckter Fehler und Hinweise auf Unklarheiten werden jederzeit dankbar angenommen. " }
  license = make ([]string, len (lic))
  for i, l:= range (lic) { str.Set (&license[i], l) }

  hintbox.Colours (col.HintF, col.HintB)
  errorbox.Colours (col.ErrorF, col.ErrorB)
  pre()
  post()
//                                           1         2         3         4         5         6         7
//                                 012345678901234567890123456789012345678901234567890123456789012345678901234567
  str.Set (&ToWait,            "bitte warten ...")
  str.Set (&ToContinue,        "weiter: Einter")
  str.Set (&ToContinueOrNot,   "weiter: Einter                                                     fertig: Esc")
  str.Set (&ToCancel,          "                                                                abbrechen: Esc")
  str.Set (&ToScroll,          "blättern: Pfeiltasten                                           abbrechen: Esc")
  str.Set (&ToSelect,          "blättern/auswählen/abbrechen: Pfeiltasten/Enter/Esc, Maus bewegen/links/rechts")
  str.Set (&ToChange,          "blättern: Pfeiltasten       ändern: Enter       abbrechen: Esc")
  str.Set (&ToSwitch,          "blättern: Pfeiltasten    auswählen: Enter    umschalten: Tab    abbrechen: Esc")
  str.Set (&ToSelectWithPrint, "blättern: Pfeiltasten    auswählen: Enter    drucken: Druck     abbrechen: Esc")
  str.Set (&ToPrint,           "ausdrucken: Druck                                         fertig: andere Taste")
}
