# 学习笔记

## 本周知识点

递归，分治和回溯。分治和回溯只是递归的细分。解题思维要点：不要人肉递归，找重复子问题，数学归纳法。

递归代码模板

```go
func recursive(level, param int) {
  // terminator
  if level > maxLevel {
    // process result
    return 
  }
  // process current logic 
  process(level, param)
  
  // drill down
  recursive(level+1,newParam)
  
  // restore current status
}
```
