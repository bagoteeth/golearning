package testCover

import (
	"golearning/testFunc"
	"testing"
)

func TestLocalCover(t *testing.T) {
	t.Logf("\nlocal sum is %d\n", TryLocalCover())
}

func TestExternalCover(t *testing.T) {
	t.Logf("\nexternal sum is %d\n", testFunc.TryCover())
}
