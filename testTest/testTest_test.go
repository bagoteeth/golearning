package testTest

import (
	"golearning/testFunc"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	t.Log("helloworld")
}

func TestByeWorld(t *testing.T) {
	t.Log("byeworld")
}

func TestByyyy(t *testing.T) {
	t.Log("byyyyyyy")
}

func TestCount(t *testing.T) {
	sum := testFunc.TryCover()
	t.Logf("sum is %d\n", sum)
}
