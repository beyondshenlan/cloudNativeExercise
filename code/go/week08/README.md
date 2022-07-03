### 作业1：redis benchmark 工具,value 大小，redis get set 性能。

```
redis-benchmark -d 10 -t get,set
```

### SET
|-| 执行次数和耗时                                   | 
|----|-------------------------------------------|
|10| 100000 requests completed in 2.79 seconds | 
|20| 100000 requests completed in 2.78 seconds | 
|50| 100000 requests completed in 2.91 seconds | 
|100| 100000 requests completed in 2.98 seconds | 
|200| 100000 requests completed in 2.80 seconds | 
|1k| 100000 requests completed in 2.99 seconds |
|5k| 100000 requests completed in 2.94 seconds | 

### GET
|-| 执行次数和耗时                                   |
|----|-------------------------------------------|
|10| 100000 requests completed in 2.74 seconds |
|20| 100000 requests completed in 2.98 seconds |
|50| 100000 requests completed in 2.87 seconds |
|100| 100000 requests completed in 2.89 seconds |
|200| 100000 requests completed in 2.88 seconds |
|1k| 100000 requests completed in 2.00 seconds |
|5k| 100000 requests completed in 2.79 seconds |

### 作业2 统计平均key 的内存占用
| key value的大小 | key的平均大小 |
|--------------|----------|
| 10 字节        | 90       |
| 20 字节        | 98       |
| 50 字节        | 129      |
| 100 字节       | 187      |
| 200 字节       | 289      |
| 1K           | 1308     |
| 5K           | 5999     |
