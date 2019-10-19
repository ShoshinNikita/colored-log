package clog_test

import (
	stdlog "log"
	"os"
	"testing"

	clog "github.com/ShoshinNikita/log/v2"
)

const file = "test.txt"
const msg = "Hello, dear world!!!"

func BenchmarkStdLogPrintlnWithPrefixes(b *testing.B) {
	f, err := os.Create(file)
	if err != nil {
		stdlog.Fatalln(err)
	}
	defer f.Close()

	l := &stdlog.Logger{}

	// github.com/ShoshinNikita/log prints it by default
	l.SetFlags(stdlog.Lshortfile | stdlog.Ltime | stdlog.Ldate)
	l.SetOutput(f)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Println(msg)
	}
}

func BenchmarkDevLogPrintln(b *testing.B) {
	f, err := os.Create(file)
	if err != nil {
		stdlog.Fatalln(err)
	}
	defer f.Close()

	l := clog.NewDevConfig().SetOutput(f).Build()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Print(msg)
	}
}

func BenchmarkDevLogErrorln(b *testing.B) {
	f, err := os.Create(file)
	if err != nil {
		stdlog.Fatalln(err)
	}
	defer f.Close()

	l := clog.NewDevConfig().SetOutput(f).Build()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Error(msg)
	}
}

func BenchmarkProdLogPrintln(b *testing.B) {
	f, err := os.Create(file)
	if err != nil {
		stdlog.Fatalln(err)
	}
	defer f.Close()

	l := clog.NewProdConfig().SetOutput(f).Build()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Print(msg)
	}
}

func BenchmarkProdLogErrorln(b *testing.B) {
	f, err := os.Create(file)
	if err != nil {
		stdlog.Fatalln(err)
	}
	defer f.Close()

	l := clog.NewProdConfig().SetOutput(f).Build()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Error(msg)
	}
}
