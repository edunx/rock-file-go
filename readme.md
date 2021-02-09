---
   磐石框架文件输入框架
---
# 配置
```lua
    local ud = rock.file{
        path = "access.log" ,
        backup = "day" , -- day , hour , off 自动分割文件
        autowrap = "\n", --\n, \r\n , "" 自动添加换行符号
    }
```

# 调用
```golang
    //满足 pub.Transport 接口
    transport := pub.CheckTransportByTable("transport" , opt)

    // 写入
    transport.Push( msg )

```