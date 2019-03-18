// Create by Yale 2019/3/18 17:51
package stream


type findFirst struct {
	value T
	count int64
	cancelRequest bool
}

func (ft *findFirst)Begin(size int64)  {
}
func (ft *findFirst)Accept(t T)  {
	ft.value = t
	ft.count++
	if ft.count == 1{
		ft.value = t
		ft.cancelRequest = true
	}
}
func (ft *findFirst)End() {

}
func (ft *findFirst)CancellationRequested() bool  {
	return  ft.cancelRequest
}
