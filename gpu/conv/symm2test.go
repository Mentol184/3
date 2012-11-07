package conv

import (
	"nimble-cube/core"
	"nimble-cube/gpu"
	"nimble-cube/mag"
)

// Internal: main function for conv test.
func TestSymm2(N0, N1, N2 int) {
	C := 1e-9
	mesh := core.NewMesh(N0, N1, N2, C, 2*C, 3*C)
	core.Log(mesh)
	N := mesh.NCell()

	gpu.LockCudaThread()
	hin := core.MakeChan3("hin", "", mesh, core.UnifiedMemory)
	hout := core.MakeChan3("hout", "", mesh, core.UnifiedMemory)

	acc := 1
	kern := mag.BruteKernel(core.ZeroPad(mesh), acc)

	arr := hin.UnsafeArray()
	initConvTestInput(arr)

	F := 10
	go func() {
		for i := 0; i < F; i++ {
			hin.WriteNext(N)
			hin.WriteDone()
		}
	}()

	//go NewSymmetricHtoD(mesh, kern, hin.NewReader(), hout).Run()
	go NewSymm2D(mesh.Size(), kern, hin, hout).Run()

	houtR := hout.NewReader()
	for i := 0; i < F; i++ {
		houtR.ReadNext(N)
		houtR.ReadDone()
	}

	outarr := hout.UnsafeData()

	ref := core.MakeVectors(mesh.Size())
	Brute(arr, ref, kern)
	checkErr(outarr, core.Contiguous3(ref))
}
