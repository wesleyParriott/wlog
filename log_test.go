package wlog

import (
	"os"
	"testing"
)

func TestLogs(t *testing.T) {
	wlog := CreateWlogWithParams(os.Stdout, 0)
	wlog.Debug("hi %s", "me")
	wlog.Info("hi %s", "me")
	wlog.Error("hi %s", "me")
	wlog.Fatal("hi %s", "me")
}
