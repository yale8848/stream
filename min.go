// Create by Yale 2019/3/19 9:47
package stream

type min struct {
	compare MinCompare
	value T
}

func (sk *min)Begin(size int64)  {
}
func (sk *min)Accept(t T)  {
	if sk.value == nil {
		sk.value = t
	}
	if sk.compare(sk.value,t) {
		sk.value=t
	}

}
func (sk *min)End() {
}
func (sk *min)CancellationRequested() bool  {
	return  false
}
