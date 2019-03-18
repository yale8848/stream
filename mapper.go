// Create by Yale 2019/3/18 15:10
package stream


type mapper struct {
	me *sinkNode
	fun Function
}

func (ft *mapper)Begin(size int64)  {
	ft.me.next.value.Begin(size)
}
func (ft *mapper)Accept(t T)  {
	ft.me.next.value.Accept(ft.fun(t))

}
func (ft *mapper)End() {
	if ft.me.next!=nil{
		ft.me.next.value.End()
	}
}
func (ft *mapper)CancellationRequested() bool  {
	return  false
}

