# [456. 132 Pattern](https://leetcode.com/problems/132-pattern/)


## 题目

Given a sequence of n integers a1, a2, ..., an, a 132 pattern is a subsequence a**i**, a**j**, a**k** such that **i** < **j** < **k** and a**i** < a**k** < a**j**. Design an algorithm that takes a list of n numbers as input and checks whether there is a 132 pattern in the list.

**Note:** n will be less than 15,000.

**Example 1:**

    Input: [1, 2, 3, 4]
    
    Output: False
    
    Explanation: There is no 132 pattern in the sequence.

**Example 2:**

    Input: [3, 1, 4, 2]
    
    Output: True
    
    Explanation: There is a 132 pattern in the sequence: [1, 4, 2].

**Example 3:**

    Input: [-1, 3, 2, 0]
    
    Output: True
    
    Explanation: There are three 132 patterns in the sequence: [-1, 3, 2], [-1, 3, 0] and [-1, 2, 0].


## 题目大意


给定一个整数序列：a1, a2, ..., an，一个 132 模式的子序列 ai, aj, ak 被定义为：当 i < j < k 时，ai < ak < aj。设计一个算法，当给定有 n 个数字的序列时，验证这个序列中是否含有 132 模式的子序列。注意：n 的值小于 15000。


## 解题思路

- 线性扫描 + 栈；使用栈保存扫描过的值，当前值与栈中的值比对获取a3的值，将当前值作为a2进行压栈，对后续扫描的值如果小于a3则直接返回结果

- 扫描是从尾部开始扫描；当前值与栈中的值比对时，依次弹出小于当前值的栈顶元素作为a3，并将当前值作为a2进行入栈；

- 定位到a2、a3后，后续的每次扫描到的值都作为a1与a3进行比较，满足条件就可以返回true