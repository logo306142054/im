package log

import (
	"fmt"
	. "im/public"
)

func Log(msg string) {
	fmt.Printf("%s: %s\n", GetCurrTime(), msg)
}
