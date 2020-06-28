/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	next1, next2 := l1, l2
	l := new(ListNode)
	next := l
	for next1 != nil || next2 != nil {
		if next1 != nil && (next2 == nil || next1.Val <= next2.Val) {
			next.Next = &ListNode{Val: next1.Val, Next: nil}
			next = next.Next
			next1 = next1.Next
		}
		if next2 != nil && (next1 == nil || next2.Val <= next1.Val) {
			next.Next = &ListNode{Val: next2.Val, Next: nil}
			next = next.Next
			next2 = next2.Next
		}
	}
	return l.Next
}
// @lc code=end

