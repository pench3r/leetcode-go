# [129. Sum Root to Leaf Numbers](https://leetcode.com/problems/sum-root-to-leaf-numbers/)


## 题目

Given a binary tree containing digits from `0-9` only, each root-to-leaf path could represent a number.

An example is the root-to-leaf path `1->2->3` which represents the number `123`.

Find the total sum of all root-to-leaf numbers.

**Note:** A leaf is a node with no children.

**Example:**

    Input: [1,2,3]
        1
       / \
      2   3
    Output: 25
    Explanation:
    The root-to-leaf path 1->2 represents the number 12.
    The root-to-leaf path 1->3 represents the number 13.
    Therefore, sum = 12 + 13 = 25.

**Example 2:**

    Input: [4,9,0,5,1]
        4
       / \
      9   0
     / \
    5   1
    Output: 1026
    Explanation:
    The root-to-leaf path 4->9->5 represents the number 495.
    The root-to-leaf path 4->9->1 represents the number 491.
    The root-to-leaf path 4->0 represents the number 40.
    Therefore, sum = 495 + 491 + 40 = 1026.

## 题目大意

给定一个二叉树，它的每个结点都存放一个 0-9 的数字，每条从根到叶子节点的路径都代表一个数字。例如，从根到叶子节点路径 1->2->3 代表数字 123。计算从根到叶子节点生成的所有数字之和。说明: 叶子节点是指没有子节点的节点。


## 解题思路

- 大致思路与0113相同，都是使用dfs遍历，当遍历到了叶子节点时保存结果

- 不同的地方在于：将节点的值转化为str类型，并以slice形式返回，最后拼接转化为整型输出结果

```
dfs(root) []string {
    var res []string
    // 左子树
    tmpLeft := dfs(root.Left)
    // 当前root节点的值与左子树返回的值拼接
    for i:=0; i<len(tmpLeft); i++ {
        res = append(res, strcon.Itoa(root.Val) + tmpLeft[i])
    }
    // 右子树
    ...
    return res
}
```