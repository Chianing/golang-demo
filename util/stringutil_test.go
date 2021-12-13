package util

import (
	"fmt"
	"testing"
)

func TestTrimEscapeSign(t *testing.T) {
	str := "{\"id\": 1,\n\"name\":\"test\"}"
	fmt.Println(TrimEscapeSign(str))
}
