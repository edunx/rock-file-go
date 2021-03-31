package file

import (
	"github.com/edunx/lua"
)

func createFileUserdata(L *lua.LState , args *lua.Args) lua.LValue {
	opt := args.CheckTable(L , 1)
	f := &File{
		C: Config{
			path: opt.CheckString("path", "access.log"),
			backup: CheckBackupByTable(opt , "backup" , "off"),
			warp: opt.CheckString("warp" , ""),
		},
	}

	if e := f.Start(); e != nil {
		L.RaiseError(" file start fail , err: %v", e)
		return lua.LNil
	}

	return f.ToLightUserData(L)
}


func (self *File) ToLightUserData(L *lua.LState) *lua.LightUserData {
	ud := &lua.LightUserData{ Value: self }
	return ud
}

func (self *File) debug(L *lua.LState , args *lua.Args) lua.LValue {
	n := args.Len()
	if n <= 0 {
		return lua.LNil
	}

	for i:=1;i<=n;i++{
		self.Write(args.CheckString(L , i))
	}

	return lua.LTrue
}

func (self *File) Index(L *lua.LState , key string) lua.LValue {
	if key == "debug" { return lua.NewGFunction( self.debug ) }
	return lua.LNil
}

func LuaInjectApi(L *lua.LState, parent *lua.UserKV) {
	parent.Set("file" , lua.NewGFunction( createFileUserdata ))
}
