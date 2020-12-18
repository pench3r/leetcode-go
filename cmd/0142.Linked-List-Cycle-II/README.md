# [142. Linked List Cycle II](https://leetcode.com/problems/linked-list-cycle-ii/)

## 题目

Given a linked list, return the node where the cycle begins. If there is no cycle, return null.

To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.

Note: Do not modify the linked list.

Example 1:

```c
Input: head = [3,2,0,-4], pos = 1
Output: tail connects to node index 1
Explanation: There is a cycle in the linked list, where tail connects to the second node.
```

Example 2:

```c
Input: head = [1,2], pos = 0
Output: tail connects to node index 0
Explanation: There is a cycle in the linked list, where tail connects to the first node.
```

Example 3:

```c
Input: head = [1], pos = -1
Output: no cycle
Explanation: There is no cycle in the linked list.
```


## 题目大意

判断链表是否有环，不能使用额外的空间。如果有环，输出环的起点指针，如果没有环，则输出空。

## 解题思路

- 这道题在141的基础上要求输出起点指针；依然使用快慢指针进行解题，通过推导公式可以确定起点、环的起点、相遇点之间的关系

- 假设从起点到环的切入点需要走X1步，从切入点到相遇点需要走X2步，相遇点再回到切入点需要走X3步；由于slow和fast会相遇，那么两者的花费的时间是相同的

```
fast T = (X1 + X2 + X3 + X2) / 2
slow T = (X1 + X2) / 1

推导可得 X1 = X3
```

结果表明，相遇点到环切入点的距离 和 起点到切入点的距离相同；

- 找到相遇点后，从头部和相遇点同时移动，直到两者相遇即为切入点