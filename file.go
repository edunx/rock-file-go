package file

import (
	pub "github.com/edunx/rock-public-go"
	"os"
)

func (self *File) Start() error {
	file, err := os.OpenFile(self.C.path, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		return err
	}

	self.Fd = file
	return nil
}

func (self *File) Close() {
	self.Fd.Close()
}

func (self *File) Reload() {
	self.Fd.Close()
	err := self.Start()
	if err != nil {
		pub.Out.Err("reload file fail , err: %v" , err)
		return
	}

	pub.Out.Err("reload file succeed")
}

func (self *File) Push(v interface{}) {

	str, ok := v.(string)
	if ok {
		self.Fd.WriteString(str)
		return
	}

	bytes, ok := v.([]byte)
	if ok {
		self.Fd.Write(bytes)
		return
	}

	msg, ok := v.(pub.Message)
	if ok {
		self.Fd.Write(msg.Byte())
		return
	}

	pub.Out.Err("file type error")
}
