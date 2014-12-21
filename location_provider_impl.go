package assert

import (
	"runtime"
	"strings"
)

func provideLocation(skip int) location {
	testName, path, line := testCallerInfo(skip + 1)
	file := fileFromPath(path)
	return location{
		Test:     testName,
		FileName: file,
		Line:     line,
	}
}

func testCallerInfo(s int) (testName string, path string, line int) {
	testName, path, line = callerInfo(s + 2)
	for skip := s + 3; !strings.HasPrefix(testName, "Test"); skip++ {
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
