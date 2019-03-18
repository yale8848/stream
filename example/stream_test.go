// Create by Yale 2019/2/25 16:08
package example

import (
	"fmt"
	"github.com/yale8848/stream"
	"strings"
	"testing"
)

var data =[]person{{age:11,name:"alice"},{age:19,name:"pig"},{age:5,name:"cat"},{age:21,name:"bob"},{age:13,name:"pig"},{age:6,name:"lili"}}

func st() stream.Stream  {
	return stream.Of(data).Filter(func(v stream.T) bool {
		p:=v.(person)
		fmt.Printf("Filter %v\r\n",v)
		return  p.age>10
	}).Peek(func(v stream.T) {
		fmt.Printf("Peek %v\r\n",v)
	}).Skip(1).Map(func(v stream.T) stream.T {
		p:=v.(person)
		p.name = strings.ToUpper(p.name)
		return p
	}).Sorted(func(v1, v2 stream.T) bool {
		s1:=v1.(person)
		s2:=v2.(person)
		return  strings.Compare(s1.name,s2.name)<0
	}).Limit(2).Distinct(func(v stream.T) stream.T {
		p:=v.(person)
		fmt.Printf("Distinct %v\r\n",v)
		return p.name
	})
}
func TestOf(t *testing.T) {

	st().ForEach(func(v stream.T) {
		fmt.Printf("ForEach %v\r\n",v)
	})

}
func TestMin(t *testing.T)  {
   fmt.Println(st().Min(func(min stream.T, v stream.T) bool {
	   m:=min.(person)
	   v1:=v.(person)
	   return m.age > v1.age
   }))
}
func TestMax(t *testing.T)  {
	fmt.Println(st().Max(func(max stream.T, v stream.T) bool {
		m:=max.(person)
		v1:=v.(person)
		return m.age < v1.age
	}))
}
func TestCollect(t *testing.T)  {
	fmt.Println(st().Collect())
}

func TestCount(t *testing.T)  {
	fmt.Println(stream.Of(data).Filter(func(v stream.T) bool {
		p:=v.(person)
		fmt.Printf("Filter %v\r\n",v)
		return  p.age>10
	}).Count())
}
func TestAnyMatch(t *testing.T)  {
	fmt.Println(st().AnyMatch(func(v stream.T) bool {
		v1:=v.(person)
		return v1.age>18
	}))
}
func TestAllMatch(t *testing.T)  {
	fmt.Println(st().AllMatch(func(v stream.T) bool {
		v1:=v.(person)
		return v1.age==19
	}))
}
func TestNoneMatch(t *testing.T)  {
	fmt.Println(st().NoneMatch(func(v stream.T) bool {
		v1:=v.(person)
		return v1.age==20
	}))
}
func TestFindFirst(t *testing.T)  {
	fmt.Println(st().FindFirst())
}
