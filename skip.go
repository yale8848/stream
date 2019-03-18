// Create by Yale 2019/3/18 20:10
package stream

type skip struct {
	me *sinkNode
	num uint64
	cancelRequest bool
	value []T

	len int64
	count int64
}

func (ft *skip)Begin(size int64)  {

	if size == 0 {
		ft.cancelRequest = true
	}else if size == -1 {
		ft.value = make([]T,ft.len)
	}else {
		if size <= int64(ft.num) {
			ft.cancelRequest = true
		}else{
			ft.len = size-int64(ft.num)
			ft.value =  make([]T,ft.len)
		}
	}
}
func (ft *skip)Accept(t T)  {
	if ft.num == 0 {
		if ft.len == 0 {
			ft.value =append(ft.value,t)
		}else{
			ft.value[ft.count] = t
		}
	}
}
func (ft *skip)End() {

	ft.me.value.Begin(int64(len(ft.value)))

	for _,v:=range ft.value{
		if ft.me.next.value.CancellationRequested() {
			break
		}
		ft.me.next.value.Accept(v)
	}
	ft.me.value.End()
}
func (ft *skip)CancellationRequested() bool  {
	return  ft.cancelRequest
}
