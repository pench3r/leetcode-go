# [86. Partition List](https://leetcode.com/problems/partition-list/)

## 题目

Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.

You should preserve the original relative order of the nodes in each of the two partitions.

Example:

```
Input: head = 1->4->3->2->5->2, x = 3
Output: 1->2->2->4->3->5
```


## 题目大意

给定一个数 x，比 x 大或等于的数字都要排列在比 x 小的数字后面，并且相对位置不能发生变化。由于相对位置不能发生变化，所以不能用类似冒泡排序的思想。

## 解题思路

- 结果分为两部分，前面的list比x小，后面的list比x大；维护before和after两个指针，当前元素小于x则直接挂在before指针后，并移动before指针，反之针对after进行相同的操作；最后将afterhead直接挂在before的后面即可

- 使用cur指针进行单次循环，before、after保存对应的链表头部，beforeHead和afterHead保存对应的链表

- before和after的初始化为一个空的链表节点：`&ListNode{}`