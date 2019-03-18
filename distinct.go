// Create by Yale 2019/3/18 20:03
package stream

type distinct struct {
	me *sinkNode
	les Less
	data []T
	mp map[T]T
	fun Function
}

func (ft *distinct)Begin(size int64)  {
	ft.mp = make(map[T]T,0)
}
func (ft *distinct)Accept(t T)  {
	ft.mp[ft.fun(t)] = t
}
func (ft *distinct)End() {


	ft.me.next.value.Begin(int64(len(ft.mp)))
	for _,v:=range ft.mp{
		ft.me.next.value.Accept(v)
	}
	ft.me.next.value.End()

}
func (ft *distinct)CancellationRequested() bool  {
	return  false
}
