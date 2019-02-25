// Create by Yale 2019/2/25 16:08
package example

import (
	"fmt"
	"github.com/yale8848/stream"
	"strings"
	"testing"
)

var data =[]person{{age:11,name:"alice"},{age:19,name:"pig"},{age:5,name:"cat"},{age:21,name:"bob"}}

func TestOf(t *testing.T) {

	stream.Of(data).Filter(func(v stream.T) bool {
		p:=v.(person)
		return  p.age>10
	}).Peek(func(v stream.T) {
		fmt.Printf("Peek %v\r\n",v)
	}).Sorted(func(v1, v2 stream.T) bool {
		s1:=v1.(person)
		s2:=v2.(person)
		return  strings.Compare(s1.name,s2.name)<0
	}).Map(func(v stream.T) stream.T {
		p:=v.(person)
		p.name = strings.ToUpper(p.name)
		return p
	}).ForEach(func(v stream.T) {
		fmt.Printf("ForEach %v\r\n",v)
	})

}


