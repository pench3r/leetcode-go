# [23. Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/)

## 题目

Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.



Example :

```
Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6

```

## 题目大意

合并 K 个有序链表

## 解题思路

- 与归并排序类似，使用分治算法进行处理；代码实现使用二分+递归的方式

- 二分：直接使用list[:mid]，mid=length/2; 边界：length小于1时返回nil，length等于1时，返回当前节点

- 递归双链表合并算法：参数为两个链表的头节点s1、s2，返回一个新的链表头; 关键点：当s1.Val <= s2.Val, s1.Next = dfs(s1.Next, s2)；反之亦然

- 边界处理：s1==nil则返回s2，s2==nil则返回s1；