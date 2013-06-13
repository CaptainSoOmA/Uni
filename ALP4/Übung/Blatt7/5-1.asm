TEXT TestAndSet(SB),7,$0
	MOVL valptr+0(FP), BP 	// BP = &b
	MOVL $1, AX				// AX = 1
	LOCK					// Bus sperren 
	XCHGL AX, 0(BP)			// AX = *b || *b =1 (=true)
	MOVL AX, ret+4(FP)		// AX ist Rueckgabewert
	RET