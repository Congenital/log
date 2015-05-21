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
	LOG_LEVEL = DEBUG_N
)

const (
	FILE_LEVEL = DEBUG_N
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

type IWrite interface {
	Write(string)
}

type ELevel struct {
	level int
}

type ILevel interface {
	GetLevel() int
	SetLevel(int)
}

type EStatus struct {
	status int
}

type IStatus interface {
	GetStatus() int
	SetStatus(int)
}

type EColor struct {
	color string
}

type IColor interface {
	GetColor() int
	SetColor(int)
}

type ELog struct {
	EColor
	ELevel
	EStatus
	IWrite
	log string
	IOn
	IOff

	statusLock sync.RWMutex
	writeLock  sync.RWMutex
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

func (this *ELog) SetLevel(level int) {
	this.level = level
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
	this.statusLock.RLock()
	defer this.statusLock.RUnlock()

	return this.status
}

func (this *ELog) SetColor(color string) {
	this.color = color
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

func (this *ELog) Format(status int, log ...interface{}) string {
	t1 := time.Now().UnixNano()
	var data string
	var f string

	if this.GetLevel() <= FILE_LEVEL {
		_, file, line, ok := runtime.Caller(3)
		if ok == true {
			files := strings.Split(file, "/src/")
			if len(files) >= 2 {
				f = files[len(files)-1]
			} else {
				f = file
			}

			f = fmt.Sprintf(" >> file: %s   line: %v", f, line)
		}
	}

	value := log[0]
	for _, v := range value.([]interface{}) {
		data += fmt.Sprintf("%v	", v)
	}

	data = this.log + " - " + time.Now().Format("2006-01-02 15:04:05") + f + "\n    " + data

	if status == 0 {
		data = LOG_START + this.GetColor() + "m" + data + LOG_END
	}

	data += "\n"

	if this.GetLevel() == FATAL_N {
		panic(data)
	}

	t2 := time.Now().UnixNano()
	fmt.Println(t2 - t1)
	return data
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

	levelLock      sync.RWMutex
	statusLock     sync.RWMutex
	writeLock      sync.RWMutex
	fileStatusLock sync.RWMutex

	file       *os.File
	fileStatus int
}

func (this *Log) GetStatus() int {
	this.statusLock.RLock()
	defer this.statusLock.RUnlock()

	return this.status
}

func (this *Log) SetStatus(status int) {
	this.statusLock.Lock()
	defer this.statusLock.Unlock()

	this.status = status
}

func (this *Log) SetLevel(level int) {
	this.levelLock.Lock()
	defer this.levelLock.Unlock()

	if level > LOG_LEVEL {
		this.level = LOG_LEVEL
		return
	}

	if level <= TRASH {
		this.level = TRASH
		return
	}

	this.level = level
}

func (this *Log) GetLevel() int {
	this.levelLock.RLock()
	defer this.levelLock.RUnlock()

	return this.level
}

func (this *Log) SetFile(file *os.File) {
	this.writeLock.Lock()
	defer this.writeLock.Unlock()

	this.file = file
	if this.file != nil {
		this.SetFileStatus(1)
	}
}

func (this *Log) GetFile() *os.File {
	this.writeLock.RLock()
	defer this.writeLock.RUnlock()

	return this.file
}

func (this *Log) SetFileStatus(status int) {
	this.fileStatusLock.Lock()
	defer this.fileStatusLock.Unlock()

	this.fileStatus = status
}

func (this *Log) GetFileStatus() int {
	this.fileStatusLock.RLock()
	defer this.fileStatusLock.RUnlock()

	return this.fileStatus
}

func (this *Log) On() {
	this.SetStatus(ON)
}

func (this *Log) Off() {
	this.SetStatus(OFF)
}

func (this *Log) Write(data string) (int, error) {

	if this.file == nil {
		return this.WriteTo(data, os.Stdout)
	}

	return this.WriteTo(data, this.file)
}

func (this *Log) WriteTo(data string, file *os.File) (int, error) {
	this.writeLock.Lock()
	defer this.writeLock.Unlock()

	return file.WriteString(data)
}

func (this *Log) Debug(log ...interface{}) {
	if this.status == OFF || this.level < this.Debug_log.GetLevel() || this.Debug_log.GetStatus() == OFF {
		return
	}

	//this.Debug_log.Log(log...)
	this.Write(this.Debug_log.Format(this.fileStatus, log))
}

func (this *Log) Info(log ...interface{}) {
	if this.status == OFF || this.level < this.Info_log.GetLevel() || this.Info_log.GetStatus() == OFF {
		return
	}

	//this.Info_log.Log(log...)
	this.Write(this.Info_log.Format(this.fileStatus, log))
}

func (this *Log) Warn(log ...interface{}) {
	if this.status == OFF || this.level < this.Warn_log.GetLevel() || this.Warn_log.GetStatus() == OFF {
		return
	}

	//this.Warn_log.Log(log...)
	this.Write(this.Warn_log.Format(this.fileStatus, log))
}

func (this *Log) Error(log ...interface{}) {
	if this.status == OFF || this.level < this.Error_log.GetLevel() || this.Warn_log.GetStatus() == OFF {
		return
	}

	//this.Error_log.Log(log...)
	this.Write(this.Error_log.Format(this.fileStatus, log))
}

func (this *Log) Fatal(log ...interface{}) {
	if this.status == OFF || this.level < this.Fatal_log.GetLevel() || this.Fatal_log.GetStatus() == OFF {
		return
	}

	//this.Fatal_log.Log(log...)
	this.Write(this.Fatal_log.Format(this.fileStatus, log))
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

func SetFile(file *os.File) {
	Loger.SetFile(file)
}

func GetFile() *os.File {
	return Loger.GetFile()
}

var Loger = &Log{
	Debug_log: NewELog(BLUE, DEBUG, DEBUG_N),
	Info_log:  NewELog(GREEN, INFO, INFO_N),
	Warn_log:  NewELog(YELLO, WARN, WARN_N),
	Error_log: NewELog(PURPLE_RED, ERROR, ERROR_N),
	Fatal_log: NewELog(RED, FATAL, FATAL_N),
}

func init() {

	Loger.SetLevel(LOG_LEVEL)
	Loger.On()
}
