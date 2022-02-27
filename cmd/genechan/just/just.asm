TEXT main.main(SB) /Users/maxim/go/src/github.com/MistaTwista/generigo/cmd/genechan/just/main.go
func main() {
  0x1054c40		493b6610		CMPQ 0x10(R14), SP	
  0x1054c44		7659			JBE 0x1054c9f		
  0x1054c46		4883ec48		SUBQ $0x48, SP		
  0x1054c4a		48896c2440		MOVQ BP, 0x40(SP)	
  0x1054c4f		488d6c2440		LEAQ 0x40(SP), BP	
	_ = genechan.Processor([]int{1, 2, 3}, 2, 2)
  0x1054c54		440f117c2428		MOVUPS X15, 0x28(SP)							
  0x1054c5a		440f117c2430		MOVUPS X15, 0x30(SP)							
  0x1054c60		48c744242801000000	MOVQ $0x1, 0x28(SP)							
  0x1054c69		48c744243002000000	MOVQ $0x2, 0x30(SP)							
  0x1054c72		48c744243803000000	MOVQ $0x3, 0x38(SP)							
  0x1054c7b		488d442428		LEAQ 0x28(SP), AX							
  0x1054c80		bb03000000		MOVL $0x3, BX								
  0x1054c85		4889d9			MOVQ BX, CX								
  0x1054c88		bf02000000		MOVL $0x2, DI								
  0x1054c8d		4889fe			MOVQ DI, SI								
  0x1054c90		e86bf9ffff		CALL github.com/MistaTwista/generigo/internal/genechan.Processor(SB)	
}
  0x1054c95		488b6c2440		MOVQ 0x40(SP), BP	
  0x1054c9a		4883c448		ADDQ $0x48, SP		
  0x1054c9e		c3			RET			
func main() {
  0x1054c9f		90			NOPL					
  0x1054ca0		e85bc5ffff		CALL runtime.morestack_noctxt.abi0(SB)	
  0x1054ca5		eb99			JMP main.main(SB)			
