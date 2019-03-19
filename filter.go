// Create by Yale 2019/3/18 14:58
package stream


type filter struct {
	me *sinkNode
	pre Predicate
}

func (sk *filter)Begin(size int64)  {
	sk.me.next.value.Begin(size)
}
func (sk *filter)Accept(t T)  {
	next:=sk.me.next.value
	if sk.pre(t) {
		if !next.CancellationRequested() {
			next.Accept(t)
		}
	}
}
func (sk *filter)End() {
	if sk.me.next!=nil{
		sk.me.next.value.End()
	}
}
func (sk *filter)CancellationRequested() bool  {
	return  false
}
