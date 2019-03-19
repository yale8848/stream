// Create by Yale 2019/3/18 15:13
package stream


type foreach struct {
	cons Consumer
}

func (sk *foreach)Begin(size int64)  {

}
func (sk *foreach)Accept(t T)  {
	sk.cons(t)
}
func (sk *foreach)End() {

}
func (sk *foreach)CancellationRequested() bool  {
	return  false
}


