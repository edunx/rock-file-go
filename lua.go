package file

import (
	"github.com/edunx/lua"
)

const (
	MT string = "ROCK_FILE_GO_MT"
)

func LuaInjectApi(L *lua.LState, parent *lua.LTable) {
	mt := L.NewTypeMetatable(MT)
	L.SetField(mt, "__index", L.NewFunction(get))
	L.SetField(mt, "__newindex", L.NewFunction(set))

	L.SetField(parent, "file", L.NewFunction(CreateFileUserdata))
}

func CreateFileUserdata(L *lua.LState) int {
	opt := L.CheckTable(1)

	f := &File{
		C: Config{
			path: opt.CheckString("path", "access.log"),
			backup: CheckBackupByTable(opt , "backup" , "off"),
		},
	}

	if e := f.Start(); e != nil {
		L.RaiseError(" file start fail , err: %v", e)
		return 0
	}

	ud := L.NewUserDataByInterface(f , MT)

	L.Push(ud)
	return 1
}

func get(L *lua.LState) int {
	return 0
}

func set(L *lua.LState) int {
	return 0
}

func (self *File) ToUserData(L *lua.LState) *lua.LUserData {
	return L.NewUserDataByInterface(self, MT)
}
