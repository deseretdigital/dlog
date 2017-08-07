package dlog

import "testing"

//
//import (
//	"bytes"
//	"testing"
//)
//
//type TestWriter struct {
//	message []byte
//}
//
//func (t *TestWriter) Write(p []byte) (n int, err error) {
//	t.message = p
//	return 0, nil
//}
//
//func TestLogger(t *testing.T) {
//
//	trace := new(TestWriter)
//	info := new(TestWriter)
//	warning := new(TestWriter)
//	err := new(TestWriter)
//
//	Init(trace, info, warning, err)
//	traceMessage := []byte("message a")
//	Trace(traceMessage)
//	infoMessage := []byte("message b")
//	Info(infoMessage)
//	warningMessage := []byte("message c")
//	Warning(warningMessage)
//	errorMessage := []byte("message d")
//	Error(errorMessage)
//
//	if bytes.HasSuffix(trace.message, traceMessage) {
//		t.Errorf("Expected %v got %v", traceMessage, trace.message)
//	}
//	if bytes.HasSuffix(info.message, traceMessage) {
//		t.Errorf("Expected %v got %v", infoMessage, info.message)
//	}
//	if bytes.HasSuffix(warning.message, warningMessage) {
//		t.Errorf("Expected %v got %v", warningMessage, warning.message)
//	}
//	if bytes.HasSuffix(err.message, errorMessage) {
//		t.Errorf("Expected %v got %v", errorMessage, err.message)
//	}
//
//}

func TestLogger2(t *testing.T) {
	config := Config{
		AppName: "A name is a name",
		Url:     "graylog.deseretdigital.com:12212",
		Verbose: true}
	InitGelf(config)

	Trace("So cool!")
	Info("Haha")
	Warning("Booyashakala")
	Error("Hey dog!")

}
