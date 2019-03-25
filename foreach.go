// Create by Yale 2019/3/18 15:13
package stream


type foreach struct {
	cons ConsumerCancel
	cancelRequest bool
}

func (sk *foreach)Begin(size int64)  {

}
func (sk *foreach)Accept(t T)  {
	if !sk.cons(t) {
		sk.cancelRequest = true
	}
}
func (sk *foreach)End() {

}
func (sk *foreach)CancellationRequested() bool  {
	return  sk.cancelRequest
}


