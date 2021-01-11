# [71. Simplify Path](https://leetcode.com/problems/simplify-path/)

## 题目

Given an absolute path for a file (Unix-style), simplify it. Or in other words, convert it to the canonical path.

In a UNIX-style file system, a period . refers to the current directory. Furthermore, a double period .. moves the directory up a level. For more information, see: Absolute path vs relative path in Linux/Unix

Note that the returned canonical path must always begin with a slash /, and there must be only a single slash / between two directory names. The last directory name (if it exists) must not end with a trailing /. Also, the canonical path must be the shortest string representing the absolute path.



Example 1:

```c
Input: "/home/"
Output: "/home"
Explanation: Note that there is no trailing slash after the last directory name.
```

Example 2:

```c
Input: "/../"
Output: "/"
Explanation: Going one level up from the root directory is a no-op, as the root level is the highest level you can go.
```
Example 3:

```c
Input: "/home//foo/"
Output: "/home/foo"
Explanation: In the canonical path, multiple consecutive slashes are replaced by a single one.
```

Example 4:

```c
Input: "/a/./b/../../c/"
Output: "/c"
```


Example 5:

```c
Input: "/a/../../b/../c//.//"
Output: "/c"
```

Example 6:

```c
Input: "/a//b////c/d//././/.."
Output: "/a/b/c"
```

## 题目大意

给出一个 Unix 的文件路径，要求简化这个路径。这道题也是考察栈的题目。

## 解题思路

- 遍历解码后的路径列表，中间使用stack来保存不同边界的处理结果，最后再还原成路径格式

- 碰到`.`时，直接进行下次遍历；碰到`..`时，当stack非空，stac=stack[:len(stack)-1],否则直接进行下次遍历

- 返回结果前，如果stack为空，则直接返回`/`