// Create by Yale 2019/3/18 17:33
package stream

import "sort"

type sortData struct {
	values []T
	les  Less
}

func (s sortData) Len() int {
	return len(s.values)
}
func (s sortData) Swap(i, j int) {
	s.values[i], s.values[j] = s.values[j], s.values[i]
}
func (s sortData) Less(i, j int) bool {
	return s.les(s.values[i], s.values[j])
}


type sorted struct {
	me *sinkNode
	les Less
	data []T
}

func (sk *sorted)Begin(size int64)  {
	sk.data = make([]T,0)
}
func (sk *sorted)Accept(t T)  {
	sk.data = append(sk.data,t)
}
func (sk *sorted)End() {
	sk.me.next.value.Begin(int64(len(sk.data)))

	st:=sortData{values:sk.data,les:sk.les}
	sort.Sort(&st)

	for _,v:=range st.values{
		if sk.me.next.value.CancellationRequested() {
			break
		}
		sk.me.next.value.Accept(v)
	}
	sk.me.next.value.End()

}
func (sk *sorted)CancellationRequested() bool  {
	return  false
}