
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



## or-channel
