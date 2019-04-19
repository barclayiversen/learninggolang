package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func changeMe(p *person) {
	fmt.Println("changing")
	(*p).First = "biggie"
	(*p).Last = "sweet"
	//fmt.Println("changed?",p)
}

func main() {

	//p := person{
	//	First: "barclay",
	//	Last:  "iversen",
	//	Age:   30,
	//}
	//p1 := person{
	//	First: "cup",
	//	Last:  "cake",
	//	Age:   25,
	//}

	//bs, err := json.Marshal(people)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(bs)
	//fmt.Println(people)
	//fmt.Println(string(bs))

	s := `[{"First":"barclay","Last":"iversen","Age":30},{"First":"cup","Last":"cake","Age":25}]`
	ds := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", ds)

	people := []person{}

	err := json.Unmarshal(ds, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("all of the datat", people)

	for i, v := range people {
		fmt.Println("----- Person Number", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
