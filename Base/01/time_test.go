package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCurrentTime(t *testing.T) {
	got := CurrentTime()
	fmt.Println(got)
	want := time.Since(got)
	fmt.Println(want)
	if want > 10*time.Second {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
