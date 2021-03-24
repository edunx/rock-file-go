# rock-file-go
## 说明
文件写入类 ， 满足transport.tunnel框架

## 函数说明
### rock.file
- 函数: rock.file( table ) 
- 说明: 参数主要为table数据类型
```lua
    local ud = rock.file{
        path = "access.log" ,
        backup = "day" , -- day , hour , off 自动分割文件
        warp = "\n", --\n, \r\n , "" 自动添加换行符号
    }
    ud.debug("www.baidu.com" , "www.google.com")
```

# 安装使用 
- 跟其他包一样 直接引入后注册服务就行 

```golang
    import (
        lfile "github.com/edunx/rock-file-go"
        tp    "github.com/edunx/rock-transport-go"
    )

    //注入 lua api
    lfile.LuaInjectionApi(L , rock)
    
    //满足 transport tunnel 接口

    //获取对象 
    obj :=  tp.CheckTunnelUserData(L , idx)

    obj.Push( msg )

```