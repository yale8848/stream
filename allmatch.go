// Create by Yale 2019/3/19 10:05
package stream

type allMatch struct {
	pre Predicate
	value bool
	cancelRequest bool
}

func (sk *allMatch)Begin(size int64)  {
	sk.value = true
}
func (sk *allMatch)Accept(t T)  {

	if !sk.pre(t) {
		sk.value = false
		sk.cancelRequest = true
	}

}
func (sk *allMatch)End() {

}
func (sk *allMatch)CancellationRequested() bool  {
	return  sk.cancelRequest
}

