# Daemon Process
```sh
$ ./a.sh

```
    运行在当前shell进程下，关闭shell或者Ctrl-c都能中断a.sh的运行

```sh
$ ./a.sh &
```
运行在后台，依旧会在前台打印，jobs 可以查询到, Ctrl-c不能中断运行，但是关闭shell可以中断运行。

```sh
$ (./a.sh &)
```
 将运行在前台的命令放在后台运行，忽略hanup信号，不在作业列表中，jobs 不能查到，只能通过ps来查询


 ```sh
$ nohup ./a.sh &
 ```
 在缺省下，此作业的所有输出都被重定向到名为nohub.out文件中, 防止shell关闭时程序停掉。

 ```sh
$ nohup ./a.sh &> myout.file &
 ```
 重定向到myout.file

    disown -h$myjob$ 让已经运行的程序忽略hangup信号
    
    jobs   -可以查看运行在后台的程序，有各种状态。增加&运行之后，jobs查看是running状态。

    Ctrl-z -将前台执行的命令放在后台，并暂停

    bg %jobnumber -将后台暂停的任务继续执行

    fg  -将后台运行的任务放在前台运行
