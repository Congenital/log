package log

import (
	l "log"
	"os"
	"runtime"
	"sync"
	"testing"
)

func BenchmarkLog(b *testing.B) {

	runtime.GOMAXPROCS(runtime.NumCPU())
	FatalOff()
	file, err := os.Create("test.log")
	if err != nil {
		Fatal(err)
	}
	defer file.Close()

	//SetFile(file)

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

	Fatal("fdasfasd", "fdasfasdfasd", "fdsfsda", 321, 321)

	var N int = 100000

	waitgroup := sync.WaitGroup{}
	for i := 0; i < N; i++ {
		waitgroup.Add(1)
		go func(i int) {
			defer func() {
				waitgroup.Done()
			}()
			s := `fjkdlsfjklsdafjsdklafjdklsafjsdlalf`

			Debug(s, i)
			DebugOff()
			DebugOn()
			Info(s, i)
			InfoOff()
			InfoOn()
			Warn(s, i)
			WarnOff()
			WarnOn()
			Error(s, i)
			ErrorOff()
			ErrorOn()

		}(i)
	}

	waitgroup.Wait()
}

func BenchmarkOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Debug("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	}
}

func BenchmarkTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l.Printf("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	}
}

func BenchmarkElog(b *testing.B) {
	elog := &ELog{}
	for i := 0; i < b.N; i++ {
		func(log ...interface{}) {
			l.Printf(elog.Format(0, log))
		}("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	}
}
