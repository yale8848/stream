// Create by Yale 2019/3/19 9:28
package stream

type limit struct {
	me *sinkNode
	num uint64
	cancelRequest bool
	value []T
	count int64
}

func (sk *limit)Begin(size int64)  {

	if sk.num == 0  {
		sk.cancelRequest = true
	}else{
		sk.value =  make([]T,0)
	}
}
func (sk *limit)Accept(t T)  {
	sk.count++
	if sk.count<=int64(sk.num) {
		sk.value = append(sk.value,t)
	}else{
		sk.cancelRequest = true
	}
}
func (sk *limit)End() {

	next:=sk.me.next.value
	next.Begin(int64(len(sk.value)))

	for _,v:=range sk.value{
		if next.CancellationRequested() {
			break
		}
		next.Accept(v)
	}
	next.End()
}
func (sk *limit)CancellationRequested() bool  {
	return  sk.cancelRequest
}
