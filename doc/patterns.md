
# 第4章 Go语言的并发模式



## 1 for-select循环 

```text

for {  // 要不就无限循环,要不就使用 range语句循环
   select {  // 使用 channel进行作业
   }
}


```

有以下几种情况你可以见到这种模式.
1. 向 channel 发送迭代变量.
看起来像这样, 通常情况下，需要将可迭代的内容转换为 channel上的值
```text

for _, s : range[]string{"a","b","c"} {
    select {
    case <- done:
      return
    case stringStream <-s :
            
        }
   }


```

2。 循环等待停止
创建循环，无限循环直到停止的 go 很常见.

第一种变体， 保持 select语句尽可能短:
```text
for {
    select {
        case <-done:
          return
        default:
        
    }
    // 进行非强制式任务    
}


```

第2种变体将工作嵌入到选择语句的默认子句中:
```text

for {
    select {
        case <-done:
          return
        default:
            // 进行非强制式任务    
    }
    
}

```


## 2 防止 go泄漏

将 父子go 进行成功整合的一种方法就是在
父子 go之间建立一个 "信号通道", 让 父 go开业向子 go发出取消信号.
按照惯例， 这个信号通常是一个名叫 done 的只读 channel .
父 go将该 channel传递给 子go, 然后在想要取消子go时，关闭该 channel.


现在我们知道如何确保 go不泄漏, 我们可以规定一个约定:
如果 go负责创建 go,它也负责它可以停止 go.


这个地方有个约定:
如果 go 负责创建 go, 那么也也要负责可以停止 go.

这个 名字都是叫 "done channel"

## or-channel

需求:
有时候，你可能发现你想要将多个 done channel 合并成一个单一的 done channel
来这个复合 done channel 中的任意一个 channel 关闭的时候,关闭整个 done channel.

一关全关.

// 这种模式通过递归和 go创建一个复合的 done channel.

这个地方第一次，出现了 channel 树.


## 错误处理


Go语言中避免了流行的错误处理异常模型. (Go语言的错误处理机制，是不同的.)

如何优雅的解决.
应该做什么，不应该做什么.

思考一下:
思考错误处理时最根本的问题是: "谁应该负责处理错误?"
在某些时候，程序需要停止并将错误输出来，并且实际上对它做了些什么.

这些地方还是得分情况来看考虑:
一般来说, 你的并发进程应该把他们的错误发送到你的程序的另一部分,它有你的程序状态的完整信息，
并可以作出更明智的决定.


