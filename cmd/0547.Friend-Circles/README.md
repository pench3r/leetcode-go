# [547. Friend Circles](https://leetcode.com/problems/friend-circles/)

## 题目

There are **N** students in a class. Some of them are friends, while some are not. Their friendship is transitive in nature. For example, if A is a **direct** friend of B, and B is a **direct**friend of C, then A is an **indirect** friend of C. And we defined a friend circle is a group of students who are direct or indirect friends.

Given a **N*N** matrix **M** representing the friend relationship between students in the class. If M[i][j] = 1, then the ith and jth students are **direct** friends with each other, otherwise not. And you have to output the total number of friend circles among all the students.

**Example 1:**

    Input: 
    [[1,1,0],
     [1,1,0],
     [0,0,1]]
    Output: 2
    Explanation:The 0th and 1st students are direct friends, so they are in a friend circle. 
    The 2nd student himself is in a friend circle. So return 2.

**Example 2:**

    Input: 
    [[1,1,0],
     [1,1,1],
     [0,1,1]]
    Output: 1
    Explanation:The 0th and 1st students are direct friends, the 1st and 2nd students are direct friends, 
    so the 0th and 2nd students are indirect friends. All of them are in the same friend circle, so return 1.

**Note:**

1. N is in range [1,200].
2. M[i][i] = 1 for all students.
3. If M[i][j] = 1, then M[j][i] = 1.


## 题目大意

班上有 N 名学生。其中有些人是朋友，有些则不是。他们的友谊具有是传递性。如果已知 A 是 B 的朋友，B 是 C 的朋友，那么我们可以认为 A 也是 C 的朋友。所谓的朋友圈，是指所有朋友的集合。

给定一个 N * N 的矩阵 M，表示班级中学生之间的朋友关系。如果 M[i][j] = 1，表示已知第 i 个和 j 个学生互为朋友关系，否则为不知道。你必须输出所有学生中的已知的朋友圈总数。


注意：

- N 在[1,200]的范围内。
- 对于所有学生，有M[i][i] = 1。
- 如果有 M[i][j] = 1，则有 M[j][i] = 1。


## 解题思路

该问题属于连通性问题，可以使用并查集作为解题思路；

使用学生的总人数作为UnionFind的初始化参数，在注意事项中可以看出矩阵的对角线为1，并且是对角线内容对称的，因此在遍历矩阵时可以使用

```
for i := 0; i < size; i++ {
    for j := 0; j <= i; j++ {
        ...
    }
}
```
当检测到矩阵对应的节点为1，标明这2个人认识，则进行Union连通操作，遍历结束后uf.count的数量就代表有多少个root，也就是多少个朋友圈