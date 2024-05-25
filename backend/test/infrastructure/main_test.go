package infrastructure

import (
	"MelkOnline/internal"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	es := internal.NewEchoServer()
	go es.Start()
	time.Sleep(1 * time.Second)
	os.Exit(m.Run())
}
