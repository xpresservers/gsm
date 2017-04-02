package logger

import (
	"github.com/fatih/color"
	"os"
	"time"
	"fmt"
)

func prependArgs(elem interface{}, a ...interface{}) []interface{} {
	return append([]interface{}{elem}, a...)
}

func getNow() string {
	return time.Now().Format(time.RFC1123Z)
}

func withTime(pattern string, a ...interface{}) (newPattern string, newParams []interface{}) {
	newPattern = fmt.Sprintf("[%%s] %s\n", pattern)
	newParams = prependArgs(getNow(), a...)
	return
}

func Info(pattern string, a ...interface{}) {
	newPattern, newParams := withTime(pattern, a...)
	color.New(color.FgHiBlue).Printf(newPattern, newParams...)
}

func Warn(pattern string, a ...interface{}) {
	newPattern, newParams := withTime(pattern, a...)
	color.New(color.FgHiRed).Printf(newPattern, newParams...)
}

func Fatal(err error) {
	color.New(color.FgRed).Println(err)
	os.Exit(1)
}
