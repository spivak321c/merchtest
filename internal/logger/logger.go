package logger

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

func Init(dev bool) error {
	if dev {
		l, err := zap.NewDevelopment()
		if err != nil { return err }
		Log = l
		return nil
	}
	l, err := zap.NewProduction()
	if err != nil { return err }
	Log = l
	return nil
}