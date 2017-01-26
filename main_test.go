package main

import (
	"testing"
	"fmt"
)

func TestMain(t *testing.T) {
	LogFatal = func(v ...interface{}) {
		fmt.Println(v...)
	}
	main()
}
