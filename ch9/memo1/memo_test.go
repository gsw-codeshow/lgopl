package memo1_test

import (
	"lgopl/ch9/memo1"
	"lgopl/ch9/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo1.New(httpGetBody)
	memotest.Sequential(t, m)
}

// NOTE: not concurrency-safe!  Test fails.
func TestConcurrent(t *testing.T) {
	m := memo1.New(httpGetBody)
	memotest.Concurrent(t, m)
}
