package easy

import "fmt"

func Example2() {

	l1 := &ListNode{9, &ListNode{8, nil}}
	l2 := &ListNode{1, nil}
	// l1 := &ListNode{1, &ListNode{8, nil}}
	// l2 := &ListNode{0, nil}
	// l1 := &ListNode{1, &ListNode{8, &ListNode{9, nil}}}
	// l2 := &ListNode{0, &ListNode{9, &ListNode{9, nil}}}
	r := addTwoNumbers(l1, l2)

	for {
		if r == nil {
			break
		}
		fmt.Println("xx: ", r.Val)
		r = r.Next
	}
	// Output:
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	f := func(i1, i2, j int) (int, *ListNode) {
		v := i1 + i2 + j
		if v > 9 {
			return 1, &ListNode{Val: v - 10, Next: nil}
		}
		return 0, &ListNode{Val: v, Next: nil}
	}

	iadd := 0
	r := &ListNode{}
	p, p1, p2 := r, l1, l2
	for {
		if p1 == nil && p2 == nil {
			if iadd != 0 {
				p.Next = &ListNode{Val: iadd, Next: nil}
			}
			break
		}
		if p1 == nil {
			iadd, p.Next = f(0, p2.Val, iadd)
			p, p2 = p.Next, p2.Next
			continue
		}
		if p2 == nil {
			iadd, p.Next = f(p1.Val, 0, iadd)
			p, p1 = p.Next, p1.Next
			continue
		}
		iadd, p.Next = f(p1.Val, p2.Val, iadd)
		p, p1, p2 = p.Next, p1.Next, p2.Next
	}
	return r.Next
}
