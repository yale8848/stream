// Create by Yale 2019/3/18 15:13
package stream


type foreach struct {
	cons Consumer
}

func (ft *foreach)Begin(size int64)  {

}
func (ft *foreach)Accept(t T)  {
	ft.cons(t)
}
func (ft *foreach)End() {

}
func (ft *foreach)CancellationRequested() bool  {
	return  false
}


