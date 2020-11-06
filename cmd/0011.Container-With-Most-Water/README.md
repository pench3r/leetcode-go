# [11. Container With Most Water](https://leetcode.com/problems/container-with-most-water/)

## 题目

Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.

![](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/17/question_11.jpg)

The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

Example 1:

```c
Input: [1,8,6,2,5,4,8,3,7]
Output: 49
```


## 题目大意

给出一个非负整数数组 a1，a2，a3，…… an，每个整数标识一个竖立在坐标轴 x 位置的一堵高度为 ai 的墙，选择两堵墙，和 x 轴构成的容器可以容纳最多的水。

## 解题思路

根据题意可知，两端的长度决定容器的大小，使用双指针从左和右进行扫描，计算更新当前容器的大小，当左边的值大于右边的值，移动右指针，反之亦然

直到左右两个指针相遇，则退出循环；

和两层循环的暴力破解思路相比，省去了大量的无效比对

