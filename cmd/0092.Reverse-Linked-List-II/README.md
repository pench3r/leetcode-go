# [92. Reverse Linked List II](https://leetcode.com/problems/reverse-linked-list-ii/)

## 题目

Reverse a linked list from position m to n. Do it in one-pass.

Note: 1 ≤ m ≤ n ≤ length of list.

Example:

```
Input: 1->2->3->4->5->NULL, m = 2, n = 4
Output: 1->4->3->2->5->NULL
```


## 题目大意

给定 2 个链表中结点的位置 m, n，反转这个两个位置区间内的所有结点。

## 解题思路

- 先获取到翻转区间的起点的前一个节点，再处理翻转区间范围内的节点；整个翻转过程可以看作是头部、翻转区间、尾部，确保遍历过程中翻转区间与头、尾的关联关系

- 边界处理：当head为空 或者 m>=n时，直接返回head节点

- 为了便于循环处理，创建一个空的节点它的Next指向head，遍历到m-1处即为我们要找的节点，定义为prev

- 翻转区间的遍历遍历为n-m，在遍历中的3个关键节点：

1. prev作为翻转区间前的一个节点，用于维护翻转区间的头部，后续节点都是依次插入该头部位置；prev不会变动只会更新prev.Next

2. prev.Next就是每次遍历待更新的部分

2. 使用cur保存prev.Next节点，也就是翻转区间的起始点而在翻转后为终点；cur不会变动，只会在遍历的过程中使用cur.Next来保存翻转区间与尾部的关联节点

翻转过程：

```
cur := prev.Next
for i:=0; i<n-m; i++ {
    next := prev.Next
    prev.Next = cur.Next
    prev.Next.Next = next
    cur.Next = cur.Next.Next
}
```