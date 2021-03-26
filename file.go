package file

import (
	"fmt"
	"github.com/edunx/lua"
	pub "github.com/edunx/rock-public-go"
	tp "github.com/edunx/rock-transport-go"
	"os"
	"time"
)

func (self *File) filename( now time.Time ) string {

	if self.C.backup == "off" { return self.C.path }
	if self.C.backup == "day" {
		name := fmt.Sprintf("%s.%d-%d-%d" , self.C.path , now.Year() , now.Month() , now.Day())
		if name == self.name {
			return fmt.Sprintf("%s.%d-%d-%d" , self.C.path , now.Year() , now.Month() , now.Day() + 1)
		}
		return name
	}

	if self.C.backup == "hour" {
		name := fmt.Sprintf("%s.%d-%d-%d.%d" , self.C.path , now.Year() , now.Month() , now.Day() , now.Hour())
		if name == self.name {
			return fmt.Sprintf("%s.%d-%d-%d.%d" , self.C.path , now.Year() , now.Month() , now.Day() , now.Hour() + 1)
		}
		return name
	}

	return self.C.path
}

func (self *File) backup( now time.Time ) {
	pub.Out.Err("start .. backup , time: %v" , now)

	filename := self.filename( now )
	file, err := os.OpenFile(filename , os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		pub.Out.Err("file backup fail , err: %v" , err)
		return
	}

	old := self.Fd
	self.Fd = file
	old.Close()
	pub.Out.Err("backup succeed , time: %v" , now)
}

func (self *File) timer() {
	if self.C.backup == "off" { return }
	var sleep *time.Ticker
	var tk *time.Ticker

	if self.C.backup == "day" {
		sleep = time.NewTicker(time.Duration(DAY - (time.Now().Unix() - TIMESTAMP )  % DAY) * time.Second)
		tk = time.NewTicker(time.Second * DAY)
		goto RUN
	}

	if self.C.backup == "hour" {
		sleep = time.NewTicker(time.Duration(HOUR - (time.Now().Unix() - TIMESTAMP )  % HOUR) * time.Second)
		tk = time.NewTicker(time.Second * HOUR)
		goto RUN
	}

	return

RUN:
	now := <-sleep.C
	self.backup( now )

	for now = range tk.C {
		self.backup( now )
	}
}

func (self *File) Start() error {
	filename := self.filename( time.Now() )

	file, err := os.OpenFile(filename , os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		return err
	}

	self.Fd = file
	go self.timer()

	return nil
}

func (self *File) Close() {
	self.Fd.Close()
}

func (self *File) Write( v interface{} ) error {
	var str string
	var bytes []byte
	var msg lua.Message

	str, ok := v.(string)
	if ok {
		self.Fd.WriteString(str)
		goto DONE
	}

	bytes, ok = v.([]byte)
	if ok {
		self.Fd.Write(bytes)
		goto DONE
	}

	msg, ok = v.(tp.Message)
	if ok {
		self.Fd.Write(msg.Byte())
		goto DONE
	}

	pub.Out.Err("file type error")

DONE:
	if self.C.warp == "\r\n" || self.C.warp == "\n" || self.C.warp == "\r" {
		self.Fd.WriteString(self.C.warp)
	}

	return nil
}

func (self *File) Type() string {
	return "file"
}

func (self *File) Name() string {
	return self.name
}

func (self *File) Proxy(info string , v interface{}) {
}