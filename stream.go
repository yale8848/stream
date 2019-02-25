// Create by Yale 2019/2/22 19:47
package stream

import (
	"reflect"
	"sort"
)

type T interface{}
type Less func(v1, v2 T) bool

type Predicate func(v T) bool
type Function func(v T) T
type Consumer func(v T)
type MinCompare func(min T, v T) bool
type MaxCompare func(max T, v T) bool

type eachFun func(index int, v T) (T, bool, bool)
type stream struct {
	values []T
	stop   bool
	items  []eachFun
	reload bool
	one    T
	find   bool
}

type sortData struct {
	values []T
	less   Less
}

func (s sortData) Len() int {
	return len(s.values)
}
func (s sortData) Swap(i, j int) {
	s.values[i], s.values[j] = s.values[j], s.values[i]
}
func (s sortData) Less(i, j int) bool {
	return s.less(s.values[i], s.values[j])
}

type Stream interface {
	Filter(p Predicate) Stream
	Map(f Function) Stream
	Sorted(less Less) Stream
	Peek(c Consumer) Stream
	Distinct(f Function) Stream
	Skip(n int) Stream
	Limit(n int) Stream

	ForEach(c Consumer)
	Collect() []T
	Min(mc MinCompare) T
	Max(mc MaxCompare) T
	Count() int
	AnyMatch(p Predicate) bool
	AllMatch(p Predicate) bool
	NoneMatch(p Predicate) bool
	FindFirst() T
}
func OfAny(arr ...T)Stream{
	return Of(arr)
}
func Of(arr T) Stream {
	stm := &stream{items: make([]eachFun, 0)}
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
func (stm *stream) AnyMatch(p Predicate) bool {

	fun := func() eachFun {

		return func(index int, v T) (T, bool, bool) {

			if index == len(stm.values)-1 {
				stm.stop = true
			}
			if p(v) {
				stm.find = true
				stm.stop = true
			}
			return nil, false, false
		}
	}
	stm.items = append(stm.items, fun())
	stm.handle()
	return stm.find
}
func (stm *stream) AllMatch(p Predicate) bool {

	fun := func() eachFun {
		find := true
		return func(index int, v T) (T, bool, bool) {

			if !p(v) {
				find = false
				stm.stop = true
			}

			if index == len(stm.values)-1 {
				stm.find = find
				stm.stop = true
			}

			return nil, false, false
		}
	}
	stm.items = append(stm.items, fun())
	stm.handle()
	return stm.find
}

func (stm *stream) NoneMatch(p Predicate) bool {

	fun := func() eachFun {
		find := false
		return func(index int, v T) (T, bool, bool) {

			if p(v) {
				find = true
				stm.stop = true
			}

			if index == len(stm.values)-1 {
				stm.find = find
				stm.stop = true
			}
			return nil, false, false
		}
	}
	stm.items = append(stm.items, fun())
	return !stm.find
}

func (stm *stream) FindFirst() T {
	fun := func() eachFun {
		return func(index int, v T) (T, bool, bool) {
			stm.one = v
			stm.stop = true
			return nil, false, false
		}
	}
	stm.items = append(stm.items, fun())
	stm.handle()
	return stm.one
}

func (stm *stream) Limit(n int) Stream {
	fun := func() eachFun {

		temp := make([]T, 0)
		return func(index int, v T) (T, bool, bool) {
			temp = append(temp, v)
			if index == len(stm.values)-1 {
				if n > len(temp) {
					n = len(temp)
				}
				if n < 0 {
					n = 0
				}
				stm.values = temp[:n]
				return nil, false, true
			}
			return nil, false, false
		}
	}
	stm.items = append(stm.items, fun())
	return stm
}
func (stm *stream) handle() {
	for index, v := range stm.values {
		if stm.stop {
			break
		}
		if stm.reload {
			break
		}
		for indexItem, itemV := range stm.items {

			vn, next, reLoad := itemV(index, v)
			if reLoad {

				if indexItem == len(stm.items)-1 {
					stm.items = make([]eachFun, 0)
				} else {
					stm.items = stm.items[indexItem+1:]
				}
				stm.reload = true
				break
			}
			if !next {
				break
			}
			if vn != nil {
				v = vn
			}
		}
	}
	if stm.reload {
		stm.reload = false
		stm.handle()
	}

}
func (stm *stream) Filter(p Predicate) Stream {

	fun := func() eachFun {
		filtered := false

		return func(index int, v T) (T, bool, bool) {
			next := false
			if p(v) {
				filtered = true
				next = true
			}
			if index == len(stm.values)-1 {
				if !filtered {
					stm.stop = true
				}
			}
			return nil, next, false
		}
	}
	stm.items = append(stm.items, fun())
	return stm
}
func (stm *stream) Map(f Function) Stream {
	fun := func() eachFun {
		return func(index int, v T) (T, bool, bool) {
			return f(v), true, false
		}
	}
	stm.items = append(stm.items, fun())
	return stm
}
func (stm *stream) Distinct(f Function) Stream {

	fun := func() eachFun {
		dMap := make(map[T]T, 0)
		return func(index int, v T) (T, bool, bool) {

			dMap[f(v)] = v
			if index == len(stm.values)-1 {
				nv := make([]T, 0)
				for _, value := range dMap {
					nv = append(nv, value)
					stm.values = nv
				}
				return nil, false, true
			}
			return nil, false, false
		}
	}
	stm.items = append(stm.items, fun())
	return stm
}
func (stm *stream) Sorted(less Less) Stream {
	fun := func() eachFun {

		dataCollect := sortData{
			values: make([]T, 0),
			less:   less,
		}
		return func(index int, v T) (T, bool, bool) {
			dataCollect.values = append(dataCollect.values, v)
			if index == len(stm.values)-1 {
				sort.Sort(&dataCollect)
				stm.values = dataCollect.values
				return nil, false, true
			}
			return nil, false, false
		}
	}
	stm.items = append(stm.items, fun())
	return stm
}
func (stm *stream) Skip(n int) Stream {
	fun := func() eachFun {

		temp := make([]T, 0)
		return func(index int, v T) (T, bool, bool) {
			temp = append(temp, v)
			if index == len(stm.values)-1 {
				if n > len(temp) {
					n = len(temp)
				}
				if n < 0 {
					n = 0
				}
				stm.values = temp[n:]
				return nil, false, true
			}
			return nil, false, false
		}
	}
	stm.items = append(stm.items, fun())
	return stm
}
func (stm *stream) ForEach(c Consumer) {
	fun := func() eachFun {
		return func(index int, v T) (T, bool, bool) {
			c(v)
			if index == len(stm.values)-1 {
				stm.stop = true
			}
			return nil, false, false
		}
	}
	stm.items = append(stm.items, fun())

	stm.handle()
}
func (stm *stream) Peek(c Consumer) Stream {
	fun := func() eachFun {
		return func(index int, v T) (T, bool, bool) {
			c(v)
			return nil, true, false
		}
	}
	stm.items = append(stm.items, fun())

	return stm
}
func (stm *stream) Collect() []T {
	fun := func() eachFun {
		temp := make([]T, 0)
		return func(index int, v T) (T, bool, bool) {
			temp = append(temp, v)
			if index == len(stm.values)-1 {
				stm.values = temp
			}
			return nil, false, false
		}
	}
	stm.items = append(stm.items, fun())

	stm.handle()

	return stm.values

}

func (stm *stream) Min(mc MinCompare) T {
	fun := func() eachFun {
		return func(index int, v T) (T, bool, bool) {
			if stm.one == nil {
				stm.one = v
			}
			if mc(stm.one, v) {
				stm.one = v
			}
			return nil, false, false
		}
	}
	stm.items = append(stm.items, fun())
	stm.handle()
	return stm.one

}
func (stm *stream) Max(mc MaxCompare) T {
	fun := func() eachFun {
		return func(index int, v T) (T, bool, bool) {
			if stm.one == nil {
				stm.one = v
			}
			if mc(stm.one, v) {
				stm.one = v
			}
			return nil, false, false
		}
	}
	stm.items = append(stm.items, fun())
	stm.handle()
	return stm.one

}
func (stm *stream) Count() int {
	stm.handle()
	return len(stm.values)
}
