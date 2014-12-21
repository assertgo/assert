package assert

import (
	"runtime"
	"strings"
)

func provideLocation() location {
	testName, path, line := testCallerInfo()
	file := fileFromPath(path)
	return location{
		Test:     testName,
		FileName: file,
		Line:     line,
	}
}

func testCallerInfo() (testName string, path string, line int) {
	testName, path, line = callerInfo(3)
	for skip := 4; !strings.HasPrefix(testName, "Test"); skip++ {
		testName, _, _ = callerInfo(skip)
	}
	return
}

func callerInfo(skip int) (testName string, path string, line int) {
	pc, path, line, _ := runtime.Caller(skip)
	testName = parseTestName(pc)
	return
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
