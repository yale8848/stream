// Create by Yale 2019/2/22 19:47
package stream

import (
	"errors"
	"reflect"
)


type sinkNode struct {
	value sink
	next *sinkNode
}

type sink interface {
	Begin(size int64)
	Accept(t T)
	End()
	CancellationRequested() bool
}

type TS func([]T)
type ErrMsg func(recoverErr interface{})
type T interface{}

type Sum func(v T) int64

type Predicate func(v T) bool

type Function func(v T) T

type Consumer func(v T)

type ConsumerCancel func(v T) bool

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
	Skip(num uint64)Stream
	Limit(num uint64)Stream



	///

	///Terminal ops
	//None Short-circuiting
	ForEach(c ConsumerCancel)
	Count() int64
	Collect()[]T
	Min(min MinCompare)T
	Max(max MaxCompare)T
	Sum(s Sum)int64

	//Short-circuiting
	FindFirst()T
	AnyMatch(pre Predicate)bool
	AllMatch(pre Predicate)bool
	NoneMatch(pre Predicate)bool

	///
	do()

	///
	Group(num uint64,fun TS)
	GroupStrings(num uint64,fun func([]string))
	GroupRoutine(num uint64,fun TS,err ErrMsg)


}
func (stm *stream)Sum(s Sum)int64{
	sk:=&sum{s:s}
	n:=&sinkNode{value:sk}
	stm.link.next = n
	stm.do()
	return  sk.num
}
func (stm *stream)GroupRoutine(num uint64,fun TS,err ErrMsg){
	sk:=&groupRoutine{num:num,fun:fun,err:err}
	n:=&sinkNode{value:sk}
	stm.link.next = n
	stm.do()
}
func (stm *stream)Group(num uint64,fun TS){
	sk:=&group{num:num,fun:fun}
	n:=&sinkNode{value:sk}
	stm.link.next = n
	stm.do()
}
func (stm *stream)GroupStrings(num uint64,fun func([]string)){
	sk:=&groupStrings{num:num,fun:fun}
	n:=&sinkNode{value:sk}
	stm.link.next = n
	stm.do()
}
func OfAny(arr ...T)(Stream,error){
	return Of(arr)
}
func Of(arr T) (Stream,error) {
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
	}else{
		return nil,errors.New("value must Array or Slice")
	}
	return stm,nil
}
func (stm *stream)AllMatch(pre Predicate)bool{
	sk:=&allMatch{pre:pre}
	n:=&sinkNode{value:sk}
	stm.link.next = n
	stm.do()
	return sk.value
}
func (stm *stream)NoneMatch(pre Predicate)bool{
	sk:=&noneMatch{pre:pre}
	n:=&sinkNode{value:sk}
	stm.link.next = n
	stm.do()
	return sk.value
}
func (stm *stream)AnyMatch(pre Predicate)bool{
	sk:=&anyMatch{pre:pre}
	n:=&sinkNode{value:sk}
	stm.link.next = n
	stm.do()
	return sk.value
}
func (stm *stream) Max(m MaxCompare)T {
	sk:=&max{compare:m}
	n:=&sinkNode{value:sk}
	stm.link.next = n
	stm.do()
	return sk.value
}
func (stm *stream) Min(m MinCompare)T {
	sk:=&min{compare:m}
	n:=&sinkNode{value:sk}
	stm.link.next = n
	stm.do()
	return sk.value
}

func (stm *stream)Limit(num uint64)Stream{
	sk:=&limit{num:num}
	n:=&sinkNode{value:sk}
	stm.link.next = n
	stm.link = n
	sk.me = stm.link
	return stm
}
func (stm *stream)Skip(num uint64)Stream{
	sk:=&skip{num:num}
	n:=&sinkNode{value:sk}
	stm.link.next = n
	stm.link = n
	sk.me = stm.link
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
func (stm *stream) ForEach(c ConsumerCancel) {
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
