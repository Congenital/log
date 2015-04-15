/*
 * author : qiang.sheng@godinsec.com
 * time : 2014-4-13
 *
 * info : log function
 * level : {
 *			DEBUG < INFO < WARN < ERROR < FATAL
 *			5		4		3		2		1
 *		}
 *
 *
 */
package log

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	TRASH = iota
	FATAL_N
	ERROR_N
	WARN_N
	INFO_N
	DEBUG_N
	LOG_LEVEL
)

const (
	FILE_LEVEL = WARN_N
)

const (
	OFF = iota
	ON
)

const (
	BLACK      = "30"
	RED        = "31" //fatal default color
	GREEN      = "32" //info default color
	YELLO      = "33" //warn default color
	BLUE       = "34" //debug default color
	PURPLE_RED = "35" //error default color
	CYAN_BLUE  = "36"
	WHITE      = "37"

	DEBUG = "[DEBUG	]"
	INFO  = "[INFO	]"
	WARN  = "[WARN	]"
	ERROR = "[ERROR	]"
	FATAL = "[FATAL	]"

	LOG_START = "\033[1;0;"
	LOG_END   = "\033[0m"
)

type IOn interface {
	On()
}

type IOff interface {
	Off()
}

type ILog interface {
	Log(...interface{})
}

type IDebug interface {
	Debug(...interface{})
}

type IInfo interface {
	Info(...interface{})
}

type IWarn interface {
	Warn(...interface{})
}

type IError interface {
	Error(...interface{})
}

type IFatal interface {
	Fatal(...interface{})
}

type ELevel struct {
	level int
}

type IWrite interface {
	Write(...interface{})
}

type EStatus struct {
	status int
}

type EColor struct {
	color string
}

type ELog struct {
	EColor
	ELevel
	EStatus
	IWrite
	log string
	IOn
	IOff

	statusLock sync.Mutex
	writeLock  sync.Mutex
}

func NewELog(color string, log string, level int) *ELog {

	return &ELog{
		EColor: EColor{
			color: color,
		},

		ELevel: ELevel{
			level: level,
		},

		EStatus: EStatus{
			status: ON,
		},

		log: log,
	}
}

func (this *ELog) Log(log ...interface{}) {
	if this.GetStatus() == OFF {
		return
	}

	this.Write(log)
}

func (this *ELog) GetLevel() int {
	return this.level
}

func (this *ELog) SetStatus(status int) {
	this.statusLock.Lock()
	defer this.statusLock.Unlock()

	this.status = status
}

func (this *ELog) GetStatus() int {
	this.statusLock.Lock()
	defer this.statusLock.Unlock()

	return this.status
}

func (this *ELog) GetColor() string {
	return this.color
}

func (this *ELog) On() {
	this.SetStatus(ON)
}

func (this *ELog) Off() {
	this.SetStatus(OFF)
}

func (this *ELog) Write(log ...interface{}) {
	var data string

	start := LOG_START + this.GetColor() + "m" + this.log + " - " + time.Now().Format("2006-01-02 15:04:05")
	var f string

	if this.GetLevel() <= FILE_LEVEL {
		_, file, line, ok := runtime.Caller(4)
		if ok == true {
			files := strings.Split(file, "/src/")
			if len(files) >= 2 {
				f = files[len(files)-1]
			} else {
				f = file
			}

			f = fmt.Sprintf(" >> file: %s	line: %v", f, line)
		}
	}

	value := log[0]
	for _, v := range value.([]interface{}) {
		data += fmt.Sprintf("%v", v)
	}

	data = fmt.Sprintf("%v", start+f+"\n	"+data+LOG_END+"\n")

	if this.GetLevel() == FATAL_N {
		panic(data)
	}

	this.StdoutWrite(data)
}

func (this *ELog) StdoutWrite(data string) (int, error) {
	this.writeLock.Lock()
	defer this.writeLock.Unlock()

	return os.Stdout.WriteString(data)
}

type Log struct {
	EStatus
	ELevel
	IOn
	IOff
	Debug_log *ELog
	Info_log  *ELog
	Warn_log  *ELog
	Error_log *ELog
	Fatal_log *ELog
	IDebug
	IInfo
	IWarn
	IError
	IFatal

	statusLock sync.Mutex
}

func (this *Log) SetLevel(level int) {
	if level >= LOG_LEVEL {
		this.level = LOG_LEVEL
		return
	}

	if level <= TRASH {
		this.level = TRASH
		return
	}

	this.level = level
}

func (this *Log) GetStatus() int {
	this.statusLock.Lock()
	defer this.statusLock.Unlock()

	return this.status
}

func (this *Log) SetStatus(status int) {
	this.statusLock.Lock()
	defer this.statusLock.Unlock()

	this.status = status
}

func (this *Log) On() {
	this.SetStatus(ON)
}

func (this *Log) Off() {
	this.SetStatus(OFF)
}

func (this *Log) Debug(log ...interface{}) {
	if this.status == OFF || this.level < this.Debug_log.GetLevel() {
		return
	}

	this.Debug_log.Log(log...)
}

func (this *Log) Info(log ...interface{}) {
	if this.status == OFF || this.level < this.Info_log.GetLevel() {
		return
	}

	this.Info_log.Log(log...)
}

func (this *Log) Warn(log ...interface{}) {
	if this.status == OFF || this.level < this.Warn_log.GetLevel() {
		return
	}

	this.Warn_log.Log(log...)
}

func (this *Log) Error(log ...interface{}) {
	if this.status == OFF || this.level < this.Error_log.GetLevel() {
		return
	}

	this.Error_log.Log(log...)
}

func (this *Log) Fatal(log ...interface{}) {
	if this.status == OFF || this.level < this.Fatal_log.GetLevel() {
		return
	}

	this.Fatal_log.Log(log...)
}

var Loger = &Log{
	Debug_log: NewELog(BLUE, DEBUG, DEBUG_N),
	Info_log:  NewELog(GREEN, INFO, INFO_N),
	Warn_log:  NewELog(YELLO, WARN, WARN_N),
	Error_log: NewELog(PURPLE_RED, ERROR, ERROR_N),
	Fatal_log: NewELog(RED, FATAL, FATAL_N),
}

func Debug(log ...interface{}) {
	Loger.Debug(log...)
}

func Info(log ...interface{}) {
	Loger.Info(log...)
}

func Warn(log ...interface{}) {
	Loger.Warn(log...)
}

func Error(log ...interface{}) {
	Loger.Error(log...)
}

func Fatal(log ...interface{}) {
	Loger.Fatal(log...)
}

func DebugOn() {
	Loger.Debug_log.On()
}

func DebugOff() {
	Loger.Debug_log.Off()
}

func InfoOff() {
	Loger.Info_log.Off()
}

func InfoOn() {
	Loger.Info_log.On()
}

func WarnOff() {
	Loger.Warn_log.Off()
}

func WarnOn() {
	Loger.Warn_log.On()
}

func ErrorOff() {
	Loger.Error_log.Off()
}

func ErrorOn() {
	Loger.Error_log.On()
}

func FatalOff() {
	Loger.Fatal_log.Off()
}

func FatalOn() {
	Loger.Fatal_log.On()
}

func On() {
	Loger.On()
}

func Off() {
	Loger.Off()
}

func SetLevel(level int) {
	Loger.SetLevel(level)
}

func init() {
	if Loger == nil {
		Loger = &Log{
			Debug_log: NewELog(BLUE, DEBUG, DEBUG_N),
			Info_log:  NewELog(GREEN, INFO, INFO_N),
			Warn_log:  NewELog(YELLO, WARN, WARN_N),
			Error_log: NewELog(PURPLE_RED, ERROR, ERROR_N),
			Fatal_log: NewELog(RED, FATAL, FATAL_N),
		}
	}

	Loger.SetLevel(LOG_LEVEL)
	Loger.On()
}
