# [102. Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/)

## 题目


Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],

```c
    3
   / \
  9  20
    /  \
   15   7
```

return its level order traversal as:

```c
[
  [3],
  [9,20],
  [15,7]
]
```
 

## 题目大意

按层序从上到下遍历一颗树。

## 解题思路

- 使用队列实现的BFS，使用curCounter记录当前层数的个数，nextLevel记录下一层的个数；如果curCounter为0，则说明当前层已经遍历结束，保存结果并使用nextLevel的值作为新的curCounter；否则弹出队列头部元素，将子节点分别加入队列，并加入nextLevel的计数，curCounter减一。

- 使用DFS，需要将层数作为参数进行传递

```
    dfs(root, levelNum, *res)
```