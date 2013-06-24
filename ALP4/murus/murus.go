package main

/* (c) 1986-2013   Christian Maurer
       dr-maurer.eu proprietary - all rights reserved

   Die nur zum Zweck der Lehre konstruierten Quellen von murus haben rein akademischem Charakter;
   sie liefern u. a. Beispiele für mein Buch "Nichtsequentielle Programmierung mit Go 1 kompakt".
   Für Lehrzwecke an Universitäten und in Schulen sind die Quelltexte uneingeschränkt verwendbar;
   jegliche Form weitergehender (insbesondere kommerzieller) Nutzung ist jedoch strikt untersagt.
   Davon abweichende Bedingungen sind der schriftlichen Vereinbarung mit dem Urheber vorbehalten.

   THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDER AND THE CONTRIBUTORS "AS IS" AND ANY EXPRESS
   OR IMPLIED WARRANTIES, INCLUDING BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
   AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL THE COPYRIGHT OWNER OR
   ANY CONTRIBUTOR BE LIABLE  FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSE-
   QUENTIAL DAMAGES  (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
   LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)  HOWEVER CAUSED  AND ON ANY THEORY OF
   LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT  (INCLUDING NEGLIGENCE OR OTHERWISE)
   ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH
   DAMAGE.

   Die Quellen von murus sind mit größter Sorgfalt entwickelt und werden laufend gepflegt - ABER:
   Es gibt keine fehlerfreie Software - auch in diesen Quelltexten stecken wahrscheinlich Fehler:
   Ihre Verwendung in Programmen könnte zu SCHÄDEN führen, z. B. zur Inbrandsetzung von Rechnern,
   zur Entgleisung von Eisenbahnzügen, zum GAU in Atomkraftwerken oder zum Absturz des Mondes ...
   Deshalb wird vor der Einbindung irgendwelcher Quelltexte von murus in Programme zu ernsthaften
   Zwecken AUSDRÜCKLICH GEWARNT ! (Ausgenommen sind nur Demo-Programme zum Einsatz in der Lehre.)

   Alle Quellen und Dokumentationen sind im weltweiten Netz unter http://murus.org/go/ zu finden.
   Meldungen entdeckter Fehler und Hinweise auf Unklarheiten werden jederzeit dankbar angenommen. */

import (
  "murus/ker"; . "murus/obj"
  "murus/col"; "murus/scr"; "murus/errh";
  "murus/integ"; "murus/lint"; "murus/brat"; "murus/real"
  "murus/buf"; "murus/bpqu"; "murus/img"; "murus/pset"
  "murus/fuday"
  "murus/gra1"
  "murus/date"
  "murus/acc"
  "murus/fig"
  "murus/eye"
  "murus/persaddr"; "murus/schol"
  "murus/audio"
  "murus/v"
  "murus/car"; "murus/chanm"
  "murus/lock"; "murus/lockp"; "murus/lock2"
  "murus/puls"; "murus/dlock"; "murus/conn" // -> ntr
  "murus/asem"; "murus/barr"; "murus/rw"; "murus/lr"
  "murus/phil"; "murus/barb"; "murus/smok"
  "murus/rob"
)


func circ (c col.Colour, x, y uint) {
//
  scr.Colour (c)
  scr.Circle (int(8*x), int(16*y/2), 16*y/2-1)
}


func dr (x0, x1, y int, c col.Colour, f bool) {
//
  const dx = 2
  nx1, ny, y1:= int(scr.NX1()), int(scr.NY()), 0
  for x:= x0; x < x1; x += dx {
    if ! f { scr.SaveGr (uint(x), uint(y), uint(x) + car.W, uint(y) + car.H) }
    car.Draw (true, c, x, y); ker.Msleep (10); car.Draw (true, col.ScreenB, x, y)
    if f && x > x0 + 26 * nx1 && x % 8 == 0 && y + car.H < ny { y1++; y += y1 }
    if ! f { scr.RestoreGr (uint(x), uint(y), uint(x) + car.W, uint(y) + car.H) }
  }
}


func moon (x0 int) {
//
  const r = 40
  x, y, y1, ny:= x0, r, 0, int(scr.NY())
  for y < ny - r {
    scr.SaveGr (uint(x - r), uint(y - r), uint(x + r), uint(y + r))
    scr.Colour (col.Sandgelb); scr.CircleFull (x, y, r)
    ker.Msleep (33)
    scr.RestoreGr (uint(x - r), uint(y - r), uint(x + r), uint(y + r))
    y1++
    y += y1
  }
}


func joke (x0, x1, y0, nx1, ny1, x, y, w int, cl col.Colour, s string, b bool) {
//
  h:= y + 8 + 1
  y1:= y0 + 8 * h; if b { y1 -= 40 }
  dr (x0, x0 + x * nx1, y0 + y * ny1, cl, false)
  scr.SaveGr    (uint(x0 + x * nx1), uint(y1), uint(x0 + (x + w) * nx1), uint(y1 + 8 * w))
  img.Get (s, uint(x0 + x * nx1), uint(y1))
  t:= uint(1); if b { t = 2 }; ker.Sleep (t)
  scr.RestoreGr (uint(x0 + x * nx1), uint(y1), uint(x0 + (x + w) * nx1), uint(y1 + 8 * w))
  if b { w = 2 * w / 3 }
  dr (x0 + (x + w) * nx1, x1, y0 + y * ny1, cl, false)
}


func drive (cc, cl, cb col.Colour, d chan bool) {
//
  nx, nx1, ny1:= int(scr.NX()), int(scr.NX1()), int(scr.NY1())
  dw:= 96 * nx1
  x0:= (nx - dw) / 2
  x1:= x0 + dw - car.W
  y0:= ((int(scr.NLines()) - 31) / 2 + 3) * ny1
  dr (x0, x1, y0,            cc, false)
  dr (x0, x1, y0 +  2 * ny1, cl, false)
  dr (x0, x1, y0 +  3 * ny1, cl, false)
  joke (x0, x1, y0, nx1, ny1, 2, 4, 48, cl, "nsp", true)
  dr (x0, x1, y0 + 19 * ny1, cl, false)
  dr (x0, x0 + 68 * nx1, y0 + 20 * ny1, cl, false)
  csb:= col.ScreenB; col.ScreenB = col.Black
  dr (x0 + 69 * nx1, nx, y0 + 20 * ny1, col.FlashRed, true)
  col.ScreenB = csb
  joke (x0, x1, y0, nx1, ny1, 67, 21, 14, cl, "fire", false)
  joke (x0, x1, y0, nx1, ny1, 48, 22, 15, cl, "mca", false)
  moon (x0 + 90 * nx1)
  dr (x0, x1, y0 + 26 * ny1, cc, false)
  d <- true
}


func main () { // just to get all stuff compiled
//
  var _ chanm.ChannelModel = chanm.New (nil)
  t:= integ.String (0)
  _ = real.String (0.0)
  var _ lint.LongInteger = lint.New (0)
  var _ brat.Rational = brat.New ()
  var _ Object = date.New()
  var _ buf.Buffer = bpqu.New (0, 1)
  var _ pset.PersistentSet = pset.New (persaddr.New())
  var _ acc.Account = acc.New()
  var _ schol.Scholar = schol.New()
  eye.Touch()
  fuday.Touch()
  gra1.Touch()
  fig.Touch()
  var _ asem.AddSemaphore = asem.New (2)
  var _ barr.Barrier = barr.New (2)
  var _ rw.ReaderWriter = rw.New()
  var _ lr.LeftRight = lr.New()
  var _ lock2.Locker2 = lock2.NewPeterson()
  var _ lockp.LockerP = phil.NewLockNaiv()
  var _ barb.Barber = barb.NewSem()
  var _ smok.Smokers = smok.NewNaiv()
  var _ lock.Locker = dlock.New (nil)
  puls.Touch()
  var _ conn.Connection = conn.New ()
  rob.Touch()
  var _ Indexer = audio.New ()
  scr.Switch (scr.MaxMode()) // >= scr.PAL
  xx, yy:= scr.NColumns(), scr.NLines()
  cf, cl, cb:= v.Colours()
  circ (cb, xx/2, yy); circ (cl, xx-yy, yy); circ (cf, yy, yy)
  t = ""; errh.MurusLicense ("murus", v.String(), "1986-2013  Christian Maurer   http://murus.org", cf, cl, cb, &t)
  col.ScreenB = cb; done:= make (chan bool); go drive (cf, cl, cb, done); <-done
  ker.Terminate()
}
