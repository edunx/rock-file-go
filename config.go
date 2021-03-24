package file

import (
	"github.com/edunx/lua"
	"os"
)

const (
	TIMESTAMP = 1609430400
	DAY = 86400
	HOUR = 3600
)

type Config struct {
	path   string
	backup string
	warp   string
}

type File struct {
	lua.Super

	C    Config
	Fd   *os.File
	name string
}
