// Create by Yale 2019/3/19 11:11
package stream

type groupStrings struct {
	me *sinkNode
	value []T
	fun func([]string)
	num uint64
}

func (sk *groupStrings)Begin(size int64)  {
	sk.value = make([]T,0)
}
func (sk *groupStrings)Accept(t T)  {
  sk.value = append(sk.value,t)
}
func (sk *groupStrings)End() {
	num:=sk.num

	toStrings:= func(ts []T) []string{
		ret:= make([]string,len(ts))

		for i,v:=range ts{
			ret[i]= v.(string)
		}
		return ret
	}
	sv:=toStrings(sk.value)
	if num <=1||len(sk.value)/int(num) == 0{
		sk.fun(sv)
	}else{
		n:=len(sk.value)/int(num)
		if len(sk.value)%int(num)!=0{
			n = n+1
		}
		for i:=0;i<n;i++{
			if i == n-1 {
				sk.fun(sv[i*int(num):len(sk.value)])
			}else{
				sk.fun(sv[i*int(num):i*int(num)+int(num)])
			}
		}
	}

}
func (sk *groupStrings)CancellationRequested() bool  {
	return  false
}
