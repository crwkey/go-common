package pipeline

import (
	"testing"
	"time"
)

func TestSortInMem(t *testing.T) {
	defer func(now time.Time) {
		t.Log(time.Since(now))
	}(time.Now())
	p := SortInMem(ArraySource(10, -1, -3, 20, 30))
	for e := range p {
		t.Log(e)
	}
}
