# 笔记

## 参考网址

- https://www.cnblogs.com/jiujuan/p/15487918.html
- https://www.cnblogs.com/sparkdev/p/10856077.html
- https://www.jianshu.com/p/99aae91587af

## 知识点

### 为 Command 添加选项(flags)
#### 持久标志
标志可以是“持久”的，这意味着该标志将可用于分配给它的命令以及该命令下的每个命令。对于全局标志，在根上指定一个标志作为持久标志。

```golang
// v flag将在rootCmd及以下的子命令上都生效
rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
```
#### 本地标志

也可以在本地分配一个标志，该标志只应用于该特定命令。
```golang
// 该命令只在localCmd上起作用
localCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")
```
### 命令行参数(arguments)

cobra 默认提供了一些验证方法：
```golang
- NoArgs - 如果存在任何位置参数，该命令将报错
- ArbitraryArgs - 该命令会接受任何位置参数 
- OnlyValidArgs - 如果有任何位置参数不在命令的
- ValidArgs 字段中，该命令将报错 MinimumNArgs(int) - 至少要有 N 个位置参数，否则报错
- MaximumNArgs(int) - 如果位置参数超过 N 个将报错
- ExactArgs(int) - 必须有 N 个位置参数，否则报错
- ExactValidArgs(int) 必须有 N 个位置参数，且都在命令的 ValidArgs 字段中，否则报错
- RangeArgs(min, max) - 如果位置参数的个数不在区间 min 和 max 之中，报错
```
### 子命令打印参数位置

`go run main.go echo times -t 3 zm`

### 在 Commnad 执行前后执行额外的操作

Command 执行的操作是通过 Command.Run 方法实现的，为了支持我们在 Run 方法执行的前后执行一些其它的操作，Command 还提供了额外的几个方法，它们的执行顺序如下：
```golang
1. PersistentPreRun
2. PreRun
3. Run
4. PostRun
5. PersistentPostRun
```