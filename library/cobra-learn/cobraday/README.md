# 笔记

## 参考网址

- https://www.cnblogs.com/jiujuan/p/15487918.html
- https://www.cnblogs.com/sparkdev/p/10856077.html
- https://www.jianshu.com/p/99aae91587af

## 知识点

### 为 Command 添加选项(flags)

### 命令行参数(arguments)

cobra 默认提供了一些验证方法：

- NoArgs - 如果存在任何位置参数，该命令将报错
- ArbitraryArgs - 该命令会接受任何位置参数 
- OnlyValidArgs - 如果有任何位置参数不在命令的
- ValidArgs 字段中，该命令将报错 MinimumNArgs(int) - 至少要有 N 个位置参数，否则报错
- MaximumNArgs(int) - 如果位置参数超过 N 个将报错
- ExactArgs(int) - 必须有 N 个位置参数，否则报错
- ExactValidArgs(int) 必须有 N 个位置参数，且都在命令的 ValidArgs 字段中，否则报错
- RangeArgs(min, max) - 如果位置参数的个数不在区间 min 和 max 之中，报错

### 子命令打印参数位置

`go run main.go echo times -t 3 zm`

### 在 Commnad 执行前后执行额外的操作

Command 执行的操作是通过 Command.Run 方法实现的，为了支持我们在 Run 方法执行的前后执行一些其它的操作，Command 还提供了额外的几个方法，它们的执行顺序如下：

1. PersistentPreRun
2. PreRun
3. Run
4. PostRun
5. PersistentPostRun