	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 10, 14	sdk_version 10, 14
	.intel_syntax noprefix
	.globl	_sum                    ## -- Begin function sum
	.p2align	4, 0x90
_sum:                                   ## @sum
## %bb.0:
	push	rbp
	mov	rbp, rsp
	test	esi, esi
	jle	LBB0_1
## %bb.2:
	mov	r8d, esi
	cmp	esi, 3
	ja	LBB0_4
## %bb.3:
	xorpd	xmm0, xmm0
	xor	ecx, ecx
	jmp	LBB0_12
LBB0_1:
	xorpd	xmm0, xmm0
	jmp	LBB0_13
LBB0_4:
	mov	ecx, r8d
	and	ecx, -4
	lea	rsi, [rcx - 4]
	mov	rax, rsi
	shr	rax, 2
	inc	rax
	mov	r9d, eax
	and	r9d, 3
	cmp	rsi, 12
	jae	LBB0_6
## %bb.5:
	xorpd	xmm0, xmm0
	xor	eax, eax
	xorpd	xmm1, xmm1
	test	r9, r9
	jne	LBB0_9
	jmp	LBB0_11
LBB0_6:
	mov	esi, 1
	sub	rsi, rax
	lea	rsi, [r9 + rsi - 1]
	xorpd	xmm0, xmm0
	xor	eax, eax
	xorpd	xmm1, xmm1
	.p2align	4, 0x90
LBB0_7:                                 ## =>This Inner Loop Header: Depth=1
	movupd	xmm2, xmmword ptr [rdi + 8*rax]
	addpd	xmm2, xmm0
	movupd	xmm0, xmmword ptr [rdi + 8*rax + 16]
	addpd	xmm0, xmm1
	movupd	xmm1, xmmword ptr [rdi + 8*rax + 32]
	movupd	xmm3, xmmword ptr [rdi + 8*rax + 48]
	movupd	xmm4, xmmword ptr [rdi + 8*rax + 64]
	addpd	xmm4, xmm1
	addpd	xmm4, xmm2
	movupd	xmm2, xmmword ptr [rdi + 8*rax + 80]
	addpd	xmm2, xmm3
	addpd	xmm2, xmm0
	movupd	xmm0, xmmword ptr [rdi + 8*rax + 96]
	addpd	xmm0, xmm4
	movupd	xmm1, xmmword ptr [rdi + 8*rax + 112]
	addpd	xmm1, xmm2
	add	rax, 16
	add	rsi, 4
	jne	LBB0_7
## %bb.8:
	test	r9, r9
	je	LBB0_11
LBB0_9:
	lea	rax, [rdi + 8*rax + 16]
	neg	r9
	.p2align	4, 0x90
LBB0_10:                                ## =>This Inner Loop Header: Depth=1
	movupd	xmm2, xmmword ptr [rax - 16]
	addpd	xmm0, xmm2
	movupd	xmm2, xmmword ptr [rax]
	addpd	xmm1, xmm2
	add	rax, 32
	inc	r9
	jne	LBB0_10
LBB0_11:
	addpd	xmm0, xmm1
	haddpd	xmm0, xmm0
	cmp	rcx, r8
	je	LBB0_13
	.p2align	4, 0x90
LBB0_12:                                ## =>This Inner Loop Header: Depth=1
	addsd	xmm0, qword ptr [rdi + 8*rcx]
	inc	rcx
	cmp	r8, rcx
	jne	LBB0_12
LBB0_13:
	movsd	qword ptr [rdx], xmm0
	pop	rbp
	ret
                                        ## -- End function

.subsections_via_symbols
