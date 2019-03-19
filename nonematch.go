// Create by Yale 2019/3/19 10:05
package stream

type noneMatch struct {
	pre Predicate
	value bool
	cancelRequest bool
}

func (sk *noneMatch)Begin(size int64)  {
	sk.value = true
}
func (sk *noneMatch)Accept(t T)  {

	if sk.pre(t) {
		sk.value = false
		sk.cancelRequest = true
	}

}
func (sk *noneMatch)End() {

}
func (sk *noneMatch)CancellationRequested() bool  {
	return  sk.cancelRequest
}

