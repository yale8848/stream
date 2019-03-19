// Create by Yale 2019/3/18 15:26
package stream

type count struct {
	num int64
}

func (sk *count)Begin(size int64)  {
}
func (sk *count)Accept(t T)  {
	sk.num++
}
func (sk *count)End() {
}
func (sk *count)CancellationRequested() bool  {
	return  false
}

