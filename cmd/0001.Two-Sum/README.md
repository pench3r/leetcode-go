# [1. Two Sum](https://leetcode.com/problems/two-sum/)

## 题目

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

```
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```



## 题目大意

在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。

## 解题思路

问题的本质是寻找2个数的映射关系，关系为 m = TARGET - n，采用map映射即可；

题目中要求返回对应的下标，因此使用map[value] = index即可，关键代码：

```
if ok, inx := map[target-n]; ok {
    return []int{i, inx}
}
map[n] = i
```