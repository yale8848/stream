// Create by Yale 2019/3/18 19:42
package stream


type collect struct {
	num int64
	value []T
}

func (ft *collect)Begin(size int64)  {
	if size == -1 {
		ft.value= make([]T,0)
	}else{
		ft.value= make([]T,size)
	}
}
func (ft *collect)Accept(t T)  {
	ft.value = append(ft.value,t)
}
func (ft *collect)End() {

}
func (ft *collect)CancellationRequested() bool  {
	return  false
}


