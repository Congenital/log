package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	Debug([]string{"Debug", "Debug"})
	Info("Info")
	Warn("Warn")
	Error("Error")
	Fatal("Fatal")

	DebugOff()
	InfoOff()
	WarnOff()
	ErrorOff()
	FatalOff()

	Debug([]string{"Debug", "Debug"})
	Info("Info")
	Warn("Warn")
	Error("Error")
	Fatal("Fatal")

	DebugOn()
	InfoOn()
	WarnOn()
	ErrorOn()
	FatalOn()

	Debug([]string{"Debug", "Debug"})
	Info("Info")
	Warn("Warn")
	Error("Error")
	Fatal("Fatal")

	Off()
	Debug([]string{"Debug", "Debug"})
	Info("Info")
	Warn("Warn")
	Error("Error")
	Fatal("Fatal")

	On()
	Debug([]string{"Debug", "Debug"})
	Info("Info")
	Warn("Warn")
	Error("Error")
	Fatal("Fatal")

	Info(struct {
		a    int
		name string
	}{10, "qiang.sheng"})

	SetLevel(1)
	Debug([]string{"Debug", "Debug"})
	Info("Info")
	Warn("Warn")
	Error("Error")
	Fatal("Fatal")
}
