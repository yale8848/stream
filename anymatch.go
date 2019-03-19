// Create by Yale 2019/3/19 10:05
package stream

type anyMatch struct {
	pre Predicate
	value bool
	cancelRequest bool
}

func (sk *anyMatch)Begin(size int64)  {
}
func (sk *anyMatch)Accept(t T)  {

	if sk.pre(t) {
		sk.value = true
		sk.cancelRequest = true
	}

}
func (sk *anyMatch)End() {

}
func (sk *anyMatch)CancellationRequested() bool  {
	return  sk.cancelRequest
}

