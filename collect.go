// Create by Yale 2019/3/18 19:42
package stream


type collect struct {
	num int64
	value []T
	index int64
	size int64

}

func (sk *collect)Begin(size int64)  {
	sk.size = size
	if size == -1 {
		sk.value= make([]T,0)
	}else{
		sk.value= make([]T,size)
	}
}
func (sk *collect)Accept(t T)  {
	if sk.size == -1 {
		sk.value = append(sk.value,t)
	}else{
		sk.value[sk.index] = t
		sk.index++
	}
}
func (sk *collect)End() {

}
func (sk *collect)CancellationRequested() bool  {
	return  false
}


