# [287. Find the Duplicate Number](https://leetcode.com/problems/find-the-duplicate-number/)

## 题目

Given an array nums containing n + 1 integers where each integer is between 1 and n (inclusive), prove that at least one duplicate number must exist. Assume that there is only one duplicate number, find the duplicate one.

Example 1:

```c
Input: [1,3,4,2,2]
Output: 2
```

Example 2:

```c
Input: [3,1,3,4,2]
Output: 3
```

Note:  

- You must not modify the array (assume the array is read only).
- You must use only constant, O(1) extra space.
- Your runtime complexity should be less than O(n^2).
- There is only one duplicate number in the array, but it could be repeated more than once.

## 题目大意

给出 n + 1 个数，这些数是在 1-n 中取值的，同一个数字可以出现多次。要求找出这些数中重复的数字。时间复杂度最好低于 O(n^2)，空间复杂度为 O(1)。

## 解题思路

- 将数组的转化成单链表，数组的每个元素相当于当前位置连接对应的数组下标位置，这样相同的元素就为指向同一个数组下标形成环，找到环的切入点即可;

- 二分法，这里二分的是针对取值空间而不是对源数据组；在初始区间[0, n]中源数组中<=mid值的个数大于mid则说明重复的数字在[0, mid]中，否则在(mid, n]中，遍历处理只有当low==hight时则找到我们所需要的值; 其中计算mid的公式为：low + (hight-low)>>1