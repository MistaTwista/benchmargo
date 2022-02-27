TEXT main.main(SB) /Users/maxim/go/src/github.com/MistaTwista/generigo/cmd/genechan/gen/main.go
func main() {
  0x1054600		493b6610		CMPQ 0x10(R14), SP	
  0x1054604		7660			JBE 0x1054666		
  0x1054606		4883ec50		SUBQ $0x50, SP		
  0x105460a		48896c2448		MOVQ BP, 0x48(SP)	
  0x105460f		488d6c2448		LEAQ 0x48(SP), BP	
	_ = genechan.GeneProcessor[int]([]int{1, 2, 3}, 2, 2)
  0x1054614		440f117c2430		MOVUPS X15, 0x30(SP)										
  0x105461a		440f117c2438		MOVUPS X15, 0x38(SP)										
  0x1054620		48c744243001000000	MOVQ $0x1, 0x30(SP)										
  0x1054629		48c744243802000000	MOVQ $0x2, 0x38(SP)										
  0x1054632		48c744244003000000	MOVQ $0x3, 0x40(SP)										
  0x105463b		488d055e2b0200		LEAQ github.com/MistaTwista/generigo/internal/genechan..dict.GeneProcessor[int](SB), AX		
  0x1054642		488d5c2430		LEAQ 0x30(SP), BX										
  0x1054647		b903000000		MOVL $0x3, CX											
  0x105464c		4889cf			MOVQ CX, DI											
  0x105464f		be02000000		MOVL $0x2, SI											
  0x1054654		4989f0			MOVQ SI, R8											
  0x1054657		e804030000		CALL github.com/MistaTwista/generigo/internal/genechan.GeneProcessor[go.shape.int_0](SB)	
}
  0x105465c		488b6c2448		MOVQ 0x48(SP), BP	
  0x1054661		4883c450		ADDQ $0x50, SP		
  0x1054665		c3			RET			
func main() {
  0x1054666		e895cbffff		CALL runtime.morestack_noctxt.abi0(SB)	
  0x105466b		eb93			JMP main.main(SB)			

  0x105466d		cc			INT $0x3		
  0x105466e		cc			INT $0x3		
  0x105466f		cc			INT $0x3		
  0x1054670		cc			INT $0x3		
  0x1054671		cc			INT $0x3		
  0x1054672		cc			INT $0x3		
  0x1054673		cc			INT $0x3		
  0x1054674		cc			INT $0x3		
  0x1054675		cc			INT $0x3		
  0x1054676		cc			INT $0x3		
  0x1054677		cc			INT $0x3		
  0x1054678		cc			INT $0x3		
  0x1054679		cc			INT $0x3		
  0x105467a		cc			INT $0x3		
  0x105467b		cc			INT $0x3		
  0x105467c		cc			INT $0x3		
  0x105467d		cc			INT $0x3		
  0x105467e		cc			INT $0x3		
  0x105467f		cc			INT $0x3		
