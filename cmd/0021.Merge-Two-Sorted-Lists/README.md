# [21. Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/)

## 题目

Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example :

```
Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4

```

## 题目大意

合并 2 个有序链表

## 解题思路

- 使用递归的方式进行链表的合并，该题目为23题的简化版

- 递归函数的参数2个链表节点，返回新的链表节点；关键逻辑s1.Val<=s2.Val时，s1.Next = dfs(s1.Next, s2),否则s2.Next=dfs(s2.Next, s1)

- 边界处理：s1或s2为nil时，则返回另外一个