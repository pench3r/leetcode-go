# [42. Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/)

## 题目

Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.

![](https://assets.leetcode.com/uploads/2018/10/22/rainwatertrap.png)


The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped. Thanks Marcos for contributing this image!


Example:

```c
Input: [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
```

## 题目大意

从 x 轴开始，给出一个数组，数组里面的数字代表从 (0,0) 点开始，宽度为 1 个单位，高度为数组元素的值。如果下雨了，问这样一个容器能装多少单位的水？

## 解题思路

对于数组中的每个元素来说，如果存在积水的话，那么左右两边的高度都比当前元素高，同时积水量是由左右两边的最小值来计算决定；使用左右双指针的方式进行扫描遍历，假设左指针元素小于右指针元素，则`当前左指针的积水量=左边最大值-当前左指针的值`,此时对于左指针元素而言，右边存在大于左边最高的值，因为左指针已经移动经过了左边最高值，而移动左指针的条件则是有指针值高的时候。

从上述推论出发，维护左右双指针、左边出现的最高高度和右边出现的最高高度，当左右指针相遇则终止循环；

当左端高度小于右端高度时，则不断移动左指针，反之亦然；在这个过程中，当前高度小于当前方向(左或右)出现的最高高度时，计算差值并累加到最终结果，否则更新当前方向出现的最高高度