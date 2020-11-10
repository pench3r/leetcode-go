# [99. Recover Binary Search Tree](https://leetcode.com/problems/recover-binary-search-tree/)


## 题目

Two elements of a binary search tree (BST) are swapped by mistake.

Recover the tree without changing its structure.

**Example 1:**

    Input: [1,3,null,null,2]
    
       1
      /
     3
      \
       2
    
    Output: [3,1,null,null,2]
    
       3
      /
     1
      \
       2

**Example 2:**

    Input: [3,1,4,null,null,2]
    
      3
     / \
    1   4
       /
      2
    
    Output: [2,1,4,null,null,3]
    
      2
     / \
    1   4
       /
      3

**Follow up:**

- A solution using O(*n*) space is pretty straight forward.
- Could you devise a constant space solution?

## 题目大意

二叉搜索树中的两个节点被错误地交换。请在不改变其结构的情况下，恢复这棵树。


## 解题思路

这里要利用二叉搜索树的中序遍历，中序遍历后的结果是一个递增序列，如果存在两个节点错误交换，那么在中序遍历的结果中存在2个错误点无法满足递增

假如我们存在一个中序遍历后的序列：1，2，3，4，5，6；我们调换5和2的位置: 1，5，3，4，2，6；这样题目就转化为在一个数组中如何找到2个错误点

使用4个变量：cur、prev、target1、target2；当prev>cur时，则表明存在异常点；target1保存第一个异常点，target2保存第二个异常点；所以当prev>cur满足第一次时target1=prev,满足第二次时target2=root

利用上述思路，只需要在中序遍历获取值的地方进行处理，即可获取到2个异常点；最后swap这两个点的值即可

题目要求使用O(1)的空间复杂度，使用中序遍历时我们还是使用了栈空间因此不满足题目的要求；这里就需要引入Morris Traversal算法，进行中序遍历即可满足要求