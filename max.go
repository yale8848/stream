// Create by Yale 2019/3/19 9:46
package stream

type max struct {
	compare MaxCompare
	value T
}

func (sk *max)Begin(size int64)  {
}
func (sk *max)Accept(t T)  {
	if sk.value == nil {
		sk.value = t
	}
	if sk.compare(sk.value,t) {
		sk.value = t
	}
}
func (sk *max)End() {
}
func (sk *max)CancellationRequested() bool  {
	return  false
}

