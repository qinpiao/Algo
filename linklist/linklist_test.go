package linklist

import (
	"testing"
	//"fmt"
)

func TestLinkList(t *testing.T) {
	s := []interface{}{1,2,3,4,5,6}
	l := NewLinkList(s...)
	l.String()
	//fmt.Println(l.length)
	//l.Append(2)
	//l.String()
	//l.Delete(1)
	//l.String()
	//if err := l.Insert(0, 1); err != nil {
	//	fmt.Println(err)
	//}
	//l.String()
	//if data, err := l.Get(7); err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(data)
	//}
}