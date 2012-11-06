package core

type Chan1 struct {
	chandata // array+list
	mutex    *RWMutex
}

func MakeChan(tag, unit string, m *Mesh, blocks ...int) Chan1 {
	data := makedata(tag, unit, m, blocks...)
	AddQuant(tag)
	return Chan1{data, NewRWMutex(data.BlockLen(), tag)}
}

// WriteDone() signals a slice obtained by WriteNext() is fully
// written and can be sent down the Chan.
func (c *Chan1) WriteDone() {
	c.mutex.WriteDone()
}

func (c *Chan1) WriteNext(n int) Slice {
	c.mutex.WriteNext(n)
	a, b := c.mutex.WRange()
	return c.slice.Slice(a, b)
}

//func (c *Chan1) WriteDelta(Δstart, Δstop int) []float32 {
//	c.mutex.WriteDelta(Δstart, Δstop)
//	a, b := c.mutex.WRange()
//	return c.slice.list[a:b]
//}

func(c Chan1)UnsafeArray()[][][]float32{
	return Reshape(c.chandata.slice.Host(), c.Size())
}
