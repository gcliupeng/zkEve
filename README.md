# zkEve
分布式事件通知lib

以往进程间通信多使用共享内存（mmap），信号量（sem），或父子进程同步（fork，wait）。分布式环境下不易维护。

本项目使用zookeeper进行进程间消息的发送和接受，主要思路是把zookeeper的WGet 封装了一层，结构简单，不值一哂。

使用方法可见sample.go，sampleFire.go 
