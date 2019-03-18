// Create by Yale 2019/3/18 17:27
package stream


type peek struct {
	me *sinkNode
	con Consumer
}

func (ft *peek)Begin(size int64)  {
	ft.me.next.value.Begin(size)
}
func (ft *peek)Accept(t T)  {
	ft.con(t)
	ft.me.next.value.Accept(t)
}
func (ft *peek)End() {
	if ft.me.next!=nil{
		ft.me.next.value.End()
	}

}
func (ft *peek)CancellationRequested() bool  {
	return  false
}
