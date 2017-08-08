package dlog

import (
	"bytes"
	"fmt"
	"github.com/Graylog2/go-gelf/gelf"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var (
	trace    *log.Logger
	info     *log.Logger
	warning  *log.Logger
	errorLog *log.Logger
	fatalLog *log.Logger
)

type Config struct {
	Verbose      bool
	Url, AppName string
}

func init() {
	InitConsole()
}

func Tracef(format string, v ...interface{}) {
	trace.Output(2, fmt.Sprintf(format, v...))
}
func Trace(v ...interface{}) {
	trace.Output(2, fmt.Sprint(v...))
}

func Infof(format string, v ...interface{}) {
	info.Output(2, fmt.Sprintf(format, v...))
}
func Info(v ...interface{}) {
	info.Output(2, fmt.Sprint(v...))
}
func Warningf(format string, v ...interface{}) {
	warning.Output(2, fmt.Sprintf(format, v...))
}
func Warning(v ...interface{}) {
	warning.Output(2, fmt.Sprint(v...))
}
func Errorf(format string, v ...interface{}) {
	errorLog.Output(2, fmt.Sprintf(format, v...))

}
func Error(v ...interface{}) {
	errorLog.Output(2, fmt.Sprint(v...))
}
func Fatalf(format string, v ...interface{}) {
	fatalLog.Output(2, fmt.Sprintf(format, v...))
}
func Fatal(v ...interface{}) {
	fatalLog.Output(2, fmt.Sprint(v...))
}

// InitDiscard is an easy default to clear out all logging
func InitDiscard() {
	trace = log.New(ioutil.Discard,
		"trace: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	info = log.New(ioutil.Discard,
		"info: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	warning = log.New(ioutil.Discard,
		"warning: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	errorLog = log.New(ioutil.Discard,
		"error: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	fatalLog = log.New(ioutil.Discard,
		"error: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

// InitConsole is a simple default to log to console
func InitConsole() {
	trace = log.New(os.Stdout,
		"trace: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	info = log.New(os.Stdout,
		"info: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	warning = log.New(os.Stderr,
		"warning: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	errorLog = log.New(os.Stderr,
		"error: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	fatalLog = log.New(os.Stderr,
		"fatal: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func Init(
	traceHandle io.Writer ,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorLogHandle io.Writer,
	fatalLogHandler io.Writer) {

	trace = log.New(traceHandle,
		"trace: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	info = log.New(infoHandle,
		"info: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	warning = log.New(warningHandle,
		"warning: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	errorLog = log.New(errorLogHandle,
		"error: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	fatalLog = log.New(fatalLogHandler,
		"error: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func InitGelf(config Config) {
	traceWriter := ioutil.Discard
	if config.Verbose {
		traceWriter = io.MultiWriter(os.Stdout, NewGelfWriter(config.Url, config.AppName, 7))
	}
	infoWriter := io.MultiWriter(os.Stdout, NewGelfWriter(config.Url, config.AppName, 6))
	warnWriter := io.MultiWriter(os.Stdout, NewGelfWriter(config.Url, config.AppName, 4))
	errorWriter := io.MultiWriter(os.Stdout, NewGelfWriter(config.Url, config.AppName, 3))
	fatalWriter := io.MultiWriter(os.Stdout, NewGelfWriter(config.Url, config.AppName, 2))
	if !config.Verbose {
		traceWriter = ioutil.Discard
	}
	Init(traceWriter, infoWriter, warnWriter, errorWriter, fatalWriter)
}

type gelfWriter struct {
	LoggingLevel int
	AppName, Url string
	GelfWriter   *gelf.Writer
}

func (t *gelfWriter) Write(p []byte) (n int, err error) {
	err = nil
	n = len(p)
	short := p
	full := []byte("")
	if i := bytes.IndexRune(p, '\n'); i > 0 {
		short = p[:i]
		full = p
	}
	m := gelf.Message{
		Version:  "-",
		Host:     t.Url,
		Short:    string(short),
		Full:     string(full),
		TimeUnix: float64(time.Now().Unix()),
		Level:    int32(t.LoggingLevel), // info
		//Facility: w.Facility,
		Extra: map[string]interface{}{
			"_AppName": t.AppName,
		},
	}
	if err = t.GelfWriter.WriteMessage(&m); err != nil {
		return 0, err
	}
	return

}

func NewGelfWriter(url, appName string, level int) *gelfWriter {
	writer := new(gelfWriter)
	writer.LoggingLevel = level
	writer.AppName = appName
	writer.Url = url
	gelfWriter, err := gelf.NewWriter(url)
	if err != nil {
		panic(err)
	}
	writer.GelfWriter = gelfWriter
	return writer
}
