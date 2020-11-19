# [3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)

## 题目

Given a string, find the length of the longest substring without repeating characters.



Example 1:

```c
Input: "abcabcbb"
Output: 3 
Explanation: The answer is "abc", with the length of 3. 
```

Example 2:

```c
Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
```

Example 3:

```c
Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3. 
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
```

## 题目大意


在一个字符串重寻找没有重复字母的最长子串。

## 解题思路

快慢指针 + 数组计数：

- 使用数组记录对应字母出现次数，未出现为0；

- 右指针字母未出现，则移动右指针并更新对应计数器；出现时，则计算更新最终结果、移动左指针并更新对应计数器

- 左指针初始化为0，边界为len(s);右指针为-1，边界为right+1 < len(s);