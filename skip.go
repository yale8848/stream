// Create by Yale 2019/3/18 20:10
package stream

type skip struct {
	me *sinkNode
	num uint64
	cancelRequest bool
	value []T
	count int64
}

func (sk *skip)Begin(size int64)  {

	if size >=0 && size <= int64(sk.num) {
		sk.cancelRequest = true
	}else{
		sk.value =  make([]T,0)
	}

}
func (sk *skip)Accept(t T)  {
	sk.count++
	if sk.count>int64(sk.num) {
		sk.value = append(sk.value,t)
	}
}
func (sk *skip)End() {

	sk.me.next.value.Begin(int64(len(sk.value)))

	for _,v:=range sk.value{
		if sk.me.next.value.CancellationRequested() {
			break
		}
		sk.me.next.value.Accept(v)
	}
	sk.me.next.value.End()
}
func (sk *skip)CancellationRequested() bool  {
	return  sk.cancelRequest
}
