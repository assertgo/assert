package assert

import (
	"runtime"
	"strings"
)

func provideLocation() location {
	pc, _, _, _ := runtime.Caller(1)
	testName := parseTestName(pc)
	return location{
		Test:     testName,
		FileName: "location_provider_test.go",
	}
}

func parseTestName(pc uintptr) (testName string) {
	testName = runtime.FuncForPC(pc).Name()
	index := strings.LastIndex(testName, ".")
	testName = testName[index+1:]
	return
}
