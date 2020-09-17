package main

import (
	"flag"
	"testing"

	"go.uber.org/goleak"
)

func TestRun(t *testing.T) {
	goleak.VerifyNone(t, goleak.IgnoreCurrent())
	err := run([]string{"", "-h"})
	if err != flag.ErrHelp {
		t.Fatalf("err should be flag.ErrHelp: %v", err)
	}
}
