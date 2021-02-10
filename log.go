package gRouter

import "fmt"

type ILog interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

type log struct {
}

func (log *log) Debug(args ...interface{}) {
	fmt.Println(args...)
}

func (log *log) Info(args ...interface{}) {
	fmt.Println(args...)
}

func (log *log) Warn(args ...interface{}) {
	fmt.Println(args...)
}

func (log *log) Error(args ...interface{}) {
	fmt.Println(args...)
}

func (log *log) Fatal(args ...interface{}) {
	fmt.Println(args...)
}

func (log *log) Debugf(format string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(format, args))
}

func (log *log) Infof(format string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(format, args...))
}

func (log *log) Warnf(format string, args ...interface{}) {
}

func (log *log) Warningf(format string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(format, args))
}

func (log *log) Errorf(format string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(format, args))
}

func (log *log) Fatalf(format string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(format, args))
}
