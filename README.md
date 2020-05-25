# go_concurrency

Go语言中的 goroutine 到底是什么?
它们不是线程，也不是绿色线程(由语言运行时管理的线程),
它们是一个更高级别的抽象，称为协程.
协程是一种非抢占式的简单并发子(函数，闭包，方法),也就是说，它们不能被中断。
**这个地方没有理解**
取而代之的是，协程有多个 point,允许暂停或重新进入.

到底怎么运行那？
goroutine 的独特之处在于它们与Go语言的运行时的深度集成.
goroutine 没有定义自己的暂停方法或者再运行点.
Go语言的运行时会观察 goroutine 的运行时行为，并在它们阻塞时自动挂起它们，
然后在它们不被阻塞时恢复它们.




## 逻辑模型 fork-join的并发模式

join point的概念.


### 细节问题
讨论一个关于 goroutine 有趣的事情: GC并没有回收被丢弃的 goroutine.

这里有个例子:
我们将利用 goroutine 不被GC的事实与运行时的自省能力结合起来，并测算
goroutine 的大小.



所有，这个地方我们的代码要注意: 防止 goroutine 泄漏.



## 关于 go 的两个问题测试
1. go 所占的内存非常小，而且是动态分配的, 非常轻量级.
2. go切换上下文的代价和时间都非常的小.



# sync 包

WaitGroup
等待一组并发操作完成.

## 互斥锁和读写锁
这个地方是传统的 内存访问同步问题.
锁这个地方的主要作用是 临界区.

Mutex

建议采用 读写锁，而不是 互斥锁.


## cond
这个地方是一个什么东西?

**带条件的锁**

这个地方 跟条件有关系.

引出
```text

如果有一种方法可以让 goroutine 有效地等待，直到它发出信号并检查它的状态，那就更好了.

这正是 Cond类型为我们所做的.


```

代码模板
```text

c := sync.NewCond(&sync.Mutex{})
c.L.Lock()


// 当某个条件成立的时候,执行 等待
for conditionTrue == false {
    c.Wait() // 等待通知，这是一个阻塞通信， goroutine 将被暂停.
}

c.L.Unlock()


```

调用 c.Wait() ， 不只是阻塞，它挂起来了当前的 goroutine, 但是运行其他的 goroutine 在 OS线程上运行.

**带条件的锁**
当然也有锁的 "临界区."


1. 是基本使用.
2. Signal() 和 Broadcast() 特别是 Broadcast 的使用，(因为有时候 它的使用场景是比 channel还要好的.)


## once 

看看标准库中 有多少使用这个关键字的.

```shell
penglonghua@plh:~/temp$ grep -ir sync.Once $(go env GOROOT)/src | wc -l
112
```

注意使用.

## 池 (pool)


这个地方可以先想想一下:

Pool的主接口是它的 Get方法.
当调用时，Get方法首先检查池中是否有可用的实例, 如果没有，调用它的 new方法来创建一个新实例。
当完成时，调用者调用 Put方法把工作的实例归还到池中,以供其他进程使用. (如果不调用 Put就不会放入!!!)




