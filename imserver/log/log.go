/*
* created by Orz on 2017/6/11
 */
package log

import "fmt"
import . "im/public"

func Log(msg string) {
	fmt.Printf("%s: %s\n", GetCurrTime(), msg)
}
