/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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
// @lc code=end

