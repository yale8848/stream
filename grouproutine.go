// Create by Yale 2019/3/19 11:42
package stream

import (
	"sync"
)

type groupRoutine struct {
	me *sinkNode
	value []T
	fun TS
	num uint64
	err ErrMsg
}

func (sk *groupRoutine)Begin(size int64)  {
	sk.value = make([]T,0)
}
func (sk *groupRoutine)Accept(t T)  {
	sk.value = append(sk.value,t)
}
func (sk *groupRoutine)End() {
	num:=sk.num

	if num <=1||len(sk.value)/int(num) == 0{
		sk.fun(sk.value)
	}else{
		n:=len(sk.value)/int(num)
		if len(sk.value)%int(num)!=0{
			n = n+1
		}

		wg:=sync.WaitGroup{}
		wg.Add(n)
		for i:=0;i<n;i++{

			go func(i int,n int) {
				defer func() {
					wg.Done()
				}()
				defer func() {
					if err:=recover();err!=nil {
						if sk.err!=nil {
							sk.err(err)
						}
					}
				}()
				if i == n-1 {
					sk.fun(sk.value[i*int(num):len(sk.value)])
				}else{
					sk.fun(sk.value[i*int(num):i*int(num)+int(num)])
				}
			}(i,n)

		}
		wg.Wait()

	}

}
func (sk *groupRoutine)CancellationRequested() bool  {
	return  false
}

