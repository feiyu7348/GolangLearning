# 笔记

## 参考网址

- https://www.cnblogs.com/jiujuan/p/15487918.html
- https://www.cnblogs.com/sparkdev/p/10856077.html
- https://www.jianshu.com/p/99aae91587af

## 知识点

### 为 Command 添加选项(flags)

### 命令行参数(arguments)

### 子命令打印参数位置

`go run main.go echo times -t 3 zm`

### 在 Commnad 执行前后执行额外的操作

Command 执行的操作是通过 Command.Run 方法实现的，为了支持我们在 Run 方法执行的前后执行一些其它的操作，Command 还提供了额外的几个方法，它们的执行顺序如下：

1. PersistentPreRun
2. PreRun
3. Run
4. PostRun
5. PersistentPostRun