package ptx

//This file is auto-generated. Editing is futile.

func init() { Code["reducemaxvecdiff2"] = REDUCEMAXVECDIFF2 }

const REDUCEMAXVECDIFF2 = `
//
// Generated by NVIDIA NVVM Compiler
// Compiler built on Sat Sep 22 02:35:14 2012 (1348274114)
// Cuda compilation tools, release 5.0, V0.2.1221
//

.version 3.1
.target sm_30
.address_size 64

	.file	1 "/tmp/tmpxft_00005656_00000000-9_reducemaxvecdiff2.cpp3.i"
	.file	2 "/home/arne/src/code.google.com/p/nimble-cube/gpu/ptx/reducemaxvecdiff2.cu"
	.file	3 "/usr/local/cuda-5.0/nvvm/ci_include.h"
	.file	4 "/usr/local/cuda/bin/../include/sm_11_atomic_functions.h"
// __cuda_local_var_35274_32_non_const_sdata has been demoted

.visible .entry reducemaxvecdiff2(
	.param .u64 reducemaxvecdiff2_param_0,
	.param .u64 reducemaxvecdiff2_param_1,
	.param .u64 reducemaxvecdiff2_param_2,
	.param .u64 reducemaxvecdiff2_param_3,
	.param .u64 reducemaxvecdiff2_param_4,
	.param .u64 reducemaxvecdiff2_param_5,
	.param .u64 reducemaxvecdiff2_param_6,
	.param .f32 reducemaxvecdiff2_param_7,
	.param .u32 reducemaxvecdiff2_param_8
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<45>;
	.reg .f32 	%f<39>;
	.reg .s64 	%rd<28>;
	// demoted variable
	.shared .align 4 .b8 __cuda_local_var_35274_32_non_const_sdata[2048];

	ld.param.u64 	%rd9, [reducemaxvecdiff2_param_0];
	ld.param.u64 	%rd10, [reducemaxvecdiff2_param_1];
	ld.param.u64 	%rd11, [reducemaxvecdiff2_param_2];
	ld.param.u64 	%rd12, [reducemaxvecdiff2_param_3];
	ld.param.u64 	%rd13, [reducemaxvecdiff2_param_4];
	ld.param.u64 	%rd14, [reducemaxvecdiff2_param_5];
	ld.param.u64 	%rd15, [reducemaxvecdiff2_param_6];
	ld.param.f32 	%f38, [reducemaxvecdiff2_param_7];
	ld.param.u32 	%r9, [reducemaxvecdiff2_param_8];
	cvta.to.global.u64 	%rd1, %rd15;
	cvta.to.global.u64 	%rd2, %rd14;
	cvta.to.global.u64 	%rd3, %rd11;
	cvta.to.global.u64 	%rd4, %rd13;
	cvta.to.global.u64 	%rd5, %rd10;
	cvta.to.global.u64 	%rd6, %rd12;
	cvta.to.global.u64 	%rd7, %rd9;
	.loc 2 14 1
	mov.u32 	%r44, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r43, %r44, %r10, %r2;
	mov.u32 	%r11, %nctaid.x;
	mul.lo.s32 	%r4, %r44, %r11;
	.loc 2 14 1
	setp.ge.s32 	%p1, %r43, %r9;
	@%p1 bra 	BB0_2;

BB0_1:
	.loc 2 14 1
	mul.wide.s32 	%rd16, %r43, 4;
	add.s64 	%rd17, %rd7, %rd16;
	add.s64 	%rd18, %rd6, %rd16;
	ld.global.f32 	%f5, [%rd18];
	ld.global.f32 	%f6, [%rd17];
	sub.f32 	%f7, %f6, %f5;
	add.s64 	%rd19, %rd5, %rd16;
	add.s64 	%rd20, %rd4, %rd16;
	ld.global.f32 	%f8, [%rd20];
	ld.global.f32 	%f9, [%rd19];
	sub.f32 	%f10, %f9, %f8;
	mul.f32 	%f11, %f10, %f10;
	fma.rn.f32 	%f12, %f7, %f7, %f11;
	add.s64 	%rd21, %rd3, %rd16;
	add.s64 	%rd22, %rd2, %rd16;
	ld.global.f32 	%f13, [%rd22];
	ld.global.f32 	%f14, [%rd21];
	sub.f32 	%f15, %f14, %f13;
	fma.rn.f32 	%f16, %f15, %f15, %f12;
	.loc 3 435 5
	max.f32 	%f38, %f38, %f16;
	.loc 2 14 1
	add.s32 	%r43, %r43, %r4;
	.loc 2 14 1
	setp.lt.s32 	%p2, %r43, %r9;
	@%p2 bra 	BB0_1;

BB0_2:
	.loc 2 14 1
	mul.wide.s32 	%rd23, %r2, 4;
	mov.u64 	%rd24, __cuda_local_var_35274_32_non_const_sdata;
	add.s64 	%rd8, %rd24, %rd23;
	st.shared.f32 	[%rd8], %f38;
	bar.sync 	0;
	.loc 2 14 1
	setp.lt.u32 	%p3, %r44, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	.loc 2 14 1
	mov.u32 	%r7, %r44;
	shr.u32 	%r44, %r7, 1;
	.loc 2 14 1
	setp.ge.u32 	%p4, %r2, %r44;
	@%p4 bra 	BB0_5;

	.loc 2 14 1
	ld.shared.f32 	%f17, [%rd8];
	add.s32 	%r20, %r44, %r2;
	mul.wide.u32 	%rd25, %r20, 4;
	add.s64 	%rd27, %rd24, %rd25;
	ld.shared.f32 	%f18, [%rd27];
	.loc 3 435 5
	max.f32 	%f19, %f17, %f18;
	.loc 2 14 1
	st.shared.f32 	[%rd8], %f19;

BB0_5:
	.loc 2 14 1
	bar.sync 	0;
	.loc 2 14 1
	setp.gt.u32 	%p5, %r7, 131;
	@%p5 bra 	BB0_3;

BB0_6:
	.loc 2 14 1
	setp.gt.s32 	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	.loc 2 14 1
	ld.volatile.shared.f32 	%f20, [%rd8];
	ld.volatile.shared.f32 	%f21, [%rd8+128];
	.loc 3 435 5
	max.f32 	%f22, %f20, %f21;
	.loc 2 14 1
	st.volatile.shared.f32 	[%rd8], %f22;
	ld.volatile.shared.f32 	%f23, [%rd8+64];
	ld.volatile.shared.f32 	%f24, [%rd8];
	.loc 3 435 5
	max.f32 	%f25, %f24, %f23;
	.loc 2 14 1
	st.volatile.shared.f32 	[%rd8], %f25;
	ld.volatile.shared.f32 	%f26, [%rd8+32];
	ld.volatile.shared.f32 	%f27, [%rd8];
	.loc 3 435 5
	max.f32 	%f28, %f27, %f26;
	.loc 2 14 1
	st.volatile.shared.f32 	[%rd8], %f28;
	ld.volatile.shared.f32 	%f29, [%rd8+16];
	ld.volatile.shared.f32 	%f30, [%rd8];
	.loc 3 435 5
	max.f32 	%f31, %f30, %f29;
	.loc 2 14 1
	st.volatile.shared.f32 	[%rd8], %f31;
	ld.volatile.shared.f32 	%f32, [%rd8+8];
	ld.volatile.shared.f32 	%f33, [%rd8];
	.loc 3 435 5
	max.f32 	%f34, %f33, %f32;
	.loc 2 14 1
	st.volatile.shared.f32 	[%rd8], %f34;
	ld.volatile.shared.f32 	%f35, [%rd8+4];
	ld.volatile.shared.f32 	%f36, [%rd8];
	.loc 3 435 5
	max.f32 	%f37, %f36, %f35;
	.loc 2 14 1
	st.volatile.shared.f32 	[%rd8], %f37;

BB0_8:
	.loc 2 14 1
	setp.ne.s32 	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	.loc 2 14 1
	ld.shared.u32 	%r41, [__cuda_local_var_35274_32_non_const_sdata];
	.loc 3 1881 5
	atom.global.max.s32 	%r42, [%rd1], %r41;

BB0_10:
	.loc 2 15 2
	ret;
}


`
