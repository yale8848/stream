// Create by Yale 2019/3/18 20:03
package stream

type distinct struct {
	me *sinkNode
	les Less
	data []T
	mp map[T]T
	fun Function
}

func (sk *distinct)Begin(size int64)  {
	sk.mp = make(map[T]T,0)
}
func (sk *distinct)Accept(t T)  {
	sk.mp[sk.fun(t)] = t
}
func (sk *distinct)End() {
    next:=sk.me.next.value
	next.Begin(int64(len(sk.mp)))
	for _,v:=range sk.mp{
		if next.CancellationRequested() {
			break
		}
		next.Accept(v)
	}
	next.End()

}
func (sk *distinct)CancellationRequested() bool  {
	return  false
}
