# clougo-data
--------------------------------------------
## 任务内容
--------------------------------------------
1. 使用 xorm 或 gorm 实现本文的程序，从编程效率、程序结构、服务性能等角度对比 database/sql 与 orm 实现的异同！   

   orm 是否就是实现了 dao 的自动化？

   使用 ab 测试性能
2. 参考 Java JdbcTemplate 的设计思想，设计 GoSqlTemplate 的原型, 使得 sql 操作对于爱写 sql 的程序猿操作数据库更容易。 

   轻量级别的扩展，程序员的最爱

   程序猿不怕写 sql ，怕的是线程安全处理和错误处理

   sql 的 CRUD 操作 database/sql 具有强烈的模板特征，适当的回调可以让程序员自己编写 sql 语句和处理 RowMapping
 
   建立在本文 SQLExecer 接口之上做包装，直观上是有利的选择

   暂时不用考虑占位符等数据库移植问题，方便使用 mysql 或 sqlite3 就可以
 
   参考资源：github.com/jmoiron/sqlx
   
## 测试结果
---------------------------------------------
ab压力测试
```scripts
$ ab -n 1000 -c 100 http://localhost:8080/service/userinfo?userid=1
```
```scripts
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /service/userinfo?userid=1
Document Length:        95 bytes

Concurrency Level:      100
Time taken for tests:   0.574 seconds
Complete requests:      1000
Failed requests:        0
Non-2xx responses:      1000
Total transferred:      227000 bytes
HTML transferred:       95000 bytes
Requests per second:    1732.66 [#/sec] (mean)
Time per request:       57.417 [ms] (mean)
Time per request:       0.574 [ms] (mean, across all concurrent requests)
Transfer rate:          386.09 [Kbytes/sec] received

Connection Times (ms)
			         min  mean[+/-sd] median   max
Connect:        0    2   3.5      1      21
Processing:     1   54  60.5     39     277
Waiting:        0   53  60.3     37     272
Total:          1   56  60.7     41     279

Percentage of the requests served within a certain time (ms)
  50%     41
  66%     50
  75%     58
  80%     65
  90%    103
  95%    218
  98%    256
  99%    272
 100%    279 (longest request)
 ```
 ## 评价
 ----------------------------------------
 从结果看，orm比database/sql要好，其提供了齐全的api，几乎可以实现数据库的全部操作，且orm更符合“entity - dao - service” 层次结构编程模型。   
 但是database/sql 和 orm 在处理请求的 web 服务上并没有对性能造成太大的影响，毕竟请求主要的时间是耗在 IO，而不是 go 语言的执行上。
 ### ps:
 1. 这个程序只实现了第一部分的内容（用xorm）
 2. 需要将userinfo表单的created字段更名为createat

 
