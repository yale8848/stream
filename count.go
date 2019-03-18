// Create by Yale 2019/3/18 15:26
package stream

type count struct {
	num int64
}

func (ft *count)Begin(size int64)  {
}
func (ft *count)Accept(t T)  {
	ft.num++
}
func (ft *count)End() {
}
func (ft *count)CancellationRequested() bool  {
	return  false
}

