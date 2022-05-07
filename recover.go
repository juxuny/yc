package yc

import "runtime/debug"

func Collect() {
	if err := recover(); err != nil {
		debug.PrintStack()
	}
}
