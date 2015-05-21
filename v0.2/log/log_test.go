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

	SetFile(file)

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

	Debug([]string{"Debug", "Debug"})
	Info("Info")
	Warn("Warn")
	Error("Error")
	Fatal("Fatal")

	Fatal("fdasfasd", "fdasfasdfasd", "fdsfsda", 321, 321)

	waitgroup := sync.WaitGroup{}
	for i := 0; i < 100000; i++ {
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

func BenchmarkWriteTo(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	for i := 0; i < b.N; i++ {
		go WriteTo(os.Stdout, "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\n")
	}
}

func BenchmarkWrite(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	for i := 0; i < b.N; i++ {
		go Write("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\n")
	}
}

func BenchmarkFormat(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	for i := 0; i < b.N; i++ {
		Format(struct {
			Name string
			Age  int
		}{
			"qiang.sheng",
			10,
		}, "fdafasd", 10)
	}
}

func BenchmarkDebug(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	for i := 0; i < b.N; i++ {
		Debug(struct {
			Name string
			Age  int
		}{
			"qiang.sheng",
			10,
		}, "fdafasd", 10)
	}
}

func BenchmarkDebugNoLine(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	for i := 0; i < b.N; i++ {
		go DebugNoLine(struct {
			Name string
			Age  int
		}{
			"qiang.sheng",
			10,
		}, "fdafasd", 10)
	}
}

func BenchmarkDebugTo(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	for i := 0; i < b.N; i++ {
		go DebugTo(os.Stdout, struct {
			Name string
			Age  int
		}{
			"qiang.sheng",
			10,
		}, "fdafasd", 10)
	}
}
