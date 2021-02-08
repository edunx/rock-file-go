package file

import "os"

type Config struct {
	path string
}

type File struct {
	C  Config
	Fd *os.File
}
