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

func (ft *sorted)Begin(size int64)  {
  ft.data = make([]T,0)
}
func (ft *sorted)Accept(t T)  {
	ft.data = append(ft.data,t)
}
func (ft *sorted)End() {
	ft.me.next.value.Begin(int64(len(ft.data)))

	st:=sortData{values:ft.data,les:ft.les}
	sort.Sort(&st)

	for _,v:=range st.values{
		if ft.me.next.value.CancellationRequested() {
			break
		}
		ft.me.next.value.Accept(v)
	}
	ft.me.next.value.End()

}
func (ft *sorted)CancellationRequested() bool  {
	return  false
}