package cs

// (c) Christian Maurer   v. 120910 - license see murus.go

//     Nichtsequentielle Programmierung mit Go 1 kompakt, S. 85 ff.

import
  . "murus/obj"

/* Universal Synchronizations for the construction of protocols to enter
   general conditioned kritical sections with use of commen ressources,
   whose state could be changed, of processes of at most two classes.
   General conditioned critical sections can be entered concurrently
   by several processes of one class, but by processes of different classes
   only under mutual exclusion.
   The classes are identified by natural numbers, starting with 0.
   Mutual exclusion of processes of different classes is guaranteed by clients
   in form of boolean expressions, under which a critical section may be entered;
   in the enter and leave protocol operations are executed,
   that control the conditions.
   The functions Enter and Leave cannot be interrupted by calls of these functions
   of other processes. */

type
  CriticalSection interface {
// x means in the following always the calling critical section.

// Pre: k < number of classes of x.
//      The function is called within the entry conditions of x (see remark).
// Returns true, iff at least one process of the k-th class of x is blocked at the moment of the call.
// Bemark: The result can be different immediately after the call.
//         To rely on it, it is necessary, that the atomicity of the call
//         is ensured by reasonable synchronization >>> maßnahmen <<<.
//         This is the case for a call within the entry conditions of x.
  Blocked (k uint) bool

// Pre: k < number of classes of the x.
//      The calling process is not in x.
// It is now in the k-th class of x, i.e. it was eventually blocked, until c(k) was true,
// and now e(a, k) is executed (where c is the entry condition of x
// and e the processing during the entry into x).
  Enter (k uint, a Any)

// Pre: k < number of classes of the x.
//      The calling process is in the k-th class of x.
// It is now not any more in x, i.e. p(a, k) has been executed
// (where p is the processing at the exit from x, 
// and k the class of x, in which the calling process was before.)
// c(k) of x is no longer ensured.
  Leave (k uint, a Any)
}
