## 题目

> 给定一个字符串`str`和长度`leftsize`, 请把`str`左侧`leftszie`的部分和右部分做整体交换,要求空间复杂度$O(1)$

## 解

### 方式一:

#### 解题思路:

1. 左边逆序
2. 右边逆序
3. 整体逆序

#### 代码:

1. [Golang](./strReverse1.go)

### 方式二:

#### 解题思路:
```

abc defghljk      [0:3][3:11]

ljk defgh abc     [0:3][3:8]

fgh de ljk abc    [0:3][3:5]

de h fg ljk abc   [2:3][3:5]

de g f h ljk abc   [2:3][3:4]

de f g h ljk abc

```

1. 短边换与长边的另一头相同长度的内容进行交换
2. 短边交换过的边就不动了
3. 剩下的内容进行相同操作

#### 代码:

1. [Golang](./strReverse2.go)
