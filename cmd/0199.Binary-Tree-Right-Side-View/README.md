# [199. Binary Tree Right Side View](https://leetcode.com/problems/binary-tree-right-side-view/)

## 题目

Given a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.

Example:

```c
Input: [1,2,3,null,5,null,4]
Output: [1, 3, 4]
Explanation:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---
```

 

## 题目大意

从右边看一个树，输出看到的数字。注意有遮挡。


## 解题思路

- 获取每一层的最右边的元素；使用队列进行BFS遍历，并获取每一层队列中最后一个元素

- 在层遍历的基础上，如果当curLevel的个数为1时保存结果即可