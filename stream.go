// Create by Yale 2019/2/22 19:47
package stream

import (
	"reflect"
)


type sink interface {
	Begin(size int64)
	Accept(t T)
	End()
	CancellationRequested() bool
}

type sinkNode struct {
	value sink
	next *sinkNode
}

type T interface{}

type Predicate func(v T) bool

type Function func(v T) T

type Consumer func(v T)

type Less func(v1, v2 T) bool

type MinCompare func(min T, v T) bool
type MaxCompare func(max T, v T) bool

type stream struct {
	values []T
	link *sinkNode
	head *sinkNode
}


type Stream interface {
	///Intermediate ops
	//Stateless
	Filter(p Predicate) Stream
	Map(f Function) Stream
	Peek(c Consumer) Stream


	//Stateful
	Sorted(les Less) Stream
	Distinct(f Function)Stream
	Skip(n int64)
	///

	///Terminal ops
	//None Short-circuiting
	ForEach(c Consumer)
	Count() int64
	Collect()[]T

	//Short-circuiting
	FindFirst()T
	///
	do()

}
func OfAny(arr ...T)Stream{
	return Of(arr)
}
func Of(arr T) Stream {
	link:=&sinkNode{}
	stm := &stream{link:link}
	stm.head = link

	tp := reflect.TypeOf(arr)
	if tp.Kind() == reflect.Array || tp.Kind() == reflect.Slice {
		arrValue := reflect.ValueOf(arr)
		values := make([]T, arrValue.Len())
		for i := 0; i < len(values); i++ {
			v := arrValue.Index(i)
			values[i] = v.Interface()
		}
		stm.values = values
	}
	return stm
}
func (stm *stream)Distinct(f Function)Stream{
	sk:=&distinct{fun:f}
	n:=&sinkNode{value:sk}
	stm.link.next = n
	stm.link = n
	sk.me = stm.link
	return stm
}
func (stm *stream)Collect()[]T{
	sk:=&collect{}
	n:=&sinkNode{value:sk}
	stm.link.next = n
	stm.do()
	return sk.value
}

func (stm *stream)FindFirst()T{
	sk:=&findFirst{}
	n:=&sinkNode{value:sk}
	stm.link.next = n
	stm.do()
	return sk.value
}
func (stm *stream)Sorted(les Less) Stream{
	sk:=&sorted{les:les}
	n:=&sinkNode{value:sk}
	stm.link.next = n
	stm.link = n
	sk.me = stm.link
	return stm
}
func (stm *stream)Peek( c Consumer) Stream{
	sk:=&peek{con:c}
	n:=&sinkNode{value:sk}
	stm.link.next = n
	stm.link = n
	sk.me = stm.link
	return stm
}
func (stm *stream) Filter(p Predicate) Stream {

	sk:=&filter{pre:p}
	n:=&sinkNode{value:sk}
	stm.link.next = n
	stm.link = n
	sk.me = stm.link
	return stm
}
func (stm *stream) Map(f Function) Stream {

	sk:=&mapper{fun:f}
	n:=&sinkNode{value:sk}
	stm.link.next = n
	stm.link = n
	sk.me = stm.link
	return stm
}
func (stm *stream)do(){

	stm.head.next.value.Begin(-1)
	for _,v:=range stm.values{
		stm.head.next.value.Accept(v)
	}
	stm.head.next.value.End()

}
func (stm *stream) ForEach(c Consumer) {
	sk:=&foreach{cons:c}
	n:=&sinkNode{value:sk}
	stm.link.next = n
	stm.do()
}

func (stm *stream) Count() int64 {
	sk:=&count{}
	n:=&sinkNode{value:sk}
	stm.link.next = n
	stm.do()
	return sk.num
}
