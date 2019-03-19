// Create by Yale 2019/3/19 11:11
package stream

type group struct {
	me *sinkNode
	value []T
	fun TS
	num uint64
}

func (sk *group)Begin(size int64)  {
	sk.value = make([]T,0)
}
func (sk *group)Accept(t T)  {
  sk.value = append(sk.value,t)
}
func (sk *group)End() {
	num:=sk.num

	if num <=1||len(sk.value)/int(num) == 0{
		sk.fun(sk.value)
	}else{
		n:=len(sk.value)/int(num)
		if len(sk.value)%int(num)!=0{
			n = n+1
		}
		for i:=0;i<n;i++{
			if i == n-1 {
				sk.fun(sk.value[i*int(num):len(sk.value)])
			}else{
				sk.fun(sk.value[i*int(num):i*int(num)+int(num)])
			}
		}
	}

}
func (sk *group)CancellationRequested() bool  {
	return  false
}
