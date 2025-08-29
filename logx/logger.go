package logx

import (
	kitlog "github.com/theplant/appkit/log"
)

func CreateLogger() *kitlog.Logger {
	l := kitlog.Default()
	return &l
}
