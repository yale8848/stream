// Create by Yale 2019/3/18 15:10
package stream


type mapper struct {
	me *sinkNode
	fun Function
}

func (sk *mapper)Begin(size int64)  {
	sk.me.next.value.Begin(size)
}
func (sk *mapper)Accept(t T)  {

	if !sk.me.next.value.CancellationRequested() {
		sk.me.next.value.Accept(sk.fun(t))
	}
}
func (sk *mapper)End() {
	if sk.me.next!=nil{
		sk.me.next.value.End()
	}
}
func (sk *mapper)CancellationRequested() bool  {
	return  false
}

