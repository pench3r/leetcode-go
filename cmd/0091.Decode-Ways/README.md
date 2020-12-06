# [91. Decode Ways](https://leetcode.com/problems/decode-ways/)


## 题目

A message containing letters from `A-Z` is being encoded to numbers using the following mapping:

    'A' -> 1
    'B' -> 2
    ...
    'Z' -> 26

Given a **non-empty** string containing only digits, determine the total number of ways to decode it.

**Example 1:**

    Input: "12"
    Output: 2
    Explanation: It could be decoded as "AB" (1 2) or "L" (12).

**Example 2:**

    Input: "226"
    Output: 3
    Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).

## 题目大意

一条包含字母 A-Z 的消息通过以下方式进行了编码：

```c
'A' -> 1
'B' -> 2
...
'Z' -> 26
```

给定一个只包含数字的非空字符串，请计算解码方法的总数。



## 解题思路

- 该题目可以使用动态规划的进行解题，动态矩阵dp[i]表示i长度的字符串的解码总数，状态转移公式分为两种情况：
  ```
  当0<s[i-1:i]<10时，dp[i] += dp[i-1]
  当10<=s[i-2:i]<=26, dp[i] += dp[i-2]
  ```
- dp[0]为1，当第一个字符为0时,dp[1]=0，否则dp[1]=1