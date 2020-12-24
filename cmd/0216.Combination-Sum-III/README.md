# [216. Combination Sum III](https://leetcode.com/problems/combination-sum-iii/)


## 题目

Find all possible combinations of **k** numbers that add up to a number **n**, given that only numbers from 1 to 9 can be used and each combination should be a unique set of numbers.

**Note:**

- All numbers will be positive integers.
- The solution set must not contain duplicate combinations.

**Example 1:**

    Input: k = 3, n = 7
    Output: [[1,2,4]]

**Example 2:**

    Input: k = 3, n = 9
    Output: [[1,2,6], [1,3,5], [2,3,4]]

## 题目大意

找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。

说明：

- 所有数字都是正整数。
- 解集不能包含重复的组合。


## 解题思路

- 使用递归的方式实现回溯；

- 回溯的关键要进行类似：压栈、递归处理、出栈的过程

- dfs函数参数：inx, k，n，c，res，其中inx代表起始遍历的值，因为每种组合中不存在重复的数字；c保存单次数组结果；res保存最终结果；当k==0 && n==0的时候保存c至res中

- dfs函数实现时，将边界判定放置在主循环的外面；