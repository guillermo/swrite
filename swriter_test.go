package swriter

import (
	"bytes"
	"testing"
	"time"
)

func TestSlowWriter(t *testing.T) {

	buf := new(bytes.Buffer)

	sw := New(buf, time.Millisecond*10)
	defer sw.Close()

	// First write
	n, err := sw.Write([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
	if n != 5 {
		t.Fatal(n)
	}

	if buf.String() != "" {
		t.Fatal(buf.String())
	}
	time.Sleep(time.Millisecond * 5)
	if buf.String() != "" {
		t.Fatal(buf.String())
	}

	// Second write
	n, err = sw.Write([]byte("world"))
	if err != nil {
		t.Fatal(err)
	}
	if n != 5 {
		t.Fatal(n)
	}
	if buf.String() != "" {
		t.Fatal(buf.String())
	}

	time.Sleep(time.Millisecond * 3)
	if buf.String() != "" {
		t.Fatal(buf.String())
	}

	time.Sleep(time.Millisecond * 3)
	if buf.String() != "helloworld" {
		t.Fatal(buf.String())
	}
}
