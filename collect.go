// Create by Yale 2019/3/18 19:42
package stream


type collect struct {
	num int64
	value []T
	index int64
	size int64
}

func (ft *collect)Begin(size int64)  {
	ft.size = size
	if size == -1 {
		ft.value= make([]T,0)
	}else{
		ft.value= make([]T,size)
	}
}
func (ft *collect)Accept(t T)  {
	if ft.size == -1 {
		ft.value = append(ft.value,t)
	}else{
		ft.value[ft.index] = t
		ft.index++
	}
}
func (ft *collect)End() {

}
func (ft *collect)CancellationRequested() bool  {
	return  false
}


