package assert

import (
	"runtime"
	"strings"
)

func provideLocation() location {
	pc, path, _, _ := runtime.Caller(1)
	testName := parseTestName(pc)
	file := fileFromPath(path)
	return location{
		Test:     testName,
		FileName: file,
	}
}

func fileFromPath(path string) (file string) {
	file = substringAfter(path, "/")
	return
}

func parseTestName(pc uintptr) (testName string) {
	fullTestName := runtime.FuncForPC(pc).Name()
	testName = substringAfter(fullTestName, ".")
	return
}

func substringAfter(str string, sep string) (substring string) {
	index := strings.LastIndex(str, sep)
	substring = str[index+1:]
	return
}
