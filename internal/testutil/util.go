package testutil

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

// JSONEq returns true if JSON strings a & b are equal.
// Returns false if any of them is not a valid json
func JSONEq(a, b string) bool {
	var am, bm interface{}
	if err := json.Unmarshal([]byte(a), &am); err != nil {
		return false
	}
	if err := json.Unmarshal([]byte(b), &bm); err != nil {
		return false
	}

	res := reflect.DeepEqual(am, bm)
	if !res {
		fmt.Printf("JSONEq mismatch:\nACTUAL:   %s\nEXPECTED: %s\n", a, b)
	}
	return res
}

func TimePtr(t time.Time) *time.Time {
	return &t
}
