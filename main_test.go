package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestHelloWorld(t *testing.T) {
	if HelloWorld() != "hello world2" {
		t.Errorf("got %s expected %s", HelloWorld(), "hello world2")
	}
}
