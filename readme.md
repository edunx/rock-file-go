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