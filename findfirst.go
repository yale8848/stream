// Create by Yale 2019/3/18 17:51
package stream


type findFirst struct {
	value T
	count int64
	cancelRequest bool
}

func (sk *findFirst)Begin(size int64)  {
}
func (sk *findFirst)Accept(t T)  {
	sk.value = t
	sk.count++
	if sk.count == 1{
		sk.value = t
		sk.cancelRequest = true
	}
}
func (sk *findFirst)End() {

}
func (sk *findFirst)CancellationRequested() bool  {
	return  sk.cancelRequest
}
