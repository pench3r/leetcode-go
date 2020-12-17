# [25. Reverse Nodes in k-Group](https://leetcode.com/problems/reverse-nodes-in-k-group/description/)

## 题目

Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.

Example:

```c
Given this linked list: 1->2->3->4->5

For k = 2, you should return: 2->1->4->3->5

For k = 3, you should return: 3->2->1->4->5
```

Note:

- Only constant extra memory is allowed.
- You may not alter the values in the list's nodes, only nodes itself may be changed.


## 题目大意

按照每 K 个元素翻转的方式翻转链表。如果不满足 K 个元素的就不翻转。

## 解题思路

- 题目要求不使用额外的内存进行处理；使用递归方式：先获取翻转区间的end节点，再进行翻转返回新的head节点，最后翻转后的尾节点指向下一个翻转区间的递归处理

- dfs的参数：head、k，head为起始节点，k理解为长度，获取到end节点；

- 翻转函数参数：first、last分别为头、尾节点；处理逻辑如下：

```
prev := last
for first != last {
    next := first.Next  // 保存下一个遍历节点
    first.Next = prev   // 与尾部节点连接
    prev = first
    first = next
}
```

其中prev保存着尾部节点，first保存当前节点