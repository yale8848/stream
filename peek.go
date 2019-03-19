// Create by Yale 2019/3/18 17:27
package stream


type peek struct {
	me *sinkNode
	con Consumer
}

func (sk *peek)Begin(size int64)  {
	sk.me.next.value.Begin(size)
}
func (sk *peek)Accept(t T)  {
	if !sk.me.next.value.CancellationRequested() {
		sk.con(t)
		sk.me.next.value.Accept(t)
	}
}
func (sk *peek)End() {
	if sk.me.next!=nil{
		sk.me.next.value.End()
	}
}
func (sk *peek)CancellationRequested() bool  {
	return  false
}
