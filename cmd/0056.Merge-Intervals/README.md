# [56. Merge Intervals](https://leetcode.com/problems/merge-intervals/)

## 题目

Given a collection of intervals, merge all overlapping intervals.

Example 1:  

```
Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
```

Example 2:  

```
Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
```

## 题目大意

合并给的多个区间，区间有重叠的要进行区间合并。


## 解题思路

先排序，后遍历

两个区间有多种组合情况，但可以分成2种：重叠和非重叠，这样处理整个逻辑就很清晰；

由于是排序后处理，所以只要两个区间重叠了，那只需要合并这两个区间，并更新end为这两个区间的最大值即可

遍历区间，使用临时变量前一个区间，判断当前区间和前一个是否重叠，如果重叠则合并更新；否则前一个区间保存到结果中，并将当前区间保存到临时变量中