package assert

import (
	"runtime"
	"strings"
)

func provideLocation(skip int) *location {
	testName, path, line := testCallerInfo(skip + 1)
	fileName := fileNameFromPath(path)
	return &location{
		Test:     testName,
		FileName: fileName,
		Line:     line,
	}
}

func testCallerInfo(skip int) (testName string, path string, line int) {
	testName, path, line = callerInfo(skip + 1)
	for i := skip + 2; !strings.HasPrefix(testName, "Test"); i++ {
		testName, _, _ = callerInfo(i)
	}
	return
}

func callerInfo(skip int) (methodName string, path string, line int) {
	pc, path, line, ok := runtime.Caller(skip + 1)
	if !ok {
		panic("you shall not pass!")
	}
	methodName = methodNameFromPC(pc)
	return
}

func methodNameFromPC(pc uintptr) string {
	fullName := runtime.FuncForPC(pc).Name()
	return substringAfterLast(fullName, ".")
}

func fileNameFromPath(path string) string {
	return substringAfterLast(path, "/")
}

func substringAfterLast(str string, separator string) string {
	index := strings.LastIndex(str, separator)
	return str[index+1:]
}
