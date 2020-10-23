# [739. Daily Temperatures](https://leetcode.com/problems/daily-temperatures/)

## 题目


Given a list of daily temperatures T, return a list such that, for each day in the input, tells you how many days you would have to wait until a warmer temperature. If there is no future day for which this is possible, put 0 instead.

For example, given the list of temperatures T = [73, 74, 75, 71, 69, 72, 76, 73], your output should be [1, 1, 4, 2, 1, 1, 0, 0].

Note: The length of temperatures will be in the range [1, 30000]. Each temperature will be an integer in the range [30, 100].


## 题目大意

给出一个温度数组，要求输出比当天温度高的在未来的哪一天，输出未来第几天的天数。例如比 73 度高的在未来第 1 天出现，比 75 度高的在未来第 4 天出现。

## 解题思路

维护一个栈保存待检查的温度索引，当栈非空并且当前遍历的元素大于栈顶元素时，开始从栈顶依次处理；

处理栈时，弹出栈顶元素计算距离并保存到结果集中