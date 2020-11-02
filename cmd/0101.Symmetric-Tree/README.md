# [101. Symmetric Tree](https://leetcode.com/problems/symmetric-tree/)

## 题目


Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:


```c
    1
   / \
  2   2
 / \ / \
3  4 4  3
```

But the following [1,2,2,null,3,null,3] is not:

```c
    1
   / \
  2   2
   \   \
   3    3
```

Note:   

Bonus points if you could solve it both recursively and iteratively.

## 题目大意

这一题要求判断 2 颗树是否是左右对称的。


## 解题思路

问题可以转化为左子树是否可以翻转为右子树；翻转的判断逻辑基本与sameTree的递归解法相同；

关键的比对从p.left 和 q.left 变成了 p.left 和 q.right

