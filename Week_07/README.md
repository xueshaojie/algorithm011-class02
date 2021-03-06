# 学习笔记

## Trie树

字典树，是一种树形结构。某个节点上存放的是单词的某个字母。从根节点出发到某一节点，判断是否存在某个单词或者是某个单词的前缀。

核心思想是：用空间换时间。占用的空间变大，查询所使用的时间变小。

## 并查集

1. 建立一个新的并查集。初始化数组。每个元素的父元素指向自己。
2. 合并两个集合。
3. 找到某个元素所在集合的代表。

## 双向BFS

左右两边同时进行扩散。核心思想在于：两个队列，那个队列的元素少就用那个队列进行扩散。

## A*算法

代码具有简单的思考能力。用来预测那个选择更优，最终导致使用更少的时间解决问题。

解题方法：一般使用优先队列。优先级通过两个点的距离来确定。
