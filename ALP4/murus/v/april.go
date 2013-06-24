package v

// (c) Christian Maurer   v. 130526 - license see murus.go

import ("strconv"; "murus/ker"; "murus/str"
        "murus/col"; "murus/scr"; "murus/box"; "murus/pseq")

const (y0 = 8; x0 = 5; f = "/tmp/funnyfile")
var ok bool

func rot (n int) { var b byte
  for i:= 0; ! ok; i++ {
    switch i % 4 { case 0: b = '|'; case 1: b = '/'; case 2: b = '-'; case 3: b = '\\' }
    scr.Write1 (b, y0 + 1, 71); ker.Msleep (50)
  }
}

func doof (y, x uint) { scr.Colour(col.LightOrange); scr.Write ("d(o,o)f", y, x0 + x) }

func doodle (c col.Colour, n int) {
  col.ScreenB = c; scr.Cls(); ker.Msleep(50)
  col.ScreenB = col.Black; scr.Cls(); scr.Colour (col.LightGreen)
  scr.Write ("The murus-tool to remove \"       \" is going to be executed, i.e.", y0, x0)
  scr.Write ("your disk will be completely reformatted, one moment please ... ", y0 + 1, x0); doof (y0, 26)
  const m = 1<<16
  x:= str.Clr (m)
  ok = false
  for i:= 0; i < 10 * n; i++ { if i == 0 { go rot (n) }
    file:= pseq.New (x); file.Name (f + strconv.Itoa(i) + ".tmp"); file.Clr(); file.Ins (x); file.Terminate()
  }
  ok = true
}

func April1st() {
  col.ScreenF, col.ScreenB = col.White, col.Black; scr.Cls()
  scr.MouseCursor (false)
  scr.Colour(col.White); scr.Write ("Found evil software: \"       \"", 7, x0); doof (7, 22)
  scr.Colour(col.White); scr.Write ("Remove (yes/no) ?", y0, x0)
  b:= box.New(); b.Wd (3); t:= "yes"; b.Edit (&t, y0, 23)
  scr.WarpMouseGr (2 * int(scr.NX()), 2 * int(scr.NY()))
  b.Colours (col.LightOrange, col.Black); b.Write ("yes", y0, 23)
  doodle (col.LightWhite, 2); doodle (col.LightWhite, 2); doodle (col.LightYellow, 5)
  doodle (col.Yellow, 3); doodle (col.LightOrange, 5); doodle (col.Orange, 8); doodle (col.LightRed, 3)
// TODO erase all (f + "*.tmp")
  t = str.Clr (70); scr.Write (t, y0, x0); scr.Write (t, y0 + 1, x0)
  col.ScreenF, col.ScreenB = col.LightGreen, col.DarkBlue; scr.Cls()
  scr.Write ("The murus-tool has removed \"       \" - your disk is reformatted :-)", y0, x0); doof (y0, 28)
  scr.Colour (col.LightWhite)
  scr.Write ("Please install Linux, TeX, mercurial, Go and murus completely new !", y0 + 2, x0)
  ker.Sleep(20); ker.Terminate()
}
