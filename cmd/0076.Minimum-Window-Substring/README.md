# [76. Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/)

## 题目

Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).

Example:

```c
Input: S = "ADOBECODEBANC", T = "ABC"
Output: "BANC"
```

Note:    

- If there is no such window in S that covers all characters in T, return the empty string "".
- If there is such window, you are guaranteed that there will always be only one unique minimum window in S.

## 题目大意

给定一个源字符串 s，再给一个字符串 T，要求在源字符串中找到一个窗口，这个窗口包含由字符串各种排列组合组成的，窗口中可以包含 T 中没有的字符，如果存在多个，在结果中输出最小的窗口，如果找不到这样的窗口，输出空字符串。

## 解题思路

快慢指针 + 数组计数 + 长度计数：

- 使用数组保存T中所有字符出现的次数

- 移动右指针, 判断当前字符是否存在于T中，存在则更新长度计数，直到计数变量的长度等于T的长度，并以此为循环更新结果、移动左指针

- 字符在数组中的计数大于0，则表明该字符存在于T中。

- 以长度计数等于T的长度为条件，不断循环处理移动左指针，更新结果、碰到处理的字符存在于T中则更新长度计数