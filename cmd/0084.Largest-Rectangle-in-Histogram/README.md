# [84. Largest Rectangle in Histogram](https://leetcode.com/problems/largest-rectangle-in-histogram/)

## 题目

Given n non-negative integers representing the histogram's bar height where the width of each bar is 1, find the area of largest rectangle in the histogram.

 ![](https://assets.leetcode.com/uploads/2018/10/12/histogram.png)


Above is a histogram where width of each bar is 1, given height = [2,1,5,6,2,3].

![](https://assets.leetcode.com/uploads/2018/10/12/histogram_area.png)


The largest rectangle is shown in the shaded area, which has area = 10 unit.

 

Example:

```c
Input: [2,1,5,6,2,3]
Output: 10
```


## 题目大意

给出每个直方图的高度，要求在这些直方图之中找到面积最大的矩形，输出矩形的面积。


## 解题思路

维护一个单调递增栈，当扫描到的节点小于栈顶元素时，再从栈顶回溯计算更新矩形区域大小。

关键点：

- 矩形区域的高度是当前节点和栈顶元素的最小值
- 从栈顶开始，依次弹出大于当前节点的元素，再将当前节点进行入栈
- 单调栈将每两个元素之间的值进行了抽离，因为矩形的高度是有两边决定，因此单调栈的特性减少了遍历的次数，提升了计算性能