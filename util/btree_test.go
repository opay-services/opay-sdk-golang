package util

import (
	"encoding/json"
	"fmt"
	"github.com/emirpasic/gods/maps/treemap"
	"testing"
)

func TestRedBlackTree(t *testing.T)  {

	person := Person{Name:"lijian", Age:12}
	str, _ := json.Marshal(person)
	fmt.Println(string(str))

	m:=treemap.NewWithStringComparator()
	m.Put("lijian", "222")
	m.Put("wangliang", 333)


	fmt.Println(m.Get("lijian"))
	fmt.Println(m.Get("wangliang"))
	fmt.Println(m.Get("lijia"))
}


type Person struct {
	Age  int
	Name string
}


