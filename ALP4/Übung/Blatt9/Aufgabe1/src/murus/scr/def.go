package scr

// (c) Christian Maurer   v. 130406 - license see murus.go

/* Pre: For use in a console:
          The framebuffer is usable, i.e. one of the options "vga=..."
          is contained in the line "kernel ..." of /boot/grub/menu.lst
          (posible values can be found at the end of imp.go).
          Users are in the group video (or world has the rights "r" and "w"
          in /dev/fb0) and world has the right "r" in /dev/input/mice.
        For use in a window on a graphical user interface:
          X ist installed.
        Programs for execution on far hosts are only called under X.
  
   Fore-/background colour of the screen and actual fore-/backgroundcolour
   are White and Black. The screen is cleared and the cursor is off.
   SIGUSR1 and SIGUSR2 are used internally and not any more available.
   No process is in the exclusive possession of the screen. */

import (
  "murus/col"; "murus/font"
)
type
  Mode byte; const (
  QVGA = Mode(iota)
         //  15 x  40 /  320 x  240
  HVGA   //  20 x  60 /  480 x  320
  TXT    //  25 x  80 /  640 x  400 // HSVGA
  VGA    //  30 x  80 /  640 x  480
  PAL    //  36 x  96 /  768 x  576
  WVGA   //  30 x 100 /  800 x  480
  SVGA   //  37 x 100 /  800 x  600
  WPAL   //  37 x 120 /  960 x  600
  WVGApp //  36 x 128 / 1024 x  576
  WSVGA  //  37 x 128 / 1024 x  600
  XGA    //  48 x 128 / 1024 x  768
  HD     //  45 x 160 / 1280 x  720
  WXGA   //  50 x 160 / 1280 x  800
  SXVGA  //  60 x 160 / 1280 x  960
  SXGA   //  64 x 160 / 1280 x 1024
  WXGA1  //  48 x 171 / 1366 x  768
  SXGAp  //  65 x 175 / 1400 x 1050
  WXGAp  //  56 x 180 / 1440 x  900
  WXGApp //  56 x 200 / 1600 x  900
  WSXGA  //  64 x 200 / 1600 x 1024
  UXGA   //  75 x 200 / 1600 x 1200
  WSXGAp //  65 x 210 / 1680 x 1050
  FHD    //  67 x 240 / 1920 x 1080
  WUXGA  //  75 x 240 / 1920 x 1200
  SUXGA  //  90 x 240 / 1920 x 1440
  QWXGA  //  72 x 256 / 2048 x 1152
  QXGA   //  96 x 256 / 2048 x 1536
  WQHD   //  90 x 320 / 2560 x 1440
  WQXGA  // 100 x 320 / 2560 x 1600
  QSXGA  // 128 x 320 / 2560 x 2048
  QSXGAp // 131 x 350 / 2800 x 2100
  WQSXGA // 128 x 400 / 3200 x 2048
  QUXGA  // 150 x 400 / 3200 x 2400
  QHD    // 135 x 440 / 3840 x 2160
  QWUXGA // 150 x 480 / 3840 x 2400 // WQUXGA
  HXGA   // 192 x 512 / 4096 x 3072
  WHXGA  // 200 x 640 / 5120 x 3200
  HSXGA  // 256 x 640 / 5120 x 4096
  WHSXGA // 256 x 800 / 6400 x 4096
  HUXGA  // 300 x 800 / 6400 x 4800
  UHDV   // 270 x 960 / 7680 x 4320
  WHUXGA // 300 x 960 / 7680 x 4800
  NModes )


// Returns true, iff the calling process runs under X (i.e., in a GUI).
func UnderX () bool { return underX }

// modes and sizes /////////////////////////////////////////////////////

// Returns the maximal possible mode of the screen.
func MaxMode () Mode { return maxMode }

// Returns the (X, Y)-resolution of the screen in pixels.
func Res () (uint, uint) { return res() }

// Returns the default mode for the screen, which is used, if
// the screen actually used does not conform to any of the above modes.
// Furthermore, this mode is the start mode under X.
func DefaultMode () Mode { return defaultMode }

// Returns true, iff NLines/NColumns of m <= NLines/NColumns of the maximal mode.
func Switchable (m Mode) bool { return switchable(m) }

// If m is not switchable, nothing has happened.
// Otherwise, the actual mode is now m; the actual fontsize is normal;
// the colours of the screen are not changed,
// actual fore-/backgroundcolour is the fore-/backgroundcolour of the screen.
// The screen is cleared.
func Switch (m Mode) { switch_(m) }

func Fullscreen() { switch_(maxMode) }

// Returns the actual mode.
func ActMode () Mode { return mode }

// Returns - depending on the actual fontsize -
// the number of textlines and -columns of the actual mode.
func NLines () uint { return nLines }
func NColumns () uint { return nColumns }

// Returns the pixelwidth/-height of the screen in the actual mode.
func NX () uint { return nX[mode] }
func NY () uint { return nY[mode] }

// Returns the pixel distance between two textlines
// = charheight/-width of the actual fontsize (s. below).
func NX1 () uint { return actualCharwidth }
func NY1 () uint { return actualCharheight }

// Return the relation Pixelwidth : Pixelheight of the actual mode.
func Proportion () float64 { return proportion() }

// colours /////////////////////////////////////////////////////////////

// The actual foregroundcolour is f, the actual backgroundcolour is b
// resp. that of the screen.
// The colours of the screen are not changed.
func Colours (f, b col.Colour) { colours(f,b) }
func Colour (f col.Colour) { colour(f) }

// Returns the number of the representable colours in the actual mode.
func NColours () uint { return nColours() }

// ranges //////////////////////////////////////////////////////////////

// Pre: S + B <= NColumns, Z + H <= NLines resp.
//      x <= x1 < NX, y <= y1 < NY.

// The screen is cleared in its backgroundcolour.
// The cursor has the position (0, 0) and is off. 
// // If there exists a mouse, its cursor has the position (?, ?) and is off.
func Cls () { cls() }

// The screen is cleared between line l and l+h and column c and c+w
// (both including) in its backgroundcolour.
func Clr (l, c, w, h uint) { clear(l,c,w,h) }

// The pixels in the rectangle between (x, y) and (x1, y1)
// (both including) have the backgroundcolour of the screen.
func ClearGr (x, y, x1, y1 uint) { clearGr(x,y,x1,y1) }

// TODO Spec
func Invert (l, c, w, h uint) { invert(l,c,w,h) }
func InvertGr (x, y, x1, y1 uint) { invertGr(x,y,x1,y1) }

// If on, then the screen buffer is cleared and
// all further output is only going to the screen buffer,
// otherwise, the screen contains the content of the screen buffer
// and all further output is going to the screen.
func Buf (on bool) { buf(on) }

// Returns true, iff the output goes only to the screen buffer.
func Buffered () bool { return only2Buf }

// The content of the screen between line l and l+h and column c and c+w
// is copied into the archive (the former content of the archive is lost).
func Save (l, c, w, h uint) { save(l,c,w,h) }
func SaveGr (x, y, x1, y1 uint) { saveGr(x,y,x1,y1) }

// The content of the screen between line l and l+h and column c and c+w
// is restored from the archive.
func Restore (l, c, w, h uint) { restore(l,c,w,h) }
func RestoreGr (x, y, x1, y1 uint) { restoreGr(x,y,x1,y1) }

// cursor //////////////////////////////////////////////////////////////

type
  Shape byte; const (
  Off = Shape(iota)
  Understroke
  Block
  NShapes )

// Pre: l < NLines, c < NColumns.
// The cursor has the position (line, coloumn) == (l, c)
// and the shape s. (0, 0) is the top left top corner.
func Warp (l, c uint, s Shape) { warp(l,c,s) }

// Pre: x <= NColumsGr - Columnwidth, y <= NY - Lineheight.
// The cursor has the graphics position (column, line) = (x, y)
// and the shape s. (0, 0) is the top left top corner.
func WarpGr (x, y uint, s Shape) { warpGr(x,y,s) }

// text ////////////////////////////////////////////////////////////////

// Returns true, iff transparency is switched on.
func TransparenceOn () bool { return transparent }

// Transparence is switched on, iff t == true.
func SwitchTransparence (t bool) { transparent = t }

// The position (0, 0) is the top left corner of the screen.
// The pixels of the characters have the actual foregroundcolour,
// the pixels in the rectangles around them have the actual backgroundcolour
// (if transparency is switched on, those pixels are not changed).

// Pre: 32 <= b < 127, l < NLines, c + 1 < NColumns. 
// b is written to the screen at position (line, colum) = (l, c). 
func Write1 (b byte, l, c uint) { write1(b,l,c) }

// Pre: l < NLines, c + len(s) < NColumns. 
// s is written to the screen starting at position (line, column) == (l, c).
func Write (s string, l, c uint ) { write(s,l,c) }

// Pre: c + number of digits of n < NColumns, l < NLines.
// n is written to the screen starting at position (line, column) == (l, c).
func WriteNat (n, l, c uint) { writeNat(n,l,c) }

// Pre: x + Columnwidth < NX resp.
//      x + Columnwidth * Länge (s) < NX,
//      y + Lineheight < NY.
// b resp. s is written to the screen within the rectangle
// with the top left corner (x, y).
func Write1Gr (b byte, x, y int) { write1Gr(b,x,y) }
func WriteGr (s string, x, y int) { writeGr(s,x,y) }

// TODO Spec
func Write1InvGr (b byte, x, y int) { write1InvGr(b,x,y) }
func WriteInvGr (s string, x, y int) { writeInvGr(s,x,y) }

// font ////////////////////////////////////////////////////////////////

// Returns the actual fontsize; at the beginning normal.
func ActFontsize () font.Size { return actualFontsize }

// f is the actual fontsize.
// NColumns and NLines are changed accordingly.
func SwitchFontsize (f font.Size) { switchFontsize(f) }

// graphics ////////////////////////////////////////////////////////////

// Position (0, 0) is the top left corner of the screen.
// All output is done in the actual foregroundcolour;
// For operations with name ...Inv all pixels have the complementary
// colour of the fgcolour; for operations with name ...Full
// also all pixels in the interior have these colours.
// The actual linewidth at the beginning is Thin.

type
  Linewidth byte; const (
  Thin = Linewidth(iota)
  Thicker
  Yetthicker )

// Returns the actual linewidth.
func ActLinewidth () Linewidth { return actualLinewidth }

// The actual linewidth is w.
func SetLinewidth (w Linewidth) { setLinewidth(w) }

// Pre: See above.
// A pixel in the actual foregroundcolour is set at position (x, y)
// on the screen resp. the colour of that pixel is inverted.
func Point (x, y int) { point(x,y) }
func PointInv (x, y int) { pointInv(x,y) }

// Pre: See above.
// At (x[i], y[i]) (i < len(x) == len(y)) a pixel is set in the actual
// foregroundcolour resp. that pixel is inverted in its colour.
func Pointset (x, y []int) { pointset(x,y) }
func PointsetInv (x, y []int) { pointsetInv(x,y) }

// Returns the colour of the pixel at (x, y).
func PointColour (x, y uint) col.Colour { return pointColour(x,y) }

// Pre: See above.
// The part of the line segment between (x, y) and (x1, y1)
// visible on the screen is drawn in the actual foregroundcolour resp.
// the pixels on that part are inverted in their colour.
func Line (x, y, x1, y1 int) { line(x,y,x1,y1) }
func LineInv (x, y, x1, y1 int) { lineInv(x,y,x1,y1) }

// Pre: See above.
// Returns true, iff the point at (x, y) has a distance of
// at most t pixels from the line segment between (x, y) to (x1, y1).
func OnLine (x, y, x1, y1, a, b int, t uint) bool { return onLine(x,y,x1,y1,a,b,t) }

// Pre: See above.
//      If the calling process runs under X:
//        -1<<15 <= x[i], x1[i], y[i], y1[i] < 1<<15
//        for all i < n:= len(x) == len(y).
//      Otherwise:
//        0 <= x[i], x1[i] < NX and
//        0 <= y[i], y1[i] < NY for all i < N.
// For all i < n the parts of the line segments between (x[i], y[i]) and (x1[i], y1[i]),
// that are visible on the screen, are drawn in the actual foregroundcolour
// resp. all points on them are inverted.
func Lines (x, y, x1, y1 []int) { lines(x,y,x1,y1) }
func LinesInv (x, y, x1, y1 []int) { linesInv(x,y,x1,y1) }

// Pre: See above.
// TODO Spec
func OnLines (x, y, x1, y1 []int, a, b int, t uint) bool { return onLines(x,y,x1,y1,a,b,t) }

// Pre: See above.
//      x[i] < NX, y[i] < NY für alle i < n:= len(x) == len(y).
// From (x[0], y[0]) over (x[1], y[1]), ... until (x[n-1], y[n-1])
// a sequence of line segments is drawn resp. all points on it are inverted.
func Segments (x, y []int) { segments(x, y) }
func SegmentsInv (x, y []int) { segmentsInv(x, y) }

// Returns true, iff the point at (a, b) has a distance of at most t pixels
// from one of the sequence of line segments defined by x and y.
func OnSegments (x, y []int, a, b int, t uint) bool { return onSegments(x,y,a,b,t) }

// Pre: See above.
// A line through (x, y) and (x1, y1) is drawn resp. all points on it are inverted.
func InfLine (x, y, x1, y1 int) { infLine(x, y, x1, y1) }
func InfLineInv (x, y, x1, y1 int) { infLineInv(x, y, x1, y1) }

// Returns true, iff the point at (a, b) has a distance of at most t pixels
// from the line through (x, y) and (x1, y1).
func OnInfLine (x, y, x1, y1, a, b int, t uint) bool { return onInfLine(x,y,x1,y1,a,b,t) }

// Pre: See above.
// Between (x, y) and (x1, y1) a rectangle (with horizontal and vertical borders)
// is drawn in the actual foregroundcolour resp. all points on it are inverted
// resp. all its interior points (including its borders) are drawn / inverted.
func Rectangle (x, y, x1, y1 int) { rectangle(x,y,x1,y1) }
func RectangleInv (x, y, x1, y1 int) { rectangleInv(x,y,x1,y1) }
func RectangleFull (x, y, x1, y1 int) { rectangleFull(x,y,x1,y1) }
func RectangleFullInv (x, y, x1, y1 int) { rectangleFullInv(x,y,x1,y1) }

// Pre: See above.
// Returns true, iff the point at (a, b) has a distance of at most t pixels
// from the border of the rectangle between (x, y) and (x1, y1).
func OnRectangle (x, y, x1, y1, a, b int, t uint) bool { return onRectangle(x,y,x1,y1,a,b,t) }

// Returns true, iff the point at (a, b) is not outside the rectangle between (x, y) and (x1, y1).
func InRectangle (x, y, x1, y1, a, b int) bool { return inRectangle(x,y,x1,y1,a,b) }

// Pre: See above. For n:= len(x) == len(y): n > 2 and
//      PolygonFull:
//        The calling process runs under X;
//        the polygon defined by x and y is convex and drawn in the same colour.
//      PolygonFull1:
//        (x0, y0) lies in the interior of the polygon defined by x and y.
//        The polygon is drawn in the same colour.
// A polygon is drawn between (x[0], y[0]), (x[1], y[1]), ... (x[n-1], y[n-1), (x[0], y[0])
// resp. all pixels on it are inverted resp. the polygon is filled.
func Polygon (x, y []int) { polygon(x, y) }
func PolygonInv (x, y []int) { polygonInv(x, y) }
// func PolygonFull (x, y []int) { polygonFull(x, y) } // TODO
// func PolygonFull1 (x, y []int, x0, y0 int, n uint) { polygonFull1(x,y,x0,y0,n) } // TODO

// Returns true, iff the point at (a, b) has a distance of at most t pixels
// from the polyon defined by x and y.
func OnPolygon (x, y []int, a, b int, t uint) bool { return onPolygon(x,y,a,b,t) }

// Pre: See above. r <= x, x + r < NX, r <= y, y + r < NY. 
// Around (x, y) a circle with radius r is drawn / inverted
// resp. all points in its interior are set / inverted.
// x + r < NX, 0 <= r <= y, y + r < NY.
func Circle (x, y int, r uint) { circle(x,y,r) }
func CircleInv (x, y int, r uint) { circleInv(x,y,r) }
func CircleFull (x, y int, r uint) { circleFull(x,y,r) }
func CircleFullInv (x, y int, r uint) { circleFullInv(x, y, r) }

// Returns true, iff the point at (x, y) has a distance of at most t pixels
// from the border of the circle around (a, b) with radius r.
func OnCircle (x, y int, r uint, a, b int, t uint) bool { return onCircle(x,y,r,a,b,t) }
// func InCircle (x, y int, r uint, a, b int) bool { return inCircle(x,y,r,a,b) } // TODO

// Pre: See above. a <= x, x + a < NX, b <= y, y + b < NY. 
// Around (x, y) an ellipse with horizontal / vertical semiaxis a / b
// is drawn / inverted resp. all points in its interior are set / inverted.
func Ellipse (x, y int, a, b uint) { ellipse(x,y,a,b) }
func EllipseInv (x, y int, a, b uint) { ellipseInv(x,y,a,b) }
func EllipseFull (x, y int, a, b uint) { ellipseFull(x,y,a,b) }
func EllipseFullInv (x, y int, a, b uint) { ellipseFullInv(x,y,a,b) }

// Returns true, iff the point at (A, B) has a distance of at most t pixels
// from the border of the ellipse around (x, y) with semiaxis a and b.
func OnEllipse (x, y int, a, b uint, A, B int, t uint) bool { return onEllipse(x,y,a,b,A,B,t) }
// func InEllipse (x, y int, a, b uint, A, B int) bool { return inEllipse(x,y,a,b,A,B) } // TODO

// Pre: See above. n:= len(x) == len(Y).
// From (x[0], y[0]) to (x[n], y[n]) a Beziercurve of order n
// with (x[1], y[1]) .. (x[n-1], y[n-1]) as nodes is drawn to the screen
// resp. all points on that curve are inverted.
// (For n == 0 the curve is the point (x[0], y[0]),
// for n == 1 the line between (x[0], y[0]) and (x[1], y[1]).
func Curve (x, y []int) { curve(x,y) }
func CurveInv (x, y []int) { curveInv(x,y) }

// Returns true, iff the point at (x, y) has a distance of at most t pixels
// from the curve defined by x and y.
func OnCurve (x, y []int, a, b int, t uint) bool { return onCurve(x,y,a,b,t) }

// mouse ///////////////////////////////////////////////////////////////////////
// Pre: Mouse is installed.

// Pre: If called in a tty-console, /dev/input/mice is readable for the world.
// Returns true, iff there is a mouse installed.
func MouseEx () bool { return mouseEx() }

// Pre: c + w <= NColumns, l + h <= NLines.
// The mouse cursor is bounded by the rectangle
// defined by (line, column, width, height) == (l, c, w, h).
func MouseDef (l, c, w, h uint) { mouseDef(l,c,w,h) }

// Pre: 0 <= x < x + w < NX, 0 <= y < y + h < NY.
// The mouse cursor is bounded by the rectangle
// defined by (column, row, width, height) == (x, y, w, h).
func MouseDefGr (x, y, w, h int) { mouseDefGr(x,y,w,h) }

// Returns the position of the mouse cursor.
// For the result (l, c) holds 0 <= l < NLines and 0 <= c < NColumns.
func MousePos () (uint, uint) { return mousePos() }

// Returns the position of the mouse cursor.
// For the result (x, y) holds 0 <= x < NX and 0 <= y < NY.
func MousePosGr () (int, int) { return mousePosGr() }

// Pre: l < NLines, c < NColumns.
// The mouse cursor has the position (line, column) = (l, c).
func WarpMouse (l, c uint) { warpMouse(l,c) }

// Pre: 0 <= x < NX, 0 <= y < NY.
// The mouse cursor has the position (row, line) = (x, y).
func WarpMouseGr (x, y int) { warpMouseGr(x,y) }

// Pre: The calling process does not run under X.
// If no mouse exists, nothing has happened.
// Otherwise, the mouse cursor is switched on, iff b (otherwise off).
func MouseCursor (b bool) { mouseCursor(b) }

// Pre: The calling process does not run under X.
// Returns true, iff the mouse cursor is switched on.
func MouseCursorOn () bool { return mouseCursorOn() }

// Pre: c + w <= NColumns, l + h <= NLines.
// Returns false, if there is no mouse; returns otherwise true,
// iff the the mouse cursor is in the interior of the rectangle
// defined by l, c, w, h.
func UnderMouse (l, c, w, h uint) bool { return underMouse(l,c,w,h) }

// Pre: 0 <= x <= x1 < NX, 0 <= y <= y1 < NY.
// Returns false, if there is no mouse; returns otherwise true,
// iff the mouse cursor is inside the rectangle between (x, y) and (x1, y1)
// or has a distance of at most t pixels from its boundary.
func UnderMouseGr (x, y, x1, y1 int, t uint) bool { return underMouseGr(x,y,x1,y1,t) }

// serialisation ///////////////////////////////////////////////////////////////

// Pre: 0 < w <= NX, 0 < h <= NY.
// Returns the number of bytes, that are needed to serialize
// the pixels of a rectangle of the size w * h uniquely invertibly.
func Codelen (w, h uint) uint { return codelen(w,h) }

// Pre: 0 < w, x + w < NX, 0 < h, y + h < NY.
// Returns the byte sequence, that serializes the pixels
// in the rectangle between (x, y) and (x + w, y + h).
func Encode (x, y, w, h uint) []byte { return encode(x,y,w,h) }

// Pre: B is the result of a call of Encode for some rectangle.
// The pixels of that rectangle are drawn on the screen;
// the rest of the screen is not changed.
func Decode (B []byte) { decode(B) }

// ppm-serialisation ///////////////////////////////////////////////////////////

// TODO Spec
func P6Codelen (w, h uint) uint { return p6Codelen(w,h) }

// TODO Spec
func P6Size (P []byte) (uint, uint) { return p6Size(P) }

// TODO Spec
func P6Encode (x, y, w, h uint) []byte { return p6Encode(x,y,w,h) }

// TODO Spec
func P6Decode (x, y uint, P6 []byte) { p6Decode(x,y,P6) }

// synchronisation /////////////////////////////////////////////////////////////

// Lock / Unlock guarantee the mutual exclusion of writing to the screen
// (e.g. to avoid, that a process after having set its colours
// is interrupted in a subsequent draw and later resumes its drawing
// in another colour, that was meanwhile set by another process).
func Lock () { lock() }
func Unlock () { unlock() }

// openGL /////////////////////////////////////////////////////////////////////

// - only for use in murus/gl
func WriteGlx () { writeGlx() }
