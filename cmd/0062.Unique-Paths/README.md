# [62. Unique Paths](https://leetcode.com/problems/unique-paths/)


## 题目

A robot is located at the top-left corner of a *m* x *n* grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

How many possible unique paths are there?

![](https://assets.leetcode.com/uploads/2018/10/22/robot_maze.png)

Above is a 7 x 3 grid. How many possible unique paths are there?

**Note:** *m* and *n* will be at most 100.

**Example 1:**

    Input: m = 3, n = 2
    Output: 3
    Explanation:
    From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
    1. Right -> Right -> Down
    2. Right -> Down -> Right
    3. Down -> Right -> Right

**Example 2:**

    Input: m = 7, n = 3
    Output: 28


## 题目大意

一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。问总共有多少条不同的路径？


## 解题思路

- 使用DP算法，状态矩阵dp[i][j]代表从起点出发到该位置的唯一路径数，由于每次移动只能向右或向下，对于dp[i][j]来说只能通过上方的点dp[i-1][j]和左边的点[i][j-1]移动才能到达，因此得出状态转移公式：dp[i][j] = dp[i-1][j] + dp[i][j-1]

- 第一行上的状态dp[0][j]只能通过不断右移才能到达，唯一路径数只能为1，所以dp[0][j]需要初始化为1，第一列同理dp[i][0]