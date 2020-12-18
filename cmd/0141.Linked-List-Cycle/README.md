# [141. Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/description/)

## 题目

Given a linked list, determine if it has a cycle in it.

Follow up:   
Can you solve it without using extra space?



## 题目大意

判断链表是否有环，不能使用额外的空间。

## 解题思路

- 使用快慢指针进行解题；slow、fast指针从头部开始，slow每次移动一步，fast每次移动两步，当两者相遇则存在环

- 循环的边界：`fast.Next != nil && slow != nil && fast != nil`