// Create by Yale 2019/3/18 14:58
package stream


type filter struct {
	me *sinkNode
	pre Predicate
}

func (ft *filter)Begin(size int64)  {
	ft.me.next.value.Begin(size)
}
func (ft *filter)Accept(t T)  {
	if ft.pre(t) {
		ft.me.next.value.Accept(t)
	}
}
func (ft *filter)End() {
	if ft.me.next!=nil{
		ft.me.next.value.End()
	}

}
func (ft *filter)CancellationRequested() bool  {
	return  false
}
