# [206. Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/description/)

## 题目

Reverse a singly linked list.

## 题目大意

翻转单链表


## 解题思路

- 基础的翻转操作；

- 遍历保存头、尾节点：创建新的节点prevHead指向头部, curNode指向翻转前的第一个节点，也是翻转后的最后一个节点

- 使用behind变量进行"插头"法

```
for head != nil {
    next := head.Next
    head.Next = behind
    behind = head
    head = next
}
```