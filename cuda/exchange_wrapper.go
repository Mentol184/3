package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var exchange_code cu.Function

type exchange_args struct {
	arg_Hx unsafe.Pointer
	arg_Hy unsafe.Pointer
	arg_Hz unsafe.Pointer
	arg_Mx unsafe.Pointer
	arg_My unsafe.Pointer
	arg_Mz unsafe.Pointer
	arg_wx float32
	arg_wy float32
	arg_wz float32
	arg_N0 int
	arg_N1 int
	arg_N2 int
	argptr [12]unsafe.Pointer
}

// Wrapper for exchange CUDA kernel, asynchronous.
func k_exchange_async(Hx unsafe.Pointer, Hy unsafe.Pointer, Hz unsafe.Pointer, Mx unsafe.Pointer, My unsafe.Pointer, Mz unsafe.Pointer, wx float32, wy float32, wz float32, N0 int, N1 int, N2 int, cfg *config, str cu.Stream) {
	if exchange_code == 0 {
		exchange_code = cu.ModuleLoadData(exchange_ptx).GetFunction("exchange")
	}

	var a exchange_args

	a.arg_Hx = Hx
	a.argptr[0] = unsafe.Pointer(&a.arg_Hx)
	a.arg_Hy = Hy
	a.argptr[1] = unsafe.Pointer(&a.arg_Hy)
	a.arg_Hz = Hz
	a.argptr[2] = unsafe.Pointer(&a.arg_Hz)
	a.arg_Mx = Mx
	a.argptr[3] = unsafe.Pointer(&a.arg_Mx)
	a.arg_My = My
	a.argptr[4] = unsafe.Pointer(&a.arg_My)
	a.arg_Mz = Mz
	a.argptr[5] = unsafe.Pointer(&a.arg_Mz)
	a.arg_wx = wx
	a.argptr[6] = unsafe.Pointer(&a.arg_wx)
	a.arg_wy = wy
	a.argptr[7] = unsafe.Pointer(&a.arg_wy)
	a.arg_wz = wz
	a.argptr[8] = unsafe.Pointer(&a.arg_wz)
	a.arg_N0 = N0
	a.argptr[9] = unsafe.Pointer(&a.arg_N0)
	a.arg_N1 = N1
	a.argptr[10] = unsafe.Pointer(&a.arg_N1)
	a.arg_N2 = N2
	a.argptr[11] = unsafe.Pointer(&a.arg_N2)

	args := a.argptr[:]
	cu.LaunchKernel(exchange_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, str, args)
}

// Wrapper for exchange CUDA kernel, synchronized.
func k_exchange(Hx unsafe.Pointer, Hy unsafe.Pointer, Hz unsafe.Pointer, Mx unsafe.Pointer, My unsafe.Pointer, Mz unsafe.Pointer, wx float32, wy float32, wz float32, N0 int, N1 int, N2 int, cfg *config) {
	str := stream()
	k_exchange_async(Hx, Hy, Hz, Mx, My, Mz, wx, wy, wz, N0, N1, N2, cfg, str)
	syncAndRecycle(str)
}

const exchange_ptx = `
.version 3.0
.target sm_30
.address_size 64


.entry exchange(
	.param .u64 exchange_param_0,
	.param .u64 exchange_param_1,
	.param .u64 exchange_param_2,
	.param .u64 exchange_param_3,
	.param .u64 exchange_param_4,
	.param .u64 exchange_param_5,
	.param .f32 exchange_param_6,
	.param .f32 exchange_param_7,
	.param .f32 exchange_param_8,
	.param .u32 exchange_param_9,
	.param .u32 exchange_param_10,
	.param .u32 exchange_param_11
)
{
	.reg .f32 	%f<155>;
	.reg .pred 	%p<16>;
	.reg .s32 	%r<124>;
	.reg .s64 	%rl<88>;


	ld.param.u32 	%r1, [exchange_param_9];
	ld.param.u32 	%r2, [exchange_param_10];
	ld.param.u32 	%r3, [exchange_param_11];
	.loc 2 12 1
	mov.u32 	%r4, %ctaid.x;
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %tid.x;
	mad.lo.s32 	%r7, %r5, %r4, %r6;
	.loc 2 13 1
	mov.u32 	%r8, %ntid.y;
	mov.u32 	%r9, %ctaid.y;
	mov.u32 	%r10, %tid.y;
	mad.lo.s32 	%r11, %r8, %r9, %r10;
	setp.lt.s32 	%p1, %r11, %r3;
	setp.lt.s32 	%p2, %r7, %r2;
	and.pred  	%p3, %p2, %p1;
	.loc 2 19 1
	setp.gt.s32 	%p4, %r1, 0;
	and.pred  	%p5, %p3, %p4;
	.loc 2 15 1
	@!%p5 bra 	BB0_26;

	.loc 2 27 1
	add.s32 	%r30, %r7, 1;
	mov.u32 	%r123, 0;
	.loc 3 238 5
	max.s32 	%r31, %r30, %r123;
	ld.param.u32 	%r107, [exchange_param_10];
	add.s32 	%r32, %r107, -1;
	.loc 3 210 5
	min.s32 	%r33, %r31, %r32;
	add.s32 	%r34, %r7, -1;
	.loc 3 238 5
	max.s32 	%r35, %r34, %r123;
	.loc 3 210 5
	min.s32 	%r36, %r35, %r32;
	.loc 2 31 1
	add.s32 	%r37, %r11, 1;
	.loc 3 238 5
	max.s32 	%r38, %r37, %r123;
	ld.param.u32 	%r111, [exchange_param_11];
	add.s32 	%r39, %r111, -1;
	.loc 3 210 5
	min.s32 	%r40, %r38, %r39;
	add.s32 	%r41, %r11, -1;
	.loc 3 238 5
	max.s32 	%r42, %r41, %r123;
	.loc 3 210 5
	min.s32 	%r43, %r42, %r39;
	mad.lo.s32 	%r122, %r36, %r111, %r11;
	mad.lo.s32 	%r121, %r33, %r111, %r11;
	mad.lo.s32 	%r120, %r111, %r7, %r43;
	mad.lo.s32 	%r119, %r111, %r7, %r40;
	mad.lo.s32 	%r118, %r111, %r7, %r11;

BB0_2:
	ld.param.u64 	%rl73, [exchange_param_3];
	cvta.to.global.u64 	%rl8, %rl73;
	.loc 2 23 1
	cvt.s64.s32 	%rl7, %r118;
	mul.wide.s32 	%rl9, %r118, 4;
	add.s64 	%rl10, %rl8, %rl9;
	ld.param.u64 	%rl80, [exchange_param_4];
	cvta.to.global.u64 	%rl11, %rl80;
	.loc 2 23 1
	add.s64 	%rl12, %rl11, %rl9;
	ld.param.u64 	%rl87, [exchange_param_5];
	cvta.to.global.u64 	%rl13, %rl87;
	.loc 2 23 1
	add.s64 	%rl14, %rl13, %rl9;
	ld.global.f32 	%f4, [%rl10];
	ld.global.f32 	%f5, [%rl12];
	mul.f32 	%f71, %f5, %f5;
	fma.rn.f32 	%f72, %f4, %f4, %f71;
	ld.global.f32 	%f6, [%rl14];
	fma.rn.f32 	%f73, %f6, %f6, %f72;
	.loc 3 991 5
	sqrt.rn.f32 	%f7, %f73;
	.loc 2 23 1
	setp.neu.f32 	%p6, %f7, 0f00000000;
	@%p6 bra 	BB0_4;

	mov.f32 	%f145, 0f00000000;
	bra.uni 	BB0_5;

BB0_4:
	.loc 2 23 1
	rcp.rn.f32 	%f145, %f7;

BB0_5:
	mul.f32 	%f10, %f145, %f4;
	mul.f32 	%f11, %f145, %f5;
	.loc 2 24 1
	mul.f32 	%f75, %f11, %f11;
	fma.rn.f32 	%f76, %f10, %f10, %f75;
	.loc 2 23 1
	mul.f32 	%f12, %f145, %f6;
	.loc 2 24 1
	fma.rn.f32 	%f77, %f12, %f12, %f76;
	.loc 3 991 5
	sqrt.rn.f32 	%f78, %f77;
	.loc 2 25 1
	setp.eq.f32 	%p7, %f78, 0f00000000;
	selp.f32 	%f13, 0f3F800000, %f78, %p7;
	ld.param.u64 	%rl72, [exchange_param_3];
	cvta.to.global.u64 	%rl15, %rl72;
	.loc 2 27 1
	mul.wide.s32 	%rl16, %r121, 4;
	add.s64 	%rl17, %rl15, %rl16;
	ld.param.u64 	%rl79, [exchange_param_4];
	cvta.to.global.u64 	%rl18, %rl79;
	.loc 2 27 1
	add.s64 	%rl19, %rl18, %rl16;
	ld.param.u64 	%rl86, [exchange_param_5];
	cvta.to.global.u64 	%rl20, %rl86;
	.loc 2 27 1
	add.s64 	%rl21, %rl20, %rl16;
	ld.global.f32 	%f14, [%rl17];
	ld.global.f32 	%f15, [%rl19];
	mul.f32 	%f79, %f15, %f15;
	fma.rn.f32 	%f80, %f14, %f14, %f79;
	ld.global.f32 	%f16, [%rl21];
	fma.rn.f32 	%f81, %f16, %f16, %f80;
	.loc 3 991 5
	sqrt.rn.f32 	%f17, %f81;
	.loc 2 27 1
	setp.neu.f32 	%p8, %f17, 0f00000000;
	@%p8 bra 	BB0_7;

	mov.f32 	%f146, 0f00000000;
	bra.uni 	BB0_8;

BB0_7:
	.loc 2 27 1
	rcp.rn.f32 	%f146, %f17;

BB0_8:
	mul.f32 	%f20, %f146, %f14;
	mul.f32 	%f21, %f146, %f15;
	mul.f32 	%f22, %f146, %f16;
	ld.param.u64 	%rl71, [exchange_param_3];
	cvta.to.global.u64 	%rl22, %rl71;
	.loc 2 28 1
	mul.wide.s32 	%rl23, %r122, 4;
	add.s64 	%rl24, %rl22, %rl23;
	ld.param.u64 	%rl78, [exchange_param_4];
	cvta.to.global.u64 	%rl25, %rl78;
	.loc 2 28 1
	add.s64 	%rl26, %rl25, %rl23;
	ld.param.u64 	%rl85, [exchange_param_5];
	cvta.to.global.u64 	%rl27, %rl85;
	.loc 2 28 1
	add.s64 	%rl28, %rl27, %rl23;
	ld.global.f32 	%f23, [%rl24];
	ld.global.f32 	%f24, [%rl26];
	mul.f32 	%f83, %f24, %f24;
	fma.rn.f32 	%f84, %f23, %f23, %f83;
	ld.global.f32 	%f25, [%rl28];
	fma.rn.f32 	%f85, %f25, %f25, %f84;
	.loc 3 991 5
	sqrt.rn.f32 	%f26, %f85;
	.loc 2 28 1
	setp.neu.f32 	%p9, %f26, 0f00000000;
	@%p9 bra 	BB0_10;

	mov.f32 	%f147, 0f00000000;
	bra.uni 	BB0_11;

BB0_10:
	.loc 2 28 1
	rcp.rn.f32 	%f147, %f26;

BB0_11:
	.loc 2 29 1
	neg.f32 	%f87, %f10;
	fma.rn.f32 	%f88, %f147, %f23, %f87;
	neg.f32 	%f89, %f11;
	fma.rn.f32 	%f90, %f147, %f24, %f89;
	neg.f32 	%f91, %f12;
	fma.rn.f32 	%f92, %f147, %f25, %f91;
	sub.f32 	%f93, %f20, %f10;
	add.f32 	%f94, %f93, %f88;
	sub.f32 	%f95, %f21, %f11;
	add.f32 	%f96, %f95, %f90;
	sub.f32 	%f97, %f22, %f12;
	add.f32 	%f98, %f97, %f92;
	ld.param.f32 	%f143, [exchange_param_7];
	.loc 4 1311 3
	div.rn.f32 	%f99, %f143, %f13;
	.loc 2 29 1
	mul.f32 	%f29, %f99, %f94;
	mul.f32 	%f30, %f99, %f96;
	mul.f32 	%f31, %f99, %f98;
	ld.param.u64 	%rl70, [exchange_param_3];
	cvta.to.global.u64 	%rl29, %rl70;
	.loc 2 31 1
	mul.wide.s32 	%rl30, %r119, 4;
	add.s64 	%rl31, %rl29, %rl30;
	ld.param.u64 	%rl77, [exchange_param_4];
	cvta.to.global.u64 	%rl32, %rl77;
	.loc 2 31 1
	add.s64 	%rl33, %rl32, %rl30;
	ld.param.u64 	%rl84, [exchange_param_5];
	cvta.to.global.u64 	%rl34, %rl84;
	.loc 2 31 1
	add.s64 	%rl35, %rl34, %rl30;
	ld.global.f32 	%f32, [%rl31];
	ld.global.f32 	%f33, [%rl33];
	mul.f32 	%f100, %f33, %f33;
	fma.rn.f32 	%f101, %f32, %f32, %f100;
	ld.global.f32 	%f34, [%rl35];
	fma.rn.f32 	%f102, %f34, %f34, %f101;
	.loc 3 991 5
	sqrt.rn.f32 	%f35, %f102;
	.loc 2 31 1
	setp.neu.f32 	%p10, %f35, 0f00000000;
	@%p10 bra 	BB0_13;

	mov.f32 	%f148, 0f00000000;
	bra.uni 	BB0_14;

BB0_13:
	.loc 2 31 1
	rcp.rn.f32 	%f148, %f35;

BB0_14:
	mul.f32 	%f38, %f148, %f32;
	mul.f32 	%f39, %f148, %f33;
	mul.f32 	%f40, %f148, %f34;
	ld.param.u64 	%rl69, [exchange_param_3];
	cvta.to.global.u64 	%rl36, %rl69;
	.loc 2 32 1
	mul.wide.s32 	%rl37, %r120, 4;
	add.s64 	%rl38, %rl36, %rl37;
	ld.param.u64 	%rl76, [exchange_param_4];
	cvta.to.global.u64 	%rl39, %rl76;
	.loc 2 32 1
	add.s64 	%rl40, %rl39, %rl37;
	ld.param.u64 	%rl83, [exchange_param_5];
	cvta.to.global.u64 	%rl41, %rl83;
	.loc 2 32 1
	add.s64 	%rl42, %rl41, %rl37;
	ld.global.f32 	%f41, [%rl38];
	ld.global.f32 	%f42, [%rl40];
	mul.f32 	%f104, %f42, %f42;
	fma.rn.f32 	%f105, %f41, %f41, %f104;
	ld.global.f32 	%f43, [%rl42];
	fma.rn.f32 	%f106, %f43, %f43, %f105;
	.loc 3 991 5
	sqrt.rn.f32 	%f44, %f106;
	.loc 2 32 1
	setp.neu.f32 	%p11, %f44, 0f00000000;
	@%p11 bra 	BB0_16;

	mov.f32 	%f149, 0f00000000;
	bra.uni 	BB0_17;

BB0_16:
	.loc 2 32 1
	rcp.rn.f32 	%f149, %f44;

BB0_17:
	.loc 2 33 1
	fma.rn.f32 	%f109, %f149, %f41, %f87;
	fma.rn.f32 	%f111, %f149, %f42, %f89;
	fma.rn.f32 	%f113, %f149, %f43, %f91;
	sub.f32 	%f114, %f38, %f10;
	add.f32 	%f115, %f114, %f109;
	sub.f32 	%f116, %f39, %f11;
	add.f32 	%f117, %f116, %f111;
	sub.f32 	%f118, %f40, %f12;
	add.f32 	%f119, %f118, %f113;
	ld.param.f32 	%f144, [exchange_param_8];
	.loc 4 1311 3
	div.rn.f32 	%f120, %f144, %f13;
	.loc 2 33 1
	fma.rn.f32 	%f152, %f120, %f115, %f29;
	fma.rn.f32 	%f153, %f120, %f117, %f30;
	fma.rn.f32 	%f154, %f120, %f119, %f31;
	ld.param.u32 	%r103, [exchange_param_9];
	.loc 2 36 1
	setp.eq.s32 	%p12, %r103, 1;
	@%p12 bra 	BB0_25;

	.loc 2 37 1
	add.s32 	%r61, %r123, 1;
	mov.u32 	%r62, 0;
	.loc 3 238 5
	max.s32 	%r63, %r61, %r62;
	ld.param.u32 	%r102, [exchange_param_9];
	add.s32 	%r64, %r102, -1;
	.loc 3 210 5
	min.s32 	%r65, %r63, %r64;
	ld.param.u32 	%r106, [exchange_param_10];
	mad.lo.s32 	%r70, %r65, %r106, %r7;
	ld.param.u32 	%r110, [exchange_param_11];
	.loc 2 37 1
	mad.lo.s32 	%r75, %r70, %r110, %r11;
	ld.param.u64 	%rl68, [exchange_param_3];
	cvta.to.global.u64 	%rl43, %rl68;
	.loc 2 37 1
	mul.wide.s32 	%rl44, %r75, 4;
	add.s64 	%rl45, %rl43, %rl44;
	ld.param.u64 	%rl75, [exchange_param_4];
	cvta.to.global.u64 	%rl46, %rl75;
	.loc 2 37 1
	add.s64 	%rl47, %rl46, %rl44;
	ld.param.u64 	%rl82, [exchange_param_5];
	cvta.to.global.u64 	%rl48, %rl82;
	.loc 2 37 1
	add.s64 	%rl49, %rl48, %rl44;
	ld.global.f32 	%f50, [%rl45];
	ld.global.f32 	%f51, [%rl47];
	mul.f32 	%f121, %f51, %f51;
	fma.rn.f32 	%f122, %f50, %f50, %f121;
	ld.global.f32 	%f52, [%rl49];
	fma.rn.f32 	%f123, %f52, %f52, %f122;
	.loc 3 991 5
	sqrt.rn.f32 	%f53, %f123;
	.loc 2 37 1
	setp.neu.f32 	%p13, %f53, 0f00000000;
	@%p13 bra 	BB0_20;

	mov.f32 	%f150, 0f00000000;
	bra.uni 	BB0_21;

BB0_20:
	.loc 2 37 1
	rcp.rn.f32 	%f150, %f53;

BB0_21:
	mul.f32 	%f56, %f150, %f50;
	mul.f32 	%f57, %f150, %f51;
	mul.f32 	%f58, %f150, %f52;
	add.s32 	%r79, %r123, -1;
	.loc 3 238 5
	max.s32 	%r81, %r79, %r62;
	ld.param.u32 	%r101, [exchange_param_9];
	add.s32 	%r82, %r101, -1;
	.loc 3 210 5
	min.s32 	%r83, %r81, %r82;
	ld.param.u32 	%r105, [exchange_param_10];
	mad.lo.s32 	%r88, %r83, %r105, %r7;
	ld.param.u32 	%r109, [exchange_param_11];
	.loc 2 38 1
	mad.lo.s32 	%r93, %r88, %r109, %r11;
	ld.param.u64 	%rl67, [exchange_param_3];
	cvta.to.global.u64 	%rl50, %rl67;
	.loc 2 38 1
	mul.wide.s32 	%rl51, %r93, 4;
	add.s64 	%rl52, %rl50, %rl51;
	ld.param.u64 	%rl74, [exchange_param_4];
	cvta.to.global.u64 	%rl53, %rl74;
	.loc 2 38 1
	add.s64 	%rl54, %rl53, %rl51;
	ld.param.u64 	%rl81, [exchange_param_5];
	cvta.to.global.u64 	%rl55, %rl81;
	.loc 2 38 1
	add.s64 	%rl56, %rl55, %rl51;
	ld.global.f32 	%f59, [%rl52];
	ld.global.f32 	%f60, [%rl54];
	mul.f32 	%f125, %f60, %f60;
	fma.rn.f32 	%f126, %f59, %f59, %f125;
	ld.global.f32 	%f61, [%rl56];
	fma.rn.f32 	%f127, %f61, %f61, %f126;
	.loc 3 991 5
	sqrt.rn.f32 	%f62, %f127;
	.loc 2 38 1
	setp.neu.f32 	%p14, %f62, 0f00000000;
	@%p14 bra 	BB0_23;

	mov.f32 	%f151, 0f00000000;
	bra.uni 	BB0_24;

BB0_23:
	.loc 2 38 1
	rcp.rn.f32 	%f151, %f62;

BB0_24:
	.loc 2 39 1
	fma.rn.f32 	%f130, %f151, %f59, %f87;
	fma.rn.f32 	%f132, %f151, %f60, %f89;
	fma.rn.f32 	%f134, %f151, %f61, %f91;
	sub.f32 	%f135, %f56, %f10;
	add.f32 	%f136, %f135, %f130;
	sub.f32 	%f137, %f57, %f11;
	add.f32 	%f138, %f137, %f132;
	sub.f32 	%f139, %f58, %f12;
	add.f32 	%f140, %f139, %f134;
	ld.param.f32 	%f142, [exchange_param_6];
	.loc 4 1311 3
	div.rn.f32 	%f141, %f142, %f13;
	.loc 2 39 1
	fma.rn.f32 	%f152, %f141, %f136, %f152;
	fma.rn.f32 	%f153, %f141, %f138, %f153;
	fma.rn.f32 	%f154, %f141, %f140, %f154;

BB0_25:
	ld.param.u64 	%rl64, [exchange_param_0];
	cvta.to.global.u64 	%rl57, %rl64;
	.loc 2 42 1
	shl.b64 	%rl58, %rl7, 2;
	add.s64 	%rl59, %rl57, %rl58;
	st.global.f32 	[%rl59], %f152;
	ld.param.u64 	%rl65, [exchange_param_1];
	cvta.to.global.u64 	%rl60, %rl65;
	.loc 2 43 1
	add.s64 	%rl61, %rl60, %rl58;
	st.global.f32 	[%rl61], %f153;
	ld.param.u64 	%rl66, [exchange_param_2];
	cvta.to.global.u64 	%rl62, %rl66;
	.loc 2 44 1
	add.s64 	%rl63, %rl62, %rl58;
	st.global.f32 	[%rl63], %f154;
	ld.param.u32 	%r104, [exchange_param_10];
	ld.param.u32 	%r108, [exchange_param_11];
	mad.lo.s32 	%r122, %r108, %r104, %r122;
	mad.lo.s32 	%r121, %r108, %r104, %r121;
	mad.lo.s32 	%r120, %r108, %r104, %r120;
	mad.lo.s32 	%r119, %r108, %r104, %r119;
	mad.lo.s32 	%r118, %r108, %r104, %r118;
	.loc 2 19 18
	add.s32 	%r123, %r123, 1;
	ld.param.u32 	%r100, [exchange_param_9];
	.loc 2 19 1
	setp.lt.s32 	%p15, %r123, %r100;
	@%p15 bra 	BB0_2;

BB0_26:
	.loc 2 46 2
	ret;
}


`
