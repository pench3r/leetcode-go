# [114. Flatten Binary Tree to Linked List](https://leetcode.com/problems/flatten-binary-tree-to-linked-list/)


## 题目

Given a binary tree, flatten it to a linked list in-place.

For example, given the following tree:

    		1
       / \
      2   5
     / \   \
    3   4   6

The flattened tree should look like:

    1
     \
      2
       \
        3
         \
          4
           \
            5
             \
              6

## 题目大意

给定一个二叉树，原地将它展开为链表。

## 解题思路

- 非递归的解法，预排序遍历算法(中序遍历)即可满足题目的要求，遍历后将值保存到数组中，输出结果前转化为TreeNode类型排序

- 递归解法，关键思路：将右子树挂到左子树下方，然后将左子树挂到右边作为新的右子树