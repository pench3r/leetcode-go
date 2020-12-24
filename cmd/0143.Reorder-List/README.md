# [143. Reorder List](https://leetcode.com/problems/reorder-list/)

## 题目

Given a singly linked list L: L0→L1→…→Ln-1→Ln,
reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…

You may not modify the values in the list's nodes, only nodes itself may be changed.

Example 1:

```c
Given 1->2->3->4, reorder it to 1->4->2->3.
```

Example 2:

```c
Given 1->2->3->4->5, reorder it to 1->5->2->4->3.
```

## 题目大意

按照指定规则重新排序链表：第一个元素和最后一个元素排列在一起，接着第二个元素和倒数第二个元素排在一起，接着第三个元素和倒数第三个元素排在一起。


## 解题思路


- 找到链表中间点、翻转后半段的链表、拼接前半段和后半段

- 定位链表中间点：快慢指针，slow每次移动一步、fast每次移动两步

- 翻转后半段的链表：找到的中间节点作为后半段的链表头，不参与翻转；使用preCurNode保存翻转前的第一个节点，也是翻转后的最后一个节点；

- 拼接前半段和后半段：从头开始遍历直到中间节点位置停止，依次将后半段的节点插入前半段中