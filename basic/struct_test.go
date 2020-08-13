package basic

import (
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

func ExampleStruct() {

	fs := []*Doc{&Doc{"a"}, &Doc{"b"}}

	wg := sync.WaitGroup{}
	wg.Add(2)
	for i := range fs {
		f := fs[i]
		go func() {
			fmt.Println("before ", f.name)
			change(f)
			fmt.Println("after ", f.name)
			wg.Done()
		}()
	}
	wg.Wait()

	// Output:
}

type Doc struct {
	name string
}

func change(f *Doc) error {
	f.name = "bbb"
	return nil
}

type Person struct {
	Name  string
	Label string
	Age   int
	Order int
}

func ExampleXxx() {

	p1 := Person{Name: "a", Label: "A"}
	p2 := Person{Name: "a1", Label: "A1"}

	p := []*Person{&p1, &p2}

	ss, err := json.Marshal(&p)
	fmt.Println(string(ss), err)

	pp := []*Person{}
	if err := json.Unmarshal([]byte(`[{"Name":"a","Label":"A"},{"Name":"a1","Label":"A1"}]`), &pp); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(pp[0].Name)
	fmt.Println(pp[1].Name)

	// Output:
}

func ExampleSort() {

	p1 := Person{Name: "a", Label: "A", Age: 4}
	p2 := Person{Name: "a1", Label: "A1", Age: 1}
	p3 := Person{Name: "a2", Label: "C1", Age: 3}
	p4 := Person{Name: "a2", Label: "C1", Age: 2}

	p := []*Person{&p1, &p2, &p3, &p4}

	m := map[int]int{3: 1, 4: 2, 1: 3}
	for i := range p {
		if v, ok := m[p[i].Age]; ok {
			p[i].Order = v
		} else {
			p[i].Order = -1
		}
	}

	sort.Slice(p, func(i, j int) bool {
		return p[i].Order > p[j].Order
	})

	for _, v := range p {
		fmt.Println(v)
	}

	// Output:
}

