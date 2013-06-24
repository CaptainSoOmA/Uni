package xker

// (c) Christian Maurer   v. 130404 - license see murus.go

// +build linux,cgo

// #cgo LDFLAGS: -lX11 -lGL
// #include <stdlib.h>
// #include <X11/X.h>
// #include <X11/Xlib.h>
// #include <X11/Xutil.h>
// #include <X11/Xatom.h>
// #include <X11/cursorfont.h>
// #include <GL/gl.h>
// #include <GL/glx.h>
/*
int typ (XEvent *E) { return (*E).type; }

void sw (Display *D, Window W, Window W0, Atom A, Atom F, int on)
{ XEvent E;
  E.type = ClientMessage;
  E.xclient.display = D;
  E.xclient.window = W;
  E.xclient.message_type = A;
  E.xclient.format = 32;
  E.xclient.data.l[0] = on; // standards.freedesktop.org/wm-spec/
  E.xclient.data.l[1] = F;
  E.xclient.data.l[2] = 0;
  if (XSendEvent (D, W0, False, SubstructureNotifyMask, &E) < 0) ;
  if (XFlush (D) < 0) ;
}

void m3 (Display *D, Window W, Atom A)
{ XEvent E;
  E.type = ClientMessage;
  E.xclient.display = D;
  E.xclient.window = W;
  E.xclient.message_type = A;
  E.xclient.send_event = False;
  E.xclient.format = 16; // doesn't matter
  if (XSendEvent (D, W, False, 0L, &E) < 0) ;
  if (XSync (D, False) < 0) ;
}

unsigned int xKeyState (XEvent *E) { return (*E).xkey.state; }

unsigned int xKeyCode (XEvent *E) { return (*E).xkey.keycode; }

unsigned int xButtonState (XEvent *E) { return (*E).xbutton.state; }

unsigned int xButtonButton (XEvent *E) { return (*E).xbutton.button; }

int xButtonX (XEvent *E) { return (*E).xbutton.x; }

int xButtonY (XEvent *E) { return (*E).xbutton.y; }

unsigned int xMotionState (XEvent *E) { return (*E).xmotion.state; }

unsigned int xMotionHint (XEvent *E) { return (*E).xmotion.is_hint; }

int xMotionX (XEvent *E) { return (*E).xmotion.x; }

int xMotionY (XEvent *E) { return (*E).xmotion.y; }

Atom mT (XEvent *E) { return (*E).xclient.message_type; }

unsigned long xGetPixel (XImage *I, int x, int y) { return XGetPixel (I, x, y); }

void xPutPixel (XImage *I, int x, int y, unsigned long p) { XPutPixel (I, x, y, p); }

void xDestroyImage (XImage *I) { XDestroyImage (I); }

void initialize (Display *d, int s, Window w)
{ int a[11];
  a[0] = GLX_RED_SIZE;     a[1] = 1;
  a[2] = GLX_GREEN_SIZE;   a[3] = 1;
  a[4] = GLX_BLUE_SIZE;    a[5] = 1;
  a[6] = GLX_DOUBLEBUFFER; a[7] = 1;
  a[8] = GLX_DEPTH_SIZE;   a[9] = 1;
  a[10] = 0;
  int n;
  GLXFBConfig config = *(glXChooseFBConfig (d, s, a, &n));
  GLXContext c = glXCreateNewContext (d, config, GLX_RGBA_TYPE, NULL, 1);
  glXMakeContextCurrent (d, w, w, c);
}

void write (Display *d, Window w) { glXSwapBuffers (d, w); }
*/
import
  "C"
import (
  "os"; "unsafe"
  "murus/obj"; "murus/env"
  "murus/col"
)
const
  _DISPLAY = "DISPLAY"
var (
  dspl string
  txt [LASTEvent+1]string
  display *C.struct_Display
  maxX, maxY,       // full screen
  xx, yy C.uint // window
  window, window0 C.Window // C.XID = C.ulong = CARD32 = uint32
  graphicsContext C.GC
  initialized bool
  screen C.int
  bitdepth C.uint
  pixmap, pixmap1 C.Pixmap
  xMin, yMin, xMax, yMax, ZB, ZH, ZA C.int
  xxb bool
  actualEvent C.XEvent
  naviFd uint
  netwm_state, fullscreen, naviAtom C.Atom
)


func Active () bool {
//
  dspl = env.Val (_DISPLAY)
  return dspl != ""
}


func Far () bool {
//
  return env.Val (_DISPLAY)[0] == 'l' // localhost
}


func cc (c col.Colour) C.ulong {
//
  return C.ulong(col.Code (c))
}


func Colours (F, B col.Colour) {
//
  if ! initialized { return }
//  C.XSetForeground (display, graphicsContext, C.ulong(col.CodeF))
//  C.XSetBackground (display, graphicsContext, C.ulong(col.CodeB))
  C.XSetForeground (display, graphicsContext, cc (col.ActualF))
  C.XSetBackground (display, graphicsContext, cc (col.ActualB))
  C.XFlush (display)
}


func InitBW (W, H uint) (uint, uint, uint) {
//
  return Init (W, H, col.White, col.Black)
}


func Screen () (C.uint, C.uint, C.uint) {
//
  d:= C.CString(dspl); defer C.free (unsafe.Pointer(d))
  display = C.XOpenDisplay (d)
  if display == nil { panic ("display == nil"); /* Terminate (); */ os.Exit (1) }
  window0 = C.XDefaultRootWindow (display)
  screen = C.XDefaultScreen (display)
  return C.uint(C.XDisplayWidth (display, screen)),
         C.uint(C.XDisplayHeight (display, screen)),
         C.uint(C.XDefaultDepth (display, screen))
}


func Init (w, h uint, cF, cB col.Colour) (uint, uint, uint) {
//
  initialized = false
//
  maxX, maxY, bitdepth = Screen ()
  bdp:= uint(bitdepth)
  switch bdp { case 15, 16, 24, 32:
    col.SetColourDepth (bdp)
  default:
    panic ("strange colourdepth"); /* Terminate (); */ os.Exit (1)
  }
  colourdepth:= (bdp + 1) / 8
  col.ScreenF, col.ScreenB = cF, cB
  xx, yy = C.uint(w), C.uint(h)
  window = C.XCreateSimpleWindow (display, C.Window(window0), 0, 0, xx, yy, C.uint(0), C.ulong(0), cc (col.ScreenB))
//  var E C.XEvent
//  C.XNextEvent (display, &E) // XCreate... did not produce an XEvent
  C.initialize (display, screen, window)
  t:= C.CString (env.Par (0)); defer C.free (unsafe.Pointer(t))
  C.XStoreName (display, window, t)
  C.XMapRaised (display, window)
  const inputmask = (C.KeyPressMask + // C.KeyReleaseMask +
                     C.ButtonPressMask + C.ButtonReleaseMask + C.PointerMotionMask +
                     C.ExposureMask + C.StructureNotifyMask)
  C.XSelectInput (display, window, inputmask)
  cursor:= C.XCreateFontCursor (display, C.XC_gumby)
  C.XDefineCursor (display, window, cursor)
  graphicsContext = C.XDefaultGC (display, screen)
//  C.XFlushGC (display graphicsContext)
  C.XSetGraphicsExposures (display, graphicsContext, C.False)
  SetFontsize (16)
  initialized = true
  Colours (cF, cB)
  C.XSetForeground (display, graphicsContext, cc (col.ScreenB))

  C.XFillRectangle (display, C.Drawable(window), graphicsContext, 0, 0, xx, yy)
  pixmap = C.XCreatePixmap (display, C.Drawable(window), xx, yy, bitdepth)
  pixmap1 = C.XCreatePixmap (display, C.Drawable(window), maxX, maxY, bitdepth)
  C.XFillRectangle (display, C.Drawable(pixmap), graphicsContext, 0, 0, xx, yy)
  C.XFillRectangle (display, C.Drawable(pixmap1), graphicsContext, 0, 0, xx, yy)
  C.XSetForeground (display, graphicsContext, cc (col.ScreenF))

  MouseDef (0, 0, int(xx - 1), int(yy - 1))
  var E C.XEvent
  C.XNextEvent (display, &E)
  var et C.int = C.typ (&E) // et == *E.type
  switch et { case C.Expose, C.ConfigureNotify: // zur Erstausgabe
    for C.XCheckTypedEvent (display, et, &E) == C.True { }
//    pp2ff ()
  case KeyPress, KeyRelease, ButtonPress, ButtonRelease, MotionNotify:
    C.XPutBackEvent (display, &E)
  case C.ReparentNotify: // at Switch (?)
    // ignore
  default: // for test purposes
//    println ("at initializing x:" + txt [et])
  }
  p:= C.CString ("WM_PROTOCOLS"); defer C.free (unsafe.Pointer(p))
  wm_protocols:= C.XInternAtom (display, p, C.False)
  C.XSetWMProtocols (display, window, &wm_protocols, 1)
  s:= C.CString ("_NET_WM_STATE"); defer C.free (unsafe.Pointer(s))
  netwm_state = C.XInternAtom (display, s, C.False)
  f:= C.CString ("_NET_WM_STATE_FULLSCREEN"); defer C.free (unsafe.Pointer(f))
  fullscreen = C.XInternAtom (display, f, C.False)
  m:= C.CString ("navi"); defer C.free (unsafe.Pointer(m))
  naviAtom = C.XInternAtom (display, m, C.False)
  Eventpipe = make (chan Event)
  go sendEvents ()
//  C.XFlush (display)
//  println ("init ok")
  return uint(maxX), uint(maxY), colourdepth
}


func catchNavi () {
//
  for {
    if false /* navi.Given () */ {
      C.m3 (display, window, naviAtom)
    }
  }
}


func terminate () {
//
//  C.XFreeGC (display, graphicsContext)
//  C.XUnmapWindow (display, window)
  C.XDestroyWindow (display, window)
  C.XDestroyWindow (display, C.Window(window0))
  C.XCloseDisplay (display)
  initialized = false
}


func MaxNoLines () uint {
//
  return uint(maxY)
}


func MaxNoColumns () uint {
//
  return uint(maxX)
}


func Switch (x, y uint) {
//
  if xx == maxX && yy == maxY {
    C.sw (display, window, window0, netwm_state, fullscreen, C.int(0))
  }
  xx, yy = C.uint(x), C.uint(y)
  if xx == maxX && yy == maxY {
    C.sw (display, window, window0, netwm_state, fullscreen, C.int(1))
  } else {
    C.XResizeWindow (display, window, xx, yy) // resizerequest.width, resizerequest.height
    C.XFlush (display)
  }
  C.XFreePixmap (display, pixmap)
  pixmap = C.XCreatePixmap (display, C.Drawable(window), xx, yy, bitdepth)
  C.XFreePixmap (display, pixmap1)
  pixmap1 = C.XCreatePixmap (display, C.Drawable(window), xx, yy, bitdepth)
  C.XSetForeground (display, graphicsContext, cc(col.ActualB))
//
  C.XFillRectangle (display, C.Drawable(window), graphicsContext, 0, 0, xx, yy)
  C.XFillRectangle (display, C.Drawable(pixmap), graphicsContext, 0, 0, xx, yy)
  C.XFillRectangle (display, C.Drawable(pixmap1), graphicsContext, 0, 0, xx, yy)
  MouseDef (0, 0, int(xx - 1), int(yy - 1))
//  var E *C.XEvent
//  C.XNextEvent (display, E)
//  et:= C.typ (&E) // et == *E.type
//  switch et { case C.Expose, C.ConfigureNotify:
//    for C.XCheckTypedEvent (display, C.int(et), E) == C.True { }
//  case C.ReparentNotify:
//    ;
//  default:
//    C.XPutBackEvent (display, E)
//  }
  WarpMouse (int(xx) / 2, int(yy) / 2)
//  C.XSync (display, C.False)
//  if navi.initialized (naviFd) {
//    go catchNavi ()
//  }
}


func pp2ff() {
//
  C.XCopyArea (display, C.Drawable(pixmap), C.Drawable(window), graphicsContext, 0, 0, xx, yy, 0, 0)
  C.XFlush (display)
}


func Buf (on bool) {
//
  if xxb == on { return }
  xxb = on
  if on {
    C.XSetForeground (display, graphicsContext, cc (col.ScreenB))
    C.XFillRectangle (display, C.Drawable(pixmap), graphicsContext, 0, 0, C.uint(xx), C.uint(yy))
    C.XSetForeground (display, graphicsContext, cc (col.ScreenF))
    C.XFlush (display)
  } else {
    pp2ff()
  }
}


func Clr (x, y, w, h uint) {
//
  C.XSetForeground (display, graphicsContext, cc (col.ScreenB))
  C.XFillRectangle (display, C.Drawable(window), graphicsContext, C.int(x), C.int(y), C.uint(w), C.uint(h))
  C.XFillRectangle (display, C.Drawable(pixmap), graphicsContext, C.int(x), C.int(y), C.uint(w), C.uint(h))
  C.XSetForeground (display, graphicsContext, cc (col.ScreenF))
  C.XFlush (display)
}


// /usr/share/fonts/misc
func SetFontsize (a uint) {
//
  var Fontname string
  switch a { case 7: // Tiny
    Fontname = "-misc-fixed-medium-r-*-*-7-*-*-*-*-*-iso8859-15"
    ZB, ZH, ZA = 5, 7, 6
  case 10: // Small
    Fontname = "-misc-fixed-medium-r-*-*-10-*-*-*-*-*-iso8859-15"
    ZB, ZH, ZA = 6, 10, 8
  case 16: // Normal
    Fontname = "-xos4-terminus-bold-r-*-*-16-*-*-*-*-*-iso8859-15"
    ZB, ZH, ZA = 8, 16, 12
  case 24: // Big
    Fontname = "-xos4-terminus-bold-r-*-*-24-*-*-*-*-*-iso8859-15"
    ZB, ZH, ZA = 12, 24, 19
  case 32: // Huge
    Fontname = "-xos4-terminus-bold-r-*-*-32-*-*-*-*-*-iso8859-15"
    ZB, ZH, ZA = 16, 32, 26
  default:
    return
  }
  F:= C.CString (Fontname); defer C.free (unsafe.Pointer(F))
  fsp:= C.XLoadQueryFont (display, F)
  if fsp == nil {
    if a < 16 {
      panic ("misc-fixed-font is not installed !"); /* Terminate (); */ os.Exit (1)
    } else {
      panic ("terminus-font is not installed !"); /* Terminate (); */ os.Exit (1)
    }
  }
  ZA = C.int(fsp.max_bounds.ascent)
  ZH = C.int(fsp.max_bounds.ascent + fsp.max_bounds.descent)
  Font:= C.Font(fsp.fid)
  C.XSetFont (display, graphicsContext, Font)
}


func Write (s string, x, y int, t bool) {
//
  n:= C.uint(len (s))
  if ! t {
    C.XSetForeground (display, graphicsContext, C.ulong(col.CodeB))
    if ! xxb { C.XFillRectangle (display, C.Drawable(window), graphicsContext, C.int(x), C.int(y), n * C.uint(ZB), C.uint(ZH)) }
    C.XFillRectangle (display, C.Drawable(pixmap), graphicsContext, C.int(x), C.int(y), n * C.uint(ZB), C.uint(ZH))
    C.XSetForeground (display, graphicsContext, C.ulong(col.CodeF))
  }
  cs:= C.CString (s)
  if ! xxb { C.XDrawString (display, C.Drawable(window), graphicsContext, C.int(x), C.int(y) + ZA, cs, C.int(n)) }
  C.XDrawString (display, C.Drawable(pixmap), graphicsContext, C.int(x), C.int(y) + ZA, cs, C.int(n))
  C.free (unsafe.Pointer (cs))
  C.XFlush (display)
}


func WriteInvert (s string, x, y int, t bool) {
//
  C.XSetFunction (display, graphicsContext, C.GXinvert)
  Write (s, x, y, t)
  C.XSetFunction (display, graphicsContext, C.GXcopy)
}


func Save (x, y, w, h uint) {
//
  C.XCopyArea (display, C.Drawable(window), C.Drawable(pixmap1), graphicsContext, C.int(x), C.int(y), C.uint(w), C.uint(h), C.int(x), C.int(y))
}


func Restore (x, y, w, h uint) {
//
  C.XCopyArea (display, C.Drawable(pixmap1), C.Drawable(window), graphicsContext, C.int(x), C.int(y), C.uint(w), C.uint(h), C.int(x), C.int(y))
  C.XCopyArea (display, C.Drawable(window), C.Drawable(pixmap), graphicsContext, C.int(x), C.int(y), C.uint(w), C.uint(h), C.int(x), C.int(y))
}


func Invert (x, y, x1, y1 uint) {
//
  if x > x1 || y > y1 { } // desaster
  C.XSetFunction (display, graphicsContext, C.GXinvert)
  if ! xxb {
    C.XFillRectangle (display, C.Drawable(window), graphicsContext, C.int(x), C.int(y), C.uint(x1 - x + 1), C.uint(y1 - y + 1))
  }
  C.XFillRectangle (display, C.Drawable(pixmap), graphicsContext, C.int(x), C.int(y), C.uint(x1 - x + 1), C.uint(y1 - y + 1))
  C.XSetFunction (display, graphicsContext, C.GXcopy)
  C.XFlush (display)
}


func SetLinewidth (b uint) {
//
  if b == 1 { b = 0 }
  C.XSetLineAttributes (display, graphicsContext, C.uint(b), C.LineSolid, C.CapRound, C.JoinRound)
}


func Colour (x, y uint) col.Colour {
//
  return col.ScreenB // TODO
}


func Point (x, y int, n bool) {
//
  if ! n { C.XSetFunction (display, graphicsContext, C.GXinvert) }
  if ! xxb { C.XDrawPoint (display, C.Drawable(window), graphicsContext, C.int(x), C.int(y)) }
  C.XDrawPoint (display, C.Drawable(pixmap), graphicsContext, C.int(x), C.int(y))
  if ! n { C.XSetFunction (display, graphicsContext, C.GXcopy) }
  C.XFlush (display)
}


func Points (X, Y []int, b bool) {
//
  l:= len (X); if len (Y) != l { return }
  p:= make ([]C.XPoint, l)
  for i:= 0; i < l; i++ {
    p[i].x, p[i].y = C.short(X[i]), C.short(Y[i])
  }
  if ! b { C.XSetFunction (display, graphicsContext, C.GXinvert) }
  if ! xxb { C.XDrawPoints (display, C.Drawable(window), graphicsContext, &p[0], C.int(l), C.CoordModeOrigin) }
  C.XDrawPoints (display, C.Drawable(pixmap), graphicsContext, &p[0], C.int(l), C.CoordModeOrigin)
  if ! b { C.XSetFunction (display, graphicsContext, C.GXcopy) }
  C.XFlush (display)
}


func Line (x, y, x1, y1 int, n bool) {
//
  if ! n { C.XSetFunction (display, graphicsContext, C.GXinvert) }
  if ! xxb { C.XDrawLine (display, C.Drawable(window), graphicsContext, C.int(x), C.int(y), C.int(x1), C.int(y1)) }
  C.XDrawLine (display, C.Drawable(pixmap), graphicsContext, C.int(x), C.int(y), C.int(x1), C.int(y1))
  if ! n { C.XSetFunction (display, graphicsContext, C.GXcopy) }
  C.XFlush (display)
}


func Lines (X, Y []int, n bool) {
//
  l:= len (X); if len (Y) != l { return }
  p:= make ([]C.XPoint, l)
  for i:= 0; i < l; i++ {
    p[i].x, p[i].y = C.short(X[i]), C.short(Y[i])
  }
  if ! n { C.XSetFunction (display, graphicsContext, C.GXinvert) }
  if ! xxb { C.XDrawLines (display, C.Drawable(window), graphicsContext, &p[0], C.int(l), C.CoordModeOrigin) }
  C.XDrawLines (display, C.Drawable(pixmap), graphicsContext, &p[0], C.int(l), C.CoordModeOrigin)
  if ! n { C.XSetFunction (display, graphicsContext, C.GXcopy) }
  C.XFlush (display)
}


func Segments (X, Y, X1, Y1 []int, n bool) {
//
  l:= len (X); if len (Y) != l { return }
  s:= make ([]C.XSegment, l)
  for i:= 0; i < l; i++ {
    s[i].x1, s[i].y1, s[i].x2, s[i].y2 = C.short(X[i]), C.short(Y[i]), C.short(X1[i]), C.short(Y1[i])
  }
  if ! n { C.XSetFunction (display, graphicsContext, C.GXinvert) }
  if ! xxb { C.XDrawSegments (display, C.Drawable(window), graphicsContext, &s[0], C.int(l)) }
  C.XDrawSegments (display, C.Drawable(pixmap), graphicsContext, &s[0], C.int(l))
  if ! n { C.XSetFunction (display, graphicsContext, C.GXcopy) }
  C.XFlush (display)
}


func PolygonFull (X, Y []int, n bool) {
//
  l:= len (X); if len (Y) != l { return }
  p:= make ([]C.XPoint, l)
  for i:= 0; i < l; i++ {
    p[i].x, p[i].y = C.short(X[i]), C.short(Y[i])
  }
  if ! n { C.XSetFunction (display, graphicsContext, C.GXcopyInverted) }
  if ! xxb { C.XFillPolygon (display, C.Drawable(window), graphicsContext, &p[0], C.int(l), C.Convex, C.CoordModeOrigin) }
  C.XFillPolygon (display, C.Drawable(pixmap), graphicsContext, &p[0], C.int(l), C.Convex, C.CoordModeOrigin)
  if ! n { C.XSetFunction (display, graphicsContext, C.GXcopy) }
  C.XFlush (display)
}


func Rectangle (x, y, w, h int, n, f bool) {
//
  if f {
    if ! n { C.XSetFunction (display, graphicsContext, C.GXinvert) } // C.GXcopyInverted ? 
    if ! xxb { C.XFillRectangle (display, C.Drawable(window), graphicsContext, C.int(x), C.int(y), C.uint(w), C.uint(h)) }
    C.XFillRectangle (display, C.Drawable(pixmap), graphicsContext, C.int(x), C.int(y), C.uint(w), C.uint(h))
  } else {
    if ! n { C.XSetFunction (display, graphicsContext, C.GXinvert) }
    if ! xxb { C.XDrawRectangle (display, C.Drawable(window), graphicsContext, C.int(x), C.int(y), C.uint(w), C.uint(h)) }
    C.XDrawRectangle (display, C.Drawable(pixmap), graphicsContext, C.int(x), C.int(y), C.uint(w), C.uint(h))
  }
  if ! n { C.XSetFunction (display, graphicsContext, C.GXcopy) }
  C.XFlush (display)
}


func Ellipse (x, y, a, b int, n, f bool) {
//
  x0, y0:= C.int(x - a), C.int(y - b)
  if f {
    if ! n { C.XSetFunction (display, graphicsContext, C.GXinvert) } // C.GXcopyInverted ?
    if ! xxb { C.XFillArc (display, C.Drawable(window), graphicsContext, C.int(x0), C.int(y0), C.uint(2 * a), C.uint(2 * b), 0, 64 * 360) }
    C.XFillArc (display, C.Drawable(pixmap), graphicsContext, C.int(x0), C.int(y0), C.uint(2 * a), C.uint(2 * b), 0, 64 * 360)
  } else {
    if ! n { C.XSetFunction (display, graphicsContext, C.GXinvert) }
    if ! xxb { C.XDrawArc (display, C.Drawable(window), graphicsContext, C.int(x0), C.int(y0), C.uint(2 * a), C.uint(2 * b), 0, 64 * 360) }
    C.XDrawArc (display, C.Drawable(pixmap), graphicsContext, C.int(x0), C.int(y0), C.uint(2 * a), C.uint(2 * b), 0, 64 * 360)
  }
  if ! n { C.XSetFunction (display, graphicsContext, C.GXcopy) }
  C.XFlush (display)
}


var (
  xM, yM int // xbutton.x, xbutton.y // for screen
)


func sendEvents () {
//
  var tst bool
  tst = true
  tst = false
//
  var typ C.int
  var ev *C.XEvent
  if C.XCheckMaskEvent (display, C.ExposureMask + C.StructureNotifyMask + C.SubstructureNotifyMask, &actualEvent) == C.True {
    aT:= C.typ (&actualEvent)
    switch aT { case C.Expose, C.ConfigureNotify:
      for C.XCheckTypedEvent (display, C.int(aT), &actualEvent) == C.True {
        ev = &actualEvent
        typ = C.typ (ev)
  if tst { println ("... got " + txt [typ] + " event") }
//      pp2ff ()
//  C.XFlush (display)
        if typ != C.ConfigureNotify { println ("init sendEvents " + txt [typ]) }
      }
//      pp2ff ()
    }
  }
  var event Event
  xM, yM = 0, 0
// println ("sendEvents before for-loop")
  for {
//    if tst { println ("x is waiting for next event ...") }
    C.XNextEvent (display, &actualEvent)
    event.C, event.S = 0, 0
    ev = &actualEvent
    typ = C.typ (ev)
    if tst { println ("... got " + txt [typ] + " event") }
    switch typ { case KeyPress:
      event.C, event.S = uint(C.xKeyCode (ev)), uint(C.xKeyState (ev))
    case KeyRelease:
      event.C, event.S = uint(C.xKeyCode (ev)), uint(C.xKeyState (ev))
    case ButtonPress:
      event.C, event.S = uint(C.xButtonButton (ev)), uint(C.xButtonState (ev))
      xM, yM = int(C.xButtonX (ev)), int(C.xButtonY (ev))
      if xM < int(xMin) || xM > int(xMax) || yM < int(yMin) || yM > int(yMax) {
        event.C, event.S, xM, yM, typ = 0, 0, 0, 0, C.LASTEvent
      }
    case ButtonRelease:
      event.C, event.S = uint(C.xButtonButton (ev)), uint(C.xButtonState (ev))
      xM, yM = int(C.xButtonX (ev)), int(C.xButtonY (ev))
      if xM < int(xMin) || xM > int(xMax) || yM < int(yMin) || yM > int(yMax) {
        event.C, event.S, xM, yM, typ = 0, 0, 0, 0, C.LASTEvent
      }
    case MotionNotify:
      event.C, event.S = uint(0), uint(C.xMotionState (ev))
      xM, yM = int(C.xMotionX (ev)), int(C.xMotionY (ev))
      if xM < int(xMin) || xM > int(xMax) || yM < int(yMin) || yM > int(yMax) {
        event.C, event.S, xM, yM, typ = 0, 0, 0, 0, C.LASTEvent
      }
    case EnterNotify:
//      pp2ff ()
//      println (txt [typ])
    case LeaveNotify:
//      println (txt [typ])
    case FocusIn:
//      println (txt [typ])
    case FocusOut:
//      println (txt [typ])
    case KeymapNotify:
//      println (txt [typ])
    case Expose:
      pp2ff ()
//      println (txt [typ])
    case GraphicsExpose:
//      println (txt [typ])
    case NoExpose:
//      println (txt [typ])
    case VisibilityNotify:
      println (txt [typ])
    case CreateNotify:
//      println (txt [typ])
    case DestroyNotify:
//      println (txt [typ])
    case UnmapNotify:
//      println (txt [typ])
    case MapNotify:
//  C.XSync (display, C.False)
//  C.XSync (display, C.True)
//      pp2ff ()
    case MapRequest:
//      println (txt [typ])
    case ReparentNotify:
//      println (txt [typ])
    case ConfigureNotify:
//      println (txt [typ])
    case ConfigureRequest:
//      println (txt [typ])
    case GravityNotify:
//      println (txt [typ])
    case ResizeRequest:
//      pp2ff ()
//      println (txt [typ])
    case CirculateNotify:
//      println (txt [typ])
    case CirculateRequest:
//      println (txt [typ])
    case PropertyNotify:
//      println (txt [typ])
    case SelectionClear:
//      println (txt [typ])
    case SelectionRequest:
//      println (txt [typ])
    case SelectionNotify:
//      println (txt [typ])
    case ColormapNotify:
//      println (txt [typ])
    case ClientMessage:
      println (txt [typ])
//      mT:= C.mT (ev)
//      if mT != naviAtom { println ("unknown xclient.message_type ", uint32(mT)) }
    case MappingNotify:
      println (txt [typ])
    case GenericEvent:
      println (txt [typ])
    default:
      println ("strange sort of XEvent: ", typ)
    }
    switch typ { case KeyPress, ButtonPress, ButtonRelease, MotionNotify:
      event.T = uint(typ)
      Eventpipe <- event
    }
  }
}


func MousePos () (int, int) {
//
  if xM < int(xMin) || xM > int(xMax) || yM < int(yMin) || yM > int(yMax) { println ("x.xM, yM: oops") }
  return xM, yM
}


func MouseDef (x, y, x1, y1 int) {
//
  xMin, yMin = C.int(x), C.int(y)
  xMax, yMax = C.int(x1), C.int(y1)
}


func WarpMouse (x, y int) {
//
  C.XWarpPointer (display, C.None, window, 0, 0, 0, 0, C.int(x), C.int(y))
  C.XFlush (display)
}


func Lock() {
//
  C.XLockDisplay (display)
}


func Unlock() {
//
  C.XUnlockDisplay (display)
}


var
  clz = int(obj.Codelen (uint(0)))


func cd (x uint) []byte {
//
  const m = uint(256)
  b:= make ([]byte, 4)
  b[0] = byte(x % m)
  x = x / m
  b[1] = byte(x % m)
  x = x / m
  b[2] = byte(x % m)
  x = x / m
  b[3] = byte(x % m)
  return b
}


func Codelen (x, y uint) uint {
//
  return 4 * x * y
}


func Encode (x, y, w, h uint, e []byte) {
//
  const M = C.ulong(1 << 32 - 1)
  ximg:= C.XGetImage (display, C.Drawable(window), C.int(x), C.int(y), C.uint(w), C.uint(h), M, C.XYPixmap)
  n:= uint(0)
//  var pixel C.ulong
  for y:= 0; y < int(h); y++ {
    for x:= 0; x < int(w); x++ {
      pixel:= C.xGetPixel (ximg, C.int(x), C.int(y)) // = C.XGetPixel (ximg, C.int(x), C.int(y))
      copy (e[n:n+4], cd (uint(pixel)))
      n += 4
    }
  }
  C.xDestroyImage (ximg) // C.XDestroyImage (ximg)
}


func Decode (x, y, w, h uint, b []byte) {
//
  const M = C.ulong(1 << 32 - 1)
//////////////////////////////////////////////////////////////////////////////////////////////  steals a lot of time
  ximg:= C.XGetImage (display, C.Drawable(window), C.int(x), C.int(y), C.uint(w), C.uint(h), M, C.XYPixmap)
//////////////////////////////////////////////////////////////////////////////////////////////  steals a lot of time
  n:= uint(0)
//  var pixel C.ulong
  for j:= uint(0); j < h; j++ {
    for i:= uint(0); i < w; i++ {
      pixel:= (C.ulong)(obj.Decode (n, b[n:n+4]).(uint))
      C.xPutPixel (ximg, C.int(i), C.int(j), pixel) // C.XPutPixel (ximg, C.int(i), C.int(j), pixel)
      n += 4
    }
  }
  C.XPutImage (display, C.Drawable(window), graphicsContext, ximg, 0, 0, C.int(x), C.int(y), C.uint(w), C.uint(h))
  C.XCopyArea (display, C.Drawable(window), C.Drawable(pixmap), graphicsContext, C.int(x), C.int(y), C.uint(w), C.uint(h), C.int(x), C.int(y))
  C.XFlush (display)
  C.xDestroyImage (ximg) // C.XDestroyImage (ximg)
}


func P6Codelen (w, h uint) uint {
//
  return w * h * col.P6
}


func P6Encode (w, h uint, b, p []byte) {
//
  i, j:= uint(0), uint(0)
  di:= uint(clz)
  for y:= uint(0); y < h; y++ {
    for x:= uint(0); x < w; x++ {
      col.P6Encode (b[i:i+di], p[j:j+col.P6])
      i += di
      j += col.P6
    }
  }
}


func P6Decode (x, y, w, h uint, p, b []byte) {
//
  var c col.Colour
  i, j:= uint(0), uint(0)
  di:= uint(clz)
  for y:= uint(0); y < h; y++ {
    for x:= uint(0); x < w; x++ {
      col.Decode (&c, p[j:j+col.P6])
      copy (b[i:i+di], obj.Encode (col.Code (c)))
      i += di
      j += col.P6
    }
  }
}


func WriteGlx() {
//
  C.write (display, window)
}


func init () {
//
  txt [KeyPress] =         "KeyPress"
  txt [KeyRelease] =       "KeyRelease"
  txt [ButtonPress] =      "ButtonPress"
  txt [ButtonRelease] =    "ButtonRelease"
  txt [MotionNotify] =     "MotionNotify"
  txt [EnterNotify] =      "EnterNotify"
  txt [LeaveNotify] =      "LeaveNotify"
  txt [FocusIn] =          "FocusIn"
  txt [FocusOut] =         "FocusOut"
  txt [KeymapNotify] =     "KeymapNotify"
  txt [Expose] =           "Expose"
  txt [GraphicsExpose] =   "GraphicsExpose"
  txt [NoExpose] =         "NoExpose"
  txt [VisibilityNotify] = "VisibilityNotify"
  txt [CreateNotify] =     "CreateNotify"
  txt [DestroyNotify] =    "DestroyNotify"
  txt [UnmapNotify] =      "UnmapNotify"
  txt [MapNotify] =        "MapNotify"
  txt [MapRequest] =       "MapRequest"
  txt [ReparentNotify] =   "ReparentNotify"
  txt [ConfigureNotify] =  "ConfigureNotify"
  txt [ConfigureRequest] = "ConfigureRequest"
  txt [GravityNotify] =    "GravityNotify"
  txt [ResizeRequest] =    "ResizeRequest"
  txt [CirculateNotify] =  "CirculateNotify"
  txt [CirculateRequest] = "CirculateRequest"
  txt [PropertyNotify] =   "PropertyNotify"
  txt [SelectionClear] =   "SelectionClear"
  txt [SelectionRequest] = "SelectionRequest"
  txt [SelectionNotify] =  "SelectionNotify"
  txt [ColormapNotify] =   "ColormapNotify"
  txt [ClientMessage] =    "ClientMessage"
  txt [MappingNotify] =    "MappingNotify"
  txt [GenericEvent] =     "GenericEvent"
  txt [LASTEvent] =        "LASTEvent"
  if C.XInitThreads() == 0 { panic ("XKern.XInitThreads error") }
}
