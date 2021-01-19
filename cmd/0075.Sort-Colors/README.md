# [75. Sort Colors](https://leetcode.com/problems/sort-colors/)

## 题目

Given an array with n objects colored red, white or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white and blue.

Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

Note: You are not suppose to use the library's sort function for this problem.

Example 1:  

```
Input: [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
```

Follow up:

- A rather straight forward solution is a two-pass algorithm using counting sort.  
First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array with total number of 0's, then 1's and followed by 2's.
- Could you come up with a one-pass algorithm using only constant space?


## 题目大意

抽象题意其实就是排序。这题可以用快排一次通过。

## 解题思路

- 题目中只会出现3个数字:0, 1, 2，所以使用3个游标进行控制。具体的做法是：当碰到0时，按照2，1，0的顺序给对应的游标进行赋值并移动；当碰到1时，则按照2，1进行处理；碰到2时则只移动2的游标即可