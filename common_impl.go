package assert

type errorLogger interface {
	Log(location *location, message string)
}

type errorLoggerImpl struct {
	writer stringWriter
}

func (logger *errorLoggerImpl) Log(location *location, message string) {
}

type stringWriter interface {
	WriteString(s string) (n int, err error)
}

type location struct {
	Test     string
	FileName string
	Line     int
}

func logIfFalse(t TestingT, condition bool, format string, args ...interface{}) {
	if !condition {
		t.Errorf(format, args...)
	}
}
