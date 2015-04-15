package log

import (
	"runtime"
	"sync"
	"testing"
)

func TestLog(t *testing.T) {

	runtime.GOMAXPROCS(runtime.NumCPU())
	FatalOff()

	Debug("Debugs : ", len("fdsafs"), []string{"Debug", "Debug"}, 10, 20)
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
	FatalOff()

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
	FatalOff()

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

	SetLevel(100)

	var N int = 100000

	waitgroup := sync.WaitGroup{}
	for i := 0; i < N; i++ {
		waitgroup.Add(1)
		go func(i int) {
			defer func() {
				waitgroup.Done()
			}()

			Debug("Debug", i)
			DebugOff()
			DebugOn()
			Info("Info", i)
			InfoOff()
			InfoOn()
			Warn("Warn", i)
			WarnOff()
			WarnOn()
			Error("Error", i)
			ErrorOff()
			ErrorOn()

		}(i)
	}

	waitgroup.Wait()
}
