package file

import "github.com/edunx/lua"

func CheckBackupByTable( tab *lua.LTable , key string , d string) string {

	v := tab.RawGetString( key )
	if v.Type() != lua.LTString { return d }

	rc := v.String()

	if rc == "day" { return "day" }
	if rc == "hour" { return "hour" }
	if rc == "off" { return   "off" }

	return d
}