package turner_test

import (
	"testing"

	"github.com/alisdairrankine/turner"
)

func TestAdd(t *testing.T) {
	v1 := turner.Vec3{1, 2, 3}
	v2 := v1.Add(turner.Vec3{3, 2, 1})

	if v2.X != 4 {
		t.Log("X is incorrect")
		t.Fail()
	}

	if v2.Y != 4 {
		t.Log("Y is incorrect")
		t.Fail()
	}

	if v2.Z != 4 {
		t.Log("Z is incorrect")
		t.Fail()
	}
}
