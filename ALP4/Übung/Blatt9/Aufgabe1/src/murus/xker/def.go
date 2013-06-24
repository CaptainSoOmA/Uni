package xker

// (c) Christian Maurer   v. 120715 - license see murus.go

// >>> This package only serves the implementations of murus/mouse, 
//     murus/kbd and murus/scr; it must not no be used elsewhere.

const ( // see /usr/include/X11/X.h
  KeyPress = 2 + iota; KeyRelease
  ButtonPress; ButtonRelease; MotionNotify
  EnterNotify; LeaveNotify; FocusIn; FocusOut
  KeymapNotify; Expose; GraphicsExpose; NoExpose
  VisibilityNotify; CreateNotify; DestroyNotify
  UnmapNotify; MapNotify; MapRequest; ReparentNotify
  ConfigureNotify; ConfigureRequest; GravityNotify
  ResizeRequest; CirculateNotify; CirculateRequest; PropertyNotify
  SelectionClear; SelectionRequest; SelectionNotify
  ColormapNotify; ClientMessage; MappingNotify
  GenericEvent; LASTEvent
)
type
  Event struct {
             T,     // type
             C,     // xkey.keycode, xbutton.button, xmotion.is_hint
             S uint // state
        }
var
  Eventpipe chan Event
