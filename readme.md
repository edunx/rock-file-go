---
   磐石框架文件输入框架
---
# 配置
```lua
    local ud = rock.file{
        path = "access.log" 
    }
```

# 调用
```golang
    //满足 pub.Transport 接口
    transport := pub.CheckTransportByTable("transport" , opt)

    // 写入
    transport.Push( msg )

```