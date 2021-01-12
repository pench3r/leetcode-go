# [53. Maximum Subarray](https://leetcode.com/problems/maximum-subarray/)


## 题目

Given an integer array `nums`, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

**Example:**


    Input: [-2,1,-3,4,-1,2,1,-5,4],
    Output: 6
    Explanation: [4,-1,2,1] has the largest sum = 6.


**Follow up:**

If you have figured out the O(*n*) solution, try coding another solution using the divide and conquer approach, which is more subtle.

## 题目大意

给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

## 解题思路

- 使用动态规划进行解题，dp[i]保存到在[0, i]区间中最大的和，状态转移公式核心原则：当前面区间的最大和小于0，则不再做累加；

```
当dp[i-1] > 0时，dp[i] = dp[i-1] + nums[i]
当dp[i-1] <= 0, dp[i] = nums[i]
```

- 边界处理：当只有一个元素时直接返回即可