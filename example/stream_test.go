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
	}).Map(func(v stream.T) stream.T {
		p:=v.(person)
		p.name = strings.ToUpper(p.name)
		fmt.Printf("Mapper %v\r\n",p)
		return p
	}).Peek(func(v stream.T) {
		fmt.Printf("Peek %v\r\n",v)
	}).Sorted(func(v1, v2 stream.T) bool {
		va:=v1.(person)
		vb:=v2.(person)
		return strings.Compare(va.name,vb.name)<0
	})
}

func TestForEach(t *testing.T) {
	st:=st()
	st.ForEach(func(v stream.T) {
		fmt.Printf("ForEach %v\r\n",v)
	})
}
func TestCount(t *testing.T)  {
	fmt.Println(st().Count())
}
func TestFindFirst(t *testing.T)  {
	fmt.Println(st().FindFirst())
}
func TestCollect(t *testing.T)  {
	fmt.Println(st().Collect())
}

