	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 10, 14	sdk_version 10, 14
	.intel_syntax noprefix
	.globl	_sum                    ## -- Begin function sum
	.p2align	4, 0x90
_sum:                                   ## @sum
	.cfi_startproc
## %bb.0:
	push	rbp
	.cfi_def_cfa_offset 16
	.cfi_offset rbp, -16
	mov	rbp, rsp
	.cfi_def_cfa_register rbp
	test	esi, esi
	jle	LBB0_1
## %bb.2:
	mov	esi, esi
	lea	rcx, [rsi - 1]
	mov	eax, esi
	and	eax, 7
	cmp	rcx, 7
	jae	LBB0_8
## %bb.3:
	vxorpd	xmm0, xmm0, xmm0
	xor	ecx, ecx
	test	rax, rax
	jne	LBB0_5
	jmp	LBB0_7
LBB0_1:
	vxorps	xmm0, xmm0, xmm0
	vmovsd	qword ptr [rdx], xmm0
	pop	rbp
	ret
LBB0_8:
	sub	rsi, rax
	vxorpd	xmm0, xmm0, xmm0
	xor	ecx, ecx
	.p2align	4, 0x90
LBB0_9:                                 ## =>This Inner Loop Header: Depth=1
	vaddsd	xmm0, xmm0, qword ptr [rdi + 8*rcx]
	vaddsd	xmm0, xmm0, qword ptr [rdi + 8*rcx + 8]
	vaddsd	xmm0, xmm0, qword ptr [rdi + 8*rcx + 16]
	vaddsd	xmm0, xmm0, qword ptr [rdi + 8*rcx + 24]
	vaddsd	xmm0, xmm0, qword ptr [rdi + 8*rcx + 32]
	vaddsd	xmm0, xmm0, qword ptr [rdi + 8*rcx + 40]
	vaddsd	xmm0, xmm0, qword ptr [rdi + 8*rcx + 48]
	vaddsd	xmm0, xmm0, qword ptr [rdi + 8*rcx + 56]
	add	rcx, 8
	cmp	rsi, rcx
	jne	LBB0_9
## %bb.4:
	test	rax, rax
	je	LBB0_7
LBB0_5:
	lea	rcx, [rdi + 8*rcx]
	xor	esi, esi
	.p2align	4, 0x90
LBB0_6:                                 ## =>This Inner Loop Header: Depth=1
	vaddsd	xmm0, xmm0, qword ptr [rcx + 8*rsi]
	inc	rsi
	cmp	rax, rsi
	jne	LBB0_6
LBB0_7:
	vmovsd	qword ptr [rdx], xmm0
	pop	rbp
	ret
	.cfi_endproc
                                        ## -- End function

.subsections_via_symbols
