// Create by Yale 2019/3/18 15:26
package stream

type sum struct {
	num int64
	s Sum
}

func (sk *sum)Begin(size int64)  {
}
func (sk *sum)Accept(t T)  {
	sk.num = sk.num+sk.s(t)
}
func (sk *sum)End() {
}
func (sk *sum)CancellationRequested() bool  {
	return  false
}

