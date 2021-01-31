package consistent

import (
	"math"
	"testing"
)

func TestGetCorrectNearestKey(t *testing.T) {

	keyList := []string{
		"8229115871",
		"2801799973",
		string(math.MaxInt32),
		"468396611",
		"822911587",
		"468396617",
	}
	rawKey := "7"

	consistent := NewConsistent(keyList)
	nearestKey := consistent.GetNearestKey(rawKey)

	if nearestKey != "2801799973" {
		t.Fail()
	}

}
