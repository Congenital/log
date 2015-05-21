package log

import (
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
	"unsafe"
)

const (
	TRASH = iota
	FATAL_N
	ERROR_N
	WARN_N
	INFO_N
	DEBUG_N
)

var LOG_LEVEL = DEBUG_N

var FILE_LEVEL = DEBUG_N

const (
	OFF = iota
	ON
)

const (
	BLACK      = "30m"
	RED        = "31m" //fatal default color
	GREEN      = "32m" //info default color
	YELLO      = "33m" //warn default color
	BLUE       = "34m" //debug default color
	PURPLE_RED = "35m" //error default color
	CYAN_BLUE  = "36m"
	WHITE      = "37m"

	LOG_START = "\033[1;0;"
	LOG_END   = "\033[0m"

	DEBUG = LOG_START + BLUE
	INFO  = LOG_START + GREEN
	WARN  = LOG_START + YELLO
	ERROR = LOG_START + PURPLE_RED
	FATAL = LOG_START + RED
)

var ioSync sync.RWMutex

var output = os.Stdout
var outputsync sync.RWMutex

var allOn = true
var allsync sync.RWMutex

var debugOn = true
var debugsync sync.RWMutex

var infoOn = true
var infosync sync.RWMutex

var warnOn = true
var warnsync sync.RWMutex

var errorOn = true
var errorsync sync.RWMutex

var fatalOn = true
var fatalsync sync.RWMutex

func SetFile(file *os.File) {
	outputsync.Lock()
	defer outputsync.Unlock()

	output = file
}

func On() {
	allsync.Lock()
	defer allsync.Unlock()

	allOn = true

	DebugOn()
	InfoOn()
	WarnOn()
	ErrorOn()
	FatalOn()
}

func Off() {
	allsync.Lock()
	defer allsync.Unlock()

	allOn = false

	DebugOff()
	InfoOff()
	WarnOff()
	ErrorOff()
	FatalOff()
}

func GetStatus() bool {
	allsync.RLock()
	defer allsync.RUnlock()

	return allOn
}

func DebugOn() {
	debugsync.Lock()
	defer debugsync.Unlock()

	debugOn = true
}

func DebugOff() {
	debugsync.Lock()
	defer debugsync.Unlock()

	debugOn = false
}

func GetDebugStatus() bool {
	debugsync.RLock()
	defer debugsync.RUnlock()

	return debugOn
}

func InfoOn() {
	infosync.Lock()
	defer infosync.Unlock()

	infoOn = true
}

func InfoOff() {
	infosync.Lock()
	defer infosync.Unlock()

	infoOn = false
}

func GetInfoStatus() bool {
	infosync.RLock()
	defer infosync.RUnlock()

	return infoOn
}

func WarnOn() {
	warnsync.Lock()
	defer warnsync.Unlock()

	warnOn = true
}

func WarnOff() {
	warnsync.Lock()
	defer warnsync.Unlock()

	warnOn = false
}

func GetWarnStatus() bool {
	warnsync.RLock()
	defer warnsync.RUnlock()

	return warnOn
}

func ErrorOn() {
	errorsync.Lock()
	defer errorsync.Unlock()

	errorOn = true
}

func ErrorOff() {
	errorsync.Lock()
	defer errorsync.Unlock()

	errorOn = false
}

func GetErrorStatus() bool {
	errorsync.RLock()
	defer errorsync.RUnlock()

	return errorOn
}

func FatalOn() {
	fatalsync.Lock()
	defer fatalsync.Unlock()

	fatalOn = true
}

func FatalOff() {
	fatalsync.Lock()
	defer fatalsync.Unlock()

	fatalOn = false
}

func GetFatalStatus() bool {
	fatalsync.RLock()
	defer fatalsync.RUnlock()

	return fatalOn
}

func WriteTo(write io.Writer, data string) (int, error) {
	ioSync.Lock()
	defer ioSync.Unlock()

	return write.Write(*(*[]byte)(unsafe.Pointer(&data)))
}

func Write(data string) (int, error) {
	return WriteTo(output, data)
}

func WriteForStatus(write io.Writer, data string, status bool) (int, error) {
	if !status {
		return 0, errors.New("No Open Status!")
	}

	return WriteTo(write, data)
}

func Format(data ...interface{}) string {
	var datas string

	for _, v := range data {
		datas += fmt.Sprintf("%v	", v)
	}

	return datas + "\n"
}

func FileLine() string {
	var f string

	_, file, line, ok := runtime.Caller(2)
	if ok == true {
		files := strings.Split(file, "/src/")
		if len(files) >= 2 {
			f = files[len(files)-1]
		} else {
			f = file
		}

		f = fmt.Sprintf(" >> file: %s   line: %v", f, line)
	}

	return f
}

func Time() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func DebugNoLine(data ...interface{}) {
	WriteForStatus(output, DEBUG+Format(data...)+LOG_END, GetDebugStatus())
}

func InfoNoLine(data ...interface{}) {
	WriteForStatus(output, INFO+Format(data...)+LOG_END, GetInfoStatus())
}

func WarnNoLine(data ...interface{}) {
	WriteForStatus(output, WARN+Format(data...)+LOG_END, GetWarnStatus())
}

func ErrorNoLine(data ...interface{}) {
	WriteForStatus(output, ERROR+Format(data...)+LOG_END, GetErrorStatus())
}

func FatalNoLine(data ...interface{}) {
	WriteForStatus(output, FATAL+Format(data...)+LOG_END, GetFatalStatus())
	panic(FATAL + "\n   " + Format(data...) + LOG_END)
}

func Debug(data ...interface{}) {
	WriteForStatus(output, DEBUG+"	-	"+Time()+FileLine()+"\n"+Format(data...)+LOG_END, GetDebugStatus())
}

func Info(data ...interface{}) {
	WriteForStatus(output, INFO+"  -   "+Time()+FileLine()+"\n"+Format(data...)+LOG_END, GetInfoStatus())
}

func Warn(data ...interface{}) {
	WriteForStatus(output, WARN+"  -   "+Time()+FileLine()+"\n"+Format(data...)+LOG_END, GetWarnStatus())
}

func Error(data ...interface{}) {
	WriteForStatus(output, ERROR+"  -   "+Time()+FileLine()+"\n"+Format(data...)+LOG_END, GetErrorStatus())
}

func Fatal(data ...interface{}) {
	WriteForStatus(output, FATAL+"  -   "+Time()+FileLine()+"\n"+Format(data...)+LOG_END, GetFatalStatus())
	if GetFatalStatus() {
		panic(FATAL + "\n   " + "  -   " + Time() + Format(data...) + LOG_END)
	}
}

func DebugTo(write io.Writer, data ...interface{}) {
	WriteForStatus(write, "  -   "+Time()+FileLine()+"\n"+Format(data...), GetDebugStatus())
}

func InfoTo(write io.Writer, data ...interface{}) {
	WriteForStatus(write, "  -   "+Time()+FileLine()+"\n"+Format(data...), GetInfoStatus())
}

func WarnTo(write io.Writer, data ...interface{}) {
	WriteForStatus(write, "  -   "+Time()+FileLine()+"\n"+Format(data...), GetWarnStatus())
}

func ErrorTo(write io.Writer, data ...interface{}) {
	WriteForStatus(write, "  -   "+Time()+FileLine()+"\n"+Format(data...), GetErrorStatus())
}

func FatalTo(write io.Writer, data ...interface{}) {
	WriteForStatus(write, "  -   "+Time()+FileLine()+"\n"+Format(data...), GetFatalStatus())
	if GetFatalStatus() {
		panic(FATAL + "\n   " + Format(data...) + LOG_END)
	}
}
