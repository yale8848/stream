package stream

type reduce struct {
	r        BinaryOperator
	identity T
	result   T
}

func (sk *reduce) Begin(size int64) {
	sk.result = sk.identity
}
func (sk *reduce) Accept(t T) {
	sk.result = sk.r(sk.result, t)
}
func (sk *reduce) End() {
}
func (sk *reduce) CancellationRequested() bool {
	return false
}
