package cdio

// (c) Christian Maurer   v. 130526 - license see murus.go

import (
  . "murus/obj"
  "murus/str"; "murus/col"; "murus/scr"; "murus/box"
  "murus/nat"; "murus/pbar"; "murus/clk";
  "murus/cdker"
)
const (
  wr = 60 // column width of controllers
  dC =  2 // column distance
  l0 =  2
  l1 = l0 + cdker.NCtrl * 3 + 2
  ct =  2 /* digits for track */ + dC
  cs = ct + 5 /* clk.Mm_ss */ + dC
  cr = 18
)
var (
  ctrlBar [cdker.NCtrl]*pbar.Imp
  timeBar [2]*pbar.Imp
  ctrlText [cdker.NCtrl]string
  timeText [2]string
  bx *box.Imp
  lv [cdker.NCtrl]uint // lines of volumebar
  lt [2]uint; // lines of trackbar
  cB, cF, timeF, trackTimeF, lengthF, ctrlF col.Colour
  done bool
)


func trackUnderMouse () (uint8, bool) {
//
  if scr.UnderMouse (l0, 0, cs + 5, uint(cdker.NTracks())) {
    l, _:= scr.MousePos()
    return uint8(l) - l0, true
  }
  return 0, false
}


func timeUnderMouse () (bool, uint, bool) {
//
  for i:= uint(0); i < 2; i++ {
    if scr.UnderMouse (lt[i], cr, wr, 2) {
      var n uint
      timeBar[i].Edit (&n)
      return i != 0, n, true
    }
  }
  return false, 0, false
}


func controlUnderMouse (c *cdker.Controller) (uint, bool) {
//
  for i:= cdker.Controller (0); i < cdker.NCtrl; i++ {
    if scr.UnderMouse (lv[i], cr, wr, 2) {
      var n uint
      ctrlBar[i].Edit (&n)
      *c = i
      return n, true
    }
  }
  return 0, false
}


func wr1 (n, l, c uint, f, b col.Colour) {
//
  nat.SetColours (f, b)
  nat.Write (n, l, c + 1)
}


func wrt (t *clk.Imp, fm Format, f, b col.Colour, l, c uint) {
//
  t.SetFormat (fm)
  t.SetColours (f, b)
  t.Write(l, c)
}


func writeMask () {
//
  n:= cdker.NTracks()
  n1:= n + 1
  for t:= uint8(0); t < n1; t++ {
    if t < n {
      wr1 (uint(t) + 1, l0 + uint(t), 0, cF, cB)
      wrt (cdker.StartTime[t], clk.Mm_ss, timeF, cB, l0 + uint(t), cs)
      wrt (cdker.Length[t], clk.Mm_ss, lengthF, cB, l0 + uint(t), ct)
    }
  }
  timeBar[1].Def (cdker.TotalTime.NSeconds())
  wrt (cdker.TotalTime, clk.Mm_ss, timeF, cB, lt[1] + 1, cr + wr - 5)
  done = true
}


func write () {
//
  n, a:= cdker.NTracks(), cdker.ActTrack()
//  cdker.Mutex.Lock()
  bx.Wd (12) // TODO
  bx.Colours (col.HintF, col.HintB)
  bx.Write (cdker.String(), l1, cr)
  var f col.Colour
  for t:= uint8(0); t < n; t++ {
    if t == a { f = trackTimeF } else { f = cF }
    wr1 (uint(t) + 1, l0 + uint(t), 0, f, cB)
  }
  for c:= cdker.Controller (0); c < cdker.NCtrl; c++ {
    ctrlBar[c].Fill (uint(cdker.Volume (c)))
    ctrlBar[c].Write ()
  }
  timeBar[0].Def (cdker.Length[a].NSeconds())
  timeBar[0].Fill (cdker.TrackTime.NSeconds())
  timeBar[0].Write ()
  timeBar[1].Fill (cdker.Time.NSeconds())
  timeBar[1].Write ()
  bx.Wd (2)
  bx.Colours (trackTimeF, cB)
  bx.Write (nat.StringFmt (uint(a) + 1, 2, false), l1, cr + wr - 2)
  wrt (cdker.TrackTime, clk.Mm_ss, trackTimeF, cB, lt[0] + 1, cr);
  wrt (cdker.Length[a], clk.Mm_ss, lengthF, cB, lt[0] + 1, cr + wr - 5);
  wrt (cdker.Time, clk.Mm_ss, trackTimeF, cB, lt[1] + 1, cr)
//  cdker.Mutex.Unlock()
}


func init () {
//
  scr.Switch (scr.TXT)
  cF, cB = col.LightMagenta, col.Black
// scr.Colours (cF, cB)
  lengthF, timeF = col.Red, col.LightBlue
  trackTimeF, ctrlF = col.Colour3 (191, 191, 255), col.Colour3 (63, 111, 255)
  for c:= cdker.Controller (0); c < cdker.NCtrl; c++ {
    ctrlText[c] = cdker.Ctrltext[c]
    str.Center (&ctrlText [c], wr)
    lv[c] = l0 + 3 * uint(c)
  }
  timeText = [2]string { "tracktime", "total time" }
  str.Center (&timeText[0], wr)
  str.Center (&timeText [1], wr)
  lt = [2]uint { l1 + 2, l1 + 2 + 3 }
  bx = box.New()
  bx.Wd (wr)
  bx.Colours (col.HintF, col.HintB)
  bx.Colours (trackTimeF, cB)
  bx.Write (" track", l1, cr + wr - 6 - 2)
  cw, lh:= scr.NX1(), scr.NY1()
  bx.Colours (cF, cB)
  bx.Wd (wr)
  for c:= cdker.Controller (0); c < cdker.NCtrl; c++ {
    ctrlBar[c] = pbar.New (true)
    ctrlBar[c].Def (cdker.MaxVol)
    ctrlBar[c].SetColours (ctrlF, cB)
    ctrlBar[c].Locate (cr * cw, lv[c] * lh, wr * cw, lh)
    bx.Write (ctrlText [c], lv[c] + 1, cr)
  }
  for i:= 0; i < 2; i++ {
    timeBar[i] = pbar.New (true)
    timeBar[i].SetColours (ctrlF, cB)
    timeBar[i].Locate (cr * cw, lt[i] * lh, wr * cw, lh)
    bx.Write (timeText[i], lt[i] + 1, cr)
  }
  scr.MouseCursor (true)
  scr.WarpMouse (lv[cdker.All] + 1, cr + wr / 4)
}
