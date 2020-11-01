# [100. Same Tree](https://leetcode.com/problems/same-tree/)

## 题目


Given two binary trees, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical and the nodes have the same value.

Example 1:


```c
Input:     1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]

Output: true
```

Example 2:

```c
Input:     1         1
          /           \
         2             2

        [1,2],     [1,null,2]

Output: false
```

Example 3:

```c
Input:     1         1
          / \       / \
         2   1     1   2

        [1,2,1],   [1,1,2]

Output: false
```

## 题目大意

这一题要求判断 2 颗树是否是完全相等的。


## 解题思路

使用递归算法，每次判断当前两个节点的值是否相同，接着递归处理左子树节点和右子树节点

特殊处理考虑当前两个节点为nil的情况

